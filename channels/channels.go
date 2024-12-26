package channels

type Signal struct {
	ID      int
	Payload string
	Context <-chan struct{} // Cancellation context.
}

var (
	InputChan   = make(chan Signal, 10) // Buffered channel.
	OutputChan1 = make(chan Signal)
	OutputChan2 = make(chan Signal)
	OutputChan3 = make(chan Signal)
)
