package sql_test

import (
	"fmt"
	"testing"

	"github.com/Briareos/rocket/sql"
)

func TestUserService_Get(t *testing.T) {
	db, err := sql.NewConnection("root", "patkatebi123", "52.24.138.122", "3306", "raketa")
	if err != nil {
		t.Fatal(err)
	}

	userService := sql.NewUserService(db)

	user, err := userService.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)
}
