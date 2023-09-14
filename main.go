package main

import (
	"fmt"
	"main/cmd"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "Dex Cli"
	app.Copyright = "Copyright in 2023"
	app.Commands = []*cli.Command{
		cmd.PowCmd,
	}
	app.Flags = []cli.Flag{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
