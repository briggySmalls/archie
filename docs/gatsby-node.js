/**
 * Implement Gatsby's Node APIs in this file.
 *
 * See: https://www.gatsbyjs.org/docs/node-apis/
 */
const axios = require('axios')
const { spawn } = require('child_process');
const { v4 } = require('uuid')
const crypto = require('crypto')
const winston = require('winston')

const logger = winston.createLogger({
  level: (process.env.NODE_ENV !== 'production') ? 'debug' : 'info',
  format: winston.format.json(),
  defaultMeta: { service: 'user-service' },
  transports: [
    new winston.transports.Console({format: winston.format.simple()}),
  ]
});

exports.onCreateNode = async ({ node, getNode, actions }) => {
  // Short-circuit if we're not considering a yaml file
  if (node.internal.type !== `DataYaml`) {
    return
  }
  // Get the archie model string
  const model = node.model
  // Get the filepath
  const fileNode = getNode(node.parent)
  // Create an archie model node
  createModel(actions.createNode, model, fileNode.name, node.id)
  // Now create diagram nodes
  for (const diagram of node.diagrams) {
    // Get the arguments
    const args = {type: diagram.type}
    args.scope = diagram.hasOwnProperty('scope') ? diagram.scope : null
    args.tag = diagram.hasOwnProperty('tag') ? diagram.tag : null
    args.name = fileNode.name
    // Get the dot graph
    const graphviz = await requestGraphviz(model, args)
    // Convert the dot graph to svg
    const svg = await convertToSvg(graphviz)
    // Create a node
    const diagramNode = createSvg(actions.createNode, svg, args, node.id)
    // Link the new node to its parent
    actions.createParentChildLink({parent: node, child: diagramNode})
  }
  logger.debug(`finished!`)
}

async function requestGraphviz(model, args) {
  const endpoint = getEndpoint(args)
  try {
    // Request the diagram
    logger.debug(`Making request...`)
    const result = await axios.post(endpoint, model)
    logger.debug(`request complete!`)
    return result.data
  } catch (error) {
    // The request failed
    throw new Error(`Request to ${endpoint} failed: ${error}`)
  }
}

async function convertToSvg(graphviz) {
  // Create a subcommand to call dot
  logger.debug(`Converting with dot...`)
  const dotProcess = spawn('dot', ['-Tsvg'], {stdio: ['pipe', 'pipe', process.stderr]}); // (A)
  // Pipe in our diagram
  dotProcess.stdin.write(graphviz)
  dotProcess.stdin.end()
  // Fetch the response
  let svg = ''
  for await (const data of dotProcess.stdout) {
    svg += data.toString()
  };
  logger.debug(`conversion complete!`)
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

function createSvg(create, svg, args, parentId) {
  return create({
    id: v4(),
    value: svg,
    args: args,
    parent: parentId,
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
}

function createModel(create, model, name, parentId) {
  return create({
    id: v4(),
    value: model,
    name: name,
    parent: parentId,
    children: [],
    internal: {
      mediaType: `text/yaml`,
      type: `ArchieModel`,
      contentDigest: crypto
        .createHash(`md5`)
        .update(JSON.stringify(model))
        .digest(`hex`),
      description: `Archie yaml model`,
    },
  })
}
