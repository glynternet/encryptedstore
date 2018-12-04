package main

import (
	"log"
	"net/http"
	"os"

	ehttp "github.com/glynternet/encryptedstore/pkg/http"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	err := run(logger, ":8080")
	if err != nil {
		logger.Println(err)
	}
}
func run(logger *log.Logger, addr string) error {
	logger.Printf("serving at %s", addr)
	return http.ListenAndServe(addr, ehttp.NewEncrypterServeMux(logger))
}
