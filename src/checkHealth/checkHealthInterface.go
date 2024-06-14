package checkHealth

type CheckHealthRepository interface {
	RetrieveVersion() (string, error)
}

type CheckHealthUsecase interface {
	GetVersion() (string, error)
}
