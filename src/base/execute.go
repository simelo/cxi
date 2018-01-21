package base

import (
	// "fmt"
	// "errors"
	"math/rand"
	"time"
	// "github.com/skycoin/skycoin/src/cipher/encoder"
)

func (prgrm *CXProgram) Run () error {
	// fmt.Println("")
	// prgrm.PrintProgram()
	rand.Seed(time.Now().UTC().UnixNano())

	if mod, err := prgrm.SelectPackage(MAIN_PKG); err == nil {
		if fn, err := mod.SelectFunction(MAIN_FUNC); err == nil {
			if len(fn.Expressions) < 1 {
				return nil
			}
			// main function
			mainCall := MakeCall(fn, nil, mod, mod.Program)
			
			// initializing program resources
			//prgrm.CallStack = append(prgrm.CallStack, mainCall)
			prgrm.CallStack[0] = mainCall
			prgrm.Stacks = append(prgrm.Stacks, MakeStack(1024))
			prgrm.Stacks[0].StackPointer = fn.Size

			var err error

			for !prgrm.Terminated {
				call := &prgrm.CallStack[prgrm.CallCounter]
				err = call.call(prgrm)
				if err != nil {
					return err
				}
			}
			return err
		} else {
			return err
		}
	} else {
		return err
	}
}

func (call *CXCall) call (prgrm *CXProgram) error {
	// CX is still single-threaded, so only one stack
	// fmt.Println(call.Line, call.Operator.Length, prgrm.CallCounter)
	if call.Line >= call.Operator.Length {
		/*
                  popping the stack
                */
		// going back to the previous call
		prgrm.CallCounter--
		if prgrm.CallCounter < 0 {
			// then the program finished
			prgrm.Terminated = true
			// fmt.Println(prgrm.Stacks[0].Stack)
			// fmt.Println("prgrm.Data", prgrm.Data)
		} else {
			// copying the outputs to the previous stack frame
			returnAddr := &prgrm.CallStack[prgrm.CallCounter]
			returnOp := returnAddr.Operator
			returnLine := returnAddr.Line
			returnFP := returnAddr.FramePointer
			fp := call.FramePointer

			expr := returnOp.Expressions[returnLine]
			outOffset := 0
			for i, out := range expr.Outputs {
				// copy byte by byte to the previous stack frame
				for c := 0; c < out.Size; c++ {
					prgrm.Stacks[0].Stack[returnFP + out.Offset + c] =
						prgrm.Stacks[0].Stack[fp + call.Operator.Outputs[i].Offset + c]
				}
				outOffset += out.Size
			}

			// return the stack pointer to its previous state
			prgrm.Stacks[0].StackPointer = call.FramePointer
			// we'll now execute the next command
			prgrm.CallStack[prgrm.CallCounter].Line++
			// calling the actual command
			prgrm.CallStack[prgrm.CallCounter].call(prgrm)
		}
	} else {
		/*
                  continue with call operator's execution
                */
		fn := call.Operator
		expr := fn.Expressions[call.Line]
		// if it's a native, then we just process the arguments with execNative
		if expr.Operator.IsNative {
			// previousFP := call.FramePointer
			// call.FramePointer = prgrm.Stacks[0].StackPointer
			// the stack pointer is moved to create room for the next call
			// prgrm.Stacks[0].StackPointer += fn.Size
			execNative(prgrm)
			// prgrm.Stacks[0].StackPointer = call.FramePointer
			// call.FramePointer = previousFP
			call.Line++
		} else {
			/*
                          It was not a native, so we need to create another call
                          with the current expression's operator
                        */






			

			/*
                  preparing inputs and outputs in the stack for the next function call
                */
			// fmt.Println("hi")
			// sp := prgrm.Stacks[0].StackPointer
			// fp := call.FramePointer
			// tmp := sp
			// for _, inp := range expr.Inputs {
			// 	// we write the input values for the next frame
			// 	size := inp.Size

			// 	var byts []byte
			// 	switch inp.MemoryType {
			// 	case MEM_STACK:
			// 		byts = prgrm.Stacks[0].Stack[fp + inp.Offset : fp + inp.Offset + inp.Size]
			// 	case MEM_DATA:
			// 		byts = inp.Program.Data[inp.Offset : inp.Offset + inp.Size]
			// 	default:
			// 		panic("implement the other mem types")
			// 	}
				
			// 	for c := 0; c < size; c++ {
			// 		// we copy each byte outside of current frame
			// 		// prgrm.Stacks[0].Stack[tmp+c] = prgrm.Stacks[0].Stack[fp+offset+c]
			// 		prgrm.Stacks[0].Stack[tmp+c] = byts[c]
			// 	}
			// 	tmp += size
			// }
			// for _, out := range expr.Outputs {
			// 	// we make room to receive the outputs
			// 	tmp += out.Size
			// }


			
			// once the subcall finishes, call next line
			// call.Line++
			// we're going to use the next call in the callstack
			prgrm.CallCounter++
			newCall := &prgrm.CallStack[prgrm.CallCounter]
			// setting the new call
			newCall.Operator = expr.Operator
			newCall.Line = 0
			newCall.FramePointer = prgrm.Stacks[0].StackPointer
			// the stack pointer is moved to create room for the next call
			prgrm.Stacks[0].StackPointer += fn.Size

			fp := newCall.FramePointer

			for i, inp := range expr.Inputs {
				var byts []byte
				switch inp.MemoryType {
				case MEM_STACK:
					byts = prgrm.Stacks[0].Stack[fp + inp.Offset : fp + inp.Offset + inp.Size]
				case MEM_DATA:
					byts = prgrm.Data[inp.Offset : inp.Offset + inp.Size]
				default:
					panic("implement the other mem types")
				}
				for c := 0; c < inp.Size; c++ {
					prgrm.Stacks[0].Stack[fp + newCall.Operator.Inputs[i].Offset + c] = 
					byts[c]
				}
			}
		}
	}
	return nil
}
