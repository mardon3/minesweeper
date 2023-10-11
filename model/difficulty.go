package model

// Difficulty is a struct  representing a board's dimensions and mines count
type Difficulty struct {
    Height          int
    Width           int
    Mines           int
    ScreenWidth     int
    ScreenHeight    int
    CellSize        int
}

// Different difficulty levels
var (
    Beginner     = Difficulty{Height: 9, Width: 9, Mines: 10, ScreenWidth: 560, ScreenHeight: 632, CellSize: 60}
    Intermediate = Difficulty{Height: 16, Width: 16, Mines: 40, ScreenWidth: 738, ScreenHeight: 810, CellSize: 44}
    Expert       = Difficulty{Height: 16, Width: 30, Mines: 99, ScreenWidth: 1382, ScreenHeight: 810, CellSize: 44}
)