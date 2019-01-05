package base

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func opLt(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) < ReadByte(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) < ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) < ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) < ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) < ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) < ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opGt(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) > ReadByte(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) > ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) > ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) > ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) > ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) > ReadF64(fp, inp2))
	}
	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opLteq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) <= ReadByte(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) <= ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) <= ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) <= ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) <= ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) <= ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opGteq(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) >= ReadByte(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) >= ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) >= ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) >= ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) >= ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) >= ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opEqual(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) == ReadByte(fp, inp2))
	case TypeBool:
		outB1 = FromBool(ReadBool(fp, inp1) == ReadBool(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) == ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) == ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) == ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) == ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) == ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opUnequal(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromBool(ReadByte(fp, inp1) != ReadByte(fp, inp2))
	case TypeBool:
		outB1 = FromBool(ReadBool(fp, inp1) != ReadBool(fp, inp2))
	case TypeStr:
		outB1 = FromBool(ReadStr(fp, inp1) != ReadStr(fp, inp2))
	case TypeI32:
		outB1 = FromBool(ReadI32(fp, inp1) != ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromBool(ReadI64(fp, inp1) != ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromBool(ReadF32(fp, inp1) != ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromBool(ReadF64(fp, inp1) != ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitand(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) & ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) & ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitor(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) | ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) | ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitxor(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) ^ ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) ^ ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opMul(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromByte(ReadByte(fp, inp1) * ReadByte(fp, inp2))
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) * ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) * ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromF32(ReadF32(fp, inp1) * ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromF64(ReadF64(fp, inp1) * ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opDiv(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromByte(ReadByte(fp, inp1) / ReadByte(fp, inp2))
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) / ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) / ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromF32(ReadF32(fp, inp1) / ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromF64(ReadF64(fp, inp1) / ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opMod(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromByte(ReadByte(fp, inp1) % ReadByte(fp, inp2))
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) % ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) % ReadI64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opAdd(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromByte(ReadByte(fp, inp1) + ReadByte(fp, inp2))
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) + ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) + ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromF32(ReadF32(fp, inp1) + ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromF64(ReadF64(fp, inp1) + ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opSub(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeByte:
		outB1 = FromByte(ReadByte(fp, inp1) - ReadByte(fp, inp2))
	case TypeI32:
		outB1 = FromI32(ReadI32(fp, inp1) - ReadI32(fp, inp2))
	case TypeI64:
		outB1 = FromI64(ReadI64(fp, inp1) - ReadI64(fp, inp2))
	case TypeF32:
		outB1 = FromF32(ReadF32(fp, inp1) - ReadF32(fp, inp2))
	case TypeF64:
		outB1 = FromF64(ReadF64(fp, inp1) - ReadF64(fp, inp2))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitshl(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) << uint32(ReadI32(fp, inp2))))
	case TypeI64:
		outB1 = FromI64(int64(uint64(ReadI64(fp, inp1)) << uint64(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitshr(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) >> uint32(ReadI32(fp, inp2))))
	case TypeI64:
		outB1 = FromI64(int64(uint32(ReadI64(fp, inp1)) >> uint32(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opBitclear(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]
	var outB1 []byte
	switch inp1.Type {
	case TypeI32:
		outB1 = FromI32(int32(uint32(ReadI32(fp, inp1)) &^ uint32(ReadI32(fp, inp2))))
	case TypeI64:
		outB1 = FromI64(int64(uint32(ReadI64(fp, inp1)) &^ uint32(ReadI64(fp, inp2))))
	}

	WriteMemory(GetFinalOffset(fp, out1), outB1)
}

func opLen(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	elt := GetAssignmentElement(inp1)

	if elt.IsSlice || elt.Type == TypeAff {
		preInp1Offset := GetFinalOffset(fp, inp1)

		var inp1Offset int32
		encoder.DeserializeAtomic(PROGRAM.Memory[preInp1Offset:preInp1Offset+TypePointerSize], &inp1Offset)

		if inp1Offset == 0 {
			// then it's nil
			WriteMemory(GetFinalOffset(fp, out1), encoder.SerializeAtomic(int32(0)))
			return
		}

		sliceHeader := PROGRAM.Memory[inp1Offset+ObjectHeaderSize : inp1Offset+ObjectHeaderSize+SliceHeaderSize]
		WriteMemory(GetFinalOffset(fp, out1), sliceHeader[:4])
	} else if elt.Type == TypeStr {
		var strOffset = int(GetStrOffset(fp, inp1))
		WriteMemory(GetFinalOffset(fp, out1), PROGRAM.Memory[strOffset:strOffset+StrHeaderSize])
	} else {
		outB1 := FromI32(int32(elt.Lengths[0]))
		WriteMemory(GetFinalOffset(fp, out1), outB1)
	}
}

func opAppend(expr *CXExpression, fp int) {
	inp1, inp2, out1 := expr.Inputs[0], expr.Inputs[1], expr.Outputs[0]

	preInp1Offset := GetFinalOffset(fp, inp1)
	inp2Offset := GetFinalOffset(fp, inp2)
	out1Offset := GetFinalOffset(fp, out1)

	var inp1Offset int32
	encoder.DeserializeAtomic(PROGRAM.Memory[preInp1Offset:preInp1Offset+TypePointerSize], &inp1Offset)

	var off int32

	byts := PROGRAM.Memory[out1Offset : out1Offset+TypePointerSize]
	encoder.DeserializeAtomic(byts, &off)

	var heapOffset int

	if off == 0 {
		// then out1 == nil
		var sliceHeader []byte
		var len1 int32

		if inp1Offset != 0 {
			// then we need to reserve for obj1 too
			// sliceHeader = PROGRAM.Memory[inp1Offset-SLICE_HEADER_SIZE : inp1Offset]
			sliceHeader = PROGRAM.Memory[inp1Offset+ObjectHeaderSize : inp1Offset+ObjectHeaderSize+SliceHeaderSize]

			encoder.DeserializeAtomic(sliceHeader[:4], &len1)
			heapOffset = AllocateSeq((int(len1) * inp2.TotalSize) + inp2.TotalSize + ObjectHeaderSize + SliceHeaderSize)
		} else {
			// then obj1 is nil and zero-sized
			heapOffset = AllocateSeq(inp2.TotalSize + ObjectHeaderSize + SliceHeaderSize)
		}

		WriteMemory(out1Offset, encoder.SerializeAtomic(int32(heapOffset)))

		var obj1 []byte
		var obj2 []byte

		obj1 = PROGRAM.Memory[inp1Offset+ObjectHeaderSize+SliceHeaderSize : ObjectHeaderSize+SliceHeaderSize+int32(inp1Offset)+len1*int32(inp2.TotalSize)]

		if inp2.Type == TypeStr || inp2.Type == TypeAff {
			var strOffset = int32(GetStrOffset(fp, inp2))
			obj2 = encoder.SerializeAtomic(strOffset)
		} else {
			obj2 = ReadMemory(inp2Offset, inp2)
		}

		var size []byte
		if inp1Offset != 0 {
			size = encoder.SerializeAtomic(int32(len(obj1)) + int32(len(obj2)+SliceHeaderSize))
		} else {
			size = encoder.SerializeAtomic(int32(len(obj2) + SliceHeaderSize))
		}

		var header = make([]byte, ObjectHeaderSize)
		for c := 5; c < ObjectHeaderSize; c++ {
			header[c] = size[c-5]
		}

		// length
		lenTotal := encoder.SerializeAtomic(len1 + 1)
		capTotal := lenTotal

		finalObj := append(header, lenTotal...)
		finalObj = append(finalObj, capTotal...)

		if inp1Offset != 0 {
			// then obj1 is not nil, and we need to append
			finalObj = append(finalObj, obj1...)
		}
		finalObj = append(finalObj, obj2...)

		WriteMemory(heapOffset, finalObj)
	} else {
		// then we have access to a size and capacity
		sliceHeader := PROGRAM.Memory[inp1Offset+ObjectHeaderSize : inp1Offset+ObjectHeaderSize+SliceHeaderSize]

		var l int32
		var c int32

		encoder.DeserializeAtomic(sliceHeader[:4], &l)
		encoder.DeserializeAtomic(sliceHeader[4:], &c)

		if l >= c {
			// then we need to increase cap and relocate slice
			var obj1 []byte
			var obj2 []byte

			obj1 = PROGRAM.Memory[inp1Offset+ObjectHeaderSize+SliceHeaderSize : int32(inp1Offset)+ObjectHeaderSize+SliceHeaderSize+l*int32(inp2.TotalSize)]

			if inp2.Type == TypeStr || inp2.Type == TypeAff {
				var strOffset = int32(GetStrOffset(fp, inp2))
				obj2 = encoder.SerializeAtomic(strOffset)
			} else {
				obj2 = ReadMemory(inp2Offset, inp2)
			}

			l++
			c = c * 2

			heapOffset = AllocateSeq(int(c)*inp2.TotalSize + ObjectHeaderSize + SliceHeaderSize)

			WriteMemory(out1Offset, encoder.SerializeAtomic(int32(heapOffset)))

			size := encoder.SerializeAtomic(int32(int(c)*inp2.TotalSize + SliceHeaderSize))

			var header = make([]byte, ObjectHeaderSize)
			for c := 5; c < ObjectHeaderSize; c++ {
				header[c] = size[c-5]
			}

			lB := encoder.SerializeAtomic(l)
			cB := encoder.SerializeAtomic(c)

			finalObj := append(header, lB...)
			finalObj = append(finalObj, cB...)
			finalObj = append(finalObj, obj1...)
			finalObj = append(finalObj, obj2...)

			WriteMemory(heapOffset, finalObj)
		} else {
			// then we can simply write the element

			// updating the length
			newL := encoder.SerializeAtomic(l + int32(1))

			for i, byt := range newL {
				PROGRAM.Memory[int(off)+ObjectHeaderSize+i] = byt
			}

			var obj2 []byte
			if inp2.Type == TypeStr || inp2.Type == TypeAff {
				var strOffset = int32(GetStrOffset(fp, inp2))
				obj2 = encoder.SerializeAtomic(strOffset)
			} else {
				obj2 = ReadMemory(inp2Offset, inp2)
			}

			// write the obj
			for i, byt := range obj2 {
				PROGRAM.Memory[off+ObjectHeaderSize+SliceHeaderSize+int32(int(l)*inp2.TotalSize+i)] = byt
			}
		}
	}
}

func buildString(expr *CXExpression, fp int) []byte {
	inp1 := expr.Inputs[0]

	fmtStr := ReadStr(fp, inp1)

	var res []byte
	var specifiersCounter int
	var lenStr = int(len(fmtStr))

	for c := 0; c < len(fmtStr); c++ {
		var nextCh byte
		ch := fmtStr[c]
		if c < lenStr-1 {
			nextCh = fmtStr[c+1]
		}
		if ch == '\\' {
			switch nextCh {
			case '%':
				c++
				res = append(res, nextCh)
				continue
			case 'n':
				c++
				res = append(res, '\n')
				continue
			default:
				res = append(res, ch)
				continue
			}
		}
		if ch == '%' {
			if specifiersCounter+1 == len(expr.Inputs) {
				res = append(res, []byte(fmt.Sprintf("%%!%c(MISSING)", nextCh))...)
				c++
				continue
			}

			inp := expr.Inputs[specifiersCounter+1]
			switch nextCh {
			case 's':
				res = append(res, []byte(checkForEscapedChars(ReadStr(fp, inp)))...)
			case 'd':
				switch inp.Type {
				case TypeI32:
					res = append(res, []byte(strconv.FormatInt(int64(ReadI32(fp, inp)), 10))...)
				case TypeI64:
					res = append(res, []byte(strconv.FormatInt(ReadI64(fp, inp), 10))...)
				}
			case 'f':
				switch inp.Type {
				case TypeF32:
					res = append(res, []byte(strconv.FormatFloat(float64(ReadF32(fp, inp)), 'f', 7, 32))...)
				case TypeF64:
					res = append(res, []byte(strconv.FormatFloat(ReadF64(fp, inp), 'f', 16, 64))...)
				}
			case 'v':
				res = append(res, []byte(GetPrintableValue(fp, inp))...)
			}
			c++
			specifiersCounter++
		} else {
			res = append(res, ch)
		}
	}

	if specifiersCounter != len(expr.Inputs)-1 {
		extra := "%!(EXTRA "
		// for _, inp := range expr.Inputs[:specifiersCounter] {
		lInps := len(expr.Inputs[specifiersCounter+1:])
		for c := 0; c < lInps; c++ {
			inp := expr.Inputs[specifiersCounter+1+c]
			elt := GetAssignmentElement(inp)
			typ := ""
			_ = typ
			if elt.CustomType != nil {
				// then it's custom type
				typ = elt.CustomType.Name
			} else {
				// then it's native type
				typ = TypeNames[elt.Type]
			}

			if c == lInps-1 {
				extra += fmt.Sprintf("%s=%s", typ, GetPrintableValue(fp, elt))
			} else {
				extra += fmt.Sprintf("%s=%s, ", typ, GetPrintableValue(fp, elt))
			}

		}

		extra += ")"

		res = append(res, []byte(extra)...)
	}

	// if specifiersCounter != len(expr.Inputs) - 1 {
	// 	panic("meow")
	// }

	return res
}

func opSprintf(expr *CXExpression, fp int) {
	out1 := expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	byts := encoder.Serialize(string(buildString(expr, fp)))
	WriteObject(out1Offset, byts)
}

func opPrintf(expr *CXExpression, fp int) {
	fmt.Print(string(buildString(expr, fp)))
}

func opRead(expr *CXExpression, fp int) {
	out1 := expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	// text = strings.Trim(text, " \n")
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	if err != nil {
		panic("")
	}
	byts := encoder.Serialize(text)
	size := encoder.Serialize(int32(len(byts)))
	heapOffset := AllocateSeq(len(byts) + ObjectHeaderSize)

	var header = make([]byte, ObjectHeaderSize)
	for c := 5; c < ObjectHeaderSize; c++ {
		header[c] = size[c-5]
	}

	obj := append(header, byts...)

	WriteMemory(heapOffset, obj)

	off := encoder.SerializeAtomic(int32(heapOffset + ObjectHeaderSize))

	WriteMemory(out1Offset, off)
}
