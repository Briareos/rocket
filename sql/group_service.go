package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
	"strconv"
	"strings"
	"time"
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

	if err := service.selectRulesQuery(groups); err != nil {
		return nil, fmt.Errorf("select rules: %v", err)
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

func (service *groupService) selectRulesQuery(groups []*rocket.Group) error {
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
		WHERE rules.group_id IN (%s)
	`, strings.Join(groupIds, ", ")))
	if err != nil {
		return fmt.Errorf("prepare query: %v", err)
	}

	rows, err := rulesQuery.Query()
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

func (service groupService) GetAvailableBodyCounts(group *rocket.Group, day time.Time) (map[time.Time]int, error) {
	availableStatuses := []string{"'available'"}

	if group.Availability.Busy {
		availableStatuses = append(availableStatuses, "'busy'")
	}
	if group.Availability.Remote {
		availableStatuses = append(availableStatuses, "'remote'")
	}

	query, err := service.db.Prepare(fmt.Sprintf(`SELECT statuses.date as date, COUNT(*) as count FROM user_group_assignments
		JOIN users ON user_group_assignments.user_id=users.id
		JOIN statuses ON users.id=statuses.user_id
		WHERE user_group_assignments.group_id=?
		AND statuses.type IN(%s)
		AND YEAR(statuses.date)=?
		AND MONTH(statuses.date)=?
		GROUP BY statuses.date`,
		strings.Join(availableStatuses, ",")))

	counts := make(map[time.Time]int)

	if err != nil {
		return counts, fmt.Errorf("prepare query: %v", err)
	}

	//rows, err := query.Query(group.ID, day.Year(), int(day.Month()))
	rows, err := query.Query(group.ID, day.Year(), int(day.Month()))
	if err != nil {
		return counts, fmt.Errorf("prepare query: %v", err)
	}

	for {
		if hasNext := rows.Next(); !hasNext {
			break
		}

		var date time.Time
		var count int

		err = rows.Scan(&date, &count)
		if err != nil {
			return nil, fmt.Errorf("scan row: %v", err)
		}

		counts[date] = count
	}

	if err != nil {
		return counts, fmt.Errorf("exec query: %v", err)
	}

	return counts, nil
}

func (service groupService) GetTotalBodyCount(group *rocket.Group) (int, error) {
	query, err := service.db.Prepare(`SELECT COUNT(*) as count FROM user_group_assignments WHERE group_id=?`)
	if err != nil {
		return 0, fmt.Errorf("prepare query: %v", err)
	}

	var count int

	err = query.QueryRow(group.ID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("exec query: %v", err)
	}

	return count, nil
}
