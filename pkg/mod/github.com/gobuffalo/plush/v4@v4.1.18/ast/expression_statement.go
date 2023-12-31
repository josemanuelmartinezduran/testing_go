package ast

type ExpressionStatement struct {
	TokenAble
	Expression Expression
}

var _ Statement = &ExpressionStatement{}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
