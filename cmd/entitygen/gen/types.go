package gen

// Column metadata
type Column struct {
	Name       string
	Type       string
	Nullable   bool
	PrimaryKey bool
}

// Table metadata
type Table struct {
	Schema    string
	Name      string
	NameUpper string
	NameLower string
	Columns   []Column
}
