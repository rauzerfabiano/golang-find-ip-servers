package app

import "github.com/urfave/cli"

// Generate returns the command line application ready to run.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line App"
	app.Usage = "Find ips and server names on the internet"
	app.Version = "1.0.0"
	return app
}
