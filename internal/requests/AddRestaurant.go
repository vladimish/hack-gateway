package requests

type AddRestaurant struct {
	Auth        string `json:"auth"`
	Name        string `json:"name"`
	Login       string `json:"login"`
	Description string `json:"description"`
}
