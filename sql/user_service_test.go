package sql_test

import (
	"fmt"
	"testing"

	"github.com/Briareos/rocket/container"
	"github.com/Briareos/rocket/sql"
	"path/filepath"
)

func TestUserService_Get(t *testing.T) {
	c := container.MustLoadFromPath(filepath.Join("..", "config-prod.yml"))

	userService := sql.NewUserService(c.DB())

	user, err := userService.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("USER: %#v\n", user)

	fmt.Println("JOINED GROUPS: ")
	for _, group := range user.JoinedGroups {
		fmt.Printf("\tGROUP: %#v\n", group)
		fmt.Println("\t\tRULES: ")
		for _, rule := range group.Rules {
			fmt.Printf("\t\tRULE: %#v\n", rule)
		}
	}

	fmt.Println("WATCHED GROUPS: ")
	for _, group := range user.WatchedGroups {
		fmt.Printf("\tGROUP: %#v\n", group)
		fmt.Println("\t\tRULES: ")
		for _, rule := range group.Rules {
			fmt.Printf("\t\tRULE: %#v\n", rule)
		}
	}

	fmt.Println("STATUSES: ")
	for _, status := range user.Statuses {
		fmt.Printf("\tSTATUS: %#v\n", status)
	}
}
