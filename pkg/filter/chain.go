package filter

type (
	Logical    int
	Relational int
)

const (
	LogicalAnd Logical = 1
	LogicalOr  Logical = 2

	RelationalEq    Relational = 1
	RelationalNotEq Relational = 2
	RelationalGt    Relational = 3
	RelationalGte   Relational = 4
	RelationalLt    Relational = 5
	RelationalLte   Relational = 6
	RelationalIn    Relational = 7
)

type Expression struct {
	AttrName      string
	Val           any
	Logical       Logical
	Relational    Relational
	subExpression []Expression
}

func (e *Expression) GetSubExpression() []Expression {
	return e.subExpression
}

func (e *Expression) Or(ex *Expression) *Expression {
	e.Logical = LogicalOr
	e.subExpression = append(e.subExpression, *ex)
	return e
}

func (e *Expression) And(ex *Expression) *Expression {
	e.Logical = LogicalAnd
	e.subExpression = append(e.subExpression, *ex)
	return e
}
func NotEq(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalNotEq}
}
func Eq(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalEq}
}

func Gt(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalGt}
}

func Gte(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalGte}
}

func Lt(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalLt}
}

func Lte(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalLte}
}

func In(attrName string, val any) *Expression {
	return &Expression{AttrName: attrName, Val: val, Relational: RelationalIn}
}
