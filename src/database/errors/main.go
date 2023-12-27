package errors

type UserAlreadyExistsError struct {
	AlreadyExistingUserUsername string
}

func (e *UserAlreadyExistsError) Error() string {
	return "cannot create user - user with username " + e.AlreadyExistingUserUsername + " already exists"
}
