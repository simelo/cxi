package base

import (
	"fmt"
	"strconv"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func opByteByte(expr *CXExpression, fp int) {
	inp1, out1 := expr.Inputs[0], expr.Outputs[0]
	out1Offset := GetFinalOffset(fp, out1)
	switch out1.Type {
	case TypeByte:
		WriteMemory(out1Offset, FromByte(ReadByte(fp, inp1)))
	case TypeStr:
		WriteObject(out1Offset, encoder.Serialize(strconv.Itoa(int(ReadByte(fp, inp1)))))
	case TypeI32:
		WriteMemory(out1Offset, FromI32(int32(ReadByte(fp, inp1))))
	case TypeI64:
		WriteMemory(out1Offset, FromI64(int64(ReadByte(fp, inp1))))
	case TypeF32:
		WriteMemory(out1Offset, FromF32(float32(ReadByte(fp, inp1))))
	case TypeF64:
		WriteMemory(out1Offset, FromF64(float64(ReadByte(fp, inp1))))
	}
}

func opBytePrint(expr *CXExpression, fp int) {
	inp1 := expr.Inputs[0]
	fmt.Println(ReadByte(fp, inp1))
}
