package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
	"strconv"
	"strings"
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

	user.WatchedGroups, err = service.selectGroupsQuery(userID, "user_group_watches")
	if err != nil {
		return nil, fmt.Errorf("select watched groups: %v", err)
	}

	user.Statuses, err = service.selectStatusesQuery(userID)
	if err != nil {
		return nil, fmt.Errorf("select statuses: %v", err)
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
			break
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

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next row: %v", err)
	}

	if err := service.selectRulesQuery(userID, groups); err != nil {
		return nil, fmt.Errorf("select rules: %v", err)
	}

	return groups, nil
}

func (service *userService) selectRulesQuery(userID int, groups []*rocket.Group) error {
	groupIds := []string{}
	indexedGroups := map[int]*rocket.Group{}

	for _, group := range groups {
		groupIds = append(groupIds, strconv.Itoa(group.ID))
		indexedGroups[group.ID] = group
	}

	rulesQuery, err := service.db.Prepare(fmt.Sprintf(`
		SELECT group_id, id, description, threshold, operation, aggregate
		FROM rules
		LEFT JOIN user_rule_mutes ON user_rule_mutes.rule_id = rules.id
		WHERE user_id != ? AND rules.group_id IN (%s)
	`, strings.Join(groupIds, ", ")))
	if err != nil {
		return fmt.Errorf("prepare query: %v", err)
	}

	rows, err := rulesQuery.Query(userID)
	if err != nil {
		return fmt.Errorf("execute query: %v", err)
	}

	for {
		if hasNext := rows.Next(); !hasNext {
			if err := rows.Err(); err != nil {
				return fmt.Errorf("next row: %v", err)
			}

			return nil
		}

		var groupId int
		rule := rocket.Rule{}

		err = rows.Scan(&groupId, &(rule.ID), &(rule.Description), &(rule.Threshold), &(rule.Operation), &(rule.Type))
		if err != nil {
			return fmt.Errorf("scan row: %v", err)
		}

		if group, ok := indexedGroups[groupId]; ok {
			group.Rules = append(group.Rules, &rule)
		}
	}

	return nil
}

func (service *userService) selectStatusesQuery(userID int) ([]*rocket.Status, error) {
	statusesQuery, err := service.db.Prepare(`SELECT id, date, reason, type FROM statuses WHERE user_id=?`)
	if err != nil {
		return nil, fmt.Errorf("prepare query: %v", err)
	}

	statuses := []*rocket.Status{}

	rows, err := statusesQuery.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("execute query: %v", err)
	}

	for {
		if hasNext := rows.Next(); !hasNext {
			if err := rows.Err(); err != nil {
				return nil, fmt.Errorf("next row: %v", err)
			}

			return statuses, nil
		}

		status := rocket.Status{}

		err = rows.Scan(&(status.ID), &(status.Date), &(status.Reason), &(status.Type))
		if err != nil {
			return nil, fmt.Errorf("scan row: %v", err)
		}

		statuses = append(statuses, &status)
	}
}
