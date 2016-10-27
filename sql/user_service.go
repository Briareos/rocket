package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
)

type userService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *userService {
	return &userService{
		db: db,
	}
}

func (service *userService) Get(userID int) (*rocket.User, error) {
	user, err := service.selectUserQuery(userID)
	if err != nil {
		return nil, fmt.Errorf("select user: %v", err)
	}

	user.JoinedGroups, err = service.selectGroupsQuery(userID, "user_group_assignments")
	if err != nil {
		return nil, fmt.Errorf("select joined groups: %v", err)
	}

	user.JoinedGroups, err = service.selectGroupsQuery(userID, "user_group_watches")
	if err != nil {
		return nil, fmt.Errorf("select watched groups: %v", err)
	}

	return user, nil
}

func (service *userService) selectUserQuery(userID int) (*rocket.User, error) {
	userQuery, err := service.db.Prepare(`SELECT google_account_id, first_name, last_name, title, email FROM users WHERE id=?`)
	if err != nil {
		return nil, fmt.Errorf("prepare query: %v", err)
	}

	user := rocket.User{
		ID: userID,
	}

	err = userQuery.QueryRow(userID).Scan(&(user.GoogleID), &(user.FirstName), &(user.LastName), &(user.Title), &(user.Email))
	if err != nil {
		return nil, fmt.Errorf("execute query: %v", err)
	}

	return &user, nil
}

func (service *userService) selectGroupsQuery(userID int, relation string) ([]*rocket.Group, error) {
	groupsQuery, err := service.db.Prepare(fmt.Sprintf(`
		SELECT id, name, description, busy_value, remote_value
		FROM groups
		INNER JOIN %s ON groups.id = %s.group_id
		WHERE user_id=?
	`, relation, relation))
	if err != nil {
		return nil, fmt.Errorf("prepare query: %v", err)
	}

	groups := []*rocket.Group{}

	rows, err := groupsQuery.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("execute query: %v", err)
	}

	for {
		if hasNext := rows.Next(); !hasNext {
			if err := rows.Err(); err != nil {
				return nil, fmt.Errorf("next row: %v", err)
			}

			return groups, nil
		}

		group := rocket.Group{
			Availability: rocket.DefaultAvailability,
		}

		err = rows.Scan(&(group.ID), &(group.Name), &(group.Description), &(group.Availability.Busy), &(group.Availability.Remote))
		if err != nil {
			return nil, fmt.Errorf("scan row: %v", err)
		}

		groups = append(groups, &group)
	}
}
