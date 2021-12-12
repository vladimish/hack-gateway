package requests

type AddTable struct {
	Auth  string `json:"auth"`
	Name  string `json:"name"`
	Login string `json:"login"`
}
