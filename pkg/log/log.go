package log

import (
	"log"
	"os"
)

var (
	Error = log.New(os.Stderr, "mybanana ERR: ", log.LstdFlags)
	Info  = log.New(os.Stdout, "mybanana INFO: ", log.LstdFlags)
)
