const visit = require(`unist-util-visit`)
const axios = require(`axios`)
const {spawn} = require('child_process');


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
    var dotData
    try {
      // Request the diagram
      console.log(`Making request`)
      const result = await axios.post(endpoint, model)
      dotData = result.data
    } catch (error) {
      // The request failed
      console.log(`Request to ${endpoint} failed: ${error}`)
      reject(error)
    }
    // Create a subcommand to call dot
    const dotProcess = spawn('dot', ['-Tsvg'], {stdio: ['pipe', 'pipe', process.stderr]}); // (A)
    // Pipe in our diagram
    dotProcess.stdin.write(dotData)
    dotProcess.stdin.end()
    // Fetch the response
    let svg = ''
    for await (const data of dotProcess.stdout) {
      svg += data.toString()
    };
    // Update the node
    node.type = 'html'
    node.children = undefined
    node.value = svg
  }
  resolve(markdownAST)
})
