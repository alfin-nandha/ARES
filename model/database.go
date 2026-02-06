package model

type Transaction struct {
}

func (*Transaction) TableName() string {
	return "public.transactions"
}

type Client struct {
	Id           int64  `gorm:"column:id"`
	Name         string `gorm:"column:name"`
	Username     string `gorm:"column:username"`
	PasswordHash string `gorm:"column:password"`
}

func (*Client) TableName() string {
	return "public.clients"
}

type Rule struct {
	Id          int64  `gorm:"column:id"`
	RuleSetId   int64  `gorm:"column:rule_set_id"`
	RuleName    string `gorm:"column:rule_name"`
	Description string `gorm:"column:description"`
	Drl_content string `gorm:"column:drl_content"`
	Priority    int    `gorm:"column:priority"`
	Version     string `gorm:"column:version"`
	IsActive    bool   `gorm:"column:is_active"`
}

func (*Rule) TableName() string {
	return "public.fraud_rules"
}

type RuleSet struct {
	Id          int64  `gorm:"column:id"`
	ClientId    int64  `gorm:"column:client_id"`
	Version     string `gorm:"column:version"`
	Description string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active"`
	Rules       []Rule `gorm:"foreignKey:RuleSetId"`
}

func (*RuleSet) TableName() string {
	return "public.rule_sets"
}

type FraudDecicion struct {
}

func (*FraudDecicion) TableName() string {
	return "public.fraud_decisions"
}
