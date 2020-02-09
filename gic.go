package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nevermosby/gic-cli-tool/config"
	gic "github.com/nevermosby/gic-cloud-sdk-go"
	"github.com/nevermosby/promptui"
	"github.com/urfave/cli/v2"
)

const (
	GICBaseUrl = "https://api2.capitalonline.net"
)

var (
	app = &cli.App{}
	p   = fmt.Println
)

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
				if c.NArg() == 0 {
					// if there is no args provided, check the legacy config
					// load the config file first, check the token is created less than one hour
					token := config.CheckToken()
					if token != "" {
						// p("token is already valid: ",token)
						p("logged already")
					} else {
						p("Token is expired, pls login mannually")
						// do the interactive input
						usernameProm := promptui.Prompt{
							Label: "Username",
						}
						username, err := usernameProm.Run()
						if err != nil {
							fmt.Printf("username input failed %v\n", err)
						}

						pwdProm := promptui.Prompt{
							Label: "Password",
							// Validate: validate,
							Mask: ' ',
						}
						pwd, err := pwdProm.Run()
						if err != nil {
							fmt.Printf("Prompt failed %v\n", err)
						}

						// use sdk to login to get token
						// then store the info into config file
						gicLogin2Save(username, pwd)
					}
				} else {
					fmt.Println("list the command args:", c.Args())
					if c.Args().Get(0) == "help" {
						p("show the help text")
					} else {
						p("pls login without args")
					}
				}
				return nil
			},
		},
		{
			Name:    "os",
			Aliases: []string{"os"},
			Usage:   "manage the available os, you can `list` `search`",
			Subcommands: []*cli.Command{
				{
					Name:     "list",
					Category: "os",
					Usage:    "list all the availabe os",
					Action: func(c *cli.Context) error {
						token := config.CheckToken()
						if token != "" {
							// p("logged already")
							var client = &gic.Client{}
							client.Init(GICBaseUrl, "")
							client.LoginWithToken(token)
							osList, err := client.ListOS()
							if err != nil {
								log.Fatal(err)
							}
							for _, os := range osList {
								p("os display name: ", os.DisplayName)
							}

						} else {
							p("Pls login first.")
						}
						return nil
					},
				},
				{
					Name:      "search",
					Category:  "os",
					Usage:     "search for all the availabe os with keyword",
					ArgsUsage: "[keyword]",
					Action: func(c *cli.Context) error {
						keyword := c.Args().First()
						token := config.CheckToken()
						var matchedOS []string
						var displayName string
						if token != "" {
							// p("logged already")
							var client = &gic.Client{}
							client.Init(GICBaseUrl, "")
							client.LoginWithToken(token)
							osList, err := client.ListOS()
							if err != nil {
								log.Fatal(err)
							}
							for _, os := range osList {
								displayName = os.DisplayName
								if strings.Contains(strings.ToLower(displayName), keyword) {
									matchedOS = append(matchedOS, displayName)
								}
							}
							for _, mos := range matchedOS {
								p("Found:", mos)
							}

						} else {
							p("Pls login first.")
						}
						return nil
					},
				},
			},
		},
		{
			// TODO: add options to specifiy the datacenter resource
			Name:    "datacenter",
			Aliases: []string{"dc"},
			Usage:   "manage the gic datacenter instances, you can `list`,`create`,`delete`,`info`",
			Subcommands: []*cli.Command{
				{
					Name:     "list",
					Category: "dc",
					Usage:    "list all the datecenter instances",
					Action: func(c *cli.Context) error {
						p("start to list all the datecenter instace: ", c.Args().First())
						token := config.CheckToken()
						if token != "" {
							// p("token is already valid: ",token)
							// p("logged already")
							var client = &gic.Client{}
							client.Init(GICBaseUrl, "")
							client.LoginWithToken(token)
							datacenters, err := client.ListDataCenter()
							if err != nil {
								log.Fatal(err)
							}
							for _, d := range datacenters {
								p("datacenter name:", d.SiteName, ",datacenter resource name:", d.Resource.Name)
							}
						} else {
							p("Pls login first.")
						}
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
	return cli.NewExitError("No such commands provided: "+"'"+c.Args().First()+"'"+". Run 'gic help' for usage", 2)
}

// use sdk to login to get token
// then store the info into config file
func gicLogin2Save(username string, pwd string) {
	var client = &gic.Client{}
	client.Init(GICBaseUrl, "")
	client.Login(username, pwd)
	p("login token insider gicLogin2Save:", client.Token)
	configFile, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}
	configFile.Url = GICBaseUrl
	configFile.Username = username
	configFile.Cred = pwd
	configFile.Token.Val = client.Token
	configFile.Token.CreatedAt = time.Now().Format(time.RFC3339)

	// then store the info into config file
	configFile.Save()
}

func main() {
	info()
	command()

	// default action behaviour
	app.Action = noArgs

	err := app.Run(os.Args)
	// for debug
	// err := app.Run([]string{"login"})

	if err != nil {
		log.Fatal(err)
	}
}
