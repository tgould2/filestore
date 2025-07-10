package p2p

// HandshakeFunc is ...?
type HandshakeFunc func(Peer) error

func NOPHandhshakeFunc(Peer) error { return nil }
