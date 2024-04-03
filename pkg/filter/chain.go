package filter

type (
	Operator int
)

const (
	OpAnd   Operator = 1
	OpOr    Operator = 2
	OpEq    Operator = 3
	OpNotEq Operator = 4
	OpGt    Operator = 5
	OpGte   Operator = 6
	OpLt    Operator = 7
	OpLte   Operator = 8
	OpIn    Operator = 9
)

type Chain struct {
	root []Expression
}

func (c *Chain) Append(ex ...Expression) *Chain {
	c.root = append(c.root, ex...)
	return c
}

func (c *Chain) And(ex ...Expression) *Chain {
	c.root = append(c.root, And(ex...))
	return c
}

func (c *Chain) Or(ex ...Expression) *Chain {
	c.root = append(c.root, Or(ex...))
	return c
}

// id == 1 and id == 2 or (id == 3 and sd == 4) and (typ == 4 or typ == 5)
// and(id = 1, id = 2).or(and(id =3 id =4))
type Expression struct {
	AttrName string
	Value    any
	Values   []any
	Op       Operator
	exps     []Expression
}

func (e *Expression) GetExps() []Expression {
	return e.exps
}

func (e *Expression) append(ex ...Expression) {
	e.exps = append(e.exps, ex...)
}

func Or(ex ...Expression) Expression {
	e := &Expression{Op: OpOr}
	e.append(ex...)
	return *e
}

func And(ex ...Expression) Expression {
	e := &Expression{Op: OpAnd}
	e.append(ex...)
	return *e
}

func NotEq(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpNotEq}
}
func Eq(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpEq}
}

func Gt(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpGt}
}

func Gte(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpGte}
}

func Lt(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpLt}
}

func Lte(attrName string, val any) Expression {
	return Expression{AttrName: attrName, Value: val, Op: OpLte}
}

func In(attrName string, values ...any) Expression {
	return Expression{AttrName: attrName, Values: values, Op: OpIn}
}
