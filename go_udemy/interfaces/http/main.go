package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("https://www.google.com/")
	resp, err := http.Get("https://indigo-dummy-api.herokuapp.com/flights")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// // fmt.Println(string(bs))
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// .data.indiGoAvailability.trips[0].flightDates[0].flights
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Length is ", len(bs))
	return len(bs), nil
}
