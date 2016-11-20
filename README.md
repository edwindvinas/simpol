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

To see the Simpol Language Specifications, please see https://github.com/edwindvinas/simpol/blob/master/SIMPOL_SPECS.md

