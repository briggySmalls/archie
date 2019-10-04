var yaml = `
elements:
  - name: user
    kind: actor
  - name: sound-system
    children:
      - name: speaker
        children:
          - name: enclosure
            technology: physical
          - name: driver
            technology: electro-mechanical
          - connector
          - cable
      - name: amplifier
        children:
          - audio in connector
          - audio out connector
          - bluetooth receiver
          - ac-dc converter
          - mixer
          - amplifier
          - name: power button
            technology: electro-mechanical
          - name: input select
            technology: electro-mechanical
  - name: app
    children:
    - bluetooth driver
    - spotify client
    - ui
associations:
  # Sound system
  - source: user
    destination: sound-system/amplifier/input select
  - source: sound-system/amplifier/input select
    destination: sound-system/amplifier/mixer
  - source: sound-system/amplifier/audio in connector
    destination: sound-system/amplifier/mixer
  - source: sound-system/amplifier/bluetooth receiver
    destination: sound-system/amplifier/mixer
  - source: sound-system/amplifier/ac-dc converter
    destination: sound-system/amplifier/mixer
  - source: sound-system/amplifier/mixer
    destination: sound-system/amplifier/amplifier
  - source: sound-system/amplifier/ac-dc converter
    destination: sound-system/amplifier/amplifier
  - source: sound-system/amplifier/amplifier
    destination: sound-system/amplifier/audio out connector
  - source: sound-system/amplifier/audio out connector
    destination: sound-system/speaker/cable
  - source: sound-system/speaker/cable
    destination: sound-system/speaker/connector
  - source: sound-system/speaker/connector
    destination: sound-system/speaker/driver
  - source: sound-system/speaker/driver
    destination: sound-system/speaker/enclosure
  - source: sound-system/amplifier/power button
    destination: sound-system/amplifier/ac-dc converter
  - source: sound-system/speaker/driver
    destination: user
  - source: user
    destination: sound-system/amplifier/power button
  # App
  - source: user
    destination: app/ui
  - source: app/ui
    destination: app/spotify client
  - source: app/spotify client
    destination: app/bluetooth driver
  - source: app/bluetooth driver
    destination: sound-system/amplifier/bluetooth receiver
`

function start() {
    // Initialise the model
    init(yaml);
    // Get the elements
    let els = elements();
    // Create the buttons
    let optionsEl = document.getElementById("options")
    Object.entries(els).forEach(([id, name]) => {
        // Create a new button
        let btn = document.createElement("button");
        btn.innerHTML = name;
        // Add a callback
        btn.addEventListener("click", function(event) {
            // Update diagram with given scope
            updateWithContext(name);
            // Run mermaid again
            mermaid.init(undefined, "#diagram");
        });
        // Now add the button to the DOM
        optionsEl.appendChild(btn);
    })
}

// Change diagram content when requested
function updateWithContext(scope) {
    let diagram = contextView(scope);
    if (diagram === null) {
        return Error("Failed to create diagram");
    }
    let diagramEl = document.getElementById("diagram");
    diagramEl.innerHTML = diagram;
    diagramEl.removeAttribute("data-processed");
}