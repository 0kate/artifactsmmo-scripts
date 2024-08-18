package dto

type DeleteCharacterRequest struct {
	Name string `json:"name"`
}

func NewDeleteCharacterRequest(name string) *DeleteCharacterRequest {
	return &DeleteCharacterRequest{
		Name: name,
	}
}
