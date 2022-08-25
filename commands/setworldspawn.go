package commands

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sentixdev/essentialsgo/commands/op"
	"github.com/sentixdev/essentialsgo/global"
)

type SetWorldSpawnXYZ struct {
	X, Y, Z float64
}

func (t SetWorldSpawnXYZ) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	w := global.Server.World()
	bp := cube.Pos{int(t.X), int(t.Y), int(t.Z)}

	if p, ok := source.(*player.Player); ok {
		w = p.World()
	}

	w.SetSpawn(bp)
	output.Printf("Set the default world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
}

type SetWorldSpawn struct {
	Sub string `optional:""`
}

func (t SetWorldSpawn) Run(source cmd.Source, output *cmd.Output) {
	if t.Sub != "" {
		return
	}
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	if p, ok := source.(*player.Player); ok {
		pos := p.Position()
		bp := cube.Pos{int(pos.X()), int(pos.Y()), int(pos.Z())}
		p.World().SetSpawn(bp)
		output.Printf("Set the world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
	} else {
		output.Error("Usage: /setworldspawn <X: float> <Y: float> <Z: float>")
	}
}
