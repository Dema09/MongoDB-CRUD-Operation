package domain

type ProfileResponse struct{
	FirstName string `json: "first_name"`
	LastName string `json: "last_name"`
	Address string `json: "address"`
	Age float64 `json: "age"`
	ProfilePicture string `json: "profile_picture"`

}
