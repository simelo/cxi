package base

// constant codes
const (
	// https://www.khronos.org/registry/OpenGL/api/GL/glcorearb.h
	// gl_1_0
	ConstGlDepthBufferBit = iota
	ConstGlStencilBufferBit
	ConstGlColorBufferBit
	ConstGlFalse
	ConstGlTrue
	ConstGlPoints
	ConstGlLines
	ConstGlLineLoop
	ConstGlLineStrip
	ConstGlTriangles
	ConstGlTriangleStrip
	ConstGlTriangleFan
	ConstGlQuads
	ConstGlNever
	ConstGlLess
	ConstGlEqual
	ConstGlLequal
	ConstGlGreater
	ConstGlNotequal
	ConstGlGequal
	ConstGlAlways
	ConstGlZero
	ConstGlOne
	ConstGlSrcAlpha
	ConstGlOneMinusSrcAlpha
	ConstGlFront
	ConstGlBack
	ConstGlFrontAndBack
	ConstGlNoError
	ConstGlInvalidEnum
	ConstGlInvalidValue
	ConstGlInvalidOperation
	ConstGlStackOverflow
	ConstGlStackUnderflow
	ConstGlOutOfMemory
	ConstGlLineSmooth
	ConstGlPolygonSmooth
	ConstGlCullFace
	ConstGlDepthRange
	ConstGlDepthTest
	ConstGlDepthWritemask
	ConstGlDepthClearValue
	ConstGlDepthFunc
	ConstGlStencilTest
	ConstGlStencilClearValue
	ConstGlStencilFunc
	ConstGlStencilValueMask
	ConstGlStencilFail
	ConstGlStencilPassDepthFail
	ConstGlStencilPassDepthPass
	ConstGlStencilRef
	ConstGlStencilWritemask
	ConstGlDither
	ConstGlBlend
	ConstGlScissorTest
	ConstGlPolygonSmoothHint
	ConstGlTexture2d
	ConstGlTextureWidth
	ConstGlTextureHeight
	ConstGlDontCare
	ConstGlUnsignedByte
	ConstGlFloat
	ConstGlInvert
	ConstGlTexture
	ConstGlColor
	ConstGlDepth
	ConstGlStencil
	ConstGlStencilIndex
	ConstGlDepthComponent
	ConstGlRgba
	ConstGlKeep
	ConstGlReplace
	ConstGlIncr
	ConstGlDecr
	ConstGlNearest
	ConstGlLinear
	ConstGlNearestMipmapNearest
	ConstGlLinearMipmapNearest
	ConstGlNearestMipmapLinear
	ConstGlLinearMipmapLinear
	ConstGlTextureMagFilter
	ConstGlTextureMinFilter
	ConstGlTextureWrapS
	ConstGlTextureWrapT
	ConstGlRepeat

	// gl_1_1
	ConstGlRgba8
	ConstGlVertexArray

	// gl_1_2
	ConstGlTextureWrapR
	ConstGlClampToEdge

	// gl_1_3
	ConstGlTexture0
	ConstGlMultisampleArb // remove _ARB
	ConstGlClampToBorder

	// gl_1_4
	ConstGlDepthComponent16
	ConstGlDepthComponent24
	ConstGlDepthComponent32
	ConstGlMirroredRepeat
	ConstGlIncrWrap
	ConstGlDecrWrap

	// gl_1_5
	ConstGlArrayBuffer
	ConstGlStreamDraw
	ConstGlStreamRead
	ConstGlStreamCopy
	ConstGlStaticDraw
	ConstGlStaticRead
	ConstGlStaticCopy
	ConstGlDynamicDraw
	ConstGlDynamicRead
	ConstGlDynamicCopy

	// gl_2_0
	ConstGlStencilBackFunc
	ConstGlStencilBackFail
	ConstGlStencilBackPassDepthFail
	ConstGlStencilBackPassDepthPass
	ConstGlFragmentShader
	ConstGlVertexShader
	ConstGlStencilBackRef
	ConstGlStencilBackValueMask
	ConstGlStencilBackWritemask

	// gl_3_0
	ConstGlDepthComponent32f
	ConstGlDepth32fStencil8
	ConstGlFramebufferUndefined
	ConstGlDepthStencilAttachment
	ConstGlDepthStencil
	ConstGlDepth24Stencil8
	ConstGlFramebufferComplete
	ConstGlFramebufferIncompleteAttachment
	ConstGlFramebufferIncompleteMissingAttachment
	ConstGlFramebufferIncompleteDrawBuffer
	ConstGlFramebufferIncompleteReadBuffer
	ConstGlFramebufferUnsupported
	ConstGlColorAttachment0
	ConstGlDepthAttachment
	ConstGlStencilAttachment
	ConstGlFramebuffer
	ConstGlRenderbuffer
	ConstGlFramebufferIncompleteMultisample
	ConstGlStencilIndex1
	ConstGlStencilIndex4
	ConstGlStencilIndex8
	ConstGlStencilIndex16

	// gl_3_2
	ConstGlFramebufferIncompleteLayerTargets

	// gl_4_4
	ConstGlMirrorClampToEdge

	// Fixed pipeline. Deprecated ?
	ConstGlPolygon
	ConstGlModelview
	ConstGlProjection
	ConstGlModelviewMatrix
	ConstGlLighting
	ConstGlLight0
	ConstGlAmbient
	ConstGlDiffuse
	ConstGlPosition
	ConstGlTextureEnv
	ConstGlTextureEnvMode
	ConstGlModulate
	ConstGlDecal
	ConstGlPointSmooth

	// glfw
	ConstGlfwFalse
	ConstGlfwTrue
	ConstGlfwPress
	ConstGlfwRelease
	ConstGlfwRepeat
	ConstGlfwKeyUnknown
	ConstGlfwCursor
	ConstGlfwStickyKeys
	ConstGlfwStickyMouseButtons
	ConstGlfwCursorNormal
	ConstGlfwCursorHidden
	ConstGlfwCursorDisabled
	ConstGlfwResizable
	ConstGlfwContextVersionMajor
	ConstGlfwContextVersionMinor
	ConstGlfwOpenglProfile
	ConstGlfwOpenglCoreprofile
	ConstGlfwOpenglForwardCompatible
	ConstGlfwMouseButtonLast
	ConstGlfwMouseButtonLeft
	ConstGlfwMouseButtonRight
	ConstGlfwMouseButtonMiddle

	// gltext
	ConstGltextLeftToRight
	ConstGltextRightToLeft
	ConstGltextTopToBottom

	// os
	ConstOsRunSuccess
	ConstOsRunEmptyCmd
	ConstOsRunPanic
	ConstOsRunStartFailed
	ConstOsRunWaitFailed
	ConstOsRunTimeout

	// cx
	ConstCxSuccess
	ConstCxRuntimeError
	ConstCxPanic
	ConstCxCompilationError
	ConstCxInternalError
	ConstCxAssert
)

// For the parser. These shouldn't be used in the runtime for performance reasons
var (
	ConstNames = map[int]string{}
	ConstCodes = map[string]int{}
	Constants  = map[int]CXConstant{}
)

// AddConstCode ...
func AddConstCode(code int, name string, typ int, value []byte) {
	ConstNames[code] = name
	ConstCodes[name] = code
	Constants[code] = CXConstant{Type: typ, Value: value}
}

// nolint typecheck
func init() {
	/* gl_1_0 */
	AddConstCode(ConstGlDepthBufferBit, "gl.DEPTH_BUFFER_BIT", TypeI32, FromI32(0x00000100))
	AddConstCode(ConstGlStencilBufferBit, "gl.STENCIL_BUFFER_BIT", TypeI32, FromI32(0x00000400))
	AddConstCode(ConstGlColorBufferBit, "gl.COLOR_BUFFER_BIT", TypeI32, FromI32(0x00004000))
	AddConstCode(ConstGlFalse, "gl.FALSE", TypeI32, FromI32(0))
	AddConstCode(ConstGlTrue, "gl.TRUE", TypeI32, FromI32(1))
	AddConstCode(ConstGlPoints, "gl.POINTS", TypeI32, FromI32(0x0000))
	AddConstCode(ConstGlLines, "gl.LINES", TypeI32, FromI32(0x0001))
	AddConstCode(ConstGlLineLoop, "gl.LINE_LOOP", TypeI32, FromI32(0x0002))
	AddConstCode(ConstGlLineStrip, "gl.LINE_STRIP", TypeI32, FromI32(0x0003))
	AddConstCode(ConstGlTriangles, "gl.TRIANGLES", TypeI32, FromI32(0x0004))
	AddConstCode(ConstGlTriangleStrip, "gl.TRIANGLE_STRIP", TypeI32, FromI32(0x0005))
	AddConstCode(ConstGlTriangleFan, "gl.TRIANGLE_FAN", TypeI32, FromI32(0x0006))
	AddConstCode(ConstGlQuads, "gl.QUADS", TypeI32, FromI32(0x0007))
	AddConstCode(ConstGlNever, "gl.NEVER", TypeI32, FromI32(0x0200))
	AddConstCode(ConstGlLess, "gl.LESS", TypeI32, FromI32(0x0201))
	AddConstCode(ConstGlEqual, "gl.EQUAL", TypeI32, FromI32(0x0202))
	AddConstCode(ConstGlLequal, "gl.LEQUAL", TypeI32, FromI32(0x0203))
	AddConstCode(ConstGlGreater, "gl.GREATER", TypeI32, FromI32(0x0204))
	AddConstCode(ConstGlNotequal, "gl.NOTEQUAL", TypeI32, FromI32(0x0205))
	AddConstCode(ConstGlGequal, "gl.GEQUAL", TypeI32, FromI32(0x0206))
	AddConstCode(ConstGlAlways, "gl.ALWAYS", TypeI32, FromI32(0x0207))
	AddConstCode(ConstGlZero, "gl.ZERO", TypeI32, FromI32(0))
	AddConstCode(ConstGlOne, "gl.ONE", TypeI32, FromI32(1))
	AddConstCode(ConstGlSrcAlpha, "gl.SRC_ALPHA", TypeI32, FromI32(0x302))
	AddConstCode(ConstGlOneMinusSrcAlpha, "gl.ONE_MINUS_SRC_ALPHA", TypeI32, FromI32(0x303))
	AddConstCode(ConstGlFront, "gl.FRONT", TypeI32, FromI32(0x404))
	AddConstCode(ConstGlBack, "gl.BACK", TypeI32, FromI32(0x405))
	AddConstCode(ConstGlFrontAndBack, "gl.FRONT_AND_BACK", TypeI32, FromI32(0x408))
	AddConstCode(ConstGlNoError, "gl.NO_ERROR", TypeI32, FromI32(0))
	AddConstCode(ConstGlInvalidEnum, "gl.INVALID_ENUM", TypeI32, FromI32(0x500))
	AddConstCode(ConstGlInvalidValue, "gl.INVALID_VALUE", TypeI32, FromI32(0x501))
	AddConstCode(ConstGlInvalidOperation, "gl.INVALID_OPERATION", TypeI32, FromI32(0x502))
	AddConstCode(ConstGlStackOverflow, "gl.STACK_OVERFLOW", TypeI32, FromI32(0x503))
	AddConstCode(ConstGlStackUnderflow, "gl.STACK_UNDERFLOW", TypeI32, FromI32(0x504))
	AddConstCode(ConstGlOutOfMemory, "gl.OUT_OF_MEMORY", TypeI32, FromI32(0x505))
	AddConstCode(ConstGlLineSmooth, "gl.LINE_SMOOTH", TypeI32, FromI32(0x0B20))
	AddConstCode(ConstGlPolygonSmooth, "gl.POLYGON_SMOOTH", TypeI32, FromI32(0x0B41))
	AddConstCode(ConstGlCullFace, "gl.CULL_FACE", TypeI32, FromI32(0x0B44))
	AddConstCode(ConstGlDepthRange, "gl.DEPTH_RANGE", TypeI32, FromI32(0x0B70))
	AddConstCode(ConstGlDepthTest, "gl.DEPTH_TEST", TypeI32, FromI32(0x0B71))
	AddConstCode(ConstGlDepthWritemask, "gl.DEPTH_WRITEMASK", TypeI32, FromI32(0x0B72))
	AddConstCode(ConstGlDepthClearValue, "gl.DEPTH_CLEAR_VALUE", TypeI32, FromI32(0x0B73))
	AddConstCode(ConstGlDepthFunc, "gl.DEPTH_FUNC", TypeI32, FromI32(0x0B74))
	AddConstCode(ConstGlStencilTest, "gl.STENCIL_TEST", TypeI32, FromI32(0x0B90))
	AddConstCode(ConstGlStencilClearValue, "gl.STENCIL_CLEAR_VALUE", TypeI32, FromI32(0x0B91))
	AddConstCode(ConstGlStencilFunc, "gl.STENCIL_FUNC", TypeI32, FromI32(0x0B92))
	AddConstCode(ConstGlStencilValueMask, "gl.STENCIL_VALUE_MASK", TypeI32, FromI32(0x0B93))
	AddConstCode(ConstGlStencilFail, "gl.STENCIL_FAIL", TypeI32, FromI32(0x0B94))
	AddConstCode(ConstGlStencilPassDepthFail, "gl.STENCIL_PASS_DEPTH_FAIL", TypeI32, FromI32(0x0B95))
	AddConstCode(ConstGlStencilPassDepthPass, "gl.STENCIL_PASS_DEPTH_PASS", TypeI32, FromI32(0x0B96))
	AddConstCode(ConstGlStencilRef, "gl.STENCIL_REF", TypeI32, FromI32(0x0B97))
	AddConstCode(ConstGlStencilWritemask, "gl.STENCIL_WRITE_MASK", TypeI32, FromI32(0x0B98))
	AddConstCode(ConstGlDither, "gl.DITHER", TypeI32, FromI32(0x0BD0))
	AddConstCode(ConstGlBlend, "gl.BLEND", TypeI32, FromI32(0x0BE2))
	AddConstCode(ConstGlScissorTest, "gl.SCISSOR_TEST", TypeI32, FromI32(0x0C11))
	AddConstCode(ConstGlPolygonSmoothHint, "gl.POLYGON_SMOOTH_HINT", TypeI32, FromI32(0x0C53))
	AddConstCode(ConstGlTexture2d, "gl.TEXTURE_2D", TypeI32, FromI32(0x0DE1))
	AddConstCode(ConstGlTextureWidth, "gl.TEXTURE_WIDTH", TypeI32, FromI32(0x1000))
	AddConstCode(ConstGlTextureHeight, "gl.TEXTURE_HEIGHT", TypeI32, FromI32(0x1001))
	AddConstCode(ConstGlDontCare, "gl.DONT_CARE", TypeI32, FromI32(0x1100))
	AddConstCode(ConstGlUnsignedByte, "gl.UNSIGNED_BYTE", TypeI32, FromI32(0x1401))
	AddConstCode(ConstGlFloat, "gl.FLOAT", TypeI32, FromI32(0x1406))
	AddConstCode(ConstGlInvert, "gl.INVERT", TypeI32, FromI32(0x150A))
	AddConstCode(ConstGlTexture, "gl.TEXTURE", TypeI32, FromI32(0x1702))
	AddConstCode(ConstGlColor, "gl.COLOR", TypeI32, FromI32(0x1800))
	AddConstCode(ConstGlDepth, "gl.DEPTH", TypeI32, FromI32(0x1801))
	AddConstCode(ConstGlStencil, "gl.STENCIL", TypeI32, FromI32(0x1802))
	AddConstCode(ConstGlStencilIndex, "gl.STENCIL_INDEX", TypeI32, FromI32(0x1901))
	AddConstCode(ConstGlDepthComponent, "gl.DEPTH_COMPONENT", TypeI32, FromI32(0x1902))
	AddConstCode(ConstGlRgba, "gl.RGBA", TypeI32, FromI32(0x1908))
	AddConstCode(ConstGlKeep, "gl.KEEP", TypeI32, FromI32(0x1E00))
	AddConstCode(ConstGlReplace, "gl.REPLACE", TypeI32, FromI32(0x1E01))
	AddConstCode(ConstGlIncr, "gl.INCR", TypeI32, FromI32(0x1E02))
	AddConstCode(ConstGlDecr, "gl.DECR", TypeI32, FromI32(0x1E03))
	AddConstCode(ConstGlNearest, "gl.NEAREST", TypeI32, FromI32(0x2600))
	AddConstCode(ConstGlLinear, "gl.LINEAR", TypeI32, FromI32(0x2601))
	AddConstCode(ConstGlNearestMipmapNearest, "gl.NEAREST_MIPMAP_NEAREST", TypeI32, FromI32(0x2700))
	AddConstCode(ConstGlLinearMipmapNearest, "gl.LINEAR_MIPMAP_NEAREST", TypeI32, FromI32(0x2701))
	AddConstCode(ConstGlNearestMipmapLinear, "gl.NEAREST_MIPMAP_LINEAR", TypeI32, FromI32(0x2702))
	AddConstCode(ConstGlLinearMipmapLinear, "gl.LINEAR_MIPMAP_LINEAR", TypeI32, FromI32(0x2703))
	AddConstCode(ConstGlTextureMagFilter, "gl.TEXTURE_MAG_FILTER", TypeI32, FromI32(0x2800))
	AddConstCode(ConstGlTextureMinFilter, "gl.TEXTURE_MIN_FILTER", TypeI32, FromI32(0x2801))
	AddConstCode(ConstGlTextureWrapS, "gl.TEXTURE_WRAP_S", TypeI32, FromI32(0x2802))
	AddConstCode(ConstGlTextureWrapT, "gl.TEXTURE_WRAP_T", TypeI32, FromI32(0x2803))
	AddConstCode(ConstGlRepeat, "gl.REPEAT", TypeI32, FromI32(0x2901))

	// gl_1_1
	AddConstCode(ConstGlRgba8, "gl.RGBA8", TypeI32, FromI32(0x8058))
	AddConstCode(ConstGlVertexArray, "gl.VERTEX_ARRAY", TypeI32, FromI32(0x8074))

	// gl_1_2
	AddConstCode(ConstGlTextureWrapR, "gl.TEXTURE_WRAP_R", TypeI32, FromI32(0x8072))
	AddConstCode(ConstGlClampToEdge, "gl.CLAMP_TO_EDGE", TypeI32, FromI32(0x812F))

	// gl_1_3
	AddConstCode(ConstGlTexture0, "gl.TEXTURE0", TypeI32, FromI32(0x84C0))
	AddConstCode(ConstGlMultisampleArb, "gl.MULTISAMPLE_ARB", TypeI32, FromI32(0x809D)) // remove _ARB
	AddConstCode(ConstGlClampToBorder, "gl.CLAMP_TO_BORDER", TypeI32, FromI32(0x812D))

	// gl_1_4
	AddConstCode(ConstGlDepthComponent16, "gl.DEPTH_COMPONENT16", TypeI32, FromI32(0x81A5))
	AddConstCode(ConstGlDepthComponent24, "gl.DEPTH_COMPONENT24", TypeI32, FromI32(0x81A6))
	AddConstCode(ConstGlDepthComponent32, "gl.DEPTH_COMPONENT32", TypeI32, FromI32(0x81A7))
	AddConstCode(ConstGlMirroredRepeat, "gl.MIRRORED_REPEAT", TypeI32, FromI32(0x8370))
	AddConstCode(ConstGlIncrWrap, "gl.INCR_WRAP", TypeI32, FromI32(0x8507))
	AddConstCode(ConstGlDecrWrap, "gl.DECR_WRAP", TypeI32, FromI32(0x8508))

	// gl_1_5
	AddConstCode(ConstGlArrayBuffer, "gl.ARRAY_BUFFER", TypeI32, FromI32(0x8892))
	AddConstCode(ConstGlStreamDraw, "gl.STREAM_DRAW", TypeI32, FromI32(0x88E0))
	AddConstCode(ConstGlStreamRead, "gl.STREAM_READ", TypeI32, FromI32(0x88E1))
	AddConstCode(ConstGlStreamCopy, "gl.STREAM_COPY", TypeI32, FromI32(0x88E2))
	AddConstCode(ConstGlStaticDraw, "gl.STATIC_DRAW", TypeI32, FromI32(0x88E4))
	AddConstCode(ConstGlStaticRead, "gl.STATIC_READ", TypeI32, FromI32(0x88E5))
	AddConstCode(ConstGlStaticCopy, "gl.STATIC_COPY", TypeI32, FromI32(0x88E6))
	AddConstCode(ConstGlDynamicDraw, "gl.DYNAMIC_DRAW", TypeI32, FromI32(0x88E8))
	AddConstCode(ConstGlDynamicRead, "gl.DYNAMIC_READ", TypeI32, FromI32(0x88E9))
	AddConstCode(ConstGlDynamicCopy, "gl.DYNAMIC_COPY", TypeI32, FromI32(0x88EA))

	// gl_2_0
	AddConstCode(ConstGlStencilBackFunc, "gl.STENCIL_BACK_FUNC", TypeI32, FromI32(0x8800))
	AddConstCode(ConstGlStencilBackFail, "gl.STENCIL_BACK_FAIL", TypeI32, FromI32(0x8801))
	AddConstCode(ConstGlStencilBackPassDepthFail, "gl.STENCIL_BACK_PASS_DEPTH_FAIL", TypeI32, FromI32(0x8802))
	AddConstCode(ConstGlStencilBackPassDepthPass, "gl.STENCIL_BACK_PASS_DEPTH_PASS", TypeI32, FromI32(0x8803))
	AddConstCode(ConstGlFragmentShader, "gl.FRAGMENT_SHADER", TypeI32, FromI32(0x8B30))
	AddConstCode(ConstGlVertexShader, "gl.VERTEX_SHADER", TypeI32, FromI32(0x8B31))
	AddConstCode(ConstGlStencilBackRef, "gl.STENCIL_BACK_REF", TypeI32, FromI32(0x8CA3))
	AddConstCode(ConstGlStencilBackValueMask, "gl.STENCIL_BACK_VALUE_MASK", TypeI32, FromI32(0x8CA4))
	AddConstCode(ConstGlStencilBackWritemask, "gl.STENCIL_BACK_WRITEMASK", TypeI32, FromI32(0x8CA5))

	// gl_3_0
	AddConstCode(ConstGlDepthComponent32f, "gl.DEPTH_COMPONENT32F", TypeI32, FromI32(0x8CAC))
	AddConstCode(ConstGlDepth32fStencil8, "gl.DEPTH32F_STENCIL8", TypeI32, FromI32(0x8CAD))
	AddConstCode(ConstGlFramebufferUndefined, "gl.FRAMEBUFFER_UNDEFINED", TypeI32, FromI32(0x8219))
	AddConstCode(ConstGlDepthStencilAttachment, "gl.DEPTH_STENCIL_ATTACHMENT", TypeI32, FromI32(0x821A))
	AddConstCode(ConstGlDepthStencil, "gl.DEPTH_STENCIL", TypeI32, FromI32(0x84F9))
	AddConstCode(ConstGlDepth24Stencil8, "gl.DEPTH24_STENCIL8", TypeI32, FromI32(0x88F0))
	AddConstCode(ConstGlFramebufferComplete, "gl.FRAMEBUFFER_COMPLETE", TypeI32, FromI32(0x8CD5))
	AddConstCode(ConstGlFramebufferIncompleteAttachment, "gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT", TypeI32, FromI32(0x8CD6))
	AddConstCode(ConstGlFramebufferIncompleteMissingAttachment, "gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT", TypeI32, FromI32(0x8CD7))
	AddConstCode(ConstGlFramebufferIncompleteDrawBuffer, "gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER", TypeI32, FromI32(0x8CDB))
	AddConstCode(ConstGlFramebufferIncompleteReadBuffer, "gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER", TypeI32, FromI32(0x8CDC))
	AddConstCode(ConstGlFramebufferUnsupported, "gl.FRAMEBUFFER_UNSUPPORTED", TypeI32, FromI32(0x8CDD))
	AddConstCode(ConstGlColorAttachment0, "gl.COLOR_ATTACHMENT0", TypeI32, FromI32(0x8CE0))
	AddConstCode(ConstGlDepthAttachment, "gl.DEPTH_ATTACHMENT", TypeI32, FromI32(0x8D00))
	AddConstCode(ConstGlStencilAttachment, "gl.STENCIL_ATTACHMENT", TypeI32, FromI32(0x8D20))
	AddConstCode(ConstGlFramebuffer, "gl.FRAMEBUFFER", TypeI32, FromI32(0x8D40))
	AddConstCode(ConstGlRenderbuffer, "gl.RENDERBUFFER", TypeI32, FromI32(0x8D41))
	AddConstCode(ConstGlFramebufferIncompleteMultisample, "gl.FRAMEBUFFER_INCOMPLETE_MULTISAMPLE", TypeI32, FromI32(0x8D56))
	AddConstCode(ConstGlStencilIndex1, "gl.STENCIL_INDEX1", TypeI32, FromI32(0x8D46))
	AddConstCode(ConstGlStencilIndex4, "gl.STENCIL_INDEX4", TypeI32, FromI32(0x8D47))
	AddConstCode(ConstGlStencilIndex8, "gl.STENCIL_INDEX8", TypeI32, FromI32(0x8D48))
	AddConstCode(ConstGlStencilIndex16, "gl.STENCIL_INDEX16", TypeI32, FromI32(0x8D49))

	// gl_3_2
	AddConstCode(ConstGlFramebufferIncompleteLayerTargets, "gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS", TypeI32, FromI32(0x8DA8))

	// gl_4_4
	AddConstCode(ConstGlMirrorClampToEdge, "gl.MIRROR_CLAMP_TO_EDGE", TypeI32, FromI32(0x8743))

	// Fixed pipeline. Deprecated ?
	AddConstCode(ConstGlPolygon, "gl.POLYGON", TypeI32, FromI32(9))
	AddConstCode(ConstGlModelview, "gl.MODELVIEW", TypeI32, FromI32(5888))
	AddConstCode(ConstGlProjection, "gl.PROJECTION", TypeI32, FromI32(5889))
	AddConstCode(ConstGlModelviewMatrix, "gl.MODELVIEW_MATRIX", TypeI32, FromI32(2982))
	AddConstCode(ConstGlLighting, "gl.LIGHTING", TypeI32, FromI32(2896))
	AddConstCode(ConstGlLight0, "gl.LIGHT0", TypeI32, FromI32(16384))
	AddConstCode(ConstGlAmbient, "gl.AMBIENT", TypeI32, FromI32(4608))
	AddConstCode(ConstGlDiffuse, "gl.DIFFUSE", TypeI32, FromI32(4609))
	AddConstCode(ConstGlPosition, "gl.POSITION", TypeI32, FromI32(4611))
	AddConstCode(ConstGlTextureEnv, "gl.TEXTURE_ENV", TypeI32, FromI32(8960))
	AddConstCode(ConstGlTextureEnvMode, "gl.TEXTURE_ENV_MODE", TypeI32, FromI32(8704))
	AddConstCode(ConstGlModulate, "gl.MODULATE", TypeI32, FromI32(8448))
	AddConstCode(ConstGlDecal, "gl.DECAL", TypeI32, FromI32(8449))
	AddConstCode(ConstGlPointSmooth, "gl.POINT_SMOOTH", TypeI32, FromI32(2832))

	// glfw
	AddConstCode(ConstGlfwFalse, "glfw.False", TypeI32, FromI32(0))
	AddConstCode(ConstGlfwTrue, "glfw.True", TypeI32, FromI32(1))
	AddConstCode(ConstGlfwPress, "glfw.Press", TypeI32, FromI32(1))
	AddConstCode(ConstGlfwRelease, "glfw.Release", TypeI32, FromI32(0))
	AddConstCode(ConstGlfwRepeat, "glfw.Repeat", TypeI32, FromI32(2))
	AddConstCode(ConstGlfwKeyUnknown, "glfw.KeyUnknown", TypeI32, FromI32(-1))
	AddConstCode(ConstGlfwCursor, "glfw.Cursor", TypeI32, FromI32(208897))
	AddConstCode(ConstGlfwStickyKeys, "glfw.StickyKeys", TypeI32, FromI32(208898))
	AddConstCode(ConstGlfwStickyMouseButtons, "glfw.StickyMouseButtons", TypeI32, FromI32(208899))
	AddConstCode(ConstGlfwCursorNormal, "glfw.CursorNormal", TypeI32, FromI32(212993))
	AddConstCode(ConstGlfwCursorHidden, "glfw.CursorHidden", TypeI32, FromI32(212994))
	AddConstCode(ConstGlfwCursorDisabled, "glfw.CursorDisabled", TypeI32, FromI32(212995))
	AddConstCode(ConstGlfwResizable, "glfw.Resizable", TypeI32, FromI32(131075))
	AddConstCode(ConstGlfwContextVersionMajor, "glfw.ContextVersionMajor", TypeI32, FromI32(139266))
	AddConstCode(ConstGlfwContextVersionMinor, "glfw.ContextVersionMinor", TypeI32, FromI32(139267))
	AddConstCode(ConstGlfwOpenglProfile, "glfw.Opengl.Profile", TypeI32, FromI32(139272))
	AddConstCode(ConstGlfwOpenglCoreprofile, "glfw.Opengl.Coreprofile", TypeI32, FromI32(204801))
	AddConstCode(ConstGlfwOpenglForwardCompatible, "glfw.Opengl.ForwardCompatible", TypeI32, FromI32(139270))
	AddConstCode(ConstGlfwMouseButtonLast, "glfw.MouseButtonLast", TypeI32, FromI32(7))
	AddConstCode(ConstGlfwMouseButtonLeft, "glfw.MouseButtonLeft", TypeI32, FromI32(0))
	AddConstCode(ConstGlfwMouseButtonRight, "glfw.MouseButtonRight", TypeI32, FromI32(1))
	AddConstCode(ConstGlfwMouseButtonMiddle, "glfw.MouseButtonMiddle", TypeI32, FromI32(2))

	// gltext
	AddConstCode(ConstGltextLeftToRight, "gltext.LeftToRight", TypeI32, FromI32(0))
	AddConstCode(ConstGltextRightToLeft, "gltext.RightToLeft", TypeI32, FromI32(1))
	AddConstCode(ConstGltextTopToBottom, "gltext.TopToBottom", TypeI32, FromI32(2))

	// os
	AddConstCode(ConstOsRunSuccess, "os.RUN_SUCCESS", TypeI32, FromI32(OS_RUN_SUCCESS))
	AddConstCode(ConstOsRunEmptyCmd, "os.RUN_EMPTY_CMD", TypeI32, FromI32(OS_RUN_EMPTY_CMD))
	AddConstCode(ConstOsRunPanic, "os.RUN_PANIC", TypeI32, FromI32(OS_RUN_PANIC))
	AddConstCode(ConstOsRunStartFailed, "os.RUN_START_FAILED", TypeI32, FromI32(OS_RUN_START_FAILED))
	AddConstCode(ConstOsRunWaitFailed, "os.RUN_WAIT_FAILED", TypeI32, FromI32(OS_RUN_WAIT_FAILED))
	AddConstCode(ConstOsRunTimeout, "os.RUN_TIMEOUT", TypeI32, FromI32(OS_RUN_TIMEOUT))

	// cx
	AddConstCode(ConstCxSuccess, "cx.SUCCESS", TypeI32, FromI32(CxSuccess))
	AddConstCode(ConstCxRuntimeError, "cx.RUNTIME_ERROR", TypeI32, FromI32(CxRuntimeError))
	AddConstCode(ConstCxPanic, "cx.PANIC", TypeI32, FromI32(CxPanic))
	AddConstCode(ConstCxCompilationError, "cx.COMPILATION_ERROR", TypeI32, FromI32(CxCompilationError))
	AddConstCode(ConstCxInternalError, "cx.INTERNAL_ERROR", TypeI32, FromI32(CxInternalError))
	AddConstCode(ConstCxAssert, "cx.ASSERT", TypeI32, FromI32(CxAssert))
}
