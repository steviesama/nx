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
	Guid        string `json:"Guid"`
	From        []byte `json:"From"`
	To          []byte `json:"To"`
	TotalSize   int    `json:"TotalSize"`
	BytesCopied int    `json:"BytesCopied"`
	Data        []byte `json:"Data"`
}
