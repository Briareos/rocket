package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
)

type groupService struct {
	db *sql.DB
}

func NewGroupService(db *sql.DB) rocket.GroupService {
	return &groupService{
		db: db,
	}
}

func (service groupService) GetAll() ([]*rocket.Group, error) {
	rows, err := service.db.Query(`SELECT id, name, description, busy_value, remote_value FROM groups`)
	if err != nil {
		return nil, fmt.Errorf("cannot connect for query: %v", err)
	}

	var groups []*rocket.Group

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

	return groups, nil
}

func (service groupService) Add(group *rocket.Group) error {
	query, err := service.db.Prepare(`
		INSERT INTO groups (name, description, busy_value, remote_value)
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare query: %v", err)
	}

	_, err = query.Exec(group.Name, group.Description, group.Availability.Busy, group.Availability.Remote)
	if err != nil {
		return fmt.Errorf("exec query: %v", err)
	}
	return nil
}

func (service groupService) AddRule(group *rocket.Group, rule *rocket.Rule) error {
	query, err := service.db.Prepare(`
		INSERT INTO rules (group_id, description, aggregate, operation, threshold)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare query: %v", err)
	}

	_, err = query.Exec(group.ID, rule.Description, rule.Type, rule.Operation, rule.Threshold)
	if err != nil {
		return fmt.Errorf("exec query: %v", err)
	}
	return nil
}
