import React from "react"
import PropTypes from "prop-types"
import axios from "axios"

class Diagram extends React.Component {
  async componentDidMount() {
    // Make the request
    const response = await axios.post(this.getEndpoint(), this.model)
    // Save the response
    this.setState({diagram: response.data})
  }

  render() {
    return <div class="dot-diagram">{this.state.diagram}</div>
  }

  getEndpoint() {
    // Get the base endpoint
    let endpoint = new URL(`http://localhost:3000/diagrams/${this.type}`)
    // Add parameters
    let params = []
    if (this.scope) {
      params.push(`scope=${encodeURIComponent(this.scope)}`)
    }
    if (this.tag) {
      params.push(`tag=${encodeURIComponent(this.tag)}`)
    }
    endpoint.search = `?${params.join('&')}`
    // Return the string
    return endpoint
  }
}

Diagram.propTypes = {
  model: PropTypes.string.isRequired, // TODO: Find out why isRequired seems ineffectual
  type: PropTypes.string.isRequired, // TODO: Find out why isRequired seems ineffectual
  scope: PropTypes.string,
  tag: PropTypes.string,
}

export default Diagram
