package a2l

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyFloatInputType struct {
	string string
}

func (d dummyFloatInputType) GetParent() antlr.Tree {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetParent(tree antlr.Tree) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetPayload() interface{} {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetChild(i int) antlr.Tree {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetChildCount() int {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetChildren() []antlr.Tree {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetSourceInterval() *antlr.Interval {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) Accept(Visitor antlr.ParseTreeVisitor) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetText() string {
	return d.string
}

func (d dummyFloatInputType) ToStringTree(strings []string, recognizer antlr.Recognizer) string {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetRuleContext() antlr.RuleContext {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetBaseRuleContext() *antlr.BaseRuleContext {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetInvokingState() int {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetInvokingState(i int) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetRuleIndex() int {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetAltNumber() int {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetAltNumber(altNumber int) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) String(strings []string, context antlr.RuleContext) string {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetException(exception antlr.RecognitionException) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) AddTokenNode(token antlr.Token) *antlr.TerminalNodeImpl {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) AddErrorNode(badToken antlr.Token) *antlr.ErrorNodeImpl {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) EnterRule(listener antlr.ParseTreeListener) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) ExitRule(listener antlr.ParseTreeListener) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetStart(token antlr.Token) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetStart() antlr.Token {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetStop(token antlr.Token) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetStop() antlr.Token {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) AddChild(child antlr.RuleContext) antlr.RuleContext {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) RemoveLastChild() {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetParser() antlr.Parser {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetF() antlr.Token {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) GetI() antlr.Token {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetF(token antlr.Token) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) SetI(token antlr.Token) {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) FLOAT() antlr.TerminalNode {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) INT() antlr.TerminalNode {
	//TODO implement me
	panic("implement me")
}

func (d dummyFloatInputType) IsNumericValueContext() {
	//TODO implement me
	panic("implement me")
}

func Test_A2lFloatToFloatType(t *testing.T) {
	type testCaseType struct {
		literalInput string
		expected     *FloatType
	}

	plus := "+"
	minus := "-"

	for _, tc := range []testCaseType{
		{literalInput: "2.3", expected: &FloatType{Value: 2.3, IntegralSize: 1, DecimalSize: 1}},
		{literalInput: "-2.3", expected: &FloatType{Value: -2.3, IntegralSign: &minus, IntegralSize: 1, DecimalSize: 1}},
		{literalInput: "+2.3", expected: &FloatType{Value: 2.3, IntegralSign: &plus, IntegralSize: 1, DecimalSize: 1}},
		{literalInput: "5.3876e4", expected: &FloatType{Value: 53876, IntegralSize: 1, DecimalSize: 4, ExponentSize: 1}},
		{literalInput: "-3.40282346639e+038", expected: &FloatType{Value: -3.40282346639e+038, IntegralSign: &minus, IntegralSize: 1, DecimalSize: 11, ExponentSign: &plus, ExponentSize: 3}},
		{literalInput: "3.40282346639e+038", expected: &FloatType{Value: 3.40282346639e+038, IntegralSize: 1, DecimalSize: 11, ExponentSign: &plus, ExponentSize: 3}},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.expected, floatToFloatType(dummyFloatInputType{string: tc.literalInput}))
		})
	}
}
