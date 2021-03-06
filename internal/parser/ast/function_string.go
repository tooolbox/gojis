// Code generated by "aststring"; DO NOT EDIT.
package ast

import (
	"bytes"
	"fmt"
)

func (node *FormalParameters) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FormalParameters (")
	_, _ = buf.WriteString("\n")
	if node.FunctionRestParameter != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionRestParameter.String(), "  "))
	}
	if node.FormalParameterList != nil {
		_, _ = buf.WriteString(PrefixToString(node.FormalParameterList.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Comma", node.Comma), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *FunctionRestParameter) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FunctionRestParameter (")
	_, _ = buf.WriteString("\n")
	if node.BindingRestElement != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingRestElement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *FormalParameterList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FormalParameterList (")
	_, _ = buf.WriteString("\n")
	for _, elem := range node.FormalParameters {
		_, _ = buf.WriteString(PrefixToString(elem.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *FormalParameter) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FormalParameter (")
	_, _ = buf.WriteString("\n")
	if node.BindingElement != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingElement.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *FunctionBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FunctionBody (")
	_, _ = buf.WriteString("\n")
	if node.FunctionStatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionStatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *FunctionStatementList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("FunctionStatementList (")
	_, _ = buf.WriteString("\n")
	if node.StatementList != nil {
		_, _ = buf.WriteString(PrefixToString(node.StatementList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *GeneratorBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("GeneratorBody (")
	_, _ = buf.WriteString("\n")
	if node.FunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ArrowFunction) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ArrowFunction (")
	_, _ = buf.WriteString("\n")
	if node.ArrowParameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.ArrowParameters.String(), "  "))
	}
	if node.ConciseBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.ConciseBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ArrowParameters) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ArrowParameters (")
	_, _ = buf.WriteString("\n")
	if node.BindingIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingIdentifier.String(), "  "))
	}
	if node.CoverParenthesizedExpressionAndArrowParameterList != nil {
		_, _ = buf.WriteString(PrefixToString(node.CoverParenthesizedExpressionAndArrowParameterList.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *ConciseBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("ConciseBody (")
	_, _ = buf.WriteString("\n")
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	if node.FunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CoverParenthesizedExpressionAndArrowParameterList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CoverParenthesizedExpressionAndArrowParameterList (")
	_, _ = buf.WriteString("\n")
	if node.Expression != nil {
		_, _ = buf.WriteString(PrefixToString(node.Expression.String(), "  "))
	}
	if node.BindingIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingIdentifier.String(), "  "))
	}
	if node.BindingPattern != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingPattern.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Comma", node.Comma), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Ellipsis", node.Ellipsis), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *MethodDefinition) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("MethodDefinition (")
	_, _ = buf.WriteString("\n")
	if node.PropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyName.String(), "  "))
	}
	if node.UniqueFormalPatameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.UniqueFormalPatameters.String(), "  "))
	}
	if node.FunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionBody.String(), "  "))
	}
	if node.GeneratorMethod != nil {
		_, _ = buf.WriteString(PrefixToString(node.GeneratorMethod.String(), "  "))
	}
	if node.AsyncMethod != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncMethod.String(), "  "))
	}
	if node.AsyncGeneratorMethod != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncGeneratorMethod.String(), "  "))
	}
	if node.PropertySetParameterList != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertySetParameterList.String(), "  "))
	}
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Get", node.Get), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Set", node.Set), "  "))
	_, _ = buf.WriteString("\n")
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *PropertySetParameterList) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("PropertySetParameterList (")
	_, _ = buf.WriteString("\n")
	if node.FormalParameter != nil {
		_, _ = buf.WriteString(PrefixToString(node.FormalParameter.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *UniqueFormalParameters) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("UniqueFormalParameters (")
	_, _ = buf.WriteString("\n")
	if node.FormalParameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.FormalParameters.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *GeneratorMethod) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("GeneratorMethod (")
	_, _ = buf.WriteString("\n")
	if node.PropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyName.String(), "  "))
	}
	if node.UniqueFormalParameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.UniqueFormalParameters.String(), "  "))
	}
	if node.GeneratorBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.GeneratorBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncMethod) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncMethod (")
	_, _ = buf.WriteString("\n")
	if node.PropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyName.String(), "  "))
	}
	if node.UniqueFormalParameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.UniqueFormalParameters.String(), "  "))
	}
	if node.AsyncFunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncFunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncFunctionBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncFunctionBody (")
	_, _ = buf.WriteString("\n")
	if node.FunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncGeneratorMethod) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncGeneratorMethod (")
	_, _ = buf.WriteString("\n")
	if node.PropertyName != nil {
		_, _ = buf.WriteString(PrefixToString(node.PropertyName.String(), "  "))
	}
	if node.UniqueFormalParameters != nil {
		_, _ = buf.WriteString(PrefixToString(node.UniqueFormalParameters.String(), "  "))
	}
	if node.AsyncGeneratorBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncGeneratorBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncGeneratorBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncGeneratorBody (")
	_, _ = buf.WriteString("\n")
	if node.FunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.FunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncArrowFunction) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncArrowFunction (")
	_, _ = buf.WriteString("\n")
	if node.AsyncArrowBindingIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncArrowBindingIdentifier.String(), "  "))
	}
	if node.AsyncConciseBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncConciseBody.String(), "  "))
	}
	if node.CoverCallExpressionAndAsyncArrowHead != nil {
		_, _ = buf.WriteString(PrefixToString(node.CoverCallExpressionAndAsyncArrowHead.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncConciseBody) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncConciseBody (")
	_, _ = buf.WriteString("\n")
	if node.AssignmentExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.AssignmentExpression.String(), "  "))
	}
	if node.AsyncFunctionBody != nil {
		_, _ = buf.WriteString(PrefixToString(node.AsyncFunctionBody.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *AsyncArrowBindingIdentifier) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("AsyncArrowBindingIdentifier (")
	_, _ = buf.WriteString("\n")
	if node.BindingIdentifier != nil {
		_, _ = buf.WriteString(PrefixToString(node.BindingIdentifier.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}

func (node *CoverCallExpressionAndAsyncArrowHead) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("CoverCallExpressionAndAsyncArrowHead (")
	_, _ = buf.WriteString("\n")
	if node.MemberExpression != nil {
		_, _ = buf.WriteString(PrefixToString(node.MemberExpression.String(), "  "))
	}
	if node.Arguments != nil {
		_, _ = buf.WriteString(PrefixToString(node.Arguments.String(), "  "))
	}
	_, _ = buf.WriteString(")")
	_, _ = buf.WriteString("\n")
	return buf.String()
}
