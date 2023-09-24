package model

// Difficulty is a custom type representing a tuple of three integers
type Difficulty struct {
    Height   int
    Width    int
    Mines    int
}

// Constants for different difficulty levels
var (
    Beginner     = Difficulty{Height: 9, Width: 9, Mines: 10}
    Intermediate = Difficulty{Height: 16, Width: 16, Mines: 40}
    Expert       = Difficulty{Height: 16, Width: 30, Mines: 99}
)
