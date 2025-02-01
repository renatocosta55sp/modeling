package bus2

type EventHandler[T Event] interface {
	Handle(event T) error
}
