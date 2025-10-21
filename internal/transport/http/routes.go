package http

type Service interface {
	DaysLeft() int64
	MongoService() MongoService
}

type MongoService interface {
	IsExist(telegramID int64) (bool, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}
