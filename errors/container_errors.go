package errors

type NoSuchContainerError struct {
	ID string
	Err error
}

func (err *NoSuchContainerError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "No such container: " + err.ID
}

type AlreadyStartedError struct {
	ID string
	Err error
}

func (err *AlreadyStartedError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Container " + err.ID + " was already started!"
}

type AlreadyStoppedError struct {
	ID string
	Err error
}

func (err *AlreadyStoppedError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Container " + err.ID + " was already stopped!"
}

type DeleteParamError struct {
	ID string
	Err error
}

func (err *DeleteParamError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Param issue when deleting " + err.ID
}

type DeleteConflictError struct {
	ID string
	Err error
}

func (err *DeleteConflictError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return "Conflict when deleting " + err.ID
}
