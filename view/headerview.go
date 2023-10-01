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
	"golang.org/x/image/font/gofont/goregular"
)

const (
	fontRobotoBoldPath   = "assets/fonts/Roboto-Bold.ttf"
	headerFlagPath       = "assets/emotes/HeaderFlag.png"
	emoteHYPERSPath      = "assets/emotes/HYPERS.png"
	emotePepeHandsPath   = "assets/emotes/PepeHands.png"
	emotePepoThinkPath   = "assets/emotes/PepoThink.png"
)

var (
	HeaderContainer *widget.Container = newHeaderContainer()
	DifficultyButton *widget.Button = newDifficultyButton()
	ResetButtonStackContainer *widget.Container = newResetButtonStackContainer()
	ResetButton *widget.Button = newResetButton()
	ResetEmoteGraphic *widget.Graphic = newResetGraphic()
	FlagsCounterText *widget.Text = newFlagsCounterText()
	HeaderFlagGraphic *widget.Graphic = newHeaderFlagGraphic()
	FlagRowContainer *widget.Container = newFlagRowContainer()
	// Reset button emote graphics
	HypersIcon, _ = loadButtonIcon(emoteHYPERSPath)
	PepeHandsIcon, _ = loadButtonIcon(emotePepeHandsPath)
	PepoThinkIcon, _ = loadButtonIcon(emotePepoThinkPath)
	// Flag graphics
	HeaderFlag, _ = loadButtonIcon(headerFlagPath)
)

func newHeaderContainer() *widget.Container {
	headerCountainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.HeaderBackgroundColor)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			//Which direction to layout children
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			//Set how much padding before displaying content
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			//Set how far apart to space the children
			widget.RowLayoutOpts.Spacing(10),
		)),
	)

	return headerCountainer
}

func newDifficultyButton() *widget.Button {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	robotoBoldFace, _ := loadAssetFont(fontRobotoBoldPath, 16)

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
		widget.ButtonOpts.Text("Difficulty: " + controller.GetDifficultyString(), robotoBoldFace, &widget.ButtonTextColor{
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
			switch controller.GetDifficultyString() {
			case "Beginner":
				controller.NewBoard(model.Intermediate)
			case "Intermediate":
				controller.NewBoard(model.Expert)
			case "Expert":
				controller.NewBoard(model.Beginner)
			}
			args.Button.Text().Label = "Difficulty: " + controller.GetDifficultyString()
		}),
	)

	return difficultyButton
}

func newResetButtonStackContainer() *widget.Container {
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

	return resetButtonStackContainer
}

func newResetButton() *widget.Button {
	// load images for emotes
	emoteImage, _ := loadBackgroundMatchingImage()
	
	resetButton  := widget.NewButton(
		widget.ButtonOpts.Image(emoteImage),
		
		// add a handler that reacts to clicking the button
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			controller.NewBoard()
		}),
	)

	return resetButton
}

func newResetGraphic() *widget.Graphic {
	resetEmoteGraphic := widget.NewGraphic(widget.GraphicOpts.Image(PepoThinkIcon))

	return resetEmoteGraphic
}

func newFlagRowContainer() *widget.Container {
	flagRowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.HeaderBackgroundColor)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			//Which direction to layout children
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			//Set how much padding before displaying content
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(2)),
			//Set how far apart to space the children
			widget.RowLayoutOpts.Spacing(2),

		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	return flagRowContainer
}


func newHeaderFlagGraphic() *widget.Graphic {
	flagGraphic := widget.NewGraphic(widget.GraphicOpts.Image(HeaderFlag))

	return flagGraphic
}

func newFlagsCounterText() *widget.Text {
	robotoBoldFace, _ := loadAssetFont(fontRobotoBoldPath, 24)

	flagsCounterText := widget.NewText(
		widget.TextOpts.Text("", robotoBoldFace, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)

	return flagsCounterText
}

func RenderHeader() *widget.Container  {
	// Difficulty Button
	HeaderContainer.AddChild(DifficultyButton)


	// Flag tracker
	FlagRowContainer.AddChild(HeaderFlagGraphic)
	FlagRowContainer.AddChild(FlagsCounterText)
	HeaderContainer.AddChild(FlagRowContainer)

	// Reset Button
	ResetButtonStackContainer.AddChild(ResetButton)
	ResetButtonStackContainer.AddChild(ResetEmoteGraphic)
	HeaderContainer.AddChild(ResetButtonStackContainer)

	return HeaderContainer
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


// Background color the same as acontainer
func loadBackgroundMatchingImage() (*widget.ButtonImage, error) {
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

// Background color the same as container, and button is not pressable
func loadNonPressableButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(model.HeaderBackgroundColor)
	hover, pressed := idle, idle

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

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}