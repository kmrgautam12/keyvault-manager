package utils

type CreateAccountInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
type CreateAccountOutput struct {
	Username string `json:"username"`
}
type GenerateJWTInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
type GenerateJWTTokenOutput struct {
	Token      string `json:"token"`
	TokenId    string `json:"token_id"`
	ValidUntil int64  `json:"valid_until"`
}
