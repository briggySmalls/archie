import React from "react"
import { Link } from "gatsby"

import Layout from "../components/layout"
import Image from "../components/image"
import SEO from "../components/seo"

const IndexPage = () => (
  <Layout>
    <SEO title="Home" />
    <h1>Systems engineering without the bloat</h1>
    <p>
      Archie is a lightweight tool for generating model-based architecture
      diagrams.
    </p>
    <div style={{ maxWidth: `300px`, marginBottom: `1.45rem` }}>
      <Image />
    </div>
    <Link to="/overview/">Overview</Link>
  </Layout>
)

export default IndexPage
