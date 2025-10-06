package test_data

type Client struct {
	Email    string
	UserID   string
	Password string
}

type ClientInfo struct {
	ClientOne Client
}

func TestInfo() ClientInfo {
	return ClientInfo{
		ClientOne: Client{
			Email:    "",
			UserID:   "",
			Password: "",
		},
	}
}
