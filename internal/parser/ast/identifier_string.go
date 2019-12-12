// Code generated by "aststring"; DO NOT EDIT.
package ast

import (
	"bytes"
	"fmt"
)

func (node *IdentifierReference) String() string {
	var buf bytes.Buffer
	buf.WriteString("IdentifierReference (")
	buf.WriteString("\n")
	if node.Identifier != nil {
		buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	buf.WriteString("\n")
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	buf.WriteString("\n")
	buf.WriteString(")")
	buf.WriteString("\n")
	return buf.String()
}

func (node *BindingIdentifier) String() string {
	var buf bytes.Buffer
	buf.WriteString("BindingIdentifier (")
	buf.WriteString("\n")
	if node.Identifier != nil {
		buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	buf.WriteString("\n")
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	buf.WriteString("\n")
	buf.WriteString(")")
	buf.WriteString("\n")
	return buf.String()
}

func (node *LabelIdentifier) String() string {
	var buf bytes.Buffer
	buf.WriteString("LabelIdentifier (")
	buf.WriteString("\n")
	if node.Identifier != nil {
		buf.WriteString(PrefixToString(node.Identifier.String(), "  "))
	}
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Yield", node.Yield), "  "))
	buf.WriteString("\n")
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "Await", node.Await), "  "))
	buf.WriteString("\n")
	buf.WriteString(")")
	buf.WriteString("\n")
	return buf.String()
}

func (node *Identifier) String() string {
	var buf bytes.Buffer
	buf.WriteString("Identifier (")
	buf.WriteString("\n")
	buf.WriteString(PrefixToString(fmt.Sprintf("%v: %v", "IdentifierName", node.IdentifierName), "  "))
	buf.WriteString("\n")
	buf.WriteString(")")
	buf.WriteString("\n")
	return buf.String()
}
