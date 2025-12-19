package order

type service struct {
	orderRepo OrderRepo
}

func NewService(or OrderRepo) *service {
	return &service{
		orderRepo: or,
	}
}
