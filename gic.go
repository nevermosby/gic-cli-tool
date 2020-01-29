package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = &cli.App{}

func init() {
	// TODO:
	// help text
}
func info() {
	app.Name = "GIC CLI "
	app.Usage = "GIC CLI is designed for CDS Cloud Platform, manage your own cloud resource easily"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Wen-Quan(David) Li",
			Email: "legendarily@gmail.com",
		},
	}
	app.Version = "v1.0.0"
	app.Copyright = "MIT (c) 2020 Wen-Quan Li"
}

func command() {
	app.Commands = []*cli.Command{
		{
			Name: "login",
			// Category: "login",
			Aliases: []string{"l"},
			Usage:   "login GIC platform to get the token and store it locally",
			Action: func(c *cli.Context) error {
				fmt.Println("Start to login: ", c.Args().First())
				return nil
			},
		},
		{
			// TODO: add options to specifiy the datacenter resource
			Name:    "datcenter",
			Aliases: []string{"dc"},
			Usage:   "manage the gic datacenter instances, you can `list`,`create`,`delete`,`info`",
			Subcommands: []*cli.Command{
				{
					Name:     "list",
					Category: "dc",
					Usage:    "list all the datecenter instances",
					Action: func(c *cli.Context) error {
						fmt.Println("start to list all the datecenter instace: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "create",
					Category: "dc",
					Usage:    "create a datacenter instance",
					Action: func(c *cli.Context) error {
						fmt.Println("start to create a datecenter instace: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "delete",
					Category: "dc",
					Usage:    "delete an existing datacenter",
					Action: func(c *cli.Context) error {
						fmt.Println("start to delete an existing datecenter: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "info",
					Category: "dc",
					Usage:    "show the detail info for an existing datacenter",
					Action: func(c *cli.Context) error {
						fmt.Println("start to reset os for an existing datacenter: ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			// TODO: add options to specifiy the vm resource
			Name:    "virtualmachine",
			Aliases: []string{"vm"},
			Usage:   "manage the gic normal vm instances, you can `list`,`create`,`delete`,`resetos`",
			Subcommands: []*cli.Command{
				{
					Name:     "list",
					Category: "vm",
					Usage:    "list all the vm instances in a datacenter",
					Action: func(c *cli.Context) error {
						fmt.Println("start to list all the vm instace in a datacenter: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "create",
					Category: "vm",
					Usage:    "create a vm instance",
					Action: func(c *cli.Context) error {
						fmt.Println("start to create a vm instace: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "delete",
					Category: "vm",
					Usage:    "delete an existing vm",
					Action: func(c *cli.Context) error {
						fmt.Println("start to delete an existing vm: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "restart",
					Category: "vm",
					Usage:    "restart an existing vm",
					Action: func(c *cli.Context) error {
						fmt.Println("start to restart an existing vm: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "shutdown",
					Category: "vm",
					Usage:    "shutdown an existing vm",
					Action: func(c *cli.Context) error {
						fmt.Println("start to shutdown an existing vm: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "poweron",
					Category: "vm",
					Usage:    "poweron an existing vm",
					Action: func(c *cli.Context) error {
						fmt.Println("start to poweron an existing vm: ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "resetos",
					Category: "vm",
					Usage:    "reset os for an existing vm",
					Action: func(c *cli.Context) error {
						fmt.Println("start to reset os for an existing vm: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
}
// the default behaviour for other commands
func noArgs(c *cli.Context) error {

	// cli.ShowAppHelp(c)
	// TODO: show the command for users
	return cli.NewExitError("No such commands provided. Run 'gic help' for usage", 2)
}

func main() {
	info()
	command()

	// default action behaviour
	app.Action = noArgs

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
