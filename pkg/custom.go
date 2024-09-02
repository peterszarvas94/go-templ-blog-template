package pkg

import (
	"github.com/a-h/templ"
)

var components = make(map[string]templ.Component)

func AddCustomComponent(name string, comp templ.Component) {
	components[name] = comp
}

func GetCustomComponent(name string) (templ.Component, bool) {
	comp, exists := components[name]
	return comp, exists
}

func GetAllRoutesForCustomComponents() []string {
	var keys []string
	for k := range components {
		keys = append(keys, k)
	}
	return keys
}
