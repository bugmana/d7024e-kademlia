package join

import (
	"kademlia/internal/node"

	"github.com/rs/zerolog/log"
)

type Join struct {
}

func (j *Join) Execute(node *node.Node) (string, error) {
	log.Debug().Msg("Executing join command")
	node.JoinNetwork()
	return "Joined network on known node", nil
}

func (j *Join) ParseOptions(options []string) error {
	return nil
}

func (j *Join) PrintUsage() {
	log.Info().Msg("Usage: join")
}