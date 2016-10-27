package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
)

type userService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) rocket.UserService {
	return &userService{
		db: db,
	}
}

func (service userService) Get(userID int) (*rocket.User, error) {
	userQuery, err := service.db.Prepare(`SELECT google_account_id, first_name, last_name, title, email FROM users WHERE id=?`)
	if err != nil {
		return nil, fmt.Errorf("prepare get user query: %v", err)
	}

	user := rocket.User{}

	err = userQuery.QueryRow(userID).Scan()
	if err != nil {
		return nil, fmt.Errorf("query get user query: %v", err)
	}

	return nil, nil
}
