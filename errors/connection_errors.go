package errors

type UnixConenctionError struct {
	Path string
	Err  error
}

func (err *UnixConenctionError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Error to connect to " + err.Path
}
