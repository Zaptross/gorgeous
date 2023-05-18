package services

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

const getElementServiceName = "getElement"

const getElementService = `
const %s = (idString) => {
	return document.getElementById(idString)
}
`

func GetElementService() string {
	g.Service(getElementServiceName, g.JavaScript(fmt.Sprintf(getElementService, getElementServiceName)))
	return getElementServiceName
}
