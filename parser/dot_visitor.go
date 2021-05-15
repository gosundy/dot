// Code generated from Dot.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dot

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DotParser.
type DotVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DotParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by DotParser#StatmentIf.
	VisitStatmentIf(ctx *StatmentIfContext) interface{}

	// Visit a parse tree produced by DotParser#StatmentExp.
	VisitStatmentExp(ctx *StatmentExpContext) interface{}

	// Visit a parse tree produced by DotParser#ifStatment.
	VisitIfStatment(ctx *IfStatmentContext) interface{}

	// Visit a parse tree produced by DotParser#ifPartStatment.
	VisitIfPartStatment(ctx *IfPartStatmentContext) interface{}

	// Visit a parse tree produced by DotParser#ContentMulDiv.
	VisitContentMulDiv(ctx *ContentMulDivContext) interface{}

	// Visit a parse tree produced by DotParser#ContentAddSub.
	VisitContentAddSub(ctx *ContentAddSubContext) interface{}

	// Visit a parse tree produced by DotParser#GetContentValue.
	VisitGetContentValue(ctx *GetContentValueContext) interface{}

	// Visit a parse tree produced by DotParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by DotParser#Compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by DotParser#AndOr.
	VisitAndOr(ctx *AndOrContext) interface{}

	// Visit a parse tree produced by DotParser#GetInt.
	VisitGetInt(ctx *GetIntContext) interface{}

	// Visit a parse tree produced by DotParser#GetString.
	VisitGetString(ctx *GetStringContext) interface{}

	// Visit a parse tree produced by DotParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by DotParser#msgContentValue.
	VisitMsgContentValue(ctx *MsgContentValueContext) interface{}

	// Visit a parse tree produced by DotParser#msgContentJsonPath.
	VisitMsgContentJsonPath(ctx *MsgContentJsonPathContext) interface{}

	// Visit a parse tree produced by DotParser#jsonPath.
	VisitJsonPath(ctx *JsonPathContext) interface{}

	// Visit a parse tree produced by DotParser#msgValueType.
	VisitMsgValueType(ctx *MsgValueTypeContext) interface{}
}
