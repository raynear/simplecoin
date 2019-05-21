package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func err(msg string, e error) {
	if e != nil {
		fmt.Println(msg, e)
		log.Fatal(msg, e)
	}
}

func main() {
	file, e := os.Open("test.config")
	err("File Error", e)

	viper.SetConfigType("prop")
	viper.ReadConfig(file)
	fmt.Println(viper.Get("lang"))

	var language string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "language for greeting",
			Destination: &language,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nobody"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	e = app.Run(os.Args)
	err("Running Error", e)
}
