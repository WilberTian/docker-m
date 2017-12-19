package errors

type NoSuchImageError struct {
	ID string
	Err error
}

func (err *NoSuchImageError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "No such image: " + err.ID
}

