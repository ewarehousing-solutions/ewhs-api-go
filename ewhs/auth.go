package ewhs

type AuthToken struct {
	Token        string `json:"token,omitempty"`
	Iat          int    `json:"iat,omitempty"`
	Exp          int    `json:"exp,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
