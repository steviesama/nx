package data

type Representer interface {
}

type Identifier interface {
}

type Serializer interface {
}

type Packet struct {
	payload []byte
}
