package rpcparser

import (
	"errors"
	"fmt"
	"kademlia/internal/contact"
	"kademlia/internal/rpc"
	"kademlia/internal/rpccommand"
	"kademlia/internal/rpccommands/ping"
	"kademlia/internal/rpccommands/pong"
	"kademlia/internal/rpccommands/store"
	"strings"

	"github.com/rs/zerolog/log"
)

// Parses a rpc and returns a rpc command.
func ParseRPC(sender *contact.Contact, rpc *rpc.RPC) (rpccommand.RPCCommand, error) {
	fields := strings.Fields(rpc.Content)
	if len(fields) == 0 {
		return nil, errors.New("Missing RPC name")
	}

	var cmd rpccommand.RPCCommand
	var err error
	rpcLog := log.Info().Str("RPCId", rpc.RPCId.String())
	switch identifier := fields[0]; identifier {
	case "PING":
		rpcLog.Msg("PING received")
		cmd = ping.New(&sender.Address, rpc.RPCId)
	case "PONG":
		rpcLog.Msg("PONG received")
		cmd = pong.New()
	case "STORE":
		rpcLog.Msg("STORE received")
		cmd = new(store.Store)
	default:
		err = errors.New(fmt.Sprintf("Received unknown RPC %s", identifier))
		cmd = nil
	}
	return cmd, err
}