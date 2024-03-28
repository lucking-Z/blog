package Filter

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

type Expression struct {
	AttrName     string
	Val          any
	LogicalOp    LogicalOperator
	RelationalOp RelationalOperator
	Nodes        []Expression
}

type Chain struct {
	ex []Expression
}

func NewFilter() *Chain {
	return &Chain{
		ex: make([]Expression, 0),
	}
}

func (f *Chain) IsEmpty() bool {
	return len(f.ex) < 1
}

func (f *Chain) GetFs() []Expression {
	return f.ex
}

// id = 1
// (id = 2 and id = 3 or (id = 4 and = id = 5))

// 同级
// 子级
func (f *Chain) Set(attrName string, relationalOp RelationalOperator, val any) *Chain {
	f.ex = append(f.ex, Expression{AttrName: attrName, Val: val, LogicalOp: And, RelationalOp: relationalOp})
	return f
}

func appendLastNode(ex Expression, mark *Expression) {
	if len(mark.Nodes) < 1 {
		mark.Nodes = append(mark.Nodes, ex)
	} else {
		appendLastNode(ex, &mark.Nodes[len(mark.Nodes)-1])
	}
}

func appendNode(ex Expression, mark *Expression) {

	appendNode(ex, &mark.Nodes[len(mark.Nodes)-1])
}

func (f *Chain) Append(attrName string, relationalOp RelationalOperator, val any) *Chain {
	if len(f.ex) < 1 {

		return f.Set(attrName, relationalOp, val)
	}
	ex := f.ex[len(f.ex)-1]
	ex.Nodes = append(ex.Nodes, Expression{AttrName: attrName, Val: val, LogicalOp: And, RelationalOp: relationalOp})
	return f
}

func (f *Chain) OrSet(attrName string, relationalOp RelationalOperator, val any) *Chain {
	f.ex = append(f.ex, Expression{AttrName: attrName, Val: val, LogicalOp: And, RelationalOp: relationalOp})
	return f
}

func (f *Chain) OrAppend(attrName string, relationalOp RelationalOperator, val any) *Chain {
	if len(f.ex) < 1 {
		return f.OrSet(attrName, relationalOp, val)
	}
	ex := f.ex[len(f.ex)-1]
	ex.Nodes = append(ex.Nodes, Expression{AttrName: attrName, Val: val, LogicalOp: Or, RelationalOp: relationalOp})
	return f
}

//id = 1 or id = 2 or (id = 3 and type = 1 and (s = 2 or s = 3 ))
//
//func (f *Chain) Eq(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: Eq})
//	return f
//}
//
//func (f *Chain) NotEq(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: NotEq})
//	return f
//}
//
//func (f *Chain) Gt(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: Gt})
//	return f
//}
//
//func (f *Chain) GtE(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: GtE})
//	return f
//}
//
//func (f *Chain) Lt(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: Lt})
//	return f
//}
//
//func (f *Chain) LtE(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: LtE})
//	return f
//}
//
//func (f *Chain) In(attrName string, val any, logicalOp LogicalOperator) *Chain {
//	f.fs = append(f.fs, Expression{AttrName: attrName, Val: val, LogicalOp: logicalOp, RelationalOp: In})
//	return f
//}
