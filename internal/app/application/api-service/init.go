package api_service

type ApiService struct {
	todoRepository todoRepository
	userRepository userRepository
}

type todoRepository interface {
	Create()
	Update()
	Get()
	Delete()
}

type userRepository interface {
	Create()
}

func New(
	todoRepository todoRepository,
	userRepository userRepository,
) *ApiService {
	return &ApiService{}
}
