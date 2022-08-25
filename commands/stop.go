package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sentixdev/essentialsgo/commands/op"
	"github.com/sentixdev/essentialsgo/global"
)

type Stop struct{}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	output.Printf("Stopping server.")
	if err := global.Server.Close(); err != nil {
		output.Printf("error shutting down server: %v", err)
	}
}
