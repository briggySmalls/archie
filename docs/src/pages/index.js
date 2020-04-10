import Layout from "../components/layout"
import SEO from "../components/seo"
import { Link } from "gatsby"
import React from 'react'
import { graphql } from 'gatsby'

const HomePage = ({data}) => {
  return (
    <Layout>
      <SEO title="Home" />
      <h1>Systems engineering without the bloat</h1>
      <p>Archie is a lightweight tool for generating model-based architecture diagrams.</p>
      <p>Define a model...</p>
      <pre>{data.model.value}</pre>
      <p>Generate diagrams..</p>
      <div>
        <h2>Context of A</h2>
        <div dangerouslySetInnerHTML={{ __html: data.scopeA.value }} />
      </div>
      <div>
        <h2>Context of B</h2>
        <div dangerouslySetInnerHTML={{ __html: data.scopeB.value }} />
      </div>
      <Link to="/overview/">Overview</Link>
    </Layout>
  )
}

export const query = graphql`
query MyQuery {
  model: archieModel(name: {eq: "simple"}) {
    value
  }
  scopeA: svgImage(args: {name: {eq: "simple"}, scope: {eq: "a"}}) {
    value
  }
  scopeB: svgImage(args: {name: {eq: "simple"}, scope: {eq: "b"}}) {
    value
  }
}
`

export default HomePage
