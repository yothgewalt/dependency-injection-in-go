package commands

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/yongyuth-chuankhuntod/libraries"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var cmds = map[string]libraries.Command{
	"app:serve": NewServeCommand(),
}

func GetSubCommands(opt fx.Option) []*cobra.Command {
	var subCommands []*cobra.Command
	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}

	return subCommands
}

func WrapSubCommand(name string, command libraries.Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: command.Short(),
		Run: func(cmd *cobra.Command, args []string) {
			logger := libraries.NewLogger()
			opts := fx.Options(
				fx.WithLogger(func() fxevent.Logger {
					return logger.NewFxLogger()
				}),
				fx.Invoke(command.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer app.Stop(ctx)

			if err != nil {
				logger.Fatal(err)
			}
		},
	}
	command.Setup(wrappedCmd)

	return wrappedCmd
}
