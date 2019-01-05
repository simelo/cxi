package base

import (
	"fmt"

	"github.com/skycoin/skycoin/src/cipher/encoder"
)

// CalculateDereferences ...
func CalculateDereferences(arg *CXArgument, finalOffset *int, fp int, dbg bool) {
	var isPointer bool
	for _, op := range arg.DereferenceOperations {
		switch op {
		case DerefArray:
			for i, idxArg := range arg.Indexes {
				var subSize = int(1)
				for _, len := range arg.Lengths[i+1:] {
					subSize *= len
				}

				var sizeToUse int
				if arg.CustomType != nil {
					sizeToUse = arg.CustomType.Size
				} else if arg.IsSlice {
					sizeToUse = arg.TotalSize
				} else {
					sizeToUse = arg.Size
				}

				*finalOffset += int(ReadI32(fp, idxArg)) * subSize * sizeToUse
			}
		case DerefPointer:
			isPointer = true
			var offset int32
			var byts []byte

			byts = PROGRAM.Memory[*finalOffset : *finalOffset+TypePointerSize]

			encoder.DeserializeAtomic(byts, &offset)
			*finalOffset = int(offset)
		}
		if dbg {
			fmt.Println("\tupdate", arg.Name, arg.DereferenceOperations, *finalOffset, PROGRAM.Memory[*finalOffset:*finalOffset+10])
		}
	}
	if dbg {
		fmt.Println("\tupdate", arg.Name, arg.DereferenceOperations, *finalOffset, PROGRAM.Memory[*finalOffset:*finalOffset+10])
	}

	// if *finalOffset >= PROGRAM.HeapStartsAt {
	if *finalOffset >= PROGRAM.HeapStartsAt && isPointer {
		// then it's an object
		*finalOffset += ObjectHeaderSize
		if arg.IsSlice {
			*finalOffset += SliceHeaderSize
		}
	}
}

// GetStrOffset ...
func GetStrOffset(fp int, arg *CXArgument) int {
	strOffset := GetFinalOffset(fp, arg)
	if arg.Name != "" {
		// then it's not a literal
		var offset = int32(0)
		encoder.DeserializeAtomic(PROGRAM.Memory[strOffset:strOffset+TypePointerSize], &offset)
		strOffset = int(offset)
	}
	return strOffset
}

// GetFinalOffset ...
func GetFinalOffset(fp int, arg *CXArgument) int {
	// defer RuntimeError(PROGRAM)
	// var elt *CXArgument
	var finalOffset = int(arg.Offset)
	// var fldIdx int

	// elt = arg

	var dbg bool
	if arg.Name != "" {
		dbg = false
	}

	if finalOffset < StackSize {
		// then it's in the stack, not in data or heap
		finalOffset += fp
	}

	if dbg {
		fmt.Println("(start", arg.Name, fmt.Sprintf("%s:%d", arg.FileName, arg.FileLine), finalOffset, arg.DereferenceOperations)
	}

	// elt = arg
	CalculateDereferences(arg, &finalOffset, fp, dbg)
	for _, fld := range arg.Fields {
		// elt = fld
		finalOffset += fld.Offset
		CalculateDereferences(fld, &finalOffset, fp, dbg)
	}

	if dbg {
		fmt.Println("\t\tresult", finalOffset, PROGRAM.Memory[finalOffset:finalOffset+10], "...)")
	}

	return finalOffset
}

// ReadMemory ...
func ReadMemory(offset int, arg *CXArgument) []byte {
	return PROGRAM.Memory[offset : offset+arg.TotalSize]
}

// Mark marks all the alive objects in the heap
func Mark(prgrm *CXProgram) {
	fp := 0
	for c := 0; c <= prgrm.CallCounter; c++ {
		op := prgrm.CallStack[c].Operator

		for _, ptr := range op.ListOfPointers {
			var heapOffset int32
			encoder.DeserializeAtomic(prgrm.Memory[fp+ptr.Offset:fp+ptr.Offset+TypePointerSize], &heapOffset)

			prgrm.Memory[heapOffset] = 1
		}

		fp += op.Size
	}
}

// MarkAndCompact ...
func MarkAndCompact() {
	var fp int
	var faddr = int32(NullHeapAddressOffset)

	// marking, setting forward addresses and updating references
	for c := 0; c <= PROGRAM.CallCounter; c++ {
		op := PROGRAM.CallStack[c].Operator

		for _, ptr := range op.ListOfPointers {
			var heapOffset int32
			encoder.DeserializeAtomic(PROGRAM.Memory[fp+ptr.Offset:fp+ptr.Offset+TypePointerSize], &heapOffset)

			if heapOffset == NullHeapAddress {
				continue
			}

			// marking as alive
			PROGRAM.Memory[heapOffset] = 1

			for i, byt := range encoder.SerializeAtomic(faddr) {
				// setting forwarding address
				PROGRAM.Memory[int(heapOffset)+MarkSize+i] = byt
				// updating reference
				PROGRAM.Memory[fp+ptr.Offset+i] = byt
			}

			var objSize int32
			encoder.DeserializeAtomic(PROGRAM.Memory[int(heapOffset)+MarkSize+TypePointerSize:int(heapOffset)+MarkSize+TypePointerSize+ObjectSize], &objSize)

			faddr += int32(ObjectHeaderSize) + objSize
		}

		fp += op.Size
	}

	// relocation of live objects
	newHeapPointer := NullHeapAddressOffset
	for c := NullHeapAddressOffset; c < PROGRAM.HeapPointer; {
		var forwardingAddress int32
		encoder.DeserializeAtomic(PROGRAM.Memory[PROGRAM.HeapStartsAt+c+MarkSize:PROGRAM.HeapStartsAt+c+MarkSize+ForwardingAddressSize], &forwardingAddress)

		var objSize int32
		encoder.DeserializeAtomic(PROGRAM.Memory[PROGRAM.HeapStartsAt+c+MarkSize+ForwardingAddressSize:PROGRAM.HeapStartsAt+c+MarkSize+ForwardingAddressSize+ObjectSize], &objSize)

		if PROGRAM.Memory[c] == 1 {
			// setting the mark back to 0
			PROGRAM.Memory[c] = 0
			// then it's alive and we'll relocate the object
			for i := int32(0); i < ObjectHeaderSize+objSize; i++ {
				PROGRAM.Memory[forwardingAddress+i] = PROGRAM.Memory[int32(c)+i]
			}
			newHeapPointer += ObjectHeaderSize + int(objSize)
		}

		c += ObjectHeaderSize + int(objSize)
	}

	PROGRAM.HeapPointer = newHeapPointer
}

// ResizeMemory ...
func ResizeMemory(newMemSize int, isExpand bool) {
	if newMemSize > MaxHeapSize {
		// heap exhausted
		panic(HeapExhaustedError)
	}

	if isExpand {
		PROGRAM.Memory = append(PROGRAM.Memory, make([]byte, MemorySize-newMemSize)...)
		MemorySize = newMemSize
	} else {
		PROGRAM.Memory = PROGRAM.Memory[:newMemSize]
		MemorySize = newMemSize
	}
}

// AllocateSeq allocates memory in the heap
func AllocateSeq(size int) (offset int) {
	result := PROGRAM.HeapStartsAt + PROGRAM.HeapPointer
	newFree := PROGRAM.HeapPointer + size

	// if newFree > MEMORY_SIZE {
	if result+size > MemorySize {
		// call GC
		MarkAndCompact()
		result = PROGRAM.HeapStartsAt + PROGRAM.HeapPointer
		newFree = PROGRAM.HeapPointer + size

		freeMemPerc := 1.0 - float32(newFree)/float32(MemorySize-PROGRAM.HeapStartsAt)

		if freeMemPerc < float32(MinHeapFreeRatio)/100.0 {
			// then we have less than MIN_HEAP_FREE_RATIO memory left. expand!
			ResizeMemory(int(float32(MinHeapFreeRatio*(MemorySize-PROGRAM.HeapStartsAt))/freeMemPerc), true)
		}

		if freeMemPerc > float32(MaxHeapFreeRatio)/100.0 {
			// then we have more than MAX_HEAP_FREE_RATIO memory left. shrink!
			ResizeMemory(int(float32(MaxHeapFreeRatio*(MemorySize-PROGRAM.HeapStartsAt))/freeMemPerc), false)
		}
	}

	PROGRAM.HeapPointer = newFree

	return result
}

// WriteMemory ...
func WriteMemory(offset int, byts []byte) {
	for c := 0; c < len(byts); c++ {
		PROGRAM.Memory[offset+c] = byts[c]
	}
}

// Utilities

// FromBool ...
func FromBool(in bool) []byte {
	if in {
		return []byte{1}
	}
	return []byte{0}

}

// FromByte ...
func FromByte(in byte) []byte {
	return encoder.SerializeAtomic(in)
}

// FromStr ...
func FromStr(in string) []byte {
	return encoder.Serialize(in)
}

// FromI8 ...
func FromI8(in int8) []byte {
	return encoder.SerializeAtomic(in)
}

// FromI32 ...
func FromI32(in int32) []byte {
	return encoder.SerializeAtomic(in)
}

// FromUI32 ...
func FromUI32(in uint32) []byte {
	return encoder.SerializeAtomic(in)
}

// FromI64 ...
func FromI64(in int64) []byte {
	return encoder.Serialize(in)
}

// FromF32 ...
func FromF32(in float32) []byte {
	return encoder.Serialize(in)
}

// FromF64 ...
func FromF64(in float64) []byte {
	return encoder.Serialize(in)
}

// func ReadArray(mem []byte, fp int, inp *CXArgument, indexes []int32) (int, int) {
// 	var offset int
// 	var size int = inp.Size
// 	for i, idx := range indexes {
// 		offset += int(idx) * inp.Lengths[i]
// 	}
// 	for _, len := range indexes {
// 		size *= int(len)
// 	}

// 	return offset, size
// }

// ReadF32A ...
func ReadF32A(fp int, inp *CXArgument) (out []float32) {
	offset := GetFinalOffset(fp, inp)
	byts := ReadMemory(offset, inp)
	byts = append(encoder.SerializeAtomic(int32(len(byts)/4)), byts...)
	encoder.DeserializeRaw(byts, &out)
	return
}

// ReadBool ...
func ReadBool(fp int, inp *CXArgument) (out bool) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeRaw(ReadMemory(offset, inp), &out)
	return
}

// ReadByte ...
func ReadByte(fp int, inp *CXArgument) (out byte) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeAtomic(ReadMemory(offset, inp), &out)
	return
}

// ReadStr ...
func ReadStr(fp int, inp *CXArgument) (out string) {
	var offset int32
	off := GetFinalOffset(fp, inp)
	if inp.Name == "" {
		// then it's a literal
		offset = int32(off)
	} else {
		encoder.DeserializeAtomic(PROGRAM.Memory[off:off+TypePointerSize], &offset)
	}

	if offset == 0 {
		// then it's nil string
		out = ""
		return
	}

	var size int32
	sizeB := PROGRAM.Memory[offset : offset+StrHeaderSize]

	encoder.DeserializeAtomic(sizeB, &size)
	encoder.DeserializeRaw(PROGRAM.Memory[offset:offset+StrHeaderSize+size], &out)

	return
}

// ReadI8 ...
func ReadI8(fp int, inp *CXArgument) (out int8) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeAtomic(ReadMemory(offset, inp), &out)
	return
}

// ReadI32 ...
func ReadI32(fp int, inp *CXArgument) (out int32) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeAtomic(ReadMemory(offset, inp), &out)
	return
}

// ReadI64 ...
func ReadI64(fp int, inp *CXArgument) (out int64) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeRaw(ReadMemory(offset, inp), &out)
	return
}

// ReadF32 ...
func ReadF32(fp int, inp *CXArgument) (out float32) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeRaw(ReadMemory(offset, inp), &out)
	return
}

// ReadF64 ...
func ReadF64(fp int, inp *CXArgument) (out float64) {
	offset := GetFinalOffset(fp, inp)
	encoder.DeserializeRaw(ReadMemory(offset, inp), &out)
	return
}
