package sender

// Sender receives a batch of spans and send it to the appropiate transport
type Sender interface {
	Send([]byte)
}
