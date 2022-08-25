package utils

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sentixdev/essentialsgo/global"
	"strings"
)

func SliceUnique(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value && entry != "" {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func PlayerByName(name string) (*player.Player, bool) {
	for _, p := range global.Server.Players() {
		if strings.ToLower(p.Name()) == strings.ToLower(name) {
			return p, true
		}
	}
	return nil, false
}
