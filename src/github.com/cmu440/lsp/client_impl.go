// Contains the implementation of a LSP client.

package lsp

import (
	"errors"
	"time"
	"fmt"
	"github.com/cmu440/lspnet"
	"container/list"
)

type client struct {
	conn *lspnet.UDPConn
	epoch_buffer list.List
}

// NewClient creates, initiates, and returns a new client. This function
// should return after a connection with the server has been established
// (i.e., the client has received an Ack message from the server in response
// to its connection request), and should return a non-nil error if a
// connection could not be made (i.e., if after K epochs, the client still
// hasn't received an Ack message from the server in response to its K
// connection requests).
//
// hostport is a colon-separated string identifying the server's host address
// and port number (i.e., "localhost:9999").
func NewClient(hostport string, params *Params) (Client, error) {
	fmt.Println("Client params: %s\n", params);
	cli := new(client);
	if addr, err := lspnet.ResolveUDPAddr("udp", hostport); err != nil {
		fmt.Println("EROR", err)
		return nil, err
	} else if conn, err := lspnet.DialUDP("udp", nil, addr); err != nil {
		fmt.Println("EROR", err)
		return nil, err
	} else {
		cli.conn = conn
	}

	epoch_ticker := time.NewTicker(time.Millisecond *
		time.Duration(params.EpochMillis));

	go func() {
		for t := range epoch_ticker.C {
			fmt.Println("Ticked: ", t)
		}
	}()

	return nil, errors.New("not yet implemented")
}

func (c *client) ConnID() int {
	return -1
}

func (c *client) Read() ([]byte, error) {
	// TODO: remove this line when you are ready to begin implementing this method.
	select {} // Blocks indefinitely.
	return nil, errors.New("not yet implemented")
}

func (c *client) Write(payload []byte) error {
	return errors.New("not yet implemented")
}

func (c *client) Close() error {
	return errors.New("not yet implemented")
}
