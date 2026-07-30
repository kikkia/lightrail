package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/kikkia/lightrail/resource"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	dinNormalFace font.Face
	nextTimes     []int64
	tick          uint16
	circleSubImg  *ebiten.Image
)

func init() {
	// Parse the OTF/TTF font using standard golang/x/image/font for Ebiten v1
	f, err := opentype.Parse(resource.Din)
	if err != nil {
		log.Fatal(err)
	}

	// Create a standard font face (size 36 at 72 DPI)
	dinNormalFace, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    36,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	circleSubImg, _ = ebiten.NewImage(40, 40, ebiten.FilterDefault)
	drawCircle(circleSubImg, 20, 20, 20, color.RGBA{0x00, 0xA0, 0xDF, 0xff})
}

func update(screen *ebiten.Image) error {
	// Initialize the glyphs for special (colorful) rendering.
	var err error
	if tick == 0 {
		nextTimes, err = fetchNextTimes()
		log.Printf("%s", nextTimes)
		log.Printf("%s", time.Now().UnixMilli())
		if err != nil {
			log.Printf("%s", err)
		}
	}

	tick++
	tick = tick % 3600

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	draw(screen)

	// ghetto force 2fps?
	time.Sleep(500 * time.Millisecond)
	return nil
}

func draw(screen *ebiten.Image) {

	var x = 60
	var y = 15
	for i := 0; i < 4; i++ {
		if nextTimes[i] == 0 {
			break
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x-50), float64(y-10))
		screen.DrawImage(circleSubImg, op)

		// Draw the number "2" inside the circle
		text.Draw(screen, "2", dinNormalFace, x-37, y+22, color.White)

		mins := fmtTime(nextTimes[i])
		var str string
		if mins == 0 {
			str = "       Now"
		} else if mins/10 > 0 {
			str = fmt.Sprintf(" %d mins", mins)
		} else {
			str = fmt.Sprintf("   %d mins", mins)
		}
		s := fmt.Sprintf("Lynnwood       %s", str)

		// Ebiten v1 text coordinates represent the *baseline* of the text,
		// so we add the font size offset to the Y coordinate to match v2's top-left layout.
		text.Draw(screen, s, dinNormalFace, x, y+30, color.White)
		y += 60
	}
}

func fmtTime(next int64) int64 {
	mins := (next - time.Now().UnixMilli()) / 60000
	return mins
}

func drawCircle(img *ebiten.Image, x, y, r int, clr color.Color) {
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				img.Set(x+dx, y+dy, clr)
			}
		}
	}
}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.Run(update, screenWidth, screenHeight, 1.0, "Text (Ebitengine Demo)"); err != nil {
		log.Fatal(err)
	}
}
