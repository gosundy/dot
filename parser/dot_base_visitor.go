// Code generated from Dot.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dot

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDotVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDotVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStatmentIf(ctx *StatmentIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStatmentExp(ctx *StatmentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitIfStatment(ctx *IfStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitIfPartStatment(ctx *IfPartStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitContentMulDiv(ctx *ContentMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitContentAddSub(ctx *ContentAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitGetContentValue(ctx *GetContentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAndOr(ctx *AndOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitGetInt(ctx *GetIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitGetString(ctx *GetStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitMsgContentValue(ctx *MsgContentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitMsgContentJsonPath(ctx *MsgContentJsonPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitJsonPath(ctx *JsonPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitMsgValueType(ctx *MsgValueTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
