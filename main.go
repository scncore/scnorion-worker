package main

import (
	"log"
	"os"

	"github.com/scncore/scnorion-worker/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "scnorion-worker",
		Commands:  getCommands(),
		Usage:     "Manage an scnorion worker",
		Authors:   []*cli.Author{{Name: "Miguel Angel Alvarez Cabrerizo", Email: "mcabrerizo@scnorion.eu"}},
		Copyright: "2025 - Miguel Angel Alvarez Cabrerizo <https://github.com/scncore>",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() []*cli.Command {
	return []*cli.Command{
		commands.AgentWorker(),
		commands.CertManagerWorker(),
		commands.NotificationsWorker(),
		commands.HealthCheck(),
	}
}
