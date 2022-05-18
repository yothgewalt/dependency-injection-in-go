package commands

import (
	"github.com/spf13/cobra"
	"github.com/yongyuth-chuankhuntod/libraries"
)

type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve todolist application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() libraries.CommandRunner {
	return func(
		environment libraries.Environment,
		logger libraries.Logger,
		echo libraries.Handler,
	) {
		logger.Info("Running server")

		if environment.EchoServerPort == "" {
			_ = echo.Engine.Start(":8080")
		} else {
			_ = echo.Engine.Start(":" + environment.EchoServerPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
