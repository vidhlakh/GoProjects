package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	// Conf field
}

func readConfig(path string) (*Config, error) {
	//Open a file using hte path. IF error return the error
	file, err := os.Open("path")
	if err != nil {
		//return nil, err
		return nil, errors.Wrap(err, "Cant open th efile ")
	}
	//IF no error close the file
	defer file.Close()
	cfg := &Config{}
	return cfg, nil

}

//Setting up Logging
func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}
func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/Config.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errorr: %s\n", err)
		log.Printf("errror:%+v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)

}
