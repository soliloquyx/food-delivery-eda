package app

type service struct{}

func NewService() *service {
	return &service{}
}
