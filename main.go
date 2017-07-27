// mainfile
package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
const (

	//iota : if confKey = 1 then mongoSessKey = 2,
	//this incrementation happen because of iota
	confKey key = iota
	mongoSessKey
)
func main() {
	router := mux.NewRouter()
	/*router.HandleFunc("/getservice1", getFunctionByName("GetServiceOne"))
	router.HandleFunc("/getservice2/{category}", GetServiceTwo)
	router.HandleFunc("/getservice3", GetServiceThree)*/

	xmlFile, err := os.Open("servers.xml")

	if err != nil {
		fmt.Println("Server file reading error")
	}

	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var middlewareServers servers

	xml.Unmarshal(xmlData, &middlewareServers)
	fmt.Println(middlewareServers)

	for i := range middlewareServers.Servers {

		for j := range middlewareServers.Servers[i].Services {

			currentService := middlewareServers.Servers[i].Services[j]
			//f := reflect.ValueOf(Foo{}).MethodByName(currentService.Servicefunction).Interface().(func())
			router.HandleFunc(("/" + currentService.Servicename), getFunctionByName(currentService.Servicefunction))

		}
	}

	c := config{
		addr:            ":80",
		mongoConnString: ":27017",
	}

	//mgoSess := newMongoSession(c.mongoConnString)

	loggerfile := handlers.LoggingHandler(os.Stdout, getHandler(logger(router, c), c))

	srvr := http.Server{
		Addr:         c.addr,
		Handler:      loggerfile,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	panic(srvr.ListenAndServe())
}

func getHandler(next http.Handler, c config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		newCtx := context.WithValue(r.Context(), confKey, c)
		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler, c config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Header.Get("X-Forwarded-For"))
		log.Println(r.Header.Get("X-Real-Ip"))
		log.Println(w.Header())
		newCtx := context.WithValue(r.Context(), confKey, c)
		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

type config struct {
	addr            string
	mongoConnString string
}

/*server xml structures*/
type service struct {
	XMLName         xml.Name `xml:"service"`
	Servicename     string   `xml:"name"`
	Servicefunction string   `xml:"function"`
}
type server struct {
	XMLName  xml.Name  `xml:"server"`
	Name     string    `xml:"servername"`
	Services []service `xml:"service"`
}

type servers struct {
	XMLName xml.Name `xml:"servers"`
	Servers []server `xml:"server"`
}

/*server xml structures end*/
