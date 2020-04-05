import React from "react"
import PropTypes from "prop-types"

const Diagram = ({ siteTitle }) => {
  return <div>{ siteTitle }</div>;
}

Diagram.propTypes = {
  siteTitle: PropTypes.string,
}

Diagram.defaultProps = {
  siteTitle: `default`,
}

export default Diagram
