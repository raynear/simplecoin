package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Config ...
type Config struct {
	LogFile string
}

func initConfig() *Config {
	conf := Config{}
	conf.LogFile = "logfile.log"

	return &conf
}

func err(msg string, e error) {
	if e != nil {
		fmt.Println(msg, e)
		log.Fatal(msg, e)
	}
}

func main() {
	conf := initConfig()
	fpLog, e := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 666)
	err("File Open Error", e)
	defer fpLog.Close()

	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	log.Println("test1")
}
