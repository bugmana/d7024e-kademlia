package cmdparser

import (
	"strings"

	. "kademlia/internal/command"
	"kademlia/internal/commands/addcontact"
	"kademlia/internal/commands/exit"
	"kademlia/internal/commands/forget"
	"kademlia/internal/commands/get"
	"kademlia/internal/commands/getcontacts"
	"kademlia/internal/commands/getid"
	"kademlia/internal/commands/initnode"
	"kademlia/internal/commands/join"
	"kademlia/internal/commands/message"
	"kademlia/internal/commands/ping"
	"kademlia/internal/commands/put"
	"kademlia/internal/commands/storage"

	"github.com/rs/zerolog/log"
)

func ParseCmd(s string) Command {
	fields := strings.Fields(s)

	var command Command

	// Assume the string has already been checked to contain a command
	switch cmd := fields[0]; cmd {
	case "ping":
		command = new(ping.Ping)
	case "msg":
		command = new(message.Message)
	case "exit":
		command = new(exit.Exit)
	case "storage":
		command = new(storage.Storage)
	case "get":
		command = new(get.Get)
	case "getid":
		command = new(getid.GetId)
	case "addcontact":
		command = new(addcontact.AddContact)
	case "getcontacts":
		command = new(getcontacts.GetContacts)
	case "init":
		command = new(initnode.InitNode)
	case "put":
		command = new(put.Put)
	case "join":
		command = new(join.Join)
	case "forget":
		command = new(forget.Forget)
	default:
		log.Error().Str("command", cmd).Msg("Received unknown command")
		return nil
	}

	err := command.ParseOptions(fields[1:]) // Extract all options
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("Failed to parse options")
		log.Info().Msg(command.PrintUsage())
		return nil
	}

	return command
}
