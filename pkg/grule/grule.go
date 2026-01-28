package grule

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/sirupsen/logrus"
)

type CustomGrule struct {
	Library *ast.KnowledgeLibrary
	Context ast.IDataContext
	KB      *ast.KnowledgeBase
}

func GetNewEngine() *CustomGrule {
	return &CustomGrule{
		Library: ast.NewKnowledgeLibrary(),
		Context: ast.NewDataContext(),
	}
}

func (cr *CustomGrule) SetIDataContext(key string, data any) {
	cr.Context.Add(key, data)
}

func (cr *CustomGrule) SetKnowledgeBase(ruleName, version, rule string) {
	rb := builder.NewRuleBuilder(cr.Library)
	res := pkg.NewBytesResource([]byte(rule))
	err := rb.BuildRuleFromResource(ruleName, version, res)
	if err != nil {
		panic(err)
	}

	var errKb error
	cr.KB, errKb = cr.Library.NewKnowledgeBaseInstance(ruleName, version)
	if errKb != nil {
		panic(errKb)
	}
}

func (cr *CustomGrule) ExecuteRule() {
	eng := engine.NewGruleEngine()
	engine.SetLogger(logrus.New())
	err := eng.Execute(cr.Context, cr.KB)
	if err != nil {
		panic(err)
	}
}
