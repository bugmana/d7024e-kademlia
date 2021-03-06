package getcontacts

import (
	"github.com/rs/zerolog/log"
	"kademlia/internal/node"
)

type GetContacts struct{}

func (g *GetContacts) Execute(node *node.Node) (string, error) {
	log.Trace().Msg("Executing getcontacts command")
	return node.RoutingTable.GetContacts(), nil
}

func (g *GetContacts) ParseOptions(options []string) error {
	return nil
}

func (g *GetContacts) PrintUsage() string {
	return "Usage: getcontacts"
}
