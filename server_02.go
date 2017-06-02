// server_02
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type key int

const (
	confKey key = iota
	mongoSessKey
)

//hello function
func GetServiceTwo(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello Buddhi")

	//c := getConf(r.Context())
	//fmt.Fprintln(w, "Server is running on :", c.addr)

	sess := getMongo(r.Context())
	col := sess.DB("myfirstdb").C("myFirstCollection")

	//retrieve data from database
	uq := student{}

	url := r.URL.Path
	urlArray := strings.Split(url, "/")
	log.Println(urlArray[2])
	//execute Map data query
	query := bson.M{"name": urlArray[2]}

	err := col.Find(query).One(&uq)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.MarshalIndent(uq, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)

}

type student struct {
	Name       string `json:"name" bson:"name"`
	University string `json:"university" bson:"university"`
	Age        int    `json:"age" bson:"age"`
	Interest   string `json:"interest" bson:"interest"`
}

//get mongo
func getMongo(ctx context.Context) *mgo.Session {
	c, ok := ctx.Value(mongoSessKey).(*mgo.Session)

	if !ok {
		panic("mongo session not in context")
	}

	return c
}

func getConf(ctx context.Context) config {
	c, ok := ctx.Value(confKey).(config)

	if !ok {
		panic("config object not in context")
	}

	return c
}

func withMongo(next http.Handler, sess *mgo.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCopy := sess.Copy()
		defer sessCopy.Close()
		newCtx := context.WithValue(r.Context(), mongoSessKey, sessCopy)
		r = r.WithContext(newCtx)
		next.ServeHTTP(w, r)
	})
}

func newMongoSession(connString string) *mgo.Session {

	sess, err := mgo.Dial(connString)
	if err != nil {
		panic(err)
	}

	return sess
}
