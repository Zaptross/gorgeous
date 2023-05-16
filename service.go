package gorgeous

import "fmt"

// TODO - utilise and document

var services = map[string]JavaScript{}

func RegisterService(name string, service JavaScript) {
	if services[name] != "" {
		panic(fmt.Sprintf(`gorgeous: service '%s' is already registered`, name))
	}
	services[name] = service
}

func collectServices() JavaScript {
	var collected JavaScript

	for _, service := range services {
		collected += service + "\n"
	}

	return collected
}
