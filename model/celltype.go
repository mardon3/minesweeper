package model

// Enum type for cell types
type CellType int

// Enums for cell types
const (
	MineCell CellType = iota
	EmptyCell 
	ValueCell 
)