package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var componentTemplate string = `import React, { Component } from 'react'

class %sComponent extends Component {
  render() {
    return (
      <div>%s</div>
    )
  }
}

export default %sComponent
`

var containerTemplate string = `import React, { Component } from 'react'
import %sComponent from './%sComponent'

class %sContainer extends Component {
  render() {
    return (
      <%sComponent {...this.props} />
    )
  }
}

export default %sContainer
`

var indexTemplate string = `import %sContainer from './%s.container'

export default %sContainer
`

func writeComponent(name string) {
	component := fmt.Sprintf(componentTemplate, name, name, name)

	outputName := fmt.Sprintf("./%s/%s.component.js", name, name)
	err := ioutil.WriteFile(outputName, []byte(component), 0644)
	if err != nil {
		panic(err)
	}
}

func writeContainer(name string) {
	container := fmt.Sprintf(containerTemplate, name, name, name, name, name)

	outputName := fmt.Sprintf("./%s/%s.container.js", name, name)
	err := ioutil.WriteFile(outputName, []byte(container), 0644)
	if err != nil {
		panic(err)
	}
}

func writeIndex(name string) {
	index := fmt.Sprintf(indexTemplate, name, name, name)

	outputName := fmt.Sprintf("./%s/index.js", name)
	err := ioutil.WriteFile(outputName, []byte(index), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	arg := os.Args[1]

	// outputName := fmt.Sprintf("./%s/%s.container.js", name)
	os.Mkdir(arg, 0755)
	writeComponent(arg)
	writeContainer(arg)
	writeIndex(arg)
}
