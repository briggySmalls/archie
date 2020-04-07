import React from "react"
import dot from "graphlib-dot"


class Diagram extends React.Component {
  componentDidMount() {
      // D3 Code to create the chart
      // using this._rootNode as container
    // Read the children as a graph
    const g = dot.read(this.props.children)
    // Create the renderer
    const render = new dot.graphlib.Graph().setGraph(g)
    // Set up an SVG group so that we can translate the final graph.
    const child = this._rootNode.appendChild(document.createElement('g'))
    // Run the renderer. This is what draws the final graph.
    render(child, g);
  }

  shouldComponentUpdate() {
    // Prevents component re-rendering
    return false;
  }

  _setRef(componentNode) {
    this._rootNode = componentNode;
  }

  render() {
    return <div className="line-container" ref={this._setRef.bind(this)} />
  }
}

export default Diagram
