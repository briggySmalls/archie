import React from "react"
import PropTypes from "prop-types"
import axios from "axios"

const Diagram = ({ type, scope, tag }) => {
  // Request the
  return <div><p>Type: { type }</p><p>Scope: { scope }</p><p>Tag: { tag }</p></div>;
}

Diagram.propTypes = {
  type: PropTypes.string.isRequired, // TODO: Find out why isRequired seems ineffectual
  scope: PropTypes.string,
  tag: PropTypes.string,
}

export default Diagram
