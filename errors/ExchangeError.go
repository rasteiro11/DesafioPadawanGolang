package errors

type ExchangeError struct {
	Type string
}

func (e *ExchangeError) Error() string {
	return e.Type
}
