package user

type userhandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *userhandler {
	return &userhandler{
		svc: svc,
	}
}
