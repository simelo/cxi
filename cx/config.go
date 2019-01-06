package base

import (
	"os"
)

// DbgGolangStackTrace ...
const DbgGolangStackTrace = true

// PROGRAM global reference to our program
var PROGRAM *CXProgram

// Var ...
var (
	CXPATH   = os.Getenv("CXPATH") + "/"
	BINPATH  = CXPATH + "bin/"
	PKGPATH  = CXPATH + "pkg/"
	SRCPATH  = CXPATH + "src/"
	COREPATH string
)

// Const ...
const (
	StackOverflowError = "stack overflow"
	HeapExhaustedError = "heap exhausted"
	MainFunc           = "main"
	SysInitFunc        = "*init"
	MainPkg            = "main"
	OsPkg              = "os"
	OsArgs             = "Args"

	NonAssignPrefix = "nonAssign"
	LocalPrefix     = "*lcl"
	LabelPrefix     = "*lbl"

	// const CORE_MODULE = "core"

	IDFn   = "identity"
	InitFn = "initDef"

	I32Size = 4
	I64Size = 8
	StrSize = 4

	MarkSize              = 1
	ObjectHeaderSize      = 9
	ObjectGcHeaderSize    = 5
	ForwardingAddressSize = 4
	ObjectSize            = 4

	CallstackSize    = 1000
	StackSize        = 1048576  // 1 Mb
	InitHeapSize     = 2097152  // 2 Mb
	MaxHeapSize      = 67108864 // 64 Mb
	MinHeapFreeRatio = 40
	MaxHeapFreeRatio = 70

	NullAddress           = StackSize
	NullHeapAddressOffset = 4
	NullHeapAddress       = 0
	StrHeaderSize         = 4
	TypePointerSize       = 4
	SliceHeaderSize       = 8

	MaxUint32 = ^uint32(0)
	MinUint32 = 0
	MaxInt32  = int(MaxUint32 >> 1)
	MinInt32  = -MaxInt32 - 1
)

//MemorySize ...
var MemorySize = StackSize + InitHeapSize + TypePointerSize

//BasicTypes ...
var BasicTypes = []string{
	"bool", "str", "byte", "i32", "i64", "f32", "f64",
	"[]bool", "[]str", "[]byte", "[]i32", "[]i64", "[]f32", "[]f64",
}

// Const CX
const (
	CxSuccess = iota
	CxRuntimeError
	CxPanic // 2
	CxCompilationError
	CxInternalError
	CxAssert
)

// Const ...
const (
	DeclPointer = iota // 0
	DeclArray          // 1
	DeclSlice          // 2
	DeclStruct         // 3
	DeclBasic          // 4
)

// what to write
const (
	PassbyValue = iota
	PassbyReference
)

// Const ...
const (
	DerefArray = iota
	DerefField
	DerefPointer
	DerefDeref
)

// Const TypeData
const (
	TypeUndefined = iota
	TypeAff
	TypeBool
	TypeByte
	TypeStr
	TypeF32
	TypeF64
	TypeI8
	TypeI16
	TypeI32
	TypeI64
	TypeUI8
	TypeUI16
	TypeUI32
	TypeUI64

	TypeThreshold

	TypeCustom
	TypePointer
	TypeIdentifier
)

// TypeDate ...
var (
	TypeCounter int

	TypeCodes = map[string]int{
		"ident": TypeIdentifier,
		"aff":   TypeAff,
		"bool":  TypeBool,
		"byte":  TypeByte,
		"str":   TypeStr,
		"f32":   TypeF32,
		"f64":   TypeF64,
		"i8":    TypeI8,
		"i16":   TypeI16,
		"i32":   TypeI32,
		"i64":   TypeI64,
		"ui8":   TypeUI8,
		"ui16":  TypeUI16,
		"ui32":  TypeUI32,
		"ui64":  TypeUI64,
		"und":   TypeUndefined,
	}

	TypeNames = map[int]string{
		TypeIdentifier: "ident",
		TypeAff:        "aff",
		TypeBool:       "bool",
		TypeByte:       "byte",
		TypeStr:        "str",
		TypeF32:        "f32",
		TypeF64:        "f64",
		TypeI8:         "i8",
		TypeI16:        "i16",
		TypeI32:        "i32",
		TypeI64:        "i64",
		TypeUI8:        "ui8",
		TypeUI16:       "ui16",
		TypeUI32:       "ui32",
		TypeUI64:       "ui64",
		TypeUndefined:  "und",
	}
)

// memory locations
const (
	MemStack = iota
	MemHeap
	MemData
)
