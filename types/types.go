package types

type Token struct {
	Name       string
	IsTerminal bool
}

type Production []Token

type Grammar struct {
	Terminals, NonTerminals []Token
	Productions             map[Token][]Production
}

type Precedence int

const (
	Lower Precedence = -1
	Equal
	Higher
)

type TokenPair struct {
	Left, Right Token
}

type OPTable struct {
	Terminals []Token
	Relations map[TokenPair]Precedence
}