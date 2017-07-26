// server_01
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//another service
func GetServiceThree(w http.ResponseWriter, r *http.Request) {

	//en := json.NewEncoder(fp)
	url := r.URL.Path
	urlArray := strings.Split(url, "/")
	parameterArray := strings.Split(urlArray[2], " ")
	ans := 0

	for i := 0; i < len(parameterArray); i++ {
		num, _ := strconv.Atoi(parameterArray[i])
		fmt.Println(num)
		ans = ans + num
		//w.Write([]byte(parameterArray[i]))
		//w.Write([]byte(strconv.Itoa(num)))
	}
	w.Write([]byte(strconv.Itoa(ans)))

}
