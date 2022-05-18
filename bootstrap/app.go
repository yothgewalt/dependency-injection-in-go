package bootstrap

import (
	"github.com/spf13/cobra"
	"github.com/yongyuth-chuankhuntod/commands"
)

var rootCmd = &cobra.Command{
	Use:   "dependency-injection",
	Short: "Dependency injection using uber-go/fx",
	Long: `
▓█████▄  ██▓    █    ██   ██████  ██▓ ███▄    █   ▄████      █████▒▒██   ██▒
▒██▀ ██▌▓██▒    ██  ▓██▒▒██    ▒ ▓██▒ ██ ▀█   █  ██▒ ▀█▒   ▓██   ▒ ▒▒ █ █ ▒░
░██   █▌▒██▒   ▓██  ▒██░░ ▓██▄   ▒██▒▓██  ▀█ ██▒▒██░▄▄▄░   ▒████ ░ ░░  █   ░
░▓█▄   ▌░██░   ▓▓█  ░██░  ▒   ██▒░██░▓██▒  ▐▌██▒░▓█  ██▓   ░▓█▒  ░  ░ █ █ ▒ 
░▒████▓ ░██░   ▒▒█████▓ ▒██████▒▒░██░▒██░   ▓██░░▒▓███▀▒   ░▒█░    ▒██▒ ▒██▒
▒▒▓  ▒ ░▓     ░▒▓▒ ▒ ▒ ▒ ▒▓▒ ▒ ░░▓  ░ ▒░   ▒ ▒  ░▒   ▒     ▒ ░    ▒▒ ░ ░▓ ░
░ ▒  ▒  ▒ ░   ░░▒░ ░ ░ ░ ░▒  ░ ░ ▒ ░░ ░░   ░ ▒░  ░   ░     ░      ░░   ░▒ ░
░ ░  ░  ▒ ░    ░░░ ░ ░ ░  ░  ░   ▒ ░   ░   ░ ░ ░ ░   ░     ░ ░     ░    ░  
  ░     ░        ░           ░   ░           ░       ░             ░    ░  
░                                                                          

This is a command runner or cli for todolist application in go.
Using this we can use underlying dependency injection container for running scripts.
Main advantage is that, we can use same services, repositories, infrastructure present in the application itself
	`,
	TraverseChildren: true,
}

type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}

	cmd.AddCommand(commands.GetSubCommands(CommomModules)...)

	return cmd
}

var RootApp = NewApp()
