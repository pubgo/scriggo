//
// Copyright (c) 2016-2017 Open2b Software Snc. All Rights Reserved.
//

package template

import (
	"fmt"
	"testing"

	"open2b/template/ast"
)

var exprTests = []struct {
	src  string
	node ast.Node
}{
	{"_", ast.NewIdentifier(0, "_")},
	{"a", ast.NewIdentifier(0, "a")},
	{"a5", ast.NewIdentifier(0, "a5")},
	{"_a", ast.NewIdentifier(0, "_a")},
	{"_5", ast.NewIdentifier(0, "_5")},
	{"0", ast.NewInt(0, 0)},
	{"3", ast.NewInt(0, 3)},
	{"2147483647", ast.NewInt(0, 2147483647)},
	{"-2147483648", ast.NewInt(0, -2147483648)},
	{"\"\"", ast.NewString(0, "")},
	{"\"a\"", ast.NewString(0, "a")},
	{`"\t"`, ast.NewString(0, "\t")},
	{`"\a\b\f\n\r\t\v\\\""`, ast.NewString(0, "\a\b\f\n\r\t\v\\\"")},
	{"``", ast.NewString(0, "")},
	{"`\\t`", ast.NewString(0, "\\t")},
	{"!a", ast.NewUnaryOperator(0, ast.OperatorNot, ast.NewIdentifier(1, "a"))},
	{"1+2", ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1-2", ast.NewBinaryOperator(1, ast.OperatorSubtraction, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1*2", ast.NewBinaryOperator(1, ast.OperatorMultiplication, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1/2", ast.NewBinaryOperator(1, ast.OperatorDivision, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1%2", ast.NewBinaryOperator(1, ast.OperatorModulo, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1==2", ast.NewBinaryOperator(1, ast.OperatorEqual, ast.NewInt(0, 1), ast.NewInt(3, 2))},
	{"1!=2", ast.NewBinaryOperator(1, ast.OperatorNotEqual, ast.NewInt(0, 1), ast.NewInt(3, 2))},
	{"1<2", ast.NewBinaryOperator(1, ast.OperatorLess, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1<=2", ast.NewBinaryOperator(1, ast.OperatorLessOrEqual, ast.NewInt(0, 1), ast.NewInt(3, 2))},
	{"1>2", ast.NewBinaryOperator(1, ast.OperatorGreater, ast.NewInt(0, 1), ast.NewInt(2, 2))},
	{"1>=2", ast.NewBinaryOperator(1, ast.OperatorGreaterOrEqual, ast.NewInt(0, 1), ast.NewInt(3, 2))},
	{"a&&b", ast.NewBinaryOperator(1, ast.OperatorAnd, ast.NewIdentifier(0, "a"), ast.NewIdentifier(3, "b"))},
	{"a||b", ast.NewBinaryOperator(1, ast.OperatorOr, ast.NewIdentifier(0, "a"), ast.NewIdentifier(3, "b"))},
	{"1+-2", ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(2, -2))},
	{"1+-2", ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(2, -2))},
	{"1+-(2)", ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1),
		ast.NewUnaryOperator(2, ast.OperatorSubtraction, ast.NewInt(4, 2)))},
	{"(a)", ast.NewIdentifier(1, "a")},
	{"a()", ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{})},
	{"a(1)", ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{ast.NewInt(2, 1)})},
	{"a(1,2)", ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{ast.NewInt(2, 1), ast.NewInt(4, 2)})},
	{"a[1]", ast.NewIndex(1, ast.NewIdentifier(0, "a"), ast.NewInt(2, 1))},
	{"a[:]", ast.NewSlice(1, ast.NewIdentifier(0, "a"), nil, nil)},
	{"a[:2]", ast.NewSlice(1, ast.NewIdentifier(0, "a"), nil, ast.NewInt(3, 2))},
	{"a[1:]", ast.NewSlice(1, ast.NewIdentifier(0, "a"), ast.NewInt(2, 1), nil)},
	{"a[1:2]", ast.NewSlice(1, ast.NewIdentifier(0, "a"), ast.NewInt(2, 1), ast.NewInt(4, 2))},
	{"a.b", ast.NewSelector(1, ast.NewIdentifier(0, "a"), "b")},
	{"1+2+3", ast.NewBinaryOperator(3, ast.OperatorAddition,
		ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(2, 2)), ast.NewInt(4, 3))},
	{"1-2-3", ast.NewBinaryOperator(3, ast.OperatorSubtraction,
		ast.NewBinaryOperator(1, ast.OperatorSubtraction, ast.NewInt(0, 1), ast.NewInt(2, 2)), ast.NewInt(4, 3))},
	{"1*2*3", ast.NewBinaryOperator(3, ast.OperatorMultiplication,
		ast.NewBinaryOperator(1, ast.OperatorMultiplication, ast.NewInt(0, 1), ast.NewInt(2, 2)), ast.NewInt(4, 3))},
	{"1+2*3", ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1),
		ast.NewBinaryOperator(3, ast.OperatorMultiplication, ast.NewInt(2, 2), ast.NewInt(4, 3)))},
	{"1-2/3", ast.NewBinaryOperator(1, ast.OperatorSubtraction, ast.NewInt(0, 1),
		ast.NewBinaryOperator(3, ast.OperatorDivision, ast.NewInt(2, 2), ast.NewInt(4, 3)))},
	{"1*2+3", ast.NewBinaryOperator(3, ast.OperatorAddition,
		ast.NewBinaryOperator(1, ast.OperatorMultiplication, ast.NewInt(0, 1), ast.NewInt(2, 2)), ast.NewInt(4, 3))},
	{"1==2+3", ast.NewBinaryOperator(1, ast.OperatorEqual, ast.NewInt(0, 1),
		ast.NewBinaryOperator(4, ast.OperatorAddition, ast.NewInt(3, 2), ast.NewInt(5, 3)))},
	{"1+2==3", ast.NewBinaryOperator(3, ast.OperatorEqual,
		ast.NewBinaryOperator(1, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(2, 2)), ast.NewInt(5, 3))},
	{"(1+2)*3", ast.NewBinaryOperator(5, ast.OperatorMultiplication, ast.NewBinaryOperator(2,
		ast.OperatorAddition, ast.NewInt(1, 1), ast.NewInt(3, 2)), ast.NewInt(6, 3))},
	{"1*(2+3)", ast.NewBinaryOperator(1, ast.OperatorMultiplication, ast.NewInt(0, 1),
		ast.NewBinaryOperator(4, ast.OperatorAddition, ast.NewInt(3, 2), ast.NewInt(5, 3)))},
	{"(1*((2)+3))", ast.NewBinaryOperator(2, ast.OperatorMultiplication, ast.NewInt(1, 1),
		ast.NewBinaryOperator(7, ast.OperatorAddition, ast.NewInt(5, 2), ast.NewInt(8, 3)))},
	{"a()*1", ast.NewBinaryOperator(3, ast.OperatorMultiplication,
		ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{}), ast.NewInt(4, 1))},
	{"1*a()", ast.NewBinaryOperator(1, ast.OperatorMultiplication,
		ast.NewInt(0, 1), ast.NewCall(3, ast.NewIdentifier(2, "a"), []ast.Expression{}))},
	{"a[1]*2", ast.NewBinaryOperator(4, ast.OperatorMultiplication,
		ast.NewIndex(1, ast.NewIdentifier(0, "a"), ast.NewInt(2, 1)), ast.NewInt(5, 2))},
	{"1*a[2]", ast.NewBinaryOperator(1, ast.OperatorMultiplication,
		ast.NewInt(0, 1), ast.NewIndex(3, ast.NewIdentifier(2, "a"), ast.NewInt(4, 2)))},
	{"a[1+2]", ast.NewIndex(1, ast.NewIdentifier(0, "a"),
		ast.NewBinaryOperator(3, ast.OperatorAddition, ast.NewInt(2, 1), ast.NewInt(4, 2)))},
	{"a[b(1)]", ast.NewIndex(1, ast.NewIdentifier(0, "a"), ast.NewCall(3,
		ast.NewIdentifier(2, "b"), []ast.Expression{ast.NewInt(4, 1)}))},
	{"a(b[1])", ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{
		ast.NewIndex(3, ast.NewIdentifier(2, "b"), ast.NewInt(4, 1))})},
	{"a.b*c", ast.NewBinaryOperator(3, ast.OperatorMultiplication, ast.NewSelector(1, ast.NewIdentifier(0, "a"), "b"),
		ast.NewIdentifier(4, "c"))},
	{"a*b.c", ast.NewBinaryOperator(1, ast.OperatorMultiplication, ast.NewIdentifier(0, "a"),
		ast.NewSelector(3, ast.NewIdentifier(2, "b"), "c"))},
	{"a.b(c)", ast.NewCall(3, ast.NewSelector(1, ast.NewIdentifier(0, "a"), "b"), []ast.Expression{ast.NewIdentifier(4, "c")})},
	{"1\t+\n2", ast.NewBinaryOperator(2, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(4, 2))},
	{"1\t\r +\n\r\n\r\t 2", ast.NewBinaryOperator(4, ast.OperatorAddition, ast.NewInt(0, 1), ast.NewInt(11, 2))},
	{"a(\n\t1\t,\n2\t)", ast.NewCall(1, ast.NewIdentifier(0, "a"), []ast.Expression{ast.NewInt(4, 1), ast.NewInt(8, 2)})},
	{"a\t\r ()", ast.NewCall(4, ast.NewIdentifier(0, "a"), []ast.Expression{})},
	{"a[\n\t1\t]", ast.NewIndex(1, ast.NewIdentifier(0, "a"), ast.NewInt(4, 1))},
	{"a\t\r [1]", ast.NewIndex(4, ast.NewIdentifier(0, "a"), ast.NewInt(5, 1))},
}

var treeTests = []struct {
	src  string
	node ast.Node
}{
	{"", ast.NewTree(nil)},
	{"a", ast.NewTree([]ast.Node{ast.NewText(0, "a")})},
	{"{{a}}", ast.NewTree([]ast.Node{ast.NewShow(0, ast.NewIdentifier(2, "a"), nil, ast.ContextHTML)})},
	{"a{{b}}", ast.NewTree([]ast.Node{
		ast.NewText(0, "a"), ast.NewShow(1, ast.NewIdentifier(3, "b"), nil, ast.ContextHTML)})},
	{"{{a}}b", ast.NewTree([]ast.Node{
		ast.NewShow(0, ast.NewIdentifier(2, "a"), nil, ast.ContextHTML), ast.NewText(5, "b")})},
	{"a{{b}}c", ast.NewTree([]ast.Node{
		ast.NewText(0, "a"), ast.NewShow(1, ast.NewIdentifier(3, "b"), nil, ast.ContextHTML), ast.NewText(6, "c")})},
	{"{% var a = 1 %}", ast.NewTree([]ast.Node{
		ast.NewVar(0, ast.NewIdentifier(8, "a"), ast.NewInt(12, 1))})},
	{"{% a = 2 %}", ast.NewTree([]ast.Node{
		ast.NewAssignment(0, ast.NewIdentifier(4, "a"), ast.NewInt(8, 2))})},
	{"{% show a %}{% end %}", ast.NewTree([]ast.Node{
		ast.NewShow(0, ast.NewIdentifier(8, "a"), nil, ast.ContextHTML)})},
	{"{% show a %}b{% end %}", ast.NewTree([]ast.Node{
		ast.NewShow(0, ast.NewIdentifier(8, "a"), ast.NewText(12, "b"), ast.ContextHTML)})},
	{"{% for v in e %}b{% end %}", ast.NewTree([]ast.Node{ast.NewFor(0,
		nil, ast.NewIdentifier(7, "v"), ast.NewIdentifier(12, "e"), []ast.Node{ast.NewText(16, "b")})})},
	{"{% for i, v in e %}b{% end %}", ast.NewTree([]ast.Node{ast.NewFor(0,
		ast.NewIdentifier(7, "i"), ast.NewIdentifier(10, "v"), ast.NewIdentifier(15, "e"), []ast.Node{ast.NewText(16, "b")})})},
	{"{% if a %}b{% end %}", ast.NewTree([]ast.Node{
		ast.NewIf(0, ast.NewIdentifier(6, "a"), []ast.Node{ast.NewText(10, "b")})})},
	{"{% extend \"/a.b\" %}", ast.NewTree([]ast.Node{ast.NewExtend(0, "/a.b", nil)})},
	{"{% include \"/a.b\" %}", ast.NewTree([]ast.Node{ast.NewInclude(0, "/a.b", nil)})},
	{"{% region \"a\" %}b{% end %}", ast.NewTree([]ast.Node{
		ast.NewRegion(0, "a", []ast.Node{ast.NewText(16, "b")})})},
}

var pageTests = map[string]struct {
	src  string
	tree *ast.Tree
}{
	"/simple.html": {
		"<!DOCTYPE html>\n<html>\n<head><title>{{ title }}</title></head>\n<body>{{ content }}</body>\n</html>",
		ast.NewTree([]ast.Node{
			ast.NewText(0, "<!DOCTYPE html>\n<html>\n<head><title>"),
			ast.NewShow(36, ast.NewIdentifier(39, "title"), nil, ast.ContextHTML),
			ast.NewText(47, "</title></head>\n<body>"),
			ast.NewShow(69, ast.NewIdentifier(72, "content"), nil, ast.ContextHTML),
			ast.NewText(82, "</body>\n</html>"),
		}),
	},
	"/simple2.html": {
		"<!DOCTYPE html>\n<html>\n<body>{% include \"/include2.html\" %}</body>\n</html>",
		ast.NewTree([]ast.Node{
			ast.NewText(0, "<!DOCTYPE html>\n<html>\n<body>"),
			ast.NewInclude(29, "/include2.html", ast.NewTree([]ast.Node{
				ast.NewText(0, "<div>"),
				ast.NewShow(5, ast.NewIdentifier(8, "content"), nil, ast.ContextHTML),
				ast.NewText(18, "</div>"),
			})),
			ast.NewText(59, "</body>\n</html>"),
		}),
	},
	"/include2.inc": {
		"<div>{{ content }}</div>",
		nil,
	},
}

func TestExpressions(t *testing.T) {
	for _, expr := range exprTests {
		var lex = newLexer([]byte("{{" + expr.src + "}}"))
		_ = <-lex.tokens
		node, tok, err := parseExpr(lex)
		if err != nil {
			t.Errorf("source: %q, %s\n", expr.src, err)
			continue
		}
		if node == nil {
			t.Errorf("source: %q, unexpected %s, expecting expression\n", expr.src, tok)
			continue
		}
		err = equals(node, expr.node, 2)
		if err != nil {
			t.Errorf("source: %q, %s\n", expr.src, err)
		}
	}
}

func TestTrees(t *testing.T) {
	for _, tree := range treeTests {
		node, err := Parse([]byte(tree.src))
		if err != nil {
			t.Errorf("source: %q, %s\n", tree.src, err)
			continue
		}
		err = equals(node, tree.node, 0)
		if err != nil {
			t.Errorf("source: %q, %s\n", tree.src, err)
		}
	}
}

func readFunc(path string) (*ast.Tree, error) {
	return Parse([]byte(pageTests[path].src))
}

func TestPages(t *testing.T) {
	// simple.html
	var parser = NewParser(readFunc)
	var p = pageTests["/simple.html"]
	var tree, err = parser.Parse("/simple.html")
	if err != nil {
		t.Errorf("source: %q, %s\n", p.src, err)
	}
	err = equals(tree, p.tree, 0)
	if err != nil {
		t.Errorf("source: %q, %s\n", p.src, err)
	}
	// simple2.html
	p = pageTests["/simple2.html"]
	tree, err = parser.Parse("/simple2.html")
	if err != nil {
		t.Errorf("source: %q, %s\n", p.src, err)
	}
	err = equals(tree, p.tree, 0)
	if err != nil {
		t.Errorf("source: %q, %s\n", p.src, err)
	}
}

func equals(n1, n2 ast.Node, p int) error {
	if n1 == nil && n2 == nil {
		return nil
	}
	if (n1 == nil) != (n2 == nil) {
		if n1 == nil {
			return fmt.Errorf("unexpected node nil, expecting %#v", n2)
		} else {
			return fmt.Errorf("unexpected node %#v, expecting nil", n1)
		}
	}
	if n1.Pos()-p != n2.Pos() {
		return fmt.Errorf("unexpected position %d, expecting %d", n1.Pos()-p, n2.Pos())
	}
	switch nn1 := n1.(type) {
	case *ast.Tree:
		nn2, ok := n2.(*ast.Tree)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		if len(nn1.Nodes) != len(nn2.Nodes) {
			return fmt.Errorf("unexpected nodes len %d, expecting %d", len(nn1.Nodes), len(nn2.Nodes))
		}
		for i, node := range nn1.Nodes {
			err := equals(node, nn2.Nodes[i], p)
			if err != nil {
				return err
			}
		}
	case *ast.Identifier:
		nn2, ok := n2.(*ast.Identifier)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", n1, n2)
		}
		if nn1.Name != nn2.Name {
			return fmt.Errorf("unexpected %q, expecting %q", nn1.Name, nn2.Name)
		}
	case *ast.Int:
		nn2, ok := n2.(*ast.Int)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", n1, n2)
		}
		if nn1.Value != nn2.Value {
			return fmt.Errorf("unexpected %d, expecting %d", nn1.Value, nn2.Value)
		}
	case *ast.String:
		nn2, ok := n2.(*ast.String)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		if nn1.Text != nn2.Text {
			return fmt.Errorf("unexpected %q, expecting %q", nn1.Text, nn2.Text)
		}
	case *ast.Parentesis:
		nn2, ok := n2.(*ast.Parentesis)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		err := equals(nn1.Expr, nn2.Expr, p)
		if err != nil {
			return err
		}
	case *ast.UnaryOperator:
		nn2, ok := n2.(*ast.UnaryOperator)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		if nn1.Op != nn2.Op {
			return fmt.Errorf("unexpected operator %d, expecting %d", nn1.Op, nn2.Op)
		}
		err := equals(nn1.Expr, nn2.Expr, p)
		if err != nil {
			return err
		}
	case *ast.BinaryOperator:
		nn2, ok := n2.(*ast.BinaryOperator)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		if nn1.Op != nn2.Op {
			return fmt.Errorf("unexpected operator %d, expecting %d", nn1.Op, nn2.Op)
		}
		err := equals(nn1.Expr1, nn2.Expr1, p)
		if err != nil {
			return err
		}
		err = equals(nn1.Expr2, nn2.Expr2, p)
		if err != nil {
			return err
		}
	case *ast.Call:
		nn2, ok := n2.(*ast.Call)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		err := equals(nn1.Func, nn2.Func, p)
		if err != nil {
			return err
		}
		if len(nn1.Args) != len(nn2.Args) {
			return fmt.Errorf("unexpected arguments len %d, expecting %d", len(nn1.Args), len(nn2.Args))
		}
		for i, arg := range nn1.Args {
			err = equals(arg, nn2.Args[i], p)
			if err != nil {
				return err
			}
		}
	case *ast.Index:
		nn2, ok := n2.(*ast.Index)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		err := equals(nn1.Expr, nn2.Expr, p)
		if err != nil {
			return err
		}
		err = equals(nn1.Index, nn2.Index, p)
		if err != nil {
			return err
		}
	case *ast.Slice:
		nn2, ok := n2.(*ast.Slice)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		err := equals(nn1.Expr, nn2.Expr, p)
		if err != nil {
			return err
		}
		err = equals(nn1.Low, nn2.Low, p)
		if err != nil {
			return err
		}
		err = equals(nn1.High, nn2.High, p)
		if err != nil {
			return err
		}
	case *ast.Show:
		nn2, ok := n2.(*ast.Show)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		err := equals(nn1.Expr, nn2.Expr, p)
		if err != nil {
			return err
		}
		if nn1.Text == nil && nn2.Text != nil {
			return fmt.Errorf("unexpected nil, expecting %#v", nn2)
		}
		if nn1.Text != nil && nn2.Text == nil {
			return fmt.Errorf("unexpected %#v, expecting nil", nn1)
		}
		if nn1.Text != nil && nn2.Text != nil {
			err = equals(nn1.Text, nn2.Text, p)
			if err != nil {
				return err
			}
		}
		if nn1.Context != nn2.Context {
			return fmt.Errorf("unexpected context %d, expecting %d", nn1.Context, nn2.Context)
		}
	case *ast.Region:
		nn2, ok := n2.(*ast.Region)
		if !ok {
			return fmt.Errorf("unexpected %#v, expecting %#v", nn1, nn2)
		}
		if nn1.Name != nn2.Name {
			return fmt.Errorf("unexpected %q, expecting %q", nn1.Name, nn2.Name)
		}
		if len(nn1.Nodes) != len(nn2.Nodes) {
			return fmt.Errorf("unexpected nodes len %d, expecting %d", len(nn1.Nodes), len(nn2.Nodes))
		}
		for i, node := range nn1.Nodes {
			err := equals(node, nn2.Nodes[i], p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
