package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

//go:generate go-bindata templates

func main() {
	app := &cli.App{
		Name:  "go-mg",
		Usage: "Go microservice generator with gRPC and net package",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "Name for microservice `NAME`",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Path to output directory `OUTPUT`",
				Value:       "",
				DefaultText: "current directory",
			},
			&cli.StringFlag{
				Name:        "network",
				Value:       "tcp",
				Aliases:     []string{"net"},
				Usage:       "Network type (tcp, tcp4, tcp6) `NETWORK`",
				DefaultText: "tcp",
			},
			&cli.StringFlag{
				Name:        "address",
				Value:       ":56001",
				Aliases:     []string{"a"},
				Usage:       "Address for tcp (\":56001\", \"127.0.0.1\", \"127.0.0.1:56001\") `ADDRESS`",
				DefaultText: ":56001",
			},
		},
		Action: func(c *cli.Context) error {
			networkFlag := c.String("network")
			addressFlag := c.String("address")

			config := struct {
				Name    string
				Network string
				Address string
				Host    string
				Port    string
			}{
				Network: networkFlag,
				Address: addressFlag,
				Name:    c.String("name"),
			}

			if networkFlag != "tcp" && networkFlag != "tcp4" && networkFlag != "tcp6" {
				return errors.New("unexpected network")
			}

			splitAddress := strings.Split(addressFlag, ":")
			if len(splitAddress) == 2 {
				config.Host = splitAddress[0]
				config.Port = splitAddress[1]
			} else if len(splitAddress) == 1 && strings.Contains(addressFlag, ":") {
				config.Port = splitAddress[0]
			} else if len(splitAddress) == 1 {
				config.Host = splitAddress[0]
				config.Port = "80"
				config.Address += ":" + config.Port
			} else {
				return errors.New("unexpected address")
			}

			return BindataRestoreAssetsWithTemplates(c.String("output"), "", config)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
