package main

import (
	"log"
	"fmt"
	"runtime"
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/NAVCoin/testnet-manager/server/nodes"
	"time"
	"math/rand"
)

func main() {
	// log out server runtime OS and Architecture
	log.Println(fmt.Sprintf("Server running in %s:%s", runtime.GOOS, runtime.GOARCH))
	log.Println(fmt.Sprintf("App pid : %d.", os.Getpid()))

	nodes.InitData()


	rand.Seed(int64(time.Now().Nanosecond()))


	// setup the router and the api
	router := mux.NewRouter()

	// load up the cache system


	nodes.InitSetupHandlers(router, "api")

	// Start http server
	port := fmt.Sprintf(":%d", 5000)
	log.Fatal(http.ListenAndServe(port, router))

}
