package errors

type NoSuchContainerError struct {
	CID string
	Err error
}

func (err *NoSuchContainerError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "No such container: " + err.CID
}

type AlreadyStartedError struct {
	CID string
	Err error
}

func (err *AlreadyStartedError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Container " + err.CID + " was already started!"
}

type AlreadyStoppedError struct {
	CID string
	Err error
}

func (err *AlreadyStoppedError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Container " + err.CID + " was already stopped!"
}

type DeleteParamError struct {
	CID string
	Err error
}

func (err *DeleteParamError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Param issue when deleting " + err.CID
}

type DeleteConflictError struct {
	CID string
	Err error
}

func (err *DeleteConflictError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Conflict when deleting " + err.CID
}
