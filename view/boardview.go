package view

import (
	"image/color"
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)


var (
	BoardContainer *widget.Container
	BoardCells [][]*widget.Button
	// Screen dimensions
	screenWidth int
	screenHeight int
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

	flagFace, _ := loadAssetFont(fontRobotoBoldPath, 16)

	boardCellButton := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition: widget.GridLayoutPositionStart,
				MaxWidth: controller.GetDifficulty().CellSize,
				MaxHeight: controller.GetDifficulty().CellSize,
			}),
		),

		widget.ButtonOpts.Image(cellImage),

		widget.ButtonOpts.Text("", flagFace, &widget.ButtonTextColor{
			Idle: color.NRGBA{255, 0, 0, 255},
		}),

		// handler for when cell is clicked
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if !controller.IsLost() && !controller.IsFlagged(r, c) {
				controller.LeftClickCell(r, c)
			}
		}),
	)
 
	return boardCellButton
}

func RenderBoard() *widget.Container {	
	BoardContainer = newBoardContainer()

	for r := 0; r < screenHeight; r++ {
		BoardCells = append(BoardCells, make([]*widget.Button, screenWidth))
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
	} else if !controller.IsRevealed(r, c) {
		idle = image.NewNineSliceColor(controller.GetColor(r, c))
		hover = image.NewNineSliceColor(color.NRGBA{130, 210, 255, 255})
		pressed = image.NewNineSliceColor(color.NRGBA{130, 210, 255, 255})
	} else {
		idle = image.NewNineSliceColor(color.White)
		hover = image.NewNineSliceColor(color.White)
		pressed = image.NewNineSliceColor(color.White)
	}	

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
