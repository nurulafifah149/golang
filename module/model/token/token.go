package token

type Tokens struct {
	AccessToken string `json:"access_token"`
}

type DefaultClaim struct {
	Expired   int    `json:"exp"`
	NotBefore int    `json:"nbf"`
	IssuedAt  int    `json:"iat"`
	Issuer    string `json:"iss"`
	Audience  string `json:"aud"`
	JTI       string `json:"jti"`
	Typ       string `json:"type"`
}

type AccessClaim struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}
