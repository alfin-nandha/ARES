package model

type Client struct {
	Id   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (*Client) TableName() string {
	return "public.clients"
}

type Rule struct {
	Id          int64  `gorm:"column:id"`
	ClientId    int64  `gorm:"column:client_id"`
	RuleName    string `gorm:"column:rule_name"`
	Description string `gorm:"column:description"`
	drl_content string `gorm:"column:drl_content"`
	priority    int    `gorm:"column:priority"`
	version     string `gorm:"column:version"`
	Isactive    bool   `gorm:"column:is_active"`
}

func (*Rule) TableName() string {
	return "public.fraud_rules"
}
