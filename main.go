package main

import (
	"github.com/codegangsta/cli"
)

const VERSION = "0.0.1"

func main() {

	INDEX_URL = getIndexURI()

	app := cli.NewApp()
	// configure app
	app.Name = "cidr"
	app.Usage = "view CIDR information for ip addresses"
	app.Version = VERSION
	app.Author = "Alex Crowe"
	// register commands with app
	app.Commands = Commands
	// run app and pass args
	app.Run(os.Args)
}
