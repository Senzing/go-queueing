package queues

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Record interface {
	GetMessage() string
	GetMessageID() string
}
