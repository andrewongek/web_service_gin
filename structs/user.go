package structs

type User struct {
	Id        int32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
