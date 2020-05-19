package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Cant open pid file (is server running? )")
	}
	if err := os.Remove(pidFile); err != nil {
		//We can go on if we fail here
		log.Printf("warning : cant remove pid file - %s", err)
	}
	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "Bad process ID")
	}
	fmt.Printf("Killing server %d\n", pid)
	return nil
}
func main() {
	if err := killServer("Server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "error in stderr:%s\n", err)
		os.Exit(1)
	}
}
