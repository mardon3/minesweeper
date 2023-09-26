package model

// Difficulty is a struct  representing a board's dimensions and mines count
type Difficulty struct {
    Height   int
    Width    int
    Mines    int
}

// Different difficulty levels
var (
    Beginner     = Difficulty{Height: 9, Width: 9, Mines: 10}
    Intermediate = Difficulty{Height: 16, Width: 16, Mines: 40}
    Expert       = Difficulty{Height: 16, Width: 30, Mines: 99}
)