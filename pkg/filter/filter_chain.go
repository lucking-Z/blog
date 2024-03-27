package filter

type (
	LogicalOperator    int
	RelationalOperator int
)

const (
	And LogicalOperator = 1
	Or  LogicalOperator = 2

	Eq    RelationalOperator = 1
	NotEq RelationalOperator = 2
	Gt    RelationalOperator = 3
	GtE   RelationalOperator = 4
	Lt    RelationalOperator = 5
	LtE   RelationalOperator = 6
	In    RelationalOperator = 7
)

type expression struct {
	attrName     string
	val          any
	logicalOp    LogicalOperator
	relationalOp RelationalOperator
}

type Filter struct {
	fs []expression
}

func NewFilter() *Filter {
	return &Filter{
		fs: make([]expression, 0),
	}
}

func (f *Filter) GetFs() []expression {
	return f.fs
}

func (f *Filter) Eq(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: Eq})
	return f
}

func (f *Filter) NotEq(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: NotEq})
	return f
}

func (f *Filter) Gt(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: Gt})
	return f
}

func (f *Filter) GtE(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: GtE})
	return f
}

func (f *Filter) Lt(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: Lt})
	return f
}

func (f *Filter) LtE(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: LtE})
	return f
}

func (f *Filter) In(attrName string, val any, logicalOp LogicalOperator) *Filter {
	f.fs = append(f.fs, expression{attrName: attrName, val: val, logicalOp: logicalOp, relationalOp: In})
	return f
}
