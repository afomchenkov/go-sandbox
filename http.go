package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://pkg.go.dev/io
// https://golang.org/pkg/os/#File

type logWriter struct{}

func LoadData() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(bs)

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these bytes:", len(bs))
	return len(bs), nil
}
