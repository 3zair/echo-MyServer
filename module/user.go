package module

type Login struct {
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Status   int        `json:"status"`
}

type Register struct {
	Status   int        `json:"status"`
}

type User struct {
	Name        string
	Password    string
}
