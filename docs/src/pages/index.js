import Layout from "../components/layout"
import SEO from "../components/seo"
import { Link } from "gatsby"
import React from 'react'
import { graphql } from 'gatsby'
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';

const HomePage = ({data}) => {
  return (
    <Layout>
      <SEO title="Home" />
      <h1>Systems engineering without the bloat</h1>
      <p>Archie is a lightweight tool for generating model-based architecture diagrams.</p>
      <Link to="/rationale/"><Button color="primary">Learn more...</Button></Link>
      <Grid container spacing={3}>
        <Grid item>
          <h2>Define a model...</h2>
          <pre>{data.model.value}</pre>
        </Grid>
        <Grid item>
          <h2>Context of A</h2>
          <div dangerouslySetInnerHTML={{ __html: data.scopeA.value }} />
        </Grid>
        <Grid item>
          <h2>Context of B</h2>
          <div dangerouslySetInnerHTML={{ __html: data.scopeB.value }} />
        </Grid>
      </Grid>
    </Layout>
  )
}

export const query = graphql`
query {
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
