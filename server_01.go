// server_01
package main

import (
	"net/http"
)

//another service
func GetServiceOne(w http.ResponseWriter, r *http.Request) {

	//en := json.NewEncoder(fp)

	w.Write([]byte("Buddhi this is server 01"))
}
