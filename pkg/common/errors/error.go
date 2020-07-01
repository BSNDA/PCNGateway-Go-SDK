package errors

func New(message string) error {
	return &BsnError{
		msg: message,
	}
}

type BsnError struct {
	msg string
}

func (b *BsnError) Error() string { return b.msg }
