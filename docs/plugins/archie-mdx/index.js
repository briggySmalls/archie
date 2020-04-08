const visit = require(`unist-util-visit`)
const axios = require(`axios`)
const dot = require('graphlib-dot')

function parseArgs(args) {
  switch (args.length) {
    case 0:
      throw new Error(`No args supplied to archie block`)
    case 1:
      return {type: args[0], scope: null, tag: null}
    case 2:
      return {type: args[0], scope: args[1], tag: null}
    case 3:
      return {type: args[0], scope: args[1], tag: args[2]}
    default:
      throw new Error(`Too many arguments supplied: ${args}`)
  }
}

function getEndpoint(args) {
  // Get the base endpoint
  let endpoint = new URL(`http://localhost:3000/diagram/${args.type}`)
  // Add parameters
  let params = []
  if (args.scope) {
    params.push(`scope=${encodeURIComponent(args.scope)}`)
  }
  if (args.tag) {
    params.push(`tag=${encodeURIComponent(args.tag)}`)
  }
  if (params.length > 0) {
    endpoint.search = `?${params.join('&')}`
  }
  // Return the string
  return endpoint.toString()
}

module.exports = ({ markdownAST }, pluginOptions) => new Promise(async (resolve, reject) => {
  // Synchronously sweep through nodes to change
  const nodesToChange = []
  // Visit all code blocks
  visit(markdownAST, `code`, node => {
    // We only care about code with the 'archie' language
    if (node.lang !== `archie`) {
      return
    }
    nodesToChange.push(node)
  })
  // Now perform asynchronous requests
  for (const node of nodesToChange) {
    // Fetch arguments from the meta
    let args
    try {
      args = parseArgs(node.meta.split(" "))
    } catch (error) {
      reject(error)
    }
    // Get the model from the code block content
    const model = node.value
    // Make the request
    const endpoint = getEndpoint(args)
    try {
      // Request the diagram
      console.log(`Making request`)
      const result = await axios.post(endpoint, model)
      // Read the content as a graphviz graph
      const g = dot.read(result.data);
      // Seriealise the graph
      const data = JSON.stringify(dot.graphlib.json.write(g))
      // Update the node
      node.type = 'html'
      node.children = undefined
      node.value = `<div class="graphviz"><script>${data}</script></div>`
    } catch (error) {
      // The request failed
      console.log(`Request to ${endpoint} failed: ${error}`)
      reject(error)
    }
  }
  resolve(markdownAST)
})
