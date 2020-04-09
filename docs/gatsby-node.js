/**
 * Implement Gatsby's Node APIs in this file.
 *
 * See: https://www.gatsbyjs.org/docs/node-apis/
 */
const axios = require('axios')
const { spawn } = require('child_process');
const { v4 } = require('uuid')
const crypto = require('crypto')

exports.onCreateNode = async ({ node, actions }) => {
  // Short-circuit if we're not considering a yaml file
  if (node.internal.type !== `DataYaml`) {
    return
  }
  // Get the archie model string
  const model = node.model
  for (const diagram of node.diagrams) {
    // Get the arguments
    const args = {type: diagram.type}
    args.scope = diagram.hasOwnProperty('scope') ? diagram.scope : null
    args.tag = diagram.hasOwnProperty('tag') ? diagram.tag : null
    // Get the dot graph
    const graphviz = await requestGraphviz(model, args)
    // Convert the dot graph to svg
    const svg = await convertToSvg(graphviz)
    // Create a new node
    const diagramNode = actions.createNode({
      id: v4(),
      value: svg,
      args: args,
      parent: node.id,
      children: [],
      internal: {
        mediaType: `image/svg+xml`,
        type: `SvgImage`,
        contentDigest: crypto
          .createHash(`md5`)
          .update(JSON.stringify(svg))
          .digest(`hex`),
        description: `Image in SVG format`,
      },
    })
    console.log(`node created`)
    // Link the new node to its parent
    actions.createParentChildLink({parent: node, child: diagramNode})
  }
  console.log(`finished!`)
}

async function requestGraphviz(model, args) {
  const endpoint = getEndpoint(args)
  try {
    // Request the diagram
    console.log(`Making request...`)
    const result = await axios.post(endpoint, model)
    console.log(`request complete!`)
    return result.data
  } catch (error) {
    // The request failed
    throw new Error(`Request to ${endpoint} failed: ${error}`)
  }
}

async function convertToSvg(graphviz) {
  // Create a subcommand to call dot
  console.log(`Converting with dot...`)
  const dotProcess = spawn('dot', ['-Tsvg'], {stdio: ['pipe', 'pipe', process.stderr]}); // (A)
  // Pipe in our diagram
  dotProcess.stdin.write(graphviz)
  dotProcess.stdin.end()
  // Fetch the response
  let svg = ''
  for await (const data of dotProcess.stdout) {
    svg += data.toString()
  };
  console.log(`conversion complete!`)
  return svg;
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
