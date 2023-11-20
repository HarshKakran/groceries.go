package models

type User struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store struct {
	Owner User   `json:"owner"`
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
}
