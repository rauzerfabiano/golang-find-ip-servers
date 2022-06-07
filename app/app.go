package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the command line application ready to run.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line App"
	app.Usage = "Find ips and server names on the internet"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find ip addresses on internet",
			Flags:  flags,
			Action: findIps,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}
}
