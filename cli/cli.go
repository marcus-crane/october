package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/marcus-crane/october/backend"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func IsCLIInvokedExplicitly(args []string) bool {
	for k, v := range os.Args {
		if k == 1 && v == "cli" {
			return true
		}
	}
}

func Invoke(isPortable bool, version string) {
	app := &cli.App{
		Name:     "october cli",
		HelpName: "october cli",
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

	// We remove the cli command so that urfave/cli doesn't try to literally parse it
	// but the help text of the cli tool still shows the user `october cli` so they don't
	// get disoriented and know that we're juggling text under the hood
	var args []string

	for k, v := range os.Args {
		if k == 1 && v == "cli" {
			continue
		}
		args = append(args, v)
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
