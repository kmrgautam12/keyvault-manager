package account

type CreateAccountInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
type CreateAccountOutput struct {
	Username string `json:"username"`
}
