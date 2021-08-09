package main

import (
	"github.com/prongbang/analyticsgen/cmd"
	"github.com/prongbang/analyticsgen/internal/pkg/common"
	"github.com/prongbang/analyticsgen/pkg/parameter"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	common.Banner()

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "platform",
				Value: "",
				Usage: "-platform [android, ios, flutter]",
			},
			&cli.StringFlag{
				Name:  "asset",
				Value: "",
				Usage: "-asset [code, key, test]",
			},
			&cli.StringFlag{
				Name:  "target",
				Value: "",
				Usage: "-target ./export",
			},
			&cli.StringFlag{
				Name:  "document",
				Value: "",
				Usage: "-document 1oBqyd7ys2GOtroqV6D4qYH6JWQjKrZiOcngmcsbq0VU",
			},
			&cli.StringFlag{
				Name:  "sheet",
				Value: "",
				Usage: "-sheet 0",
			},
		},
		Action: func(c *cli.Context) error {
			return cmd.Run(&parameter.Parameter{
				Platform: c.String(parameter.Platform),
				Asset:    c.String(parameter.Asset),
				Target:   c.String(parameter.Target),
				Document: c.String(parameter.Document),
				Sheet:    c.String(parameter.Sheet),
			})
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
