import React from 'react'
// import Helmet from 'react-helmet'

const Diagram = () => (
  <>
    <script type="text/javascript" src="https://dagrejs.github.io/project/dagre-d3/latest/dagre-d3.min.js"></script>
      <script type="text/javascript" src="https://d3js.org/d3.v5.min.js"></script>
      <script
        dangerouslySetInnerHTML={{
          __html: `
            window.onload = setTimeout(function() {
              // Create a renderer
              const render = new dagreD3.render();
              // Iterate through the graphs on the page
              const els = d3.selectAll("div.graphviz");
              console.log(els);
              for (const el of els.nodes()) {
                // Get the data
                const dataEl = d3.select(el).select('script');
                console.log(dataEl.html())
                const data = JSON.parse(dataEl.html());
                // Parse the graph
                const g = dagreD3.graphlib.json.read(data);
                // Add the svg
                const svg = d3.select(el).append('svg');
                const svgGroup = svg.append('g');
                // Render
                render(svgGroup, g);
                // Set width/height
                svg.attr("width", g.graph().width);
                svg.attr("height", g.graph().height);
              }
            }, 1000);
          `,
        }}
      />
  </>
)

export default Diagram
