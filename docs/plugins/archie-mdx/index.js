const visit = require(`unist-util-visit`)



module.exports = ({ markdownAST }, pluginOptions) => {
  // Visit all code blocks
  visit(markdownAST, `code`, node => {
    // We only care about code with the 'archie' language
    if (node.lang !== `archie`) {
      return
    }
    const args = node.meta.split(" ")
    console.log(`Found an archie node with args ${args}`)
  })
  return markdownAST
}
