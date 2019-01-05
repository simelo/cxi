// +build extra full

package base

// "os"

const (
	// gogl
	OP_GL_INIT = iota + END_OF_BASE_OPS
	OP_GL_STRS
	OP_GL_FREE
	OP_GL_NEW_TEXTURE
	OP_GL_NEW_GIF
	OP_GL_FREE_GIF
	OP_GL_GIF_FRAME_TO_TEXTURE

	// gl_0_0
	OP_GL_MATRIX_MODE
	OP_GL_ROTATEF
	OP_GL_TRANSLATEF
	OP_GL_LOAD_IDENTITY
	OP_GL_PUSH_MATRIX
	OP_GL_POP_MATRIX
	OP_GL_ENABLE_CLIENT_STATE
	OP_GL_COLOR3F
	OP_GL_COLOR4F
	OP_GL_BEGIN
	OP_GL_END
	OP_GL_NORMAL3F
	OP_GL_VERTEX_2F
	OP_GL_VERTEX_3F
	OP_GL_LIGHTFV
	OP_GL_FRUSTUM
	OP_GL_TEX_ENVI
	OP_GL_ORTHO
	OP_GL_SCALEF
	OP_GL_TEX_COORD_2D
	OP_GL_TEX_COORD_2F

	// gl_1_0
	OP_GL_CULL_FACE
	OP_GL_HINT
	OP_GL_SCISSOR
	OP_GL_TEX_PARAMETERI
	OP_GL_TEX_IMAGE_2D
	OP_GL_CLEAR
	OP_GL_CLEAR_COLOR
	OP_GL_CLEAR_STENCIL
	OP_GL_CLEAR_DEPTH
	OP_GL_STENCIL_MASK
	OP_GL_COLOR_MASK
	OP_GL_DEPTH_MASK
	OP_GL_DISABLE
	OP_GL_ENABLE
	OP_GL_BLEND_FUNC
	OP_GL_STENCIL_FUNC
	OP_GL_STENCIL_OP
	OP_GL_DEPTH_FUNC
	OP_GL_GET_ERROR
	OP_GL_GET_TEX_LEVEL_PARAMETERIV
	OP_GL_DEPTH_RANGE
	OP_GL_VIEWPORT

	// gl_1_1
	OP_GL_DRAW_ARRAYS
	OP_GL_BIND_TEXTURE
	OP_GL_DELETE_TEXTURES
	OP_GL_GEN_TEXTURES

	// gl_1_3
	OP_GL_ACTIVE_TEXTURE

	// gl_1_5
	OP_GL_BIND_BUFFER
	OP_GL_DELETE_BUFFERS
	OP_GL_GEN_BUFFERS
	OP_GL_BUFFER_DATA
	OP_GL_BUFFER_SUB_DATA

	// gl_2_0
	OP_GL_STENCIL_OP_SEPARATE
	OP_GL_STENCIL_FUNC_SEPARATE
	OP_GL_STENCIL_MASK_SEPARATE
	OP_GL_ATTACH_SHADER
	OP_GL_BIND_ATTRIB_LOCATION
	OP_GL_COMPILE_SHADER
	OP_GL_CREATE_PROGRAM
	OP_GL_CREATE_SHADER
	OP_GL_DELETE_PROGRAM
	OP_GL_DELETE_SHADER
	OP_GL_DETACH_SHADER
	OP_GL_ENABLE_VERTEX_ATTRIB_ARRAY
	OP_GL_GET_ATTRIB_LOCATION
	OP_GL_GET_SHADERIV
	OP_GL_GET_UNIFORM_LOCATION
	OP_GL_LINK_PROGRAM
	OP_GL_SHADER_SOURCE
	OP_GL_USE_PROGRAM
	OP_GL_UNIFORM_1F
	OP_GL_UNIFORM_1I
	OP_GL_VERTEX_ATTRIB_POINTER
	OP_GL_VERTEX_ATTRIB_POINTER_I32

	// gl_3_0
	OP_GL_BIND_RENDERBUFFER
	OP_GL_DELETE_RENDERBUFFERS
	OP_GL_GEN_RENDERBUFFERS
	OP_GL_RENDERBUFFER_STORAGE
	OP_GL_BIND_FRAMEBUFFER
	OP_GL_DELETE_FRAMEBUFFERS
	OP_GL_GEN_FRAMEBUFFERS
	OP_GL_CHECK_FRAMEBUFFER_STATUS
	OP_GL_FRAMEBUFFER_TEXTURE_2D
	OP_GL_FRAMEBUFFER_RENDERBUFFER
	OP_GL_BIND_VERTEX_ARRAY
	OP_GL_DELETE_VERTEX_ARRAYS
	OP_GL_GEN_VERTEX_ARRAYS

	// glfw
	OP_GLFW_INIT
	OP_GLFW_WINDOW_HINT
	OP_GLFW_CREATE_WINDOW
	OP_GLFW_MAKE_CONTEXT_CURRENT
	OP_GLFW_SHOULD_CLOSE
	OP_GLFW_SET_SHOULD_CLOSE
	OP_GLFW_POLL_EVENTS
	OP_GLFW_SWAP_BUFFERS
	OP_GLFW_GET_FRAMEBUFFER_SIZE
	OP_GLFW_SWAP_INTERVAL
	OP_GLFW_SET_KEY_CALLBACK
	OP_GLFW_SET_KEY_CALLBACK_EX
	OP_GLFW_GET_TIME
	OP_GLFW_SET_MOUSE_BUTTON_CALLBACK
	OP_GLFW_SET_MOUSE_BUTTON_CALLBACK_EX
	OP_GLFW_SET_CURSOR_POS_CALLBACK
	OP_GLFW_SET_CURSOR_POS_CALLBACK_EX
	OP_GLFW_GET_CURSOR_POS
	OP_GLFW_SET_INPUT_MODE
	OP_GLFW_SET_WINDOW_POS
	OP_GLFW_GET_KEY
	OP_GLFW_FUNC_I32_I32
	OP_GLFW_CALL_I32_I32

	// gltext
	OP_GLTEXT_LOAD_TRUE_TYPE
	OP_GLTEXT_LOAD_TRUE_TYPE_EX
	OP_GLTEXT_PRINTF
	OP_GLTEXT_METRICS
	OP_GLTEXT_TEXTURE
	OP_GLTEXT_NEXT_GLYPH
	OP_GLTEXT_GLYPH_BOUNDS
	OP_GLTEXT_GLYPH_METRICS
	OP_GLTEXT_GLYPH_INFO
)

var execNativeExtra func(*CXProgram)

func init() {
	// gogl
	AddOpCode(OP_GL_INIT, "gl.Init", []int{}, []int{})
	AddOpCode(OP_GL_STRS, "gl.Strs", []int{TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GL_FREE, "gl.Free", []int{TypeStr}, []int{})
	AddOpCode(OP_GL_NEW_TEXTURE, "gl.NewTexture", []int{TypeStr}, []int{TypeI32})
	AddOpCode(OP_GL_NEW_GIF, "gl.NewGIF", []int{TypeStr}, []int{TypeI32, TypeI32, TypeI32, TypeI32})
	AddOpCode(OP_GL_FREE_GIF, "gl.FreeGIF", []int{TypeStr}, []int{})
	AddOpCode(OP_GL_GIF_FRAME_TO_TEXTURE, "gl.GIFFrameToTexture", []int{TypeStr, TypeI32, TypeI32}, []int{TypeI32, TypeI32})

	// gl_0.0
	AddOpCode(OP_GL_MATRIX_MODE, "gl.MatrixMode", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_ROTATEF, "gl.Rotatef", []int{TypeF32, TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_TRANSLATEF, "gl.Translatef", []int{TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_LOAD_IDENTITY, "gl.LoadIdentity", []int{}, []int{})
	AddOpCode(OP_GL_PUSH_MATRIX, "gl.PushMatrix", []int{}, []int{})
	AddOpCode(OP_GL_POP_MATRIX, "gl.PopMatrix", []int{}, []int{})
	AddOpCode(OP_GL_ENABLE_CLIENT_STATE, "gl.EnableClientState", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_COLOR3F, "gl.Color3f", []int{TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_COLOR4F, "gl.Color4f", []int{TypeF32, TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_BEGIN, "gl.Begin", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_END, "gl.End", []int{}, []int{})
	AddOpCode(OP_GL_NORMAL3F, "gl.Normal3f", []int{TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_VERTEX_2F, "gl.Vertex2f", []int{TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_VERTEX_3F, "gl.Vertex3f", []int{TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_LIGHTFV, "gl.Lightfv", []int{TypeI32, TypeI32, TypeF32}, []int{})
	AddOpCode(OP_GL_FRUSTUM, "gl.Frustum", []int{TypeF64, TypeF64, TypeF64, TypeF64, TypeF64, TypeF64}, []int{})
	AddOpCode(OP_GL_TEX_ENVI, "gl.TexEnvi", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_ORTHO, "gl.Ortho", []int{TypeF64, TypeF64, TypeF64, TypeF64, TypeF64, TypeF64}, []int{})
	AddOpCode(OP_GL_SCALEF, "gl.Scalef", []int{TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_TEX_COORD_2D, "gl.TexCoord2d", []int{TypeF64, TypeF64}, []int{})
	AddOpCode(OP_GL_TEX_COORD_2F, "gl.TexCoord2f", []int{TypeF32, TypeF32}, []int{})

	// gl_1_0
	AddOpCode(OP_GL_CULL_FACE, "gl.CullFace", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_HINT, "gl.Hint", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_SCISSOR, "gl.Scissor", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_TEX_PARAMETERI, "gl.TexParameteri", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_TEX_IMAGE_2D, "gl.TexImage2D", []int{TypeI32, TypeI32, TypeI32, TypeI32, TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_CLEAR, "gl.Clear", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_CLEAR_COLOR, "gl.ClearColor", []int{TypeF32, TypeF32, TypeF32, TypeF32}, []int{})
	AddOpCode(OP_GL_CLEAR_STENCIL, "gl.ClearStencil", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_CLEAR_DEPTH, "gl.ClearDepth", []int{TypeF64}, []int{})
	AddOpCode(OP_GL_STENCIL_MASK, "gl.StencilMask", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_COLOR_MASK, "gl.ColorMask", []int{TypeBool, TypeBool, TypeBool, TypeBool}, []int{})
	AddOpCode(OP_GL_DEPTH_MASK, "gl.DepthMask", []int{TypeBool}, []int{})
	AddOpCode(OP_GL_DISABLE, "gl.Disable", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_ENABLE, "gl.Enable", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_BLEND_FUNC, "gl.BlendFunc", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_STENCIL_FUNC, "gl.StencilFunc", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_STENCIL_OP, "gl.StencilOp", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_DEPTH_FUNC, "gl.DepthFunc", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_GET_ERROR, "gl.GetError", []int{}, []int{TypeI32})
	AddOpCode(OP_GL_GET_TEX_LEVEL_PARAMETERIV, "gl.GetTexLevelParameteriv", []int{TypeI32, TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_DEPTH_RANGE, "gl.DepthRange", []int{TypeF64, TypeF64}, []int{})
	AddOpCode(OP_GL_VIEWPORT, "gl.Viewport", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})

	// gl_1_1
	AddOpCode(OP_GL_DRAW_ARRAYS, "gl.DrawArrays", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_BIND_TEXTURE, "gl.BindTexture", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_TEXTURES, "gl.DeleteTextures", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GEN_TEXTURES, "gl.GenTextures", []int{TypeI32, TypeI32}, []int{TypeI32})

	// gl_1_3
	AddOpCode(OP_GL_ACTIVE_TEXTURE, "gl.ActiveTexture", []int{TypeI32}, []int{})

	// gl_1_5
	AddOpCode(OP_GL_BIND_BUFFER, "gl.BindBuffer", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_BUFFERS, "gl.DeleteBuffers", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GEN_BUFFERS, "gl.GenBuffers", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_BUFFER_DATA, "gl.BufferData", []int{TypeI32, TypeI32, TypeF32, TypeI32}, []int{})
	AddOpCode(OP_GL_BUFFER_SUB_DATA, "gl.BufferSubData", []int{TypeI32, TypeI32, TypeI32, TypeF32}, []int{})

	//gl_2_0
	AddOpCode(OP_GL_STENCIL_OP_SEPARATE, "gl.StencilOpSeparate", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_STENCIL_FUNC_SEPARATE, "gl.StencilFuncSeparate", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_STENCIL_MASK_SEPARATE, "gl.StencilMaskSeparate", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_ATTACH_SHADER, "gl.AttachShader", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_BIND_ATTRIB_LOCATION, "gl.BindAttribLocation", []int{TypeI32, TypeI32, TypeStr}, []int{})
	AddOpCode(OP_GL_COMPILE_SHADER, "gl.CompileShader", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_CREATE_PROGRAM, "gl.CreateProgram", []int{}, []int{TypeI32})
	AddOpCode(OP_GL_CREATE_SHADER, "gl.CreateShader", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_DELETE_PROGRAM, "gl.DeleteProgram", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_SHADER, "gl.DeleteShader", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_DETACH_SHADER, "gl.DetachShader", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_ENABLE_VERTEX_ATTRIB_ARRAY, "gl.EnableVertexAttribArray", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_GET_ATTRIB_LOCATION, "gl.GetAttribLocation", []int{TypeI32, TypeStr}, []int{TypeI32})
	AddOpCode(OP_GL_GET_SHADERIV, "gl.GetShaderiv", []int{TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GET_UNIFORM_LOCATION, "gl.GetUniformLocation", []int{TypeI32, TypeStr}, []int{TypeI32})
	AddOpCode(OP_GL_LINK_PROGRAM, "gl.LinkProgram", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_SHADER_SOURCE, "gl.ShaderSource", []int{TypeI32, TypeI32, TypeStr}, []int{})
	AddOpCode(OP_GL_USE_PROGRAM, "gl.UseProgram", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_UNIFORM_1F, "gl.Uniform1f", []int{TypeI32, TypeF32}, []int{})
	AddOpCode(OP_GL_UNIFORM_1I, "gl.Uniform1i", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_VERTEX_ATTRIB_POINTER, "gl.VertexAttribPointer", []int{TypeI32, TypeI32, TypeI32, TypeBool, TypeI32}, []int{})
	AddOpCode(OP_GL_VERTEX_ATTRIB_POINTER_I32, "gl.VertexAttribPointerI32", []int{TypeI32, TypeI32, TypeI32, TypeBool, TypeI32, TypeI32}, []int{})

	// gl_3_0
	AddOpCode(OP_GL_BIND_RENDERBUFFER, "gl.BindRenderbuffer", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_RENDERBUFFERS, "gl.DeleteRenderbuffers", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GEN_RENDERBUFFERS, "gl.GenRenderbuffers", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_RENDERBUFFER_STORAGE, "gl.RenderbufferStorage", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_BIND_FRAMEBUFFER, "gl.BindFramebuffer", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_FRAMEBUFFERS, "gl.DeleteFramebuffers", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GEN_FRAMEBUFFERS, "gl.GenFramebuffers", []int{TypeI32, TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_CHECK_FRAMEBUFFER_STATUS, "gl.CheckFramebufferStatus", []int{TypeI32}, []int{TypeI32})
	AddOpCode(OP_GL_FRAMEBUFFER_TEXTURE_2D, "gl.FramebufferTexture2D", []int{TypeI32, TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_FRAMEBUFFER_RENDERBUFFER, "gl.FramebufferRenderbuffer", []int{TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_BIND_VERTEX_ARRAY, "gl.BindVertexArray", []int{TypeI32}, []int{})
	AddOpCode(OP_GL_DELETE_VERTEX_ARRAYS, "gl.DeleteVertexArrays", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GL_GEN_VERTEX_ARRAYS, "gl.GenVertexArrays", []int{TypeI32, TypeI32}, []int{TypeI32})

	// glfw
	AddOpCode(OP_GLFW_INIT, "glfw.Init", []int{}, []int{})
	AddOpCode(OP_GLFW_WINDOW_HINT, "glfw.WindowHint", []int{TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GLFW_CREATE_WINDOW, "glfw.CreateWindow", []int{TypeStr, TypeI32, TypeI32, TypeStr}, []int{})
	AddOpCode(OP_GLFW_MAKE_CONTEXT_CURRENT, "glfw.MakeContextCurrent", []int{TypeStr}, []int{})
	AddOpCode(OP_GLFW_SHOULD_CLOSE, "glfw.ShouldClose", []int{TypeStr}, []int{TypeBool})
	AddOpCode(OP_GLFW_SET_SHOULD_CLOSE, "glfw.SetShouldClose", []int{TypeStr, TypeBool}, []int{})
	AddOpCode(OP_GLFW_POLL_EVENTS, "glfw.PollEvents", []int{}, []int{})
	AddOpCode(OP_GLFW_SWAP_BUFFERS, "glfw.SwapBuffers", []int{TypeStr}, []int{})
	AddOpCode(OP_GLFW_GET_FRAMEBUFFER_SIZE, "glfw.GetFramebufferSize", []int{TypeStr}, []int{TypeI32, TypeI32})
	AddOpCode(OP_GLFW_SWAP_INTERVAL, "glfw.SwapInterval", []int{TypeI32}, []int{})
	AddOpCode(OP_GLFW_SET_KEY_CALLBACK, "glfw.SetKeyCallback", []int{TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_SET_KEY_CALLBACK_EX, "glfw.SetKeyCallbackEx", []int{TypeStr, TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_GET_TIME, "glfw.GetTime", []int{}, []int{TypeF64})
	AddOpCode(OP_GLFW_SET_MOUSE_BUTTON_CALLBACK, "glfw.SetMouseButtonCallback", []int{TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_SET_MOUSE_BUTTON_CALLBACK_EX, "glfw.SetMouseButtonCallbackEx", []int{TypeStr, TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_SET_CURSOR_POS_CALLBACK, "glfw.SetCursorPosCallback", []int{TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_SET_CURSOR_POS_CALLBACK_EX, "glfw.SetCursorPosCallbackEx", []int{TypeStr, TypeStr, TypeStr}, []int{})
	AddOpCode(OP_GLFW_GET_CURSOR_POS, "glfw.GetCursorPos", []int{TypeStr}, []int{TypeF64, TypeF64})
	AddOpCode(OP_GLFW_SET_INPUT_MODE, "glfw.SetInputMode", []int{TypeStr, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GLFW_SET_WINDOW_POS, "glfw.SetWindowPos", []int{TypeStr, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GLFW_GET_KEY, "glfw.GetKey", []int{TypeStr, TypeI32}, []int{TypeI32})
	AddOpCode(OP_GLFW_FUNC_I32_I32, "glfw.func_i32_i32", []int{TypeStr, TypeStr}, []int{TypeI32})
	AddOpCode(OP_GLFW_CALL_I32_I32, "glfw.call_i32_i32", []int{TypeI32, TypeI32, TypeI32}, []int{})

	// gltext
	AddOpCode(OP_GLTEXT_LOAD_TRUE_TYPE, "gltext.LoadTrueType", []int{TypeStr, TypeStr, TypeI32, TypeI32, TypeI32, TypeI32}, []int{})
	AddOpCode(OP_GLTEXT_PRINTF, "gltext.Printf", []int{TypeStr, TypeF32, TypeF32, TypeStr}, []int{})
	AddOpCode(OP_GLTEXT_METRICS, "gltext.Metrics", []int{TypeStr, TypeStr}, []int{TypeI32, TypeI32})
	AddOpCode(OP_GLTEXT_TEXTURE, "gltext.Texture", []int{TypeStr}, []int{TypeI32})
	AddOpCode(OP_GLTEXT_NEXT_GLYPH, "gltext.NextGlyph", []int{TypeStr, TypeStr, TypeI32}, []int{TypeI32, TypeI32, TypeI32, TypeI32, TypeI32, TypeI32, TypeI32})
	AddOpCode(OP_GLTEXT_GLYPH_BOUNDS, "gltext.GlyphBounds", []int{}, []int{TypeI32, TypeI32})
	AddOpCode(OP_GLTEXT_GLYPH_METRICS, "gltext.GlyphMetrics", []int{TypeStr, TypeI32}, []int{TypeI32, TypeI32})
	AddOpCode(OP_GLTEXT_GLYPH_INFO, "gltext.GlyphInfo", []int{TypeStr, TypeI32}, []int{TypeI32, TypeI32, TypeI32, TypeI32, TypeI32})

	// exec
	execNativeExtra = func(prgrm *CXProgram) {
		defer RuntimeError()
		call := &prgrm.CallStack[prgrm.CallCounter]
		expr := call.Operator.Expressions[call.Line]
		opCode := expr.Operator.OpCode
		fp := call.FramePointer

		if opCode < END_OF_BARE_OPS {
			execNativeBare(prgrm)
		} else if opCode < END_OF_BASE_OPS {
			execNativeBase(prgrm)
		} else {
			switch opCode {
			// gogl
			case OP_GL_INIT:
				op_gl_Init()
			case OP_GL_STRS:
				op_gl_Strs(expr, fp)
			case OP_GL_FREE:
				op_gl_Free(expr, fp)
			case OP_GL_NEW_TEXTURE:
				op_gl_NewTexture(expr, fp)
			case OP_GL_NEW_GIF:
				op_gl_NewGIF(expr, fp)
			case OP_GL_FREE_GIF:
				op_gl_FreeGIF(expr, fp)
			case OP_GL_GIF_FRAME_TO_TEXTURE:
				op_gl_GIFFrameToTexture(expr, fp)

			// gl_0_0
			case OP_GL_MATRIX_MODE:
				op_gl_MatrixMode(expr, fp)
			case OP_GL_ROTATEF:
				op_gl_Rotatef(expr, fp)
			case OP_GL_TRANSLATEF:
				op_gl_Translatef(expr, fp)
			case OP_GL_LOAD_IDENTITY:
				op_gl_LoadIdentity()
			case OP_GL_PUSH_MATRIX:
				op_gl_PushMatrix()
			case OP_GL_POP_MATRIX:
				op_gl_PopMatrix()
			case OP_GL_ENABLE_CLIENT_STATE:
				op_gl_EnableClientState(expr, fp)
			case OP_GL_COLOR3F:
				op_gl_Color3f(expr, fp)
			case OP_GL_COLOR4F:
				op_gl_Color4f(expr, fp)
			case OP_GL_BEGIN:
				op_gl_Begin(expr, fp)
			case OP_GL_END:
				op_gl_End()
			case OP_GL_NORMAL3F:
				op_gl_Normal3f(expr, fp)
			case OP_GL_VERTEX_2F:
				op_gl_Vertex2f(expr, fp)
			case OP_GL_VERTEX_3F:
				op_gl_Vertex3f(expr, fp)
			case OP_GL_LIGHTFV:
				op_gl_Lightfv(expr, fp)
			case OP_GL_FRUSTUM:
				op_gl_Frustum(expr, fp)
			case OP_GL_TEX_ENVI:
				op_gl_TexEnvi(expr, fp)
			case OP_GL_ORTHO:
				op_gl_Ortho(expr, fp)
			case OP_GL_SCALEF:
				op_gl_Scalef(expr, fp)
			case OP_GL_TEX_COORD_2D:
				op_gl_TexCoord2d(expr, fp)
			case OP_GL_TEX_COORD_2F:
				op_gl_TexCoord2f(expr, fp)

			// gl_1_0
			case OP_GL_CULL_FACE:
				op_gl_CullFace(expr, fp)
			case OP_GL_HINT:
				op_gl_Hint(expr, fp)
			case OP_GL_SCISSOR:
				op_gl_Scissor(expr, fp)
			case OP_GL_TEX_PARAMETERI:
				op_gl_TexParameteri(expr, fp)
			case OP_GL_TEX_IMAGE_2D:
				op_gl_TexImage2D(expr, fp)
			case OP_GL_CLEAR:
				op_gl_Clear(expr, fp)
			case OP_GL_CLEAR_COLOR:
				op_gl_ClearColor(expr, fp)
			case OP_GL_CLEAR_STENCIL:
				op_gl_ClearStencil(expr, fp)
			case OP_GL_CLEAR_DEPTH:
				op_gl_ClearDepth(expr, fp)
			case OP_GL_STENCIL_MASK:
				op_gl_StencilMask(expr, fp)
			case OP_GL_COLOR_MASK:
				op_gl_ColorMask(expr, fp)
			case OP_GL_DEPTH_MASK:
				op_gl_DepthMask(expr, fp)
			case OP_GL_DISABLE:
				op_gl_Disable(expr, fp)
			case OP_GL_ENABLE:
				op_gl_Enable(expr, fp)
			case OP_GL_BLEND_FUNC:
				op_gl_BlendFunc(expr, fp)
			case OP_GL_STENCIL_FUNC:
				op_gl_StencilFunc(expr, fp)
			case OP_GL_STENCIL_OP:
				op_gl_StencilOp(expr, fp)
			case OP_GL_DEPTH_FUNC:
				op_gl_DepthFunc(expr, fp)
			case OP_GL_GET_ERROR:
				op_gl_GetError(expr, fp)
			case OP_GL_GET_TEX_LEVEL_PARAMETERIV:
				op_gl_GetTexLevelParameteriv(expr, fp)
			case OP_GL_DEPTH_RANGE:
				op_gl_DepthRange(expr, fp)
			case OP_GL_VIEWPORT:
				op_gl_Viewport(expr, fp)

			// gl_1_1
			case OP_GL_DRAW_ARRAYS:
				op_gl_DrawArrays(expr, fp)
			case OP_GL_BIND_TEXTURE:
				op_gl_BindTexture(expr, fp)
			case OP_GL_DELETE_TEXTURES:
				op_gl_DeleteTextures(expr, fp)
			case OP_GL_GEN_TEXTURES:
				op_gl_GenTextures(expr, fp)

			// gl_1_3
			case OP_GL_ACTIVE_TEXTURE:
				op_gl_ActiveTexture(expr, fp)

			// gl_1_5
			case OP_GL_BIND_BUFFER:
				op_gl_BindBuffer(expr, fp)
			case OP_GL_DELETE_BUFFERS:
				op_gl_DeleteBuffers(expr, fp)
			case OP_GL_GEN_BUFFERS:
				op_gl_GenBuffers(expr, fp)
			case OP_GL_BUFFER_DATA:
				op_gl_BufferData(expr, fp)
			case OP_GL_BUFFER_SUB_DATA:
				op_gl_BufferSubData(expr, fp)

			// gl_2_0
			case OP_GL_STENCIL_OP_SEPARATE:
				op_gl_StencilOpSeparate(expr, fp)
			case OP_GL_STENCIL_FUNC_SEPARATE:
				op_gl_StencilFuncSeparate(expr, fp)
			case OP_GL_STENCIL_MASK_SEPARATE:
				op_gl_StencilMaskSeparate(expr, fp)
			case OP_GL_ATTACH_SHADER:
				op_gl_AttachShader(expr, fp)
			case OP_GL_BIND_ATTRIB_LOCATION:
				op_gl_BindAttribLocation(expr, fp)
			case OP_GL_COMPILE_SHADER:
				op_gl_CompileShader(expr, fp)
			case OP_GL_CREATE_PROGRAM:
				op_gl_CreateProgram(expr, fp)
			case OP_GL_CREATE_SHADER:
				op_gl_CreateShader(expr, fp)
			case OP_GL_DELETE_PROGRAM:
				op_gl_DeleteProgram(expr, fp)
			case OP_GL_DELETE_SHADER:
				op_gl_DeleteShader(expr, fp)
			case OP_GL_DETACH_SHADER:
				op_gl_DetachShader(expr, fp)
			case OP_GL_ENABLE_VERTEX_ATTRIB_ARRAY:
				op_gl_EnableVertexAttribArray(expr, fp)
			case OP_GL_GET_ATTRIB_LOCATION:
				op_gl_GetAttribLocation(expr, fp)
			case OP_GL_GET_SHADERIV:
				op_gl_GetShaderiv(expr, fp)
			case OP_GL_GET_UNIFORM_LOCATION:
				op_gl_GetUniformLocation(expr, fp)
			case OP_GL_LINK_PROGRAM:
				op_gl_LinkProgram(expr, fp)
			case OP_GL_SHADER_SOURCE:
				op_gl_ShaderSource(expr, fp)
			case OP_GL_USE_PROGRAM:
				op_gl_UseProgram(expr, fp)
			case OP_GL_UNIFORM_1F:
				op_gl_Uniform1f(expr, fp)
			case OP_GL_UNIFORM_1I:
				op_gl_Uniform1i(expr, fp)
			case OP_GL_VERTEX_ATTRIB_POINTER:
				op_gl_VertexAttribPointer(expr, fp)
			case OP_GL_VERTEX_ATTRIB_POINTER_I32:
				op_gl_VertexAttribPointerI32(expr, fp)

			// gl_3_0
			case OP_GL_BIND_RENDERBUFFER:
				op_gl_BindRenderbuffer(expr, fp)
			case OP_GL_DELETE_RENDERBUFFERS:
				op_gl_DeleteRenderbuffers(expr, fp)
			case OP_GL_GEN_RENDERBUFFERS:
				op_gl_GenRenderbuffers(expr, fp)
			case OP_GL_RENDERBUFFER_STORAGE:
				op_gl_RenderbufferStorage(expr, fp)
			case OP_GL_BIND_FRAMEBUFFER:
				op_gl_BindFramebuffer(expr, fp)
			case OP_GL_DELETE_FRAMEBUFFERS:
				op_gl_DeleteFramebuffers(expr, fp)
			case OP_GL_GEN_FRAMEBUFFERS:
				op_gl_GenFramebuffers(expr, fp)
			case OP_GL_CHECK_FRAMEBUFFER_STATUS:
				op_gl_CheckFramebufferStatus(expr, fp)
			case OP_GL_FRAMEBUFFER_TEXTURE_2D:
				op_gl_FramebufferTexture2D(expr, fp)
			case OP_GL_FRAMEBUFFER_RENDERBUFFER:
				op_gl_FramebufferRenderbuffer(expr, fp)
			case OP_GL_BIND_VERTEX_ARRAY:
				op_gl_BindVertexArray(expr, fp)
			case OP_GL_DELETE_VERTEX_ARRAYS:
				op_gl_DeleteVertexArrays(expr, fp)
			case OP_GL_GEN_VERTEX_ARRAYS:
				op_gl_GenVertexArrays(expr, fp)

			// glfw
			case OP_GLFW_INIT:
				op_glfw_Init(expr, fp)
			case OP_GLFW_WINDOW_HINT:
				op_glfw_WindowHint(expr, fp)
			case OP_GLFW_CREATE_WINDOW:
				op_glfw_CreateWindow(expr, fp)
			case OP_GLFW_MAKE_CONTEXT_CURRENT:
				op_glfw_MakeContextCurrent(expr, fp)
			case OP_GLFW_SHOULD_CLOSE:
				op_glfw_ShouldClose(expr, fp)
			case OP_GLFW_SET_SHOULD_CLOSE:
				op_glfw_SetShouldClose(expr, fp)
			case OP_GLFW_POLL_EVENTS:
				op_glfw_PollEvents()
			case OP_GLFW_SWAP_BUFFERS:
				op_glfw_SwapBuffers(expr, fp)
			case OP_GLFW_GET_FRAMEBUFFER_SIZE:
				op_glfw_GetFramebufferSize(expr, fp)
			case OP_GLFW_SWAP_INTERVAL:
				op_glfw_SwapInterval(expr, fp)
			case OP_GLFW_SET_KEY_CALLBACK:
				op_glfw_SetKeyCallback(expr, fp)
			case OP_GLFW_SET_KEY_CALLBACK_EX:
				op_glfw_SetKeyCallbackEx(expr, fp)
			case OP_GLFW_GET_TIME:
				op_glfw_GetTime(expr, fp)
			case OP_GLFW_SET_MOUSE_BUTTON_CALLBACK:
				op_glfw_SetMouseButtonCallback(expr, fp)
			case OP_GLFW_SET_MOUSE_BUTTON_CALLBACK_EX:
				op_glfw_SetMouseButtonCallbackEx(expr, fp)
			case OP_GLFW_SET_CURSOR_POS_CALLBACK:
				op_glfw_SetCursorPosCallback(expr, fp)
			case OP_GLFW_SET_CURSOR_POS_CALLBACK_EX:
				op_glfw_SetCursorPosCallbackEx(expr, fp)
			case OP_GLFW_GET_CURSOR_POS:
				op_glfw_GetCursorPos(expr, fp)
			case OP_GLFW_SET_INPUT_MODE:
				op_glfw_SetInputMode(expr, fp)
			case OP_GLFW_SET_WINDOW_POS:
				op_glfw_SetWindowPos(expr, fp)
			case OP_GLFW_GET_KEY:
				op_glfw_GetKey(expr, fp)
			case OP_GLFW_FUNC_I32_I32:
				op_glfw_func_i32_i32(expr, fp)
			case OP_GLFW_CALL_I32_I32:
				op_glfw_call_i32_i32(expr, fp)

			// gltext
			case OP_GLTEXT_LOAD_TRUE_TYPE:
				op_gltext_LoadTrueType(expr, fp)
			case OP_GLTEXT_PRINTF:
				op_gltext_Printf(expr, fp)
			case OP_GLTEXT_METRICS:
				op_gltext_Metrics(expr, fp)
			case OP_GLTEXT_TEXTURE:
				op_gltext_Texture(expr, fp)
			case OP_GLTEXT_NEXT_GLYPH:
				op_gltext_NextGlyph(expr, fp)
			case OP_GLTEXT_GLYPH_BOUNDS:
				op_gltext_GlyphBounds(expr, fp)
			case OP_GLTEXT_GLYPH_METRICS:
				op_gltext_GlyphMetrics(expr, fp)
			case OP_GLTEXT_GLYPH_INFO:
				op_gltext_GlyphInfo(expr, fp)
			default:
				// DumpOpCodes(opCode)
				panic("invalid extra opcode")
			}
		}
	}

	execNative = execNativeExtra
}
