package casdoor

type AuthUser struct {
	Sub      string `json:"sub"`
	Username string `json:"preferred_username"`
}
