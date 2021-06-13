package main

import (
	"errors"
	"fmt"
	"github.com/gogen/commands/add"
	"github.com/gogen/commands/create"
	"github.com/gogen/commands/generate"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:    "GoGen",
		Usage:   "Generator for go",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:        "new",
				Category:    "New",
				Description: "Create new application in GOPATH workspace",
				Before: func(context *cli.Context) error {
					if context.NArg() < 1 {
						return errors.New("please provide project name")
					}
					return nil
				},
				Action: func(context *cli.Context) error {
					wd, _ := os.Getwd()
					log.Println("make sure working directory is part of GOPATH")
					return create.Project(wd, context.Args().Get(0))
				},
			},
			{
				Name:        "regenerate",
				Category:    "regenerate",
				Description: "regenerate project files from meta file",
				Before: func(context *cli.Context) error {
					_, e := ioutil.ReadFile("./project.json")
					if e != nil {
						return e
					}
					return nil
				},
				Action: func(context *cli.Context) error {
					args := make([]string, context.NArg())
					for i := 0; i < context.NArg(); i++ {
						args[i] = context.Args().Get(i)
					}
					return generate.All("./", args...)
				},
			},
			{
				Name:        "generate",
				Category:    "generate",
				Description: "generate",
				Subcommands: []*cli.Command{
					{
						Name:  "model",
						Usage: "add a model to current project",
						Before: func(context *cli.Context) error {
							if context.NArg() < 1 {
								return errors.New("please model project name")
							}
							if context.NArg()%2 == 0 {
								return errors.New("need even number of params")
							}

							// Dependencies
							// 		DB context
							//		DB config
							_ = add.AppContext("./", "db", "db")
							_ = add.Config("./", "db", map[string]string{"DSN": "string"})
							return nil
						},
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().Get(0))
							params := make(map[string]string)
							tailArgs := c.Args().Tail()
							for i := 0; i < len(tailArgs); i = i + 2 {
								params[tailArgs[i]] = tailArgs[i+1]
							}
							//return create.Model("./",)
							err := add.Model("./", c.Args().Get(0), params)
							if err != nil {
								return err
							}
							return generate.All("./", "models", "appcontext", "configs")
						},
					},
					{
						Name:  "crud",
						Usage: "add a crud to current project",
						Before: func(context *cli.Context) error {
							if context.NArg() < 1 {
								return errors.New("please model name")
							}

							// Dependencies
							_ = add.Controller("./", context.Args().Get(0))
							_ = add.Model("./", context.Args().Get(0), map[string]string{})
							return nil
						},
						Action: func(c *cli.Context) error {
							err := add.CRUD("./", c.Args().Get(0))
							if err != nil {
								return err
							}
							return generate.All("./", "router", "controllers", "models")
						},
					},
					{
						Name:  "controller",
						Usage: "add a controller to current project",
						Before: func(context *cli.Context) error {
							if context.NArg() < 1 {
								return errors.New("please model name")
							}

							// Dependencies
							params := map[string]string{}
							_ = add.Model("./", context.Args().Get(0), params)
							return nil
						},
						Action: func(c *cli.Context) error {
							//return create.Model("./",)
							err := add.Controller("./", c.Args().Get(0))
							if err != nil {
								return err
							}
							return generate.All("./", "controllers", "models")
						},
					},
					{
						Name:  "authcontroller",
						Usage: "add a auth controller to current project",
						Before: func(context *cli.Context) error {

							// Dependency
							params := map[string]string{
								"Name":     "string",
								"Email":    "string",
								"Phone":    "string",
								"Password": "string",
								"Role":     "string",
							}
							_ = add.Model("./", "User", params)
							return nil
						},
						Action: func(c *cli.Context) error {
							err := add.AuthController("./")
							if err != nil {
								return err
							}
							return generate.All("./", "controllers")
						},
					},
					{
						Name:  "context",
						Usage: "add appcontext to current project",
						Before: func(context *cli.Context) error {
							if context.NArg() < 2 {
								return errors.New("please model name and type")
							}
							return nil
						},
						Action: func(c *cli.Context) error {
							//return create.Model("./",)
							err := add.AppContext("./", c.Args().Get(0), c.Args().Get(1))
							if err != nil {
								return err
							}
							return generate.All("./", "appcontext")
						},
					},
					{
						Name:  "config",
						Usage: "add a config to current project",
						Before: func(context *cli.Context) error {
							if context.NArg() < 1 {
								return errors.New("please provide config name")
							}
							if context.NArg()%2 == 0 {
								return errors.New("need even number of params")
							}
							return nil
						},
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().Get(0))
							params := make(map[string]string)
							tailArgs := c.Args().Tail()
							for i := 0; i < len(tailArgs); i = i + 2 {
								params[tailArgs[i]] = tailArgs[i+1]
							}
							//return create.Model("./",)
							err := add.Config("./", c.Args().Get(0), params)
							if err != nil {
								return err
							}
							return generate.All("./", "configs")
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
