package view

import (
	"image/color"
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

const (
	fontRobotoBoldPath = "assets/fonts/Roboto-Bold.ttf"
	emoteHYPERSPath    = "assets/emotes/HYPERS.png"
	emotePepeHandsPath = "assets/emotes/PepeHands.png"
	emotePepoThinkPath = "assets/emotes/PepoThink.png"
)

func NewHeaderContainer() *widget.Container {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()
	// load images for emotes
	emoteImage, _ := loadEmoteImage()

	robotoBoldFace, _ := loadAssetFont(fontRobotoBoldPath, 14)

	var resetEmoteIcon *ebiten.Image
	if controller.IsSolved() {
		resetEmoteIcon, _ = loadButtonIcon(emoteHYPERSPath)
	} else if controller.IsLost() {
		resetEmoteIcon, _ = loadButtonIcon(emotePepeHandsPath)
	} else {
		resetEmoteIcon, _ = loadButtonIcon(emotePepoThinkPath)
	}

	headerCountainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.HeaderBackgroundColor)),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			//Which direction to layout children
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			//Set how much padding before displaying content
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			//Set how far apart to space the children
			widget.RowLayoutOpts.Spacing(10),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(60, 60),
		),
	)

	difficultyButton := widget.NewButton(
		// set general widget options
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),

		// specify the images to use
		widget.ButtonOpts.Image(buttonImage),

		// specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Toggle Difficulty", robotoBoldFace, &widget.ButtonTextColor{
			Idle: color.NRGBA{0, 0, 0, 255},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   10,
			Right:  10,
			Top:    5,
			Bottom: 5,
		}),

		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			switch controller.GetDifficulty() {
			case model.Beginner:
				NewUI(model.Intermediate)
			case model.Intermediate:
				NewUI(model.Expert)
			case model.Expert:
				NewUI(model.Beginner)
			}
		}),
	)
	headerCountainer.AddChild(difficultyButton)

	resetButtonStackContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
		// instruct the container's anchor layout to center the button both horizontally and vertically;
		// since our button is a 2-widget object, we add the anchor info to the wrapping container
		// instead of the button
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)

	resetButton  := widget.NewButton(
		widget.ButtonOpts.Image(emoteImage),
		
		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
		}),
	)
	resetButtonStackContainer.AddChild(resetButton)
	resetButtonStackContainer.AddChild(widget.NewGraphic(widget.GraphicOpts.Image(resetEmoteIcon)))
	
	headerCountainer.AddChild(resetButtonStackContainer)

	return headerCountainer
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 255, G: 255, B: 255, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 200, G: 200, B: 200, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadEmoteImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(model.HeaderBackgroundColor)
	// Darker version of HeaderBackgroundColor to simulate a hover
	hover := image.NewNineSliceColor(color.NRGBA{10, 27, 52, 255})
	// Even darker version of HeaderBackgroundColor to simulate a press
	pressed := image.NewNineSliceColor(color.NRGBA{5, 13, 26, 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadAssetFont(path string, size float64) (font.Face, error) {
	fontData, err := embeddedAssets.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

func loadButtonIcon(path string) (*ebiten.Image, error) {
	f, err := embeddedAssets.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := ebitenutil.NewImageFromReader(f)

	return image, err
}