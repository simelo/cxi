package actions

import (
	. "github.com/skycoin/cx/cx"
)

func SetCorrectArithmeticOp(expr *CXExpression) {
	if expr.Operator == nil || len(expr.Outputs) < 1 {
		return
	}
	op := expr.Operator
	typ := expr.Outputs[0].Type

	if CheckArithmeticOp(expr) {
		// if !CheckSameNativeType(expr) {
		//      panic("wrong types")
		// }
		switch op.OpCode {
		case OP_I32_MUL:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_MUL]
			case TypeF32:
				expr.Operator = Natives[OP_F32_MUL]
			case TypeF64:
				expr.Operator = Natives[OP_F64_MUL]
			}
		case OP_I32_DIV:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_DIV]
			case TypeF32:
				expr.Operator = Natives[OP_F32_DIV]
			case TypeF64:
				expr.Operator = Natives[OP_F64_DIV]
			}
		case OP_I32_MOD:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_MOD]
			}

		case OP_I32_ADD:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_ADD]
			case TypeF32:
				expr.Operator = Natives[OP_F32_ADD]
			case TypeF64:
				expr.Operator = Natives[OP_F64_ADD]
			}
		case OP_I32_SUB:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_ADD]
			case TypeF32:
				expr.Operator = Natives[OP_F32_ADD]
			case TypeF64:
				expr.Operator = Natives[OP_F64_ADD]
			}

		case OP_I32_BITSHL:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_BITSHL]
			}
		case OP_I32_BITSHR:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_BITSHR]
			}

		case OP_I32_LT:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_LT]
			case TypeF32:
				expr.Operator = Natives[OP_F32_LT]
			case TypeF64:
				expr.Operator = Natives[OP_F64_LT]
			}
		case OP_I32_GT:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_GT]
			case TypeF32:
				expr.Operator = Natives[OP_F32_GT]
			case TypeF64:
				expr.Operator = Natives[OP_F64_GT]
			}
		case OP_I32_LTEQ:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_LTEQ]
			case TypeF32:
				expr.Operator = Natives[OP_F32_LTEQ]
			case TypeF64:
				expr.Operator = Natives[OP_F64_LTEQ]
			}
		case OP_I32_GTEQ:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_GTEQ]
			case TypeF32:
				expr.Operator = Natives[OP_F32_GTEQ]
			case TypeF64:
				expr.Operator = Natives[OP_F64_GTEQ]
			}

		case OP_I32_EQ:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_EQ]
			case TypeF32:
				expr.Operator = Natives[OP_F32_EQ]
			case TypeF64:
				expr.Operator = Natives[OP_F64_EQ]
			}
		case OP_I32_UNEQ:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_UNEQ]
			case TypeF32:
				expr.Operator = Natives[OP_F32_UNEQ]
			case TypeF64:
				expr.Operator = Natives[OP_F64_UNEQ]
			}

		case OP_I32_BITAND:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_BITAND]
			}

		case OP_I32_BITXOR:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_BITXOR]
			}

		case OP_I32_BITOR:
			switch typ {
			case TypeI32:
			case TypeI64:
				expr.Operator = Natives[OP_I64_BITOR]
			}
		}
	}
}

// This function writes those bytes to PRGRM.Data
func WritePrimary(typ int, byts []byte, isGlobal bool) []*CXExpression {
	if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
		arg := MakeArgument("", CurrentFile, LineNo)
		arg.AddType(TypeNames[typ])
		arg.Package = pkg
		// arg.Program = PRGRM

		var size int

		size = len(byts)

		arg.Size = GetArgSize(typ)
		arg.TotalSize = size
		arg.Offset = DataOffset

		if arg.Type == TypeStr || arg.Type == TypeAff {
			arg.PassBy = PassbyReference
			arg.Size = TypePointerSize
			arg.TotalSize = TypePointerSize
		}

		for i, byt := range byts {
			PRGRM.Memory[DataOffset+i] = byt
		}
		DataOffset += size

		expr := MakeExpression(nil, CurrentFile, LineNo)
		expr.Package = pkg
		expr.Outputs = append(expr.Outputs, arg)
		return []*CXExpression{expr}
	} else {
		panic(err)
	}
}

func CompilationError(currentFile string, lineNo int) string {
	FoundCompileErrors = true
	return ErrorHeader(currentFile, lineNo)
}

func TotalLength(lengths []int) int {
	var total = int(1)
	for _, i := range lengths {
		total *= i
	}
	return total
}

func StructLiteralFields(ident string) *CXExpression {
	if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
		arg := MakeArgument("", CurrentFile, LineNo)
		arg.AddType(TypeNames[TypeIdentifier])
		arg.Name = ident
		arg.Package = pkg

		expr := MakeExpression(nil, CurrentFile, LineNo)
		expr.Outputs = []*CXArgument{arg}
		expr.Package = pkg

		return expr
	} else {
		panic(err)
	}
}

func AffordanceStructs(pkg *CXPackage, currentFile string, lineNo int) {
	// Argument type
	argStrct := MakeStruct("Argument")
	// argStrct.Size = GetArgSize(TypeStr) + GetArgSize(TypeStr)

	argFldName := MakeField("Name", TypeStr, "", 0)
	argFldName.TotalSize = GetArgSize(TypeStr)
	argFldIndex := MakeField("Index", TypeI32, "", 0)
	argFldIndex.TotalSize = GetArgSize(TypeI32)
	argFldType := MakeField("Type", TypeStr, "", 0)
	argFldType.TotalSize = GetArgSize(TypeStr)

	argStrct.AddField(argFldName)
	argStrct.AddField(argFldIndex)
	argStrct.AddField(argFldType)

	pkg.AddStruct(argStrct)

	// Expression type
	exprStrct := MakeStruct("Expression")
	// exprStrct.Size = GetArgSize(TypeStr)

	exprFldOperator := MakeField("Operator", TypeStr, "", 0)

	exprStrct.AddField(exprFldOperator)

	pkg.AddStruct(exprStrct)

	// Function type
	fnStrct := MakeStruct("Function")
	// fnStrct.Size = GetArgSize(TypeStr) + GetArgSize(TypeStr) + GetArgSize(TypeStr)

	fnFldName := MakeField("Name", TypeStr, "", 0)
	fnFldName.TotalSize = GetArgSize(TypeStr)

	fnFldInpSig := MakeField("InputSignature", TypeStr, "", 0)
	fnFldInpSig.Size = GetArgSize(TypeStr)
	fnFldInpSig = DeclarationSpecifiers(fnFldInpSig, 0, DeclSlice)

	fnFldOutSig := MakeField("OutputSignature", TypeStr, "", 0)
	fnFldOutSig.Size = GetArgSize(TypeStr)
	fnFldOutSig = DeclarationSpecifiers(fnFldOutSig, 0, DeclSlice)

	fnStrct.AddField(fnFldName)
	fnStrct.AddField(fnFldInpSig)

	fnStrct.AddField(fnFldOutSig)

	pkg.AddStruct(fnStrct)

	// Structure type
	strctStrct := MakeStruct("Structure")
	// strctStrct.Size = GetArgSize(TypeStr)

	strctFldName := MakeField("Name", TypeStr, "", 0)
	strctFldName.TotalSize = GetArgSize(TypeStr)

	strctStrct.AddField(strctFldName)

	pkg.AddStruct(strctStrct)

	// Package type
	pkgStrct := MakeStruct("Structure")
	// pkgStrct.Size = GetArgSize(TypeStr)

	pkgFldName := MakeField("Name", TypeStr, "", 0)

	pkgStrct.AddField(pkgFldName)

	pkg.AddStruct(pkgStrct)

	// Caller type
	callStrct := MakeStruct("Caller")
	// callStrct.Size = GetArgSize(TypeStr) + GetArgSize(TypeI32)

	callFldFnName := MakeField("FnName", TypeStr, "", 0)
	callFldFnName.TotalSize = GetArgSize(TypeStr)
	callFldFnSize := MakeField("FnSize", TypeI32, "", 0)
	callFldFnSize.TotalSize = GetArgSize(TypeI32)

	callStrct.AddField(callFldFnName)
	callStrct.AddField(callFldFnSize)

	pkg.AddStruct(callStrct)

	// Program type
	prgrmStrct := MakeStruct("Program")
	// prgrmStrct.Size = GetArgSize(TypeI32) + GetArgSize(TypeI64)

	prgrmFldCallCounter := MakeField("CallCounter", TypeI32, "", 0)
	prgrmFldCallCounter.TotalSize = GetArgSize(TypeI32)
	prgrmFldFreeHeap := MakeField("HeapUsed", TypeI64, "", 0)
	prgrmFldFreeHeap.TotalSize = GetArgSize(TypeI64)

	// prgrmFldCaller := MakeField("Caller", TypeCUSTOM, "", 0)
	prgrmFldCaller := DeclarationSpecifiersStruct(callStrct.Name, callStrct.Package.Name, false, currentFile, lineNo)
	prgrmFldCaller.Name = "Caller"

	prgrmStrct.AddField(prgrmFldCallCounter)
	prgrmStrct.AddField(prgrmFldFreeHeap)
	prgrmStrct.AddField(prgrmFldCaller)

	pkg.AddStruct(prgrmStrct)
}

func PrimaryIdentifier(ident string) []*CXExpression {
	pkg, err := PRGRM.GetCurrentPackage()
	if err != nil {
		panic(err)
	}
	arg := MakeArgument(ident, CurrentFile, LineNo)
	arg.AddType(TypeNames[TypeIdentifier])
	// arg.Typ = "ident"
	arg.Name = ident
	arg.Package = pkg

	// expr := &CXExpression{Outputs: []*CXArgument{arg}}
	expr := MakeExpression(nil, CurrentFile, LineNo)
	expr.Outputs = []*CXArgument{arg}
	expr.Package = pkg

	return []*CXExpression{expr}
}
