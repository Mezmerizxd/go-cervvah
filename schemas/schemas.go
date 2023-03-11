package schemas

type AccountSchema struct {
	AccountID string `json:"account_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type AccountForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileSchema struct {
	AccountID      string `json:"account_id"`
	Token          string `json:"token"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture"`
}
type ProfileForm struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture"`
}

type LocalDatabaseSchema struct {
	Accounts []AccountSchema
	Profiles []ProfileSchema
}