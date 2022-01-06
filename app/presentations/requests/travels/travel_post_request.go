package requests

type member struct {
	Name string `json:"name" binding:"required"`
}

type travel struct {
	Name string `json:"name" binding:"required"`
}

type TravelPostRequest struct {
	Members []member `json:"members" binding:"required,dive"`
	Travel  travel   `json:"travel" binding:"required,dive"`
}
