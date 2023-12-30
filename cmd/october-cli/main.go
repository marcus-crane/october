package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/marcus-crane/october/backend"
)

var (
	version       = "DEV"
	portablebuild = "false"
)

func main() {
	isPortable := false
	isPortable, _ = strconv.ParseBool(portablebuild)

	app := &cli.App{
		Name:     "october",
		HelpName: "october-cli",
		Version:  version,
		Authors: []*cli.Author{
			{
				Name:  "Marcus Crane",
				Email: "october@utf9k.net",
			},
		},
		Usage: "sync your kobo highlights to readwise from your terminal",
		Commands: []*cli.Command{
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Usage:   "sync kobo highlights to readwise",
				Action: func(c *cli.Context) error {
					ctx := context.Background()
					b := backend.StartBackend(&ctx, version, isPortable)
					if b.Settings.ReadwiseToken == "" {
						return fmt.Errorf("no readwise token was configured. please set this up using the gui as the cli does not support this yet")
					}
					kobos := b.DetectKobos()
					if len(kobos) == 0 {
						return fmt.Errorf("no kobo was found. have you plugged one in and accepted the connection request?")
					}
					if len(kobos) > 1 {
						return fmt.Errorf("cli only supports one connected kobo at a time")
					}
					err := b.SelectKobo(kobos[0].MntPath)
					if err != nil {
						return fmt.Errorf("an error occurred trying to connect to the kobo at %s", kobos[0].MntPath)
					}
					num, err := b.ForwardToReadwise()
					if err != nil {
						return err
					}
					logrus.Infof("Successfully synced %d highlights to Readwise", num)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
