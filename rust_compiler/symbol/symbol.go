package symbol

import "errors"

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
	symbols map[string]*Symbol
	// Parent scope.
	outer *SymbolTable
}

// SYMBOL MANAGEMENT
func (st *SymbolTable) Insert(name string, sym *Symbol) {
	st.symbols[name] = sym
}

func (st *SymbolTable) Lookup(name string) Symbol {
	if symbol, ok := st.symbols[name]; ok {
		return symbol
	}

	if st.outer == nil {
		return UndefinedSymbol
	}
	return st.outer.Lookup(name)
}

// SCOPE MANAGEMENT
func (st *SymbolTable) EnterScope() *SymbolTable {
	// Make new inner scope
	return &SymbolTable{symbols: make(map[string]*Symbol), outer: st}
}

func (st *SymbolTable) LeaveScope() (*SymbolTable, error) {
	if st.outer == nil {
		return nil, errors.New("Tried to leave scope but no outer scope was found.")
	}
	return st.outer, nil
}
