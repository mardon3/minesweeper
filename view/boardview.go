package view

import (
	"image/color"
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)


var (
	BoardContainer *widget.Container
	BoardCells [][]*widget.Button
	// Screen dimensions
	screenWidth int
	screenHeight int
	// Revealed font
	RobotoBoldFace, _  = loadAssetFont(fontRobotoBoldPath, 16)
	// Not revealed font
	FlagFace, _ = loadAssetFont(fontFlagPath, 30)
	// Mine font
	MineFace, _ = loadAssetFont(fontMinePath, 30)
)

func newBoardContainer() *widget.Container {
	screenWidth = controller.GetBoardWidth()
	screenHeight = controller.GetBoardHeight()

	columnStretchSlice := make([]bool, screenWidth)
	for i := range columnStretchSlice {
		columnStretchSlice[i] = true			
	}

	rowStretchSlice := make([]bool, screenHeight)
	for i := range rowStretchSlice {
		rowStretchSlice[i] = true
	}

	boardContainer := widget.NewContainer(
		// the container will use assigned background color
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(BoardBackgroundColor)),
		// container will generate as many columns as the screen width requires
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(screenWidth),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(2)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(2, 2),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch(columnStretchSlice, rowStretchSlice),
		)),
	)

	return boardContainer
}

func newBoardCellButton(r, c int) *widget.Button {
	cellImage, _ := LoadCellImage(r, c)

	cellSize := controller.GetDifficulty().CellSize

	boardCellButton := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionStart,
				VerticalPosition: widget.GridLayoutPositionStart,
				MaxWidth: cellSize,
				MaxHeight: cellSize,
			}),

			widget.WidgetOpts.MouseButtonPressedHandler(func(args *widget.WidgetMouseButtonPressedEventArgs) {
				if args.Button == ebiten.MouseButtonRight && !controller.IsLost() && !controller.IsSolved() {
					controller.RightClickCell(r, c)	
				}
			}),
		),

		widget.ButtonOpts.Image(cellImage),

		widget.ButtonOpts.Text("", FlagFace, &widget.ButtonTextColor{
			Idle: color.NRGBA{255, 0, 0, 255},
		}),

		// handler for when cell is clicked
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if !controller.IsLost() && !controller.IsFlagged(r, c) && !controller.IsSolved() {
				controller.LeftClickCell(r, c)
			}
		}),
	)

	return boardCellButton
}

func RenderBoard() *widget.Container {	
	BoardContainer = newBoardContainer()

	// Create new BoardCells
	BoardCells = make([][]*widget.Button, screenHeight)
	for r := 0; r < screenHeight; r++ {
		BoardCells[r] = make([]*widget.Button, screenWidth)
		for c := 0; c < screenWidth; c++ {
			BoardCells[r][c] = newBoardCellButton(r, c)
			BoardContainer.AddChild(BoardCells[r][c])
		}
	}

	return BoardContainer
}

func LoadCellImage(r, c int) (*widget.ButtonImage, error) {
	var (
		idle     *image.NineSlice
		hover    *image.NineSlice
		pressed  *image.NineSlice
	)
	
	if controller.IsRevealed(r, c) && controller.GetCellType(r, c) == model.MineCell {
		idle = image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})
		hover = image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})
		pressed = image.NewNineSliceColor(color.NRGBA{255, 0, 0, 255})
	} else if !controller.IsRevealed(r, c) && (controller.IsSolved() || controller.IsLost()) {
		idle = image.NewNineSliceColor(controller.GetColor(r, c))
		hover = image.NewNineSliceColor(controller.GetColor(r, c))
		pressed = image.NewNineSliceColor(controller.GetColor(r, c))
	} else if !controller.IsRevealed(r, c) {
		idle = image.NewNineSliceColor(controller.GetColor(r, c))
		hover = image.NewNineSliceColor(color.NRGBA{130, 210, 255, 255})
		pressed = image.NewNineSliceColor(color.NRGBA{130, 210, 255, 255})
	} else {
		idle = image.NewNineSliceColor(color.NRGBA{235, 235, 210, 255})
		hover = image.NewNineSliceColor(color.NRGBA{235, 235, 210, 255})
		pressed = image.NewNineSliceColor(color.NRGBA{235, 235, 210, 255})
	}	

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
