package sql

import (
	"database/sql"
	"fmt"
	"github.com/Briareos/rocket"
)

type ruleService struct {
	db *sql.DB
	groupService rocket.GroupService
}

func NewRuleService(db *sql.DB, groupService rocket.GroupService) *ruleService {
	return &ruleService{
		db: db,
		groupService: groupService,
	}
}

func (service *ruleService) Get(ruleID int) (*rocket.Rule, error) {
	ruleQuery, err := service.db.Prepare(`SELECT description, aggregate, operation, threshold FROM rules WHERE id=?`)
	if err != nil {
		return nil, fmt.Errorf("prepare query: %v", err)
	}

	rule := rocket.Rule{
		ID: ruleID,
	}

	err = ruleQuery.QueryRow(ruleID).Scan(&(rule.Description), &(rule.Type), &(rule.Operation), &(rule).Threshold)
	if err != nil {
		return nil, fmt.Errorf("execute query: %v", err)
	}

	return &rule, nil
}


