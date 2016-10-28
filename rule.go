package rocket

const (
	RuleTypeCount      = "count"
	RuleTypePercentage = "percentage"

	RuleOperatorLessThan           = "<"
	RuleOperatorLessThanOrEqual    = "<="
	RuleOperatorGreaterThan        = ">"
	RuleOperatorGreaterThanOrEqual = ">="
	RuleOperatorEqual              = "="
)

type Rule struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Operation   string `json:"operation"`
	Threshold   int    `json:"threshold"`
}

type RuleService interface {
	Get(int) (*Rule, error)
}
