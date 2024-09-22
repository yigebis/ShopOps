package UseCase

type IErrorService interface {
	NoError() (int, error)
	InternalServer() (int, error)

	UserExists() (int, error)
	PendingVerification() (int, error)
	InvalidToken() (int, error)
	UserNotFound() (int, error)
	InvalidEmailPassword() (int, error)
	InvalidPhonePassword() (int, error)
	InvalidEmailRefresher() (int, error)
	NotVerified() (int, error)

	ExamNotFound() (int, error)
}