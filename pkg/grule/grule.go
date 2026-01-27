package grule

import (
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type UserProfile struct {
	Age        int
	Category   string
	ServiceFee float64
}

func GetUserCategory() {
	user := &UserProfile{Age: 65}

	// 1. Initialize Library and Builder
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	// 2. Load the rule from the .grl file
	fileRes := pkg.NewFileResource("./grule/rules.grl")

	// 3. Build the rule into the library under a specific name and version
	err := ruleBuilder.BuildRuleFromResource("UserRules", "1.0.0", fileRes)
	if err != nil {
		panic(fmt.Sprintf("Failed to load rules: %v", err))
	}

	// 4. Prepare Context and Engine
	dataContext := ast.NewDataContext()
	dataContext.Add("User", user)

	kb, err := lib.NewKnowledgeBaseInstance("UserRules", "1.0.0")
	if err != nil {
		panic(fmt.Sprintf("Failed to load new instance: %v", err))
	}
	eng := engine.NewGruleEngine()

	// 5. Execute
	err = eng.Execute(dataContext, kb)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User Category: %s, Fee: $%.2f\n", user.Category, user.ServiceFee)
}
