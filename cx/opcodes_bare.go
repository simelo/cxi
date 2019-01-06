package base

// CorePackages ...
var CorePackages = []string{
	// temporary solution until we can implement these packages in pure CX I guess
	"gl", "glfw", "time", "http", "os", "explorer", "aff", "gltext", "cx",
}

// op codes
// nolint golint
const (
	OP_IDENTITY = iota
	OP_JMP
	OP_DEBUG

	OP_SERIALIZE
	OP_DESERIALIZE

	OP_UND_EQUAL
	OP_UND_UNEQUAL
	OP_UND_BITAND
	OP_UND_BITXOR
	OP_UND_BITOR
	OP_UND_BITCLEAR
	OP_UND_MUL
	OP_UND_DIV
	OP_UND_MOD
	OP_UND_ADD
	OP_UND_SUB
	OP_UND_BITSHL
	OP_UND_BITSHR
	OP_UND_LT
	OP_UND_GT
	OP_UND_LTEQ
	OP_UND_GTEQ
	OP_UND_LEN
	OP_UND_PRINTF
	OP_UND_SPRINTF
	OP_UND_READ

	OP_BOOL_PRINT

	OP_BOOL_EQUAL
	OP_BOOL_UNEQUAL
	OP_BOOL_NOT
	OP_BOOL_OR
	OP_BOOL_AND

	OP_BYTE_BYTE
	OP_BYTE_STR
	OP_BYTE_I32
	OP_BYTE_I64
	OP_BYTE_F32
	OP_BYTE_F64

	OP_BYTE_PRINT

	OP_I32_BYTE
	OP_I32_STR
	OP_I32_I32
	OP_I32_I64
	OP_I32_F32
	OP_I32_F64

	OP_I32_PRINT
	OP_I32_ADD
	OP_I32_SUB
	OP_I32_MUL
	OP_I32_DIV
	OP_I32_ABS
	OP_I32_POW
	OP_I32_GT
	OP_I32_GTEQ
	OP_I32_LT
	OP_I32_LTEQ
	OP_I32_EQ
	OP_I32_UNEQ
	OP_I32_MOD
	OP_I32_RAND
	OP_I32_BITAND
	OP_I32_BITOR
	OP_I32_BITXOR
	OP_I32_BITCLEAR
	OP_I32_BITSHL
	OP_I32_BITSHR
	OP_I32_SQRT
	OP_I32_LOG
	OP_I32_LOG2
	OP_I32_LOG10
	OP_I32_MAX
	OP_I32_MIN

	OP_I64_BYTE
	OP_I64_STR
	OP_I64_I32
	OP_I64_I64
	OP_I64_F32
	OP_I64_F64

	OP_I64_PRINT
	OP_I64_ADD
	OP_I64_SUB
	OP_I64_MUL
	OP_I64_DIV
	OP_I64_ABS
	OP_I64_POW
	OP_I64_GT
	OP_I64_GTEQ
	OP_I64_LT
	OP_I64_LTEQ
	OP_I64_EQ
	OP_I64_UNEQ
	OP_I64_MOD
	OP_I64_RAND
	OP_I64_BITAND
	OP_I64_BITOR
	OP_I64_BITXOR
	OP_I64_BITCLEAR
	OP_I64_BITSHL
	OP_I64_BITSHR
	OP_I64_SQRT
	OP_I64_LOG
	OP_I64_LOG10
	OP_I64_LOG2
	OP_I64_MAX
	OP_I64_MIN

	OP_F32_IS_NAN
	OP_F32_BYTE
	OP_F32_STR
	OP_F32_I32
	OP_F32_I64
	OP_F32_F32
	OP_F32_F64

	OP_F32_PRINT
	OP_F32_ADD
	OP_F32_SUB
	OP_F32_MUL
	OP_F32_DIV
	OP_F32_ABS
	OP_F32_POW
	OP_F32_GT
	OP_F32_GTEQ
	OP_F32_LT
	OP_F32_LTEQ
	OP_F32_EQ
	OP_F32_UNEQ
	OP_F32_COS
	OP_F32_SIN
	OP_F32_SQRT
	OP_F32_LOG
	OP_F32_LOG2
	OP_F32_LOG10
	OP_F32_MAX
	OP_F32_MIN

	OP_F64_BYTE
	OP_F64_STR
	OP_F64_I32
	OP_F64_I64
	OP_F64_F32
	OP_F64_F64

	OP_F64_PRINT
	OP_F64_ADD
	OP_F64_SUB
	OP_F64_MUL
	OP_F64_DIV
	OP_F64_ABS
	OP_F64_POW
	OP_F64_GT
	OP_F64_GTEQ
	OP_F64_LT
	OP_F64_LTEQ
	OP_F64_EQ
	OP_F64_UNEQ
	OP_F64_COS
	OP_F64_SIN

	OP_F64_SQRT
	OP_F64_LOG
	OP_F64_LOG2
	OP_F64_LOG10
	OP_F64_MAX
	OP_F64_MIN

	OP_STR_PRINT
	OP_STR_CONCAT
	OP_STR_SUBSTR
	OP_STR_INDEX
	OP_STR_TRIM_SPACE
	OP_STR_EQ

	OP_STR_BYTE
	OP_STR_STR
	OP_STR_I32
	OP_STR_I64
	OP_STR_F32
	OP_STR_F64

	OP_MAKE
	OP_READ
	OP_WRITE
	OP_LEN
	OP_CONCAT
	OP_APPEND
	OP_COPY
	OP_CAST
	OP_EQ
	OP_UNEQ
	OP_RAND
	OP_AND
	OP_OR
	OP_NOT
	OP_SLEEP
	OP_HALT
	OP_GOTO
	OP_REMCX
	OP_ADDCX
	OP_QUERY
	OP_EXECUTE
	OP_INDEX
	OP_NAME
	OP_EVOLVE

	OP_ASSERT
	OP_TEST
	OP_PANIC

	// affordances
	OP_AFF_PRINT
	OP_AFF_QUERY
	OP_AFF_ON
	OP_AFF_OF
	OP_AFF_INFORM
	OP_AFF_REQUEST

	END_OF_BARE_OPS
)

// For the parser. These shouldn't be used in the runtime for performance reasons
var (
	OpNames        = map[int]string{}
	OpCodes        = map[string]int{}
	Natives        = map[int]*CXFunction{}
	execNativeBare func(*CXProgram)
	execNative     func(*CXProgram)
)

// AddOpCode ...
func AddOpCode(code int, name string, inputs []int, outputs []int) {
	OpNames[code] = name
	OpCodes[name] = code
	Natives[code] = MakeNative(code, inputs, outputs)
}

/*
// debug helper
func DumpOpCodes(opCode int) () {
	var keys []int
	for k := range OpNames {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%5d : %s\n", k, OpNames[k])
	}

	fmt.Printf("opCode : %d\n", opCode)
}*/

func init() {
	AddOpCode(OP_IDENTITY, "identity", []int{TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_JMP, "jmp", []int{TypeBool}, []int{})
	AddOpCode(OP_DEBUG, "debug", []int{}, []int{})

	AddOpCode(OP_SERIALIZE, "serialize", []int{TypeAff}, []int{TypeByte})
	AddOpCode(OP_DESERIALIZE, "deserialize", []int{TypeByte}, []int{})

	AddOpCode(OP_UND_EQUAL, "eq", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_UNEQUAL, "uneq", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_BITAND, "bitand", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_BITXOR, "bitxor", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_BITOR, "bitor", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_BITCLEAR, "bitclear", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_MUL, "mul", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_DIV, "div", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_MOD, "mod", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_ADD, "add", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_SUB, "sub", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_BITSHL, "bitshl", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_BITSHR, "bitshr", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_UND_LT, "lt", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_GT, "gt", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_LTEQ, "lteq", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_GTEQ, "gteq", []int{TypeUndefined, TypeUndefined}, []int{TypeBool})
	AddOpCode(OP_UND_LEN, "len", []int{TypeUndefined}, []int{TypeI32})
	AddOpCode(OP_UND_PRINTF, "printf", []int{TypeUndefined}, []int{})
	AddOpCode(OP_UND_SPRINTF, "sprintf", []int{TypeUndefined}, []int{TypeStr})
	AddOpCode(OP_UND_READ, "read", []int{}, []int{TypeStr})

	AddOpCode(OP_BYTE_BYTE, "byte.byte", []int{TypeByte}, []int{TypeByte})
	AddOpCode(OP_BYTE_STR, "byte.str", []int{TypeByte}, []int{TypeStr})
	AddOpCode(OP_BYTE_I32, "byte.i32", []int{TypeByte}, []int{TypeI32})
	AddOpCode(OP_BYTE_I64, "byte.i64", []int{TypeByte}, []int{TypeI64})
	AddOpCode(OP_BYTE_F32, "byte.f32", []int{TypeByte}, []int{TypeF32})
	AddOpCode(OP_BYTE_F64, "byte.f64", []int{TypeByte}, []int{TypeF64})

	AddOpCode(OP_BYTE_PRINT, "byte.print", []int{TypeByte}, []int{})

	AddOpCode(OP_BOOL_PRINT, "bool.print", []int{TypeBool}, []int{})
	AddOpCode(OP_BOOL_EQUAL, "bool.eq", []int{TypeBool, TypeBool}, []int{TypeBool})
	AddOpCode(OP_BOOL_UNEQUAL, "bool.uneq", []int{TypeBool, TypeBool}, []int{TypeBool})
	AddOpCode(OP_BOOL_NOT, "bool.not", []int{TypeBool}, []int{TypeBool})
	AddOpCode(OP_BOOL_OR, "bool.or", []int{TypeBool, TypeBool}, []int{TypeBool})
	AddOpCode(OP_BOOL_AND, "bool.and", []int{TypeBool, TypeBool}, []int{TypeBool})

	AddOpCode(OP_I32_BYTE, "i32.byte", []int{TypeI32}, []int{TypeByte})
	AddOpCode(OP_I32_STR, "i32.str", []int{TypeI32}, []int{TypeStr})
	AddOpCode(OP_I32_I32, "i32.i32", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_I64, "i32.i64", []int{TypeI32}, []int{TypeI64})
	AddOpCode(OP_I32_F32, "i32.f32", []int{TypeI32}, []int{TypeF32})
	AddOpCode(OP_I32_F64, "i32.f64", []int{TypeI32}, []int{TypeF64})

	AddOpCode(OP_I32_PRINT, "i32.print", []int{TypeI32}, []int{})
	AddOpCode(OP_I32_ADD, "i32.add", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_SUB, "i32.sub", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_MUL, "i32.mul", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_DIV, "i32.div", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_ABS, "i32.abs", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_POW, "i32.pow", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_GT, "i32.gt", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_GTEQ, "i32.gteq", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_LT, "i32.lt", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_LTEQ, "i32.lteq", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_EQ, "i32.eq", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_UNEQ, "i32.uneq", []int{TypeI32, TypeI32}, []int{TypeBool})
	AddOpCode(OP_I32_MOD, "i32.mod", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_RAND, "i32.rand", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITAND, "i32.bitand", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITOR, "i32.bitor", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITXOR, "i32.bitxor", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITCLEAR, "i32.bitclear", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITSHL, "i32.bitshl", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_BITSHR, "i32.bitshr", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_SQRT, "i32.sqrt", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_LOG, "i32.log", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_LOG2, "i32.log2", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_LOG10, "i32.log10", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_MAX, "i32.max", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_I32_MIN, "i32.min", []int{TypeI32, TypeI32}, []int{TypeI32})

	AddOpCode(OP_I64_BYTE, "i64.byte", []int{TypeI64}, []int{TypeByte})
	AddOpCode(OP_I64_STR, "i64.str", []int{TypeI64}, []int{TypeStr})
	AddOpCode(OP_I64_I32, "i64.i32", []int{TypeI64}, []int{TypeI32})
	AddOpCode(OP_I64_I64, "i64.i64", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_F32, "i64.f32", []int{TypeI64}, []int{TypeF32})
	AddOpCode(OP_I64_F64, "i64.f64", []int{TypeI64}, []int{TypeF64})

	AddOpCode(OP_I64_PRINT, "i64.print", []int{TypeI64}, []int{})
	AddOpCode(OP_I64_ADD, "i64.add", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_SUB, "i64.sub", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_MUL, "i64.mul", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_DIV, "i64.div", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_ABS, "i64.abs", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_POW, "i64.pow", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_GT, "i64.gt", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_GTEQ, "i64.gteq", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_LT, "i64.lt", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_LTEQ, "i64.lteq", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_EQ, "i64.eq", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_UNEQ, "i64.uneq", []int{TypeI64, TypeI64}, []int{TypeBool})
	AddOpCode(OP_I64_MOD, "i64.mod", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_RAND, "i64.rand", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITAND, "i64.bitand", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITOR, "i64.bitor", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITXOR, "i64.bitxor", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITCLEAR, "i64.bitclear", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITSHL, "i64.bitshl", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_BITSHR, "i64.bitshr", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_SQRT, "i64.sqrt", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_LOG, "i64.log", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_LOG2, "i64.log2", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_LOG10, "i64.log10", []int{TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_MAX, "i64.max", []int{TypeI64, TypeI64}, []int{TypeI64})
	AddOpCode(OP_I64_MIN, "i64.min", []int{TypeI64, TypeI64}, []int{TypeI64})

	AddOpCode(OP_F32_IS_NAN, "f32.isnan", []int{TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_BYTE, "f32.byte", []int{TypeF32}, []int{TypeByte})
	AddOpCode(OP_F32_STR, "f32.str", []int{TypeF32}, []int{TypeStr})
	AddOpCode(OP_F32_I32, "f32.i32", []int{TypeF32}, []int{TypeI32})
	AddOpCode(OP_F32_I64, "f32.i64", []int{TypeF32}, []int{TypeI64})
	AddOpCode(OP_F32_F32, "f32.f32", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_F64, "f32.f64", []int{TypeF32}, []int{TypeF64})
	AddOpCode(OP_F32_PRINT, "f32.print", []int{TypeF32}, []int{})
	AddOpCode(OP_F32_ADD, "f32.add", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_SUB, "f32.sub", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_MUL, "f32.mul", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_DIV, "f32.div", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_ABS, "f32.abs", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_POW, "f32.pow", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_GT, "f32.gt", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_GTEQ, "f32.gteq", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_LT, "f32.lt", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_LTEQ, "f32.lteq", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_EQ, "f32.eq", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_UNEQ, "f32.uneq", []int{TypeF32, TypeF32}, []int{TypeBool})
	AddOpCode(OP_F32_COS, "f32.cos", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_SIN, "f32.sin", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_SQRT, "f32.sqrt", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_LOG, "f32.log", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_LOG2, "f32.log2", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_LOG10, "f32.log10", []int{TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_MAX, "f32.max", []int{TypeF32, TypeF32}, []int{TypeF32})
	AddOpCode(OP_F32_MIN, "f32.min", []int{TypeF32, TypeF32}, []int{TypeF32})

	AddOpCode(OP_F64_BYTE, "f64.byte", []int{TypeF64}, []int{TypeByte})
	AddOpCode(OP_F64_STR, "f64.str", []int{TypeF64}, []int{TypeStr})
	AddOpCode(OP_F64_I32, "f64.i32", []int{TypeF64}, []int{TypeI32})
	AddOpCode(OP_F64_I64, "f64.i64", []int{TypeF64}, []int{TypeI64})
	AddOpCode(OP_F64_F32, "f64.f32", []int{TypeF64}, []int{TypeF32})
	AddOpCode(OP_F64_F64, "f64.f64", []int{TypeF64}, []int{TypeF64})

	AddOpCode(OP_F64_PRINT, "f64.print", []int{TypeF64}, []int{})
	AddOpCode(OP_F64_ADD, "f64.add", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_SUB, "f64.sub", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_MUL, "f64.mul", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_DIV, "f64.div", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_ABS, "f64.abs", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_POW, "f64.pow", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_GT, "f64.gt", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_GTEQ, "f64.gteq", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_LT, "f64.lt", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_LTEQ, "f64.lteq", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_EQ, "f64.eq", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_UNEQ, "f64.uneq", []int{TypeF64, TypeF64}, []int{TypeBool})
	AddOpCode(OP_F64_COS, "f64.cos", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_SIN, "f64.sin", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_SQRT, "f64.sqrt", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_LOG, "f64.log", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_LOG2, "f64.log2", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_LOG10, "f64.log10", []int{TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_MAX, "f64.max", []int{TypeF64, TypeF64}, []int{TypeF64})
	AddOpCode(OP_F64_MIN, "f64.min", []int{TypeF64, TypeF64}, []int{TypeF64})

	AddOpCode(OP_STR_PRINT, "str.print", []int{TypeStr}, []int{})
	AddOpCode(OP_STR_CONCAT, "str.concat", []int{TypeStr, TypeStr}, []int{TypeStr})
	AddOpCode(OP_STR_SUBSTR, "str.substr", []int{TypeStr, TypeI32, TypeI32}, []int{TypeStr})
	AddOpCode(OP_STR_INDEX, "str.index", []int{TypeStr, TypeStr}, []int{TypeI32})
	AddOpCode(OP_STR_TRIM_SPACE, "str.trimspace", []int{TypeStr}, []int{TypeStr})
	AddOpCode(OP_STR_EQ, "str.eq", []int{TypeStr, TypeStr}, []int{TypeBool})

	AddOpCode(OP_STR_BYTE, "str.byte", []int{TypeStr}, []int{TypeByte})
	AddOpCode(OP_STR_STR, "str.str", []int{TypeStr}, []int{TypeStr})
	AddOpCode(OP_STR_I32, "str.i32", []int{TypeStr}, []int{TypeI32})
	AddOpCode(OP_STR_I64, "str.i64", []int{TypeStr}, []int{TypeI64})
	AddOpCode(OP_STR_F32, "str.f32", []int{TypeStr}, []int{TypeF32})
	AddOpCode(OP_STR_F64, "str.f64", []int{TypeStr}, []int{TypeF64})

	AddOpCode(OP_APPEND, "append", []int{TypeUndefined, TypeUndefined}, []int{TypeUndefined})
	AddOpCode(OP_ASSERT, "assert", []int{TypeUndefined, TypeUndefined, TypeStr}, []int{TypeBool})
	AddOpCode(OP_TEST, "test", []int{TypeUndefined, TypeUndefined, TypeStr}, []int{})
	AddOpCode(OP_PANIC, "panic", []int{TypeUndefined, TypeUndefined, TypeStr}, []int{})

	// affordances
	AddOpCode(OP_AFF_PRINT, "aff.print", []int{TypeAff}, []int{})
	AddOpCode(OP_AFF_QUERY, "aff.query", []int{TypeAff}, []int{TypeAff})
	AddOpCode(OP_AFF_ON, "aff.on", []int{TypeAff, TypeAff}, []int{})
	AddOpCode(OP_AFF_OF, "aff.of", []int{TypeAff, TypeAff}, []int{})
	AddOpCode(OP_AFF_INFORM, "aff.inform", []int{TypeAff, TypeI32, TypeAff}, []int{})
	AddOpCode(OP_AFF_REQUEST, "aff.request", []int{TypeAff, TypeI32, TypeAff}, []int{})

	// exec
	execNativeBare = func(prgrm *CXProgram) {
		defer RuntimeError()
		call := &prgrm.CallStack[prgrm.CallCounter]
		expr := call.Operator.Expressions[call.Line]
		opCode := expr.Operator.OpCode
		fp := call.FramePointer

		switch opCode {
		case OP_IDENTITY:
			opIdentity(expr, fp)
		case OP_JMP:
			opJmp(expr, fp, call)
		case OP_DEBUG:
			prgrm.PrintStack()

		case OP_SERIALIZE:
			opSerialize(expr, fp)
		case OP_DESERIALIZE:
			opDeserialize(expr, fp)

		case OP_UND_EQUAL:
			opEqual(expr, fp)
		case OP_UND_UNEQUAL:
			opUnequal(expr, fp)
		case OP_UND_BITAND:
			opBitand(expr, fp)
		case OP_UND_BITXOR:
			opBitxor(expr, fp)
		case OP_UND_BITOR:
			opBitor(expr, fp)
		case OP_UND_BITCLEAR:
			opBitclear(expr, fp)
		case OP_UND_MUL:
			opMul(expr, fp)
		case OP_UND_DIV:
			opDiv(expr, fp)
		case OP_UND_MOD:
			opMod(expr, fp)
		case OP_UND_ADD:
			opAdd(expr, fp)
		case OP_UND_SUB:
			opSub(expr, fp)
		case OP_UND_BITSHL:
			opBitshl(expr, fp)
		case OP_UND_BITSHR:
			opBitshr(expr, fp)
		case OP_UND_LT:
			opLt(expr, fp)
		case OP_UND_GT:
			opGt(expr, fp)
		case OP_UND_LTEQ:
			opLteq(expr, fp)
		case OP_UND_GTEQ:
			opGteq(expr, fp)
		case OP_UND_LEN:
			opLen(expr, fp)
		case OP_UND_PRINTF:
			opPrintf(expr, fp)
		case OP_UND_SPRINTF:
			opSprintf(expr, fp)
		case OP_UND_READ:
			opRead(expr, fp)

		case OP_BYTE_BYTE:
			opByteByte(expr, fp)
		case OP_BYTE_STR:
			opByteByte(expr, fp)
		case OP_BYTE_I32:
			opByteByte(expr, fp)
		case OP_BYTE_I64:
			opByteByte(expr, fp)
		case OP_BYTE_F32:
			opByteByte(expr, fp)
		case OP_BYTE_F64:
			opByteByte(expr, fp)

		case OP_BYTE_PRINT:
			opBytePrint(expr, fp)

		case OP_BOOL_PRINT:
			opBoolPrint(expr, fp)
		case OP_BOOL_EQUAL:
			opBoolEqual(expr, fp)
		case OP_BOOL_UNEQUAL:
			opBoolUnequal(expr, fp)
		case OP_BOOL_NOT:
			opBoolNot(expr, fp)
		case OP_BOOL_OR:
			opBoolOr(expr, fp)
		case OP_BOOL_AND:
			opBoolAnd(expr, fp)

		case OP_I32_BYTE:
			opI32I32(expr, fp)
		case OP_I32_STR:
			opI32I32(expr, fp)
		case OP_I32_I32:
			opI32I32(expr, fp)
		case OP_I32_I64:
			opI32I32(expr, fp)
		case OP_I32_F32:
			opI32I32(expr, fp)
		case OP_I32_F64:
			opI32I32(expr, fp)

		case OP_I32_PRINT:
			opI32Print(expr, fp)
		case OP_I32_ADD:
			opI32Add(expr, fp)
		case OP_I32_SUB:
			opI32Sub(expr, fp)
		case OP_I32_MUL:
			opI32Mul(expr, fp)
		case OP_I32_DIV:
			opI32Div(expr, fp)
		case OP_I32_ABS:
			opI32Abs(expr, fp)
		case OP_I32_POW:
			opI32Pow(expr, fp)
		case OP_I32_GT:
			opI32Gt(expr, fp)
		case OP_I32_GTEQ:
			opI32Gteq(expr, fp)
		case OP_I32_LT:
			opI32Lt(expr, fp)
		case OP_I32_LTEQ:
			opI32Lteq(expr, fp)
		case OP_I32_EQ:
			opI32Eq(expr, fp)
		case OP_I32_UNEQ:
			opI32Uneq(expr, fp)
		case OP_I32_MOD:
			opI32Mod(expr, fp)
		case OP_I32_RAND:
			opI32Rand(expr, fp)
		case OP_I32_BITAND:
			opI32Bitand(expr, fp)
		case OP_I32_BITOR:
			opI32Bitor(expr, fp)
		case OP_I32_BITXOR:
			opI32Bitxor(expr, fp)
		case OP_I32_BITCLEAR:
			opI32Bitclear(expr, fp)
		case OP_I32_BITSHL:
			opI32Bitshl(expr, fp)
		case OP_I32_BITSHR:
			opI32Bitshr(expr, fp)
		case OP_I32_SQRT:
			opI32Sqrt(expr, fp)
		case OP_I32_LOG:
			opI32Log(expr, fp)
		case OP_I32_LOG2:
			opI32Log2(expr, fp)
		case OP_I32_LOG10:
			opI32Log10(expr, fp)

		case OP_I32_MAX:
			opI32Max(expr, fp)
		case OP_I32_MIN:
			opI32Min(expr, fp)

		case OP_I64_BYTE:
			opI64I64(expr, fp)
		case OP_I64_STR:
			opI64I64(expr, fp)
		case OP_I64_I32:
			opI64I64(expr, fp)
		case OP_I64_I64:
			opI64I64(expr, fp)
		case OP_I64_F32:
			opI64I64(expr, fp)
		case OP_I64_F64:
			opI64I64(expr, fp)

		case OP_I64_PRINT:
			opI64Print(expr, fp)
		case OP_I64_ADD:
			opI64Add(expr, fp)
		case OP_I64_SUB:
			opI64Sub(expr, fp)
		case OP_I64_MUL:
			opI64Mul(expr, fp)
		case OP_I64_DIV:
			opI64Div(expr, fp)
		case OP_I64_ABS:
			opI64Abs(expr, fp)
		case OP_I64_POW:
			opI64Pow(expr, fp)
		case OP_I64_GT:
			opI64Gt(expr, fp)
		case OP_I64_GTEQ:
			opI64Gteq(expr, fp)
		case OP_I64_LT:
			opI64Lt(expr, fp)
		case OP_I64_LTEQ:
			opI64Lteq(expr, fp)
		case OP_I64_EQ:
			opI64Eq(expr, fp)
		case OP_I64_UNEQ:
			opI64Uneq(expr, fp)
		case OP_I64_MOD:
			opI64Mod(expr, fp)
		case OP_I64_RAND:
			opI64Rand(expr, fp)
		case OP_I64_BITAND:
			opI64Bitand(expr, fp)
		case OP_I64_BITOR:
			opI64Bitor(expr, fp)
		case OP_I64_BITXOR:
			opI64Bitxor(expr, fp)
		case OP_I64_BITCLEAR:
			opI64Bitclear(expr, fp)
		case OP_I64_BITSHL:
			opI64Bitshl(expr, fp)
		case OP_I64_BITSHR:
			opI64Bitshr(expr, fp)
		case OP_I64_SQRT:
			opI64Sqrt(expr, fp)
		case OP_I64_LOG:
			opI64Log(expr, fp)
		case OP_I64_LOG2:
			opI64Log2(expr, fp)
		case OP_I64_LOG10:
			opI64Log10(expr, fp)
		case OP_I64_MAX:
			opI64Max(expr, fp)
		case OP_I64_MIN:
			opI64Min(expr, fp)

		case OP_F32_IS_NAN:
			opF32Isnan(expr, fp)
		case OP_F32_BYTE:
			opF32F32(expr, fp)
		case OP_F32_STR:
			opF32F32(expr, fp)
		case OP_F32_I32:
			opF32F32(expr, fp)
		case OP_F32_I64:
			opF32F32(expr, fp)
		case OP_F32_F32:
			opF32F32(expr, fp)
		case OP_F32_F64:
			opF32F32(expr, fp)
		case OP_F32_PRINT:
			opF32Print(expr, fp)
		case OP_F32_ADD:
			opF32Add(expr, fp)
		case OP_F32_SUB:
			opF32Sub(expr, fp)
		case OP_F32_MUL:
			opF32Mul(expr, fp)
		case OP_F32_DIV:
			opF32Div(expr, fp)
		case OP_F32_ABS:
			opF32Abs(expr, fp)
		case OP_F32_POW:
			opF32Pow(expr, fp)
		case OP_F32_GT:
			opF32Gt(expr, fp)
		case OP_F32_GTEQ:
			opF32Gteq(expr, fp)
		case OP_F32_LT:
			opF32Lt(expr, fp)
		case OP_F32_LTEQ:
			opF32Lteq(expr, fp)
		case OP_F32_EQ:
			opF32Eq(expr, fp)
		case OP_F32_UNEQ:
			opF32Uneq(expr, fp)
		case OP_F32_COS:
			opF32Cos(expr, fp)
		case OP_F32_SIN:
			opF32Sin(expr, fp)
		case OP_F32_SQRT:
			opF32Sqrt(expr, fp)
		case OP_F32_LOG:
			opF32Log(expr, fp)
		case OP_F32_LOG2:
			opF32Log2(expr, fp)
		case OP_F32_LOG10:
			opF32Log10(expr, fp)
		case OP_F32_MAX:
			opF32Max(expr, fp)
		case OP_F32_MIN:
			opF32Min(expr, fp)

		case OP_F64_BYTE:
			opF64F64(expr, fp)
		case OP_F64_STR:
			opF64F64(expr, fp)
		case OP_F64_I32:
			opF64F64(expr, fp)
		case OP_F64_I64:
			opF64F64(expr, fp)
		case OP_F64_F32:
			opF64F64(expr, fp)
		case OP_F64_F64:
			opF64F64(expr, fp)

		case OP_F64_PRINT:
			opF64Print(expr, fp)
		case OP_F64_ADD:
			opF64Add(expr, fp)
		case OP_F64_SUB:
			opF64Sub(expr, fp)
		case OP_F64_MUL:
			opF64Mul(expr, fp)
		case OP_F64_DIV:
			opF64Div(expr, fp)
		case OP_F64_ABS:
			opF64Abs(expr, fp)
		case OP_F64_POW:
			opF64Pow(expr, fp)
		case OP_F64_GT:
			opF64Gt(expr, fp)
		case OP_F64_GTEQ:
			opF64Gteq(expr, fp)
		case OP_F64_LT:
			opF64Lt(expr, fp)
		case OP_F64_LTEQ:
			opF64Lteq(expr, fp)
		case OP_F64_EQ:
			opF64Eq(expr, fp)
		case OP_F64_UNEQ:
			opF64Uneq(expr, fp)
		case OP_F64_COS:
			opF64Cos(expr, fp)
		case OP_F64_SIN:
			opF64Sin(expr, fp)
		case OP_F64_SQRT:
			opF64Sqrt(expr, fp)
		case OP_F64_LOG:
			opF64Log(expr, fp)
		case OP_F64_LOG2:
			opF64Log2(expr, fp)
		case OP_F64_LOG10:
			opF64Log10(expr, fp)
		case OP_F64_MAX:
			opF64Max(expr, fp)
		case OP_F64_MIN:
			opF64Min(expr, fp)

		case OP_STR_PRINT:
			opStrPrint(expr, fp)
		case OP_STR_EQ:
			opStrEq(expr, fp)
		case OP_STR_CONCAT:
			opStrConcat(expr, fp)
		case OP_STR_SUBSTR:
			opStrSubstr(expr, fp)
		case OP_STR_INDEX:
			opStrIndex(expr, fp)
		case OP_STR_TRIM_SPACE:
			opStrTrimSpace(expr, fp)

		case OP_STR_BYTE:
			opStrStr(expr, fp)
		case OP_STR_STR:
			opStrStr(expr, fp)
		case OP_STR_I32:
			opStrStr(expr, fp)
		case OP_STR_I64:
			opStrStr(expr, fp)
		case OP_STR_F32:
			opStrStr(expr, fp)
		case OP_STR_F64:
			opStrStr(expr, fp)

		case OP_MAKE:
		case OP_READ:
		case OP_WRITE:
		case OP_LEN:
		case OP_CONCAT:
		case OP_APPEND:
			opAppend(expr, fp)
		case OP_COPY:
		case OP_CAST:
		case OP_EQ:
		case OP_UNEQ:
		case OP_AND:
		case OP_OR:
		case OP_NOT:
		case OP_SLEEP:
		case OP_HALT:
		case OP_GOTO:
		case OP_REMCX:
		case OP_ADDCX:
		case OP_QUERY:
		case OP_EXECUTE:
		case OP_INDEX:
		case OP_NAME:
		case OP_EVOLVE:
		case OP_ASSERT:
			opAssertValue(expr, fp)
		case OP_TEST:
			opTest(expr, fp)
		case OP_PANIC:
			opPanic(expr, fp)

		// affordances
		case OP_AFF_PRINT:
			opAffPrint(expr, fp)
		case OP_AFF_QUERY:
			opAffQuery(expr, fp)
		case OP_AFF_ON:
			opAffOn(expr, fp)
		case OP_AFF_OF:
			opAffOf(expr, fp)
		case OP_AFF_INFORM:
			opAffInform(expr, fp)
		case OP_AFF_REQUEST:
			opAffRequest(expr, fp)
		default:
			// DumpOpCodes(opCode) // debug helper
			panic("invalid bare opcode")
		}
	}

	execNative = execNativeBare
}
