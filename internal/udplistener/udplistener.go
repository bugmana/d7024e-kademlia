package udplistener

import (
	"kademlia/internal/contact"
	"kademlia/internal/node"
	"kademlia/internal/rpc"
	"kademlia/internal/rpc/parser"
	"net"
	"strings"

	"github.com/rs/zerolog/log"
)

// Listen initiates a UDP server
func Listen(ip string, port int) {
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: port}
	ln, err := net.ListenUDP("udp4", &addr)
	defer ln.Close()
	if err != nil {
		log.Error().Msgf("Failed to listen on UDP Address: %s", err)
	}
	log.Info().Str("Address", addr.String()).Msg("Listening on UDP packets on address")

	waitForMessages(ln)
}

func waitForMessages(c *net.UDPConn) {
	for {
		buf := make([]byte, 512)
		nr, addr, err := c.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		data := buf[0:nr]
		rpcMsg, err := rpc.Deserialize(string(data))
		if err == nil {
			log.Info().
				Str("Content", rpcMsg.Content).
				Str("SenderId", rpcMsg.SenderId.String()).
				Str("RPCId", rpcMsg.RPCId.String()).
				Msg("Received message")

			c := contact.NewContact(rpcMsg.SenderId, addr.String())
			node.KadNode.RoutingTable.AddContact(c)

			cmd, err := rpcparser.ParseRPC(&c, &rpcMsg)
			if err != nil {
				log.Warn().Str("Error", err.Error()).Msg("Failed to parse RPC")
				continue
			}

			options := strings.Split(rpcMsg.Content, " ")[1:]
			if err = cmd.ParseOptions(&options); err == nil {
				cmd.Execute()
			} else {
				log.Warn().
					Str("Error", err.Error()).
					Msg("Failed to parse RPC options")
			}

			log.Debug().Str("Id", c.ID.String()).Str("Address", c.Address).Msg("Updating bucket")
		} else {
			log.Warn().Str("Error", err.Error()).Msg("Failed to deserialize message")
		}
	}
}