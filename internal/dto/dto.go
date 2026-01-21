package dto

type Response struct {
	UserServiceStatus       int `json:"user_service_status"`
	VectorMemoryStatus      int `json:"vector_memory_status"`
	PermissionServiceStatus int `json:"permission_service_status"`
}
