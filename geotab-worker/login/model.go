package login

// authenticate request
type AuthParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// authenticate response
type Credentials struct {
	Database  string `json:"database"`
	UserName  string `json:"userName"`
	SessionId string `json:"sessionId"`
}

type LoginResult struct {
	Result struct {
		Credentials Credentials `json:"credentials"`
		Path        string      `json:"path"`
	} `json:"result"`
}