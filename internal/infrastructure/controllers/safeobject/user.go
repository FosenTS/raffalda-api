package safeobject

type Policy struct {
	Login      string `json:"login"`
	Permission uint   `json:"permission"`
}

func NewUser(login string, permission uint) *Policy {
	return &Policy{Login: login, Permission: permission}
}
