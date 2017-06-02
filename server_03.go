// server_01
package main

import (
	"net/http"
)

//another service
func GetServiceThree(w http.ResponseWriter, r *http.Request) {

	//en := json.NewEncoder(fp)

	w.Write([]byte("Buddhi this is server 03"))
	w.Write([]byte(r.URL.Path))
}
