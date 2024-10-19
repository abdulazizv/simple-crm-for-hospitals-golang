package models

type AdminReq struct {
	Name         string `json:"name,omitempty"`
	UserName     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
type ListAdmin struct {
	Admins []AdminRes `json:"admins"`
}

type AdminRes struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	UserName     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
}

type AdminLoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
