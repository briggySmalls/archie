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
const { createFilePath } = require(`gatsby-source-filesystem`)

const logger = winston.createLogger({
  level: (process.env.NODE_ENV !== 'production') ? 'debug' : 'info',
  format: winston.format.json(),
  defaultMeta: { service: 'user-service' },
  transports: [
    new winston.transports.Console({format: winston.format.simple()}),
  ]
});

exports.createSchemaCustomization = ({ actions }) => {
  const { createTypes } = actions
  const typeDefs = `
    """
    Archie Node
    """
    type ArchieModel implements Node @infer {
      name: String!
      value: String!
    }

    """
    SVG Node
    """
    type SvgImage implements Node @infer {
      name: String!
      args: Args!
      value: String!
    }

    """
    Args Node
    """
    type Args {
      name: String!
      scope: String
      tag: String
    }
  `
  createTypes(typeDefs)
}

exports.onCreateNode = async ({ node, getNode, actions }) => {
  // Pull out the actions we care about
  const { createNode, createNodeField, createParentChildLink } = actions
  // Process nodes of interest
  switch (node.internal.type) {
    case `DataYaml`:
      // Convert yaml files into archie nodes
      await processArchie(node, getNode, createNode, createParentChildLink)
      return
    case `Mdx`:
      // Add a slug to Mdx pages
      addSlugToMdx(node, getNode, createNodeField)
      return
    default:
      return
  }
}

function addSlugToMdx(node, getNode, createNodeField) {
  // Detemine the slug from the node
  const slug = createFilePath({ node, getNode, basePath: `pages` })
  console.log(slug)
  // Add the slug to the Mdx node
  createNodeField({
    node,
    name: `slug`,
    value: slug,
  })
}

async function processArchie(node, getNode, createNode, createParentChildLink) {
  // Get the archie model string
  const model = node.model
  // Get the filepath
  const fileNode = getNode(node.parent)
  // Create an archie model node
  createModel(createNode, model, fileNode.name, node.id)
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
    const diagramNode = createSvg(createNode, svg, args, node.id)
    // Link the new node to its parent
    createParentChildLink({parent: node, child: diagramNode})
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
  const url = (process.env.ARCHIE_API) ? process.env.ARCHIE_API : `http://localhost:3000`
  const endpoint = new URL(`${url}/diagram/${args.type}`)
  // Add parameters
  const params = []
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
      mediaType: `text/archie`,
      type: `ArchieModel`,
      contentDigest: crypto
        .createHash(`md5`)
        .update(JSON.stringify(model))
        .digest(`hex`),
      description: `Archie yaml model`,
    },
  })
}
