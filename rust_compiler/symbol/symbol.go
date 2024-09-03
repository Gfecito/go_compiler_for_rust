package symbol

type Symbol interface {
}

type Int struct {
}

type Undefined struct {
}

// Singleton global Symbols
var (
	IntSymbol       = &Int{}
	UndefinedSymbol = &Undefined{}
)

type SymbolTable struct {
	symbols map[string]Symbol
}

func (st *SymbolTable) Lookup(name string) Symbol {
	if symbol, ok := st.symbols[name]; ok {
		return symbol
	}
	return UndefinedSymbol
}
