import React from "react"
import PropTypes from "prop-types"

const Diagram = ({ type, scope, tag }) => {
  return <div><p>Type: { type }</p><p>Scope: { scope }</p><p>Tag: { tag }</p></div>;
}

Diagram.propTypes = {
  type: PropTypes.string.isRequired,
  scope: PropTypes.string,
  tag: PropTypes.string,
}

export default Diagram
