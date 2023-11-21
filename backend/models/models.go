package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	Store    *Store `json:"store"`
}

type Store struct {
	// ID    int    `json:"id"`
	// Owner User   `json:"owner"`
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
}

type Orders []Item

// TODO: Will be useful when using JWT or any other type of authentication
// var UserStoreMapping = make(map[User]Store)
var Users []User
var ConsumerOrdersMapping = make(map[int]Orders)

func Save[T []User | *map[int]Orders](filepath string, target T) error {
	data, err := json.Marshal(target)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

func Load[T []User | map[int]Orders](filePath string, target *T) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	fmt.Printf("File Size: %v\n", len(f))

	if len(f) == 0 {
		return nil
	}

	return json.Unmarshal(f, target)
}

func GetIndexAndUsername(itemId string) (int, string, error) {
	idParts := strings.Split(itemId, "_")
	if len(idParts) != 2 {
		return 0, "", errors.New("invalid item id")
	}
	itemIdx, err := strconv.Atoi(idParts[0])
	if err != nil {
		log.Panicf("error while converting string to int. %v", err)
	}

	username := idParts[1]

	return itemIdx, username, nil
}
