%{
	package main
	import (
		// "fmt"
		"strconv"
		"github.com/skycoin/skycoin/src/cipher/encoder"
		. "github.com/skycoin/cx/cx"
		. "github.com/skycoin/cx/cxgo/actions"
	)

	// var PRGRM = MakeProgram(CALLSTACK_SIZE, STACK_SIZE, INIT_HEAP_SIZE)
	
%}

%union{
	i int
	byt byte
	i32 int32
	i64 int64
	f32 float32
	f64 float64
	tok string
	bool bool
	string string
	stringA []string

	line int

	argument *CXArgument
	arguments []*CXArgument

	expression *CXExpression
	expressions []*CXExpression

	SelectStatement SelectStatement
	SelectStatements []SelectStatement

	arrayArguments [][]*CXExpression

        function *CXFunction
}

%token  <byt>           BYTE_LITERAL
%token  <bool>          BOOLEAN_LITERAL
%token  <i32>           INT_LITERAL
%token  <i64>           LONG_LITERAL
%token  <f32>           FLOAT_LITERAL
%token  <f64>           DOUBLE_LITERAL
%token  <tok>           FUNC OP LPAREN RPAREN LBRACE RBRACE LBRACK RBRACK IDENTIFIER
                        VAR COMMA PERIOD COMMENT STRING_LITERAL PACKAGE IF ELSE FOR TYPSTRUCT STRUCT
                        SEMICOLON NEWLINE
                        ASSIGN CASSIGN IMPORT RETURN GOTO GT_OP LT_OP GTEQ_OP LTEQ_OP EQUAL COLON NEW
                        EQUALWORD GTHANWORD LTHANWORD
                        GTHANEQ LTHANEQ UNEQUAL AND OR
                        ADD_OP SUB_OP MUL_OP DIV_OP MOD_OP REF_OP NEG_OP AFFVAR
                        PLUSPLUS MINUSMINUS REMAINDER LEFTSHIFT RIGHTSHIFT EXP
                        NOT
                        BITXOR_OP BITOR_OP BITCLEAR_OP
                        PLUSEQ MINUSEQ MULTEQ DIVEQ REMAINDEREQ EXPEQ
                        LEFTSHIFTEQ RIGHTSHIFTEQ BITANDEQ BITXOREQ BITOREQ

                        DEC_OP INC_OP PTR_OP LEFT_OP RIGHT_OP
                        GE_OP LE_OP EQ_OP NE_OP AND_OP OR_OP
                        ADD_ASSIGN AND_ASSIGN LEFT_ASSIGN MOD_ASSIGN
                        MUL_ASSIGN DIV_ASSIGN OR_ASSIGN RIGHT_ASSIGN
                        SUB_ASSIGN XOR_ASSIGN
                        BOOL BYTE F32 F64
                        I8 I16 I32 I64
                        STR
                        UI8 UI16 UI32 UI64
                        UNION ENUM CONST CASE DEFAULT SWITCH BREAK CONTINUE
                        TYPE
                        
                        /* Types */
                        BASICTYPE
                        /* Selectors */
                        SPACKAGE SSTRUCT SFUNC
                        /* Removers */
                        REM DEF EXPR FIELD INPUT OUTPUT CLAUSES OBJECT OBJECTS
                        /* Stepping */
                        STEP PSTEP TSTEP
                        /* Debugging */
                        DSTACK DPROGRAM DSTATE
                        /* Affordances */
                        AFF CAFF TAG INFER VALUE
                        /* Pointers */
                        ADDR

%type   <tok>           after_period
%type   <tok>           unary_operator
%type   <tok>           assignment_operator
%type   <i>             type_specifier
%type   <argument>      declaration_specifiers
%type   <argument>      declarator
%type   <argument>      direct_declarator
%type   <argument>      parameter_declaration
%type   <arguments>     parameter_type_list
%type   <arguments>     function_parameters
%type   <arguments>     parameter_list
%type   <arguments>     fields
%type   <arguments>     struct_fields

/* %type   <stringA>       package_identifier */
                                
%type   <expressions>   assignment_expression
%type   <expressions>   constant_expression
%type   <expressions>   conditional_expression
%type   <expressions>   logical_or_expression
%type   <expressions>   logical_and_expression
%type   <expressions>   exclusive_or_expression
%type   <expressions>   inclusive_or_expression
%type   <expressions>   and_expression
%type   <expressions>   equality_expression
%type   <expressions>   relational_expression
%type   <expressions>   shift_expression
%type   <expressions>   additive_expression
%type   <expressions>   multiplicative_expression
%type   <expressions>   unary_expression
%type   <expressions>   argument_expression_list
%type   <expressions>   postfix_expression
%type   <expressions>   primary_expression

%type   <expressions>   struct_literal_expression
                        
%type   <expressions>   array_literal_expression_list
%type   <expressions>   array_literal_expression

%type   <expressions>   slice_literal_expression_list
%type   <expressions>   slice_literal_expression

%type   <expressions>   selector

%type   <expressions>   struct_literal_fields
%type   <SelectStatement>   elseif
%type   <SelectStatements>   elseif_list

%type   <expressions>   declaration
//                      %type   <expressions>   init_declarator_list
//                      %type   <expressions>   init_declarator

%type   <expressions>   initializer

%type   <expressions>   expression
%type   <expressions>   block_item
%type   <expressions>   block_item_list
%type   <expressions>   compound_statement
%type   <expressions>   else_statement
%type   <expressions>   labeled_statement
%type   <expressions>   expression_statement
%type   <expressions>   selection_statement
%type   <expressions>   iteration_statement
%type   <expressions>   jump_statement
%type   <expressions>   statement

%type   <function>      function_header

//                      %type   <stringA>       infer_action, infer_actions
%type   <string>        infer_action_arg
%type   <stringA>       infer_action, infer_actions
%type   <expressions>   infer_clauses
                        
                        // for struct literals
%right                   IDENTIFIER LBRACE
// %right                  IDENTIFIER
                        
/* %start                  translation_unit */
%%



translation_unit:
                external_declaration
        |       translation_unit external_declaration
        ;

external_declaration:
                package_declaration
        |       global_declaration
        |       function_declaration
        |       import_declaration
        |       struct_declaration
                
        |       stepping
        |       selector
        |       debugging
        ;

debugging:      
                DPROGRAM
                {
			PRGRM.PrintProgram()
                }
        ;

stepping:       TSTEP INT_LITERAL INT_LITERAL
                {
			Stepping(int($2), int($3), true)
                }
        |       STEP INT_LITERAL
                {
			Stepping(int($2), 0, false)
                }
        ;

selector:
                SPACKAGE IDENTIFIER SEMICOLON
                {
			$<string>$ = Selector($2, SELECT_TYP_PKG)
                }
        |       SFUNC IDENTIFIER
                {
			$<string>$ = Selector($2, SELECT_TYP_FUNC)
                }
                compound_statement
                {
			if len($4) > 0 {
				if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
					if fn, err := PRGRM.GetFunction($<string>3, pkg.Name); err == nil {
						for _, expr := range $4 {
							fn.AddExpression(expr)
						}
						FunctionDeclaration(fn, nil, nil, nil)
					} else {
						panic(err)
					}
				} else {
					panic(err)
				}
			}
			
			// if $<bool>4 {
				
			// 	if _, err := PRGRM.SelectFunction($<string>3); err == nil {
			// 	}
			// }
                }
        |       SSTRUCT IDENTIFIER SEMICOLON
                {
			$<string>$ = Selector($2, SELECT_TYP_STRCT)
                }
        |       SSTRUCT IDENTIFIER
                {
			$<string>$ = Selector($2, SELECT_TYP_STRCT)
                }
                struct_fields
                {
			if len($4) > 0 {
				if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
					if strct, err := PRGRM.GetStruct($<string>3, pkg.Name); err == nil {
						for _, fld := range $4 {
							strct.AddField(fld)
						}
						// FunctionDeclaration(fn, nil, nil, nil)
					} else {
						panic(err)
					}
				} else {
					panic(err)
				}
			}
			/* if $<bool>4 { */
			/* 	if _, err := PRGRM.SelectStruct($<string>3); err == nil { */
			/* 		//fmt.Println(fmt.Sprintf("== Changed to struct '%s' ==", strct.Name)) */
			/* 	} */
			/* } */
                }
        ;

global_declaration:
                VAR declarator declaration_specifiers SEMICOLON
                {
			DeclareGlobal($2, $3, nil, false)
                }
        |       VAR declarator declaration_specifiers ASSIGN initializer SEMICOLON
                {
			DeclareGlobal($2, $3, $5, true)
                }
                ;

struct_declaration:
                TYPE IDENTIFIER STRUCT struct_fields
                {
			DeclareStruct($2, $4)
                }
                ;

struct_fields:
                LBRACE RBRACE SEMICOLON
                { $$ = nil }
        |       LBRACE fields RBRACE SEMICOLON
                { $$ = $2 }
        ;

fields:         parameter_declaration SEMICOLON
                {
			$$ = []*CXArgument{$1}
                }
        |       fields parameter_declaration SEMICOLON
                {
			$$ = append($1, $2)
                }
        ;

package_declaration:
                PACKAGE IDENTIFIER SEMICOLON
                {
			DeclarePackage($2)
                }
                ;

import_declaration:
                IMPORT STRING_LITERAL SEMICOLON
                {
			// DeclareImport($2)
                }
        ;

function_header:
                FUNC IDENTIFIER
                {
			yylval.line = 0
			$$ = FunctionHeader($2, nil, false)
			InFn = true
                }
        |       FUNC LPAREN parameter_type_list RPAREN IDENTIFIER
                {
			$$ = FunctionHeader($5, $3, true)
			InFn = true
                }
        ;

function_parameters:
                LPAREN RPAREN
                { $$ = nil }
        |       LPAREN parameter_type_list RPAREN
                { $$ = $2 }
                ;

function_declaration:
                function_header function_parameters compound_statement
                {
			FunctionDeclaration($1, $2, nil, $3)
			InFn = false
                }
        |       function_header function_parameters function_parameters compound_statement
                {
			FunctionDeclaration($1, $2, $3, $4)
			InFn = false
                }
        ;

parameter_type_list:
		parameter_list
                ;

parameter_list:
                parameter_declaration
                {
			$$ = []*CXArgument{$1}
                }
	|       parameter_list COMMA parameter_declaration
                {
			$$ = append($1, $3)
                }
                ;

parameter_declaration:
                declarator declaration_specifiers
                {
			$2.Name = $1.Name
			$2.Package = $1.Package
			$$ = $2
                }
                ;

declarator:     direct_declarator
                ;

direct_declarator:
                IDENTIFIER
                {
			if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
				arg := MakeArgument("", CurrentFile, LineNo)
                                arg.AddType(TypeNames[TypeUndefined])
				arg.Name = $1
				arg.Package = pkg
				$$ = arg
			} else {
				panic(err)
			}
                }
	|       LPAREN declarator RPAREN
                { $$ = $2 }
	// |       direct_declarator '[' ']'
        //         {
	// 		$1.IsArray = true
	// 		$$ = $1
        //         }
        //	|direct_declarator '[' MUL_OP ']'
        //              	|direct_declarator '[' type_qualifier_list MUL_OP ']'
        //              	|direct_declarator '[' type_qualifier_list assignment_expression ']'
        //              	|direct_declarator '[' type_qualifier_list ']'
        //              	|direct_declarator '[' assignment_expression ']'
	// |    direct_declarator LPAREN parameter_type_list RPAREN
	// |    direct_declarator LPAREN RPAREN
	// |    direct_declarator LPAREN identifier_list RPAREN
                ;

// check
/* pointer:        /\* MUL_OP   type_qualifier_list pointer // check *\/ */
/*         /\* |       MUL_OP   type_qualifier_list // check *\/ */
/*         /\* |       MUL_OP   pointer *\/ */
/*         /\* |        *\/MUL_OP */
/*                 ; */

/* type_qualifier_list: */
/*                 type_qualifier */
/* 	|       type_qualifier_list type_qualifier */
/*                 ; */








declaration_specifiers:
                MUL_OP declaration_specifiers
                {
			$$ = DeclarationSpecifiers($2, 0, DerefPointer)
                }
        |       LBRACK INT_LITERAL RBRACK declaration_specifiers
                {
			$$ = DeclarationSpecifiers($4, int($2), DeclArray)
                }
        |       LBRACK RBRACK declaration_specifiers
                {
			$$ = DeclarationSpecifiers($3, 0, DeclSlice)
                }
        |       type_specifier
                {
			$$ = DeclarationSpecifiersBasic($1)
                }
        |       IDENTIFIER
                {
			$$ = DeclarationSpecifiersStruct($1, "", false, CurrentFile, LineNo)
                }
        |       IDENTIFIER PERIOD IDENTIFIER
                {
			$$ = DeclarationSpecifiersStruct($3, $1, true, CurrentFile, LineNo)
                }
	|       type_specifier PERIOD IDENTIFIER
                {
			$$ = DeclarationSpecifiersStruct($3, TypeNames[$1], true, CurrentFile, LineNo)
                }
        /* |       package_identifier */
        /*         { */
	/* 		$$ = DeclarationSpecifiersStruct($1[1], $1[0], true) */
        /*         } */
		/* type_specifier declaration_specifiers */
	/* |       type_specifier */
	/* |       type_qualifier declaration_specifiers */
	/* |       type_qualifier */
                ;

type_specifier:
                AFF
                { $$ = TypeAff }
        |       BOOL
                { $$ = TypeBool }
        |       BYTE
                { $$ = TypeByte }
        |       STR
                { $$ = TypeStr }
        |       F32
                { $$ = TypeF32 }
        |       F64
                { $$ = TypeF64 }
        |       I8
                { $$ = TypeI8 }
        |       I16
                { $$ = TypeI16 }
        |       I32
                { $$ = TypeI32 }
        |       I64
                { $$ = TypeI64 }
        |       UI8
                { $$ = TypeUI8 }
        |       UI16
                { $$ = TypeUI16 }
        |       UI32
                { $$ = TypeUI32 }
        |       UI64
                { $$ = TypeUI64 }
                ;


struct_literal_fields:
                // empty
                { $$ = nil }
        |       IDENTIFIER COLON constant_expression
                {
			if $3[0].IsStructLiteral {
				$$ = StructLiteralAssignment([]*CXExpression{StructLiteralFields($1)}, $3)
			} else {
				$$ = Assignment([]*CXExpression{StructLiteralFields($1)}, "=", $3)
			}
                }
        |       struct_literal_fields COMMA IDENTIFIER COLON constant_expression
                {
			if $5[0].IsStructLiteral {
				$$ = append($1, StructLiteralAssignment([]*CXExpression{StructLiteralFields($3)}, $5)...)
			} else {
				$$ = append($1, Assignment([]*CXExpression{StructLiteralFields($3)}, "=", $5)...)
			}
                }
                ;

array_literal_expression_list:
                assignment_expression
                {
			$1[len($1) - 1].IsArrayLiteral = true
			$$ = $1
                }
	|       array_literal_expression_list COMMA assignment_expression
                {
			$3[len($3) - 1].IsArrayLiteral = true
			$$ = append($1, $3...)
                }
                ;

// expressions
array_literal_expression:
                LBRACK INT_LITERAL RBRACK IDENTIFIER LBRACE array_literal_expression_list RBRACE
                {
			$$ = $6
                }
        |       LBRACK INT_LITERAL RBRACK IDENTIFIER LBRACE RBRACE
                {
			$$ = nil
                }
        |       LBRACK INT_LITERAL RBRACK type_specifier LBRACE array_literal_expression_list RBRACE
                {
			$$ = ArrayLiteralExpression(int($2), $4, $6)
                }
        |       LBRACK INT_LITERAL RBRACK type_specifier LBRACE RBRACE
                {
			$$ = nil
                }
        |       LBRACK INT_LITERAL RBRACK array_literal_expression
                {
			for _, expr := range $4 {
				if expr.Outputs[0].Name == $4[len($4) - 1].Inputs[0].Name {
					expr.Outputs[0].Lengths = append([]int{int($2)}, expr.Outputs[0].Lengths[:len(expr.Outputs[0].Lengths) - 1]...)
					expr.Outputs[0].TotalSize = expr.Outputs[0].Size * TotalLength(expr.Outputs[0].Lengths)
				}
			}

			$$ = $4
                }
                ;







slice_literal_expression_list:
                assignment_expression
                {
			$1[len($1) - 1].IsArrayLiteral = true
			$$ = $1
                }
	|       slice_literal_expression_list COMMA assignment_expression
                {
			$3[len($3) - 1].IsArrayLiteral = true
			$$ = append($1, $3...)
                }
                ;

slice_literal_expression:
                LBRACK RBRACK IDENTIFIER LBRACE slice_literal_expression_list RBRACE
                {
			$$ = $5
                }
        |       LBRACK RBRACK IDENTIFIER LBRACE RBRACE
                {
			$$ = nil
                }
        |       LBRACK RBRACK type_specifier LBRACE slice_literal_expression_list RBRACE
                {
			$$ = SliceLiteralExpression($3, $5)
                }
        |       LBRACK RBRACK type_specifier LBRACE RBRACE
                {
			$$ = nil
                }
        |       LBRACK RBRACK slice_literal_expression
                {
			for _, expr := range $3 {
				if expr.Outputs[0].Name == $3[len($3) - 1].Inputs[0].Name {
					expr.Outputs[0].Lengths = append([]int{0}, expr.Outputs[0].Lengths[:len(expr.Outputs[0].Lengths) - 1]...)
					expr.Outputs[0].TotalSize = expr.Outputs[0].Size * TotalLength(expr.Outputs[0].Lengths)
				}
			}

			$$ = $3
                }
                ;

/* package_identifier: */
/*                 IDENTIFIER PERIOD IDENTIFIER */
/*                 { */
/* 			$$ = []string{$1, $2} */
/*                 } */
/*                 ; */




/* infer_action_arg: */
/*                 MUL_OP GT_OP assignment_expression */
/*                 { */
/* 			if $3[len($3) - 1].Outputs[0].Name != "" { */
/* 				$$ = []string{$1, $3[len($3) - 1].Outputs[0].Name, $2} */
/* 			} else { */
/* 				$$ = []string{$1, strconv.Itoa($3[len($3) - 1].Outputs[0].Offset), $2} */
/* 			} */
			
/*                 } */
/*         |       MUL_OP GT_OP MUL_OP */
/*                 { */
/* 			$$ = []string{$1, $1, $2} */
/*                 } */
/*         ; */

infer_action_arg:
                IDENTIFIER
                {
			$$ = $1
                }
        |       INT_LITERAL
                {
			$$ = strconv.Itoa(int($1))
                }
        ;

infer_action:
                IDENTIFIER LPAREN infer_action_arg COMMA IDENTIFIER RPAREN
		{
			res := append([]string{$3}, $5)
			res = append(res, $1)
			$$ = res
		}
	|	IDENTIFIER LPAREN infer_action_arg RPAREN
		{
			$$ = append([]string{$1}, $3)
		}
	|	IDENTIFIER LPAREN infer_action RPAREN
		{
			$$ = append($3, $1)
		}
	|	IDENTIFIER LPAREN infer_action COMMA infer_action RPAREN
		{
			res := append($3, $5...)
			$$ = append(res, $1)
		}
        ;

infer_actions:
                infer_action SEMICOLON
                {
			$$ = $1
                }
        |       infer_actions infer_action SEMICOLON
                {
			$1 = append($1, $2...)
			$$ = $1
                }
                ;

/* infer_target: */
/*                 IDENTIFIER LPAREN IDENTIFIER RPAREN SEMICOLON */
/*                 { */
/* 			$$ = []string{$3, $1} */
/*                 } */
        ;

/* infer_targets: */
/*                 infer_target */
/*                 { */
/* 			$$ = $1 */
/*                 } */
/*         |       infer_targets infer_target */
/*                 { */
/* 			$1 = append($1, $2...) */
/* 			$$ = $1 */
/*                 } */
/*         ; */

infer_clauses:
                {
			$$ = SliceLiteralExpression(TypeAff, nil)
                }
        |       infer_actions
                {
			var exprs []*CXExpression
			for _, str := range $1 {
				expr := WritePrimary(TypeAff, encoder.Serialize(str), false)
				expr[len(expr) - 1].IsArrayLiteral = true
				exprs = append(exprs, expr...)
			}
			
			$$ = SliceLiteralExpression(TypeAff, exprs)
                }
        /* |       infer_targets */
        /*         { */
	/* 		var exprs []*CXExpression */
	/* 		for _, str := range $1 { */
	/* 			expr := WritePrimary(TypeAff, encoder.Serialize(str), false) */
	/* 			expr[len(expr) - 1].IsArrayLiteral = true */
	/* 			exprs = append(exprs, expr...) */
	/* 		} */
			
	/* 		// $$ = ArrayLiteralExpression(len(exprs), TypeStr, exprs) */
	/* 		$$ = SliceLiteralExpression(TypeAff, exprs) */
        /*         } */
                ;




primary_expression:
                IDENTIFIER
                {
			$$ = PrimaryIdentifier($1)
                }
        /* |       IDENTIFIER LBRACE struct_literal_fields RBRACE */
        /*         { */
	/* 		$$ = PrimaryStructLiteral($1, $3) */
        /*         } */
	|	FUNC LPAREN RPAREN
		{
			$$ = nil
		}
        |       INFER LBRACE infer_clauses RBRACE
                {
			$$ = $3
                }
        |       STRING_LITERAL
                {
			$$ = WritePrimary(TypeStr, encoder.Serialize($1), false)
                }
        |       BOOLEAN_LITERAL
                {
			exprs := WritePrimary(TypeBool, encoder.Serialize($1), false)
			$$ = exprs
                }
        |       BYTE_LITERAL
                {
			$$ = WritePrimary(TypeByte, encoder.Serialize($1), false)
                }
        |       INT_LITERAL
                {
			$$ = WritePrimary(TypeI32, encoder.Serialize($1), false)
                }
        |       FLOAT_LITERAL
                {
			$$ = WritePrimary(TypeF32, encoder.Serialize($1), false)
                }
        |       DOUBLE_LITERAL
                {
			$$ = WritePrimary(TypeF64, encoder.Serialize($1), false)
                }
        |       LONG_LITERAL
                {
			$$ = WritePrimary(TypeI64, encoder.Serialize($1), false)
                }
        |       LPAREN expression RPAREN
                { $$ = $2 }
        |       array_literal_expression
                {
			$$ = $1
                }
        |       slice_literal_expression
                {
			$$ = $1
                }
                ;

after_period:   type_specifier
                {
			$$ = TypeNames[$1]
                }
        |       IDENTIFIER
        ;

postfix_expression:
                primary_expression
	|       postfix_expression LBRACK expression RBRACK
                {
			$$ = PostfixExpressionArray($1, $3)
                }
        |       type_specifier PERIOD after_period
                {
			$$ = PostfixExpressionNative(int($1), $3)
                }
	|       postfix_expression LPAREN RPAREN
                {
			$$ = PostfixExpressionEmptyFunCall($1)
                }
	|       postfix_expression LPAREN argument_expression_list RPAREN
                {
			$$ = PostfixExpressionFunCall($1, $3)
                }
	|       postfix_expression INC_OP
                {
			$$ = PostfixExpressionIncDec($1, true)
                }
        |       postfix_expression DEC_OP
                {
			$$ = PostfixExpressionIncDec($1, false)
                }

        |       postfix_expression PERIOD IDENTIFIER
                {
			PostfixExpressionField($1, $3)
                }
        // |       postfix_expression PERIOD IDENTIFIER LBRACE struct_literal_fields RBRACE
        //         {
	// 		$$ = PrimaryStructLiteralExternal($1[0].Outputs[0].Name, $3, $5)
        //         }
                ;

argument_expression_list:
                assignment_expression
	|       argument_expression_list COMMA assignment_expression
                {
			$$ = append($1, $3...)
                }
                ;

unary_expression:
                postfix_expression
	|       INC_OP unary_expression
                {
			// TODO
			$$ = $2
                }
	|       DEC_OP unary_expression
                {
			// TODO
			$$ = $2
                }
	|       unary_operator unary_expression
                {
			$$ = UnaryExpression($1, $2)
                }
                ;

unary_operator:
                REF_OP
	|       MUL_OP
	|       ADD_OP
	|       SUB_OP
	|       NEG_OP
                ;

multiplicative_expression:
                unary_expression
        |       multiplicative_expression MUL_OP unary_expression
                {
			$$ = ShorthandExpression($1, $3, OP_MUL)
                }
        |       multiplicative_expression DIV_OP unary_expression
                {
			$$ = ShorthandExpression($1, $3, OP_DIV)
                }
        |       multiplicative_expression MOD_OP unary_expression
                {
			$$ = ShorthandExpression($1, $3, OP_MOD)
                }
                ;

additive_expression:
                multiplicative_expression
        |       additive_expression ADD_OP multiplicative_expression
                {
			$$ = ShorthandExpression($1, $3, OP_ADD)
                }
	|       additive_expression SUB_OP multiplicative_expression
                {
			$$ = ShorthandExpression($1, $3, OP_SUB)
                }
                ;

shift_expression:
                additive_expression
        |       shift_expression LEFT_OP additive_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITSHL)
                }
        |       shift_expression RIGHT_OP additive_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITSHR)
                }
        |       shift_expression BITCLEAR_OP additive_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITCLEAR)
                }
                ;

relational_expression:
                shift_expression
        |       relational_expression LT_OP shift_expression
                {
			$$ = ShorthandExpression($1, $3, OP_LT)
                }
        |       relational_expression GT_OP shift_expression
                {
			$$ = ShorthandExpression($1, $3, OP_GT)
                }
        |       relational_expression LTEQ_OP shift_expression
                {
			$$ = ShorthandExpression($1, $3, OP_LTEQ)
                }
        |       relational_expression GTEQ_OP shift_expression
                {
			$$ = ShorthandExpression($1, $3, OP_GTEQ)
                }
                ;

equality_expression:
                relational_expression
        |       equality_expression EQ_OP relational_expression
                {
			$$ = ShorthandExpression($1, $3, OP_EQUAL)
                }
        |       equality_expression NE_OP relational_expression
                {
			$$ = ShorthandExpression($1, $3, OP_UNEQUAL)
                }
                ;

and_expression: equality_expression
        |       and_expression REF_OP equality_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITAND)
                }
                ;

exclusive_or_expression:
                and_expression
        |       exclusive_or_expression BITXOR_OP and_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITXOR)
                }
                ;

inclusive_or_expression:
                exclusive_or_expression
        |       inclusive_or_expression BITOR_OP exclusive_or_expression
                {
			$$ = ShorthandExpression($1, $3, OP_BITOR)
                }
                ;

logical_and_expression:
                inclusive_or_expression
	|       logical_and_expression AND_OP inclusive_or_expression
                {
			$$ = UndefinedTypeOperation($1, $3, Natives[OP_BOOL_AND])
                }
                ;

logical_or_expression:
                logical_and_expression
	|       logical_or_expression OR_OP logical_and_expression
                {
			$$ = UndefinedTypeOperation($1, $3, Natives[OP_BOOL_OR])
                }
                ;

conditional_expression:
                logical_or_expression
	|       logical_or_expression '?' expression COLON conditional_expression
                ;

struct_literal_expression:
                conditional_expression
	|       IDENTIFIER LBRACE struct_literal_fields RBRACE
                {
			$$ = PrimaryStructLiteral($1, $3)
                }
        |       postfix_expression PERIOD IDENTIFIER LBRACE struct_literal_fields RBRACE
                {
			$$ = PrimaryStructLiteralExternal($1[0].Outputs[0].Name, $3, $5)
                }
                ;

assignment_expression:
                /* conditional_expression */
                struct_literal_expression
	|       unary_expression assignment_operator assignment_expression
                {
			if $3[0].IsArrayLiteral {
				if $2 != "=" && $2 != ":=" {
					panic("")
				}
				if $2 == ":=" {
					for _, from := range $3 {
						from.Outputs[0].IsShortDeclaration = true
						from.Outputs[0].PreviouslyDeclared = true
					}
				}
				$$ = ArrayLiteralAssignment($1, $3)
			} else if $3[len($3) - 1].IsStructLiteral {
				if $2 != "=" && $2 != ":=" {
					panic("")
				}
				if $2 == ":=" {
					for _, from := range $3 {
						from.Outputs[0].IsShortDeclaration = true
						from.Outputs[0].PreviouslyDeclared = true
					}
				}
				$$ = StructLiteralAssignment($1, $3)
			} else {
				$$ = Assignment($1, $2, $3)
			}
                }
                ;

assignment_operator:
                ASSIGN
        |       CASSIGN
	|       MUL_ASSIGN
	|       DIV_ASSIGN
	|       MOD_ASSIGN
	|       ADD_ASSIGN
	|       SUB_ASSIGN
	|       LEFT_ASSIGN
	|       RIGHT_ASSIGN
	|       AND_ASSIGN
	|       XOR_ASSIGN
	|       OR_ASSIGN
                ;

expression:     assignment_expression
	|       expression COMMA assignment_expression
                {
			$3[len($3) - 1].Outputs = append($1[len($1) - 1].Outputs, $3[len($3) - 1].Outputs...)
			$$ = $3
                }
                ;

constant_expression:
                conditional_expression
                ;

declaration:
                VAR declarator declaration_specifiers SEMICOLON
                {
			$$ = DeclareLocal($2, $3, nil, false)
                }
        |       VAR declarator declaration_specifiers ASSIGN initializer SEMICOLON
                {
			$$ = DeclareLocal($2, $3, $5, true)
                }
                ;

initializer:    assignment_expression
                ;

// statements
statement:      labeled_statement
	|       compound_statement
	|       expression_statement
	|       selection_statement
	|       iteration_statement
        |       selector
        |       debugging
                { $$ = nil }
	|       jump_statement
                ;

labeled_statement:
                IDENTIFIER COLON block_item
                {
			// it has to be the first expression so all the nested expressions are executed
			// instead of only executing the last one
			// UPDATE: I need to label all expressions. `goto` will jump to first occurrance anyway, so no problem
			// I need this behavior for affordances
			for _, expr := range $3 {
				expr.Label = $1
			}

			$$ = $3
                }
	|       CASE constant_expression COLON statement
                { $$ = nil }
	|       DEFAULT COLON statement
                { $$ = nil }
                ;

compound_statement:
                LBRACE RBRACE SEMICOLON
                { $$ = nil }
	|       LBRACE block_item_list RBRACE SEMICOLON
                {
                    $$ = $2
                }
                ;

block_item_list:
                block_item
	|       block_item_list block_item
                {
			$$ = append($1, $2...)
                }
                ;

block_item:     declaration
        |       statement
        |       stepping
                { $$ = nil }
        /* |       debugging */
        /*         { $$ = nil } */
        /* |       selector */
        /*         { $$ = nil} */
                ;

expression_statement:
                SEMICOLON
                { $$ = nil }
	|       expression SEMICOLON
                {
			if $1[len($1) - 1].Operator == nil && !$1[len($1) - 1].IsMethodCall {
				$$ = nil
			} else {
				$$ = $1
			}
			// $$ = $1
                }
                ;

selection_statement:
                IF conditional_expression LBRACE block_item_list RBRACE elseif_list else_statement SEMICOLON
                {
			$$ = SelectionStatement($2, $4, $6, $7, SEL_ELSEIFELSE)
                }
        |       IF conditional_expression LBRACE block_item_list RBRACE else_statement SEMICOLON
                {
			$$ = SelectionExpressions($2, $4, $6)
                }
        |       IF conditional_expression LBRACE RBRACE else_statement SEMICOLON
                {
			$$ = SelectionExpressions($2, nil, $5)
                }
        |       IF conditional_expression LBRACE block_item_list RBRACE elseif_list SEMICOLON
                {
			$$ = SelectionStatement($2, $4, $6, nil, SEL_ELSEIF)
                }
        |       IF conditional_expression LBRACE RBRACE elseif_list SEMICOLON
                {
			//
			$$ = SelectionStatement($2, nil, $5, nil, SEL_ELSEIF)
                }
        |       IF conditional_expression compound_statement
                {
			$$ = SelectionExpressions($2, $3, nil)
                }
	|       SWITCH LPAREN expression RPAREN statement
                { $$ = nil }
                ;

elseif:         ELSE IF expression LBRACE block_item_list RBRACE
                {
			$$ = SelectStatement{
				Condition: $3,
				Then: $5,
			}
                }
                ;

elseif_list:    elseif
                {
			$$ = []SelectStatement{$1}
                }
        |       elseif_list elseif
                {
			$$ = append($1, $2)
                }
        ;

else_statement:
                ELSE LBRACE block_item_list RBRACE
                {
			$$ = $3
                }
        ;

iteration_statement:
                FOR expression compound_statement
                {
			$$ = IterationExpressions(nil, $2, nil, $3)
                }
        |       FOR expression_statement expression_statement compound_statement
                {			
			$$ = IterationExpressions($2, $3, nil, $4)
                }
        |       FOR expression_statement expression_statement expression compound_statement
                {
			$$ = IterationExpressions($2, $3, $4, $5)
                }
                ;

jump_statement: GOTO IDENTIFIER SEMICOLON
                {
			if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
				expr := MakeExpression(Natives[OP_JMP], CurrentFile, LineNo)
				expr.Package = pkg
				expr.Label = $2

				arg := MakeArgument("", CurrentFile, LineNo).AddType("bool")
				arg.Package = pkg

				expr.AddInput(arg)
					
				$$ = []*CXExpression{expr}
			} else {
				panic(err)
			}
                }
	|       CONTINUE SEMICOLON
		{
			$$ = ContinueExpressions()
		}
	|       BREAK SEMICOLON
		{
			$$ = BreakExpressions()
		}
	|       RETURN SEMICOLON
                {
			if pkg, err := PRGRM.GetCurrentPackage(); err == nil {
				expr := MakeExpression(Natives[OP_JMP], CurrentFile, LineNo)

				// simulating a label so it gets executed without evaluating a predicate
				expr.Label = MakeGenSym(LabelPrefix)
				expr.ThenLines = MaxInt32
				expr.Package = pkg

				arg := MakeArgument("", CurrentFile, LineNo).AddType("bool")
				arg.Package = pkg

				expr.AddInput(arg)

				$$ = []*CXExpression{expr}
			} else {
				panic(err)
			}
                }
	|       RETURN expression SEMICOLON
                {
			$$ = nil
                }
                ;
%%
