import React from "react"
import PropTypes from "prop-types"
import axios from "axios"

class Diagram extends React.Component {
  async componentDidMount() {
    // Build the request URL
    let endpoint = new URL(`http://localhost:3000/diagrams/${this.type}`)
    let params = {}
    if (this.scope) {
      params['scope'] = this.scope
    }
    if (this.tag) {
      params['tag'] = this.tag
    }
    // Make the request
    const response = await axios.post(endpoint, model)
    // Save the response

  }
}

const Diagram = ({ model, type, scope, tag }) => {
  // Request the
  return <div><p>Type: { type }</p><p>Scope: { scope }</p><p>Tag: { tag }</p></div>;
}

Diagram.propTypes = {
  model: PropTypes.string.isRequired, // TODO: Find out why isRequired seems ineffectual
  type: PropTypes.string.isRequired, // TODO: Find out why isRequired seems ineffectual
  scope: PropTypes.string,
  tag: PropTypes.string,
}

export default Diagram
