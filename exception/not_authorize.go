package exception

type NotAuthorize struct {
	Error string
}

func NewNotAuthorize(error string) NotAuthorize {
	return NotAuthorize{Error: error}
}
