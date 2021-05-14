package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lazywhite/sample-scheduler-extender/pkg/scheduler"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	router := httprouter.New()
	router.GET("/", scheduler.Index)
	router.POST("/filter", scheduler.Filter)
	router.POST("/prioritize", scheduler.Prioritize)

	fmt.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8888", router))
}
