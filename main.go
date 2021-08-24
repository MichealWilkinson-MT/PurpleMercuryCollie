package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/michealwilkinson-mt/PurpleMercuryCollie/maths"
)

func main() {
	http.HandleFunc("/add", addController)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func addController(w http.ResponseWriter, req *http.Request) {
	// This is terrible but we're learning how to be async
	channel := make(chan maths.Equation, 5)
	var wg sync.WaitGroup
	x, _ := strconv.ParseFloat(req.URL.Query().Get("x"), 64)
	y, _ := strconv.ParseFloat(req.URL.Query().Get("y"), 64)
	wg.Add(1)
	go func(x float64, y float64) {
		maths.Add(x, y, channel)
		defer wg.Done()
	}(x, y)
	wg.Wait()
	close(channel)
	for rtn := range channel {
		// rtn := <-channel
		fmt.Fprint(w, rtn.Sum)
		fmt.Fprint(w, " = ")
		fmt.Fprintln(w, rtn.Result)
	}

}
