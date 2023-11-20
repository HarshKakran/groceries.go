package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store struct {
	// ID    int    `json:"id"`
	Owner User   `json:"owner"`
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
}

// TODO: Will be useful when using JWT or any other type of authentication
// var UserStoreMapping = make(map[User]Store)
