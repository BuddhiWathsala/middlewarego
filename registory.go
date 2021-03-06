// registory
package main

import (
	"net/http"
)

func getFunctionByName(funcName string) func(w http.ResponseWriter, r *http.Request) {
	registry := map[string]func(w http.ResponseWriter, r *http.Request){
		"GetServiceOne":   GetServiceOne,
		"GetServiceThree": GetServiceThree,
	}

	return registry[funcName]
}
