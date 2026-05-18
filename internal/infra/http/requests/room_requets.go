package requests

type RoomRequest struct{
	Name string `json:"name"`
	Description *string `json:"description"`
} 