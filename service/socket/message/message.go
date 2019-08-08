package message

import (
	"io"
)

type Sender interface {
	SendMessage(w io.Writer) error
}

type Receiver interface {
	ReceiveMessage(r io.Reader) error
}

type SendReceiver interface {
	Sender
	Receiver
}

type Message struct {
	from []byte `json:"from"`
	to   []byte `json:"to"`
	data []byte `json:"data"`
}
