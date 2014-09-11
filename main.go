package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/xyproto/simpleredis"
)

var pool *simpleredis.ConnectionPool

func ListRangeHandler(rw http.ResponseWriter, req *http.Request) {
	members, err := simpleredis.NewList(pool, mux.Vars(req)["key"]).GetAll()
	if err != nil {
		panic(err)
	}

	membersJSON, err := json.Marshal(members)
	if err != nil {
		panic(err)
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(membersJSON))
}

func ListPushHandler(rw http.ResponseWriter, req *http.Request) {
	set := simpleredis.NewList(pool, mux.Vars(req)["key"])
	err := set.Add(mux.Vars(req)["value"])
	if err != nil {
		panic(err)
	}

	members, err := set.GetAll()
	if err != nil {
		panic(err)
	}

	membersJSON, err := json.Marshal(members)
	if err != nil {
		panic(err)
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(membersJSON))
}

func InfoHandler(rw http.ResponseWriter, req *http.Request) {
	info, err := pool.Get(0).Do("INFO")
	if err != nil {
		panic(err)
	}

	infoString := string(info.([]uint8))

	rw.WriteHeader(200)
	rw.Write([]byte(infoString))
}

func main() {
	pool = simpleredis.NewConnectionPoolHost(
		os.Getenv("SERVICE_HOST") + ":" + os.Getenv("REDIS_MASTER_SERVICE_PORT"))
	defer pool.Close()

	r := mux.NewRouter()
	r.Path("/lrange/{key}").Methods("GET").HandlerFunc(ListRangeHandler)
	r.Path("/rpush/{key}/{value}").Methods("GET").HandlerFunc(ListPushHandler)
	r.Path("/info").Methods("GET").HandlerFunc(InfoHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}
