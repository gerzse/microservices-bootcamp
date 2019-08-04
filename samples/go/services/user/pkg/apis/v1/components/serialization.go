package components

type Name struct{

	First string `json:"first"`
	Last  string `json:"last"`
}

type User struct{

	Name Name `json:"name"`
	Gender string `json:"gender"`
	BirthYear uint16 `json:"born"`
}