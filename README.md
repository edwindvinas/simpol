# Simpol

Simpol is a Golang interpreter for the sample language called Simpol. It was patterned from github.com/edwindvinas/simpol.

## Installation
Requires Go.
```
$ go get -u github.com/edwindvinas/simpol
```

## Examples

```bash
variable {
STG str
STG name
INT num1
INT num2
INT num3
BLN bol1
BLN bol2
}
 
code {
PUT $The result is: $ IN str
ASK name
PUT true IN bol1
PUT false IN bol2
PUT ADD 1 2 IN num1
PUT 100 IN num2
 
PRT $Your name is $
PRT name
PRT OHR true AND bol1 bol2
PUT MUL 10 ADD num1 num2 IN num3
PRT num3
PRT DIV MUL 10 ADD num1 num2 MUL 10 ADD num1 num2
PRT $Goodbye!$
}
```

See `_examples/scripts` for more examples.

Running scripts using simpol command-line tool:

```
$ simpol simpol.sim
```

To try this online, see https://github.com/edwindvinas/simpol-online

To see the Simpol Language Specifications, please see below.

### What is the grammar of SIMPOL language?

Knowing Lex and Yacc is useless if you don't know the language you are working on. You need to understand how the language is structured so we can define the patterns to be used in lexer and also the productions to be used in Yacc.

#### Structure
SIMPOL has 2 main sections, one for declaring variables and the other for coding. 

    variable {
    ...
    }
    
    code {
    ...
    }
 
This programming language has 3 main data types namely integer, Boolean, and string.

    Integer
    Boolean
    String

#### Formatting
• Spaces are used to demarcate tokens in the language.
• Multiple spaces and tabs are treated as single spaces and are otherwise considered irrelevant.
• Indentation is also irrelevant.
• SIMPOL is case sensitive.

#### Section: variable { }
• All variables are needed to be declared inside this segment.
• No variable declaration is allowed in code { }.
• You cannot use a variable unless declared.
• All variables are global in scope.
• Example:

    variable {
    INT a
    BLN b
    STG s
    }

#### Section: code { }
• This is where the actual instructions for a SIMPOL code are written.
• Example:

    code {
    PRT ADD 1 SUB 8 9
    ASK a
    PRT a
    }

#### Variables and Data Types 
• There are 3 main data types in SIMPOL: integer (INT), Boolean (BLN) and string (STG).
• A string is enclosed by 2 dollar signs ($)
• Boolean has 2 possible values: true and false.
• To declare a variable in variable { }, the format is <data_type> <variable_name>

#### User Input and Output
• To ask for input, use the ASK operator. The format is ASK <variable_name>.
• To print values, use the PRT operator. The format is PRT <expression>

#### Operations
• Take Note: An expression can be composed of several adjacent operators. It is not only
limited to 1 operator. Thus, nested expressions are allowed.

##### Assignment Operation
Assignment 

    PUT <expression> IN <variable>

##### Arithmetic Operations
Addition 

    ADD <expression1> <expression2>

Subtraction 

    SUB <expression1> <expression2>

Multiplication 

    MUL <expression1> <expression2>

Division 

    DIV <expression1> <expression2>

Modulo 

    MOD <expression1> <expression2>

##### Numeric Predicates
Greater Than 

    GRT <expression1> <expression2>

Greater Than or Equal 

    GRE <expression1> <expression2>

Less Than 

    LET <expression1> <expression2>

Less Than or Equal 

    LEE <expression1> <expression2>

Equal 

    EQL <expression1> <expression2>

##### Logical Operations
And 

    AND <expression1> <expression2>

Or `

    OHR <expression1> <expression2>`

Not 

    NON <expression1>

#### Comments
Comments are identified by double-slash and it is only one line comments allowed

    // Here is a simpol comment

#### Limitations
* To make things simple, an identifier should consist of letters only.
* No initializing of variables inside the variable { }
* There are no operations to be performed on strings. If you try to perform any operations on them, it should throw an error.
* You can perform operations on Boolean using logical operations.
* code { } may be empty.
* It is required that variable{ } should be declared first before code{ } 
* Nested expressions should be allowed

##### Sample Empty SIMPOL Program

    variable { }
    code { }

##### Simple SIMPOL Program

    variable {
    BLN a
    }
    code {
    PUT ADD 3 SUB 8 9 IN a
    PRT a
    }

##### Sample Non-nested SIMPOL Program

	variable {
	STG str
	STG name
	INT num1
	INT num2
	INT num3
	INT num4
	INT num5
	BLN bol1
	BLN bol2
	}
	 
	code {
	PUT $The result is: $ IN str
	ASK name
	PUT true IN bol1
	PUT false IN bol2
	PUT ADD 1 2 IN num1
	PUT 100 IN num2
	 
	PRT $Your name is $
	PRT name
	PRT OHR true AND bol1 bol2
	PUT ADD num1 num2 IN num3
	PUT MUL 10 num3
	PRT num3
	PUT MUL 10 ADD num1 num2 IN num4
	PUT MUL 10 ADD num1 num2 IN num5
	PRT DIV num4 num5
	PRT $Goodbye!$
	}

##### Sample Nested SIMPOL Program
	variable {
	STG str
	STG name
	INT num1
	INT num2
	INT num3
	BLN bol1
	BLN bol2
	}
	 
	code {
	PUT $The result is: $ IN str
	ASK name
	PUT true IN bol1
	PUT false IN bol2
	PUT ADD 1 2 IN num1
	PUT 100 IN num2
	 
	PRT $Your name is $
	PRT name
	PRT OHR true AND bol1 bol2
	PUT MUL 10 ADD num1 num2 IN num3
	PRT num3
	PRT DIV MUL 10 ADD num1 num2 MUL 10 ADD num1 num2
	PRT $Goodbye!$
	}

#### Formulate Grammar

Out of the specs of the SIMPOL language, we can formulate the grammar or syntax.

    	variable {	/* <-- T_BEGIN_VARIABLE_BLOCK */
    	STG str		
    	STG name	
    	INT num1
    	INT num2
    	INT num3
    	BLN bol1
    	BLN bol2
    	}	/* <-- T_END_BLOCK */
    	
    /* 	Inside Variable Block:
        ----------------------
    		<variable_block_lines> := <data_type> <identifier>
    		<data_type> := STG | INT | BLN
    		<identifier> := <alpha-literals>
    		<alpha-literals> := [a-zA-Z]
    */
    	 
    	code {	/* <-- T_BEGIN_CODE_BLOCK */
    	PUT $The result is: $ IN str
    	ASK name
    	PUT true IN bol1
    	PUT false IN bol2
    	PUT ADD 1 2 IN num1
    	PUT 100 IN num2
    	 
    	PRT $Your name is $
    	PRT name
    	PRT OHR true AND bol1 bol2
    	PUT MUL 10 ADD num1 num2 IN num3
    	PRT num3
    	PRT DIV MUL 10 ADD num1 num2 MUL 10 ADD num1 num2
    	PRT $Goodbye!$
    	} /* <-- T_END_BLOCK */
    	
    /*
    	Inside Code Block:
    	------------------
    		<code_block_lines> := <expr>
    		<expr> := <put-expr> | <ask-expr> | <print-expr>
    		<ask-expr> := ASK <identifier>
    		<print-expr> := <identifier> | $<hard-coded-str>$
    		<hard-coded-str> := [a-zA-Z]*
    		<put-expr> := <source-expr> IN <dest-expr>
    		<dest-expr> := <identifier>
    		<source-expr> := <bool-literals> | <numeric-literals> | <arith-expr> 
    		<arith-expr> := <add-expr> | <sub-expr> | <mul-expr> | <div-expr> | <mod-expr>
    		<add-expr> := ADD <identifier> <identifier> | ADD <numeric-literals> <identifier> | ADD <numeric-literals> <numeric-literals>
    		<sub-expr> := SUB <identifier> <identifier> | SUB <numeric-literals> <identifier> | SUB <numeric-literals> <numeric-literals>
    		<mul-expr> := MUL <identifier> <identifier> | MUL <numeric-literals> <identifier> | MUL <numeric-literals> <numeric-literals>
    		<div-expr> := DIV <identifier> <identifier> | DIV <numeric-literals> <identifier> | DIV <numeric-literals> <numeric-literals>
    		<mod-expr> := MOD <identifier> <identifier> | MOD <numeric-literals> <identifier> | MOD <numeric-literals> <numeric-literals>
    		<bool-literals> := true | false
    		<arith-operators> := ADD | SUB | MUL | DIV | MOD
    		<num-predicates> := GRT | GRE | LET | LEE | EQL
    		<logic-operators> := AND | OHR | NON 
    		<numeric-literals> := [0-9] 
    */

