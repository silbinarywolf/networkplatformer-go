// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/examples/resources/images/platformer"
)

const (
	// Settings
	screenWidth  = 1024
	screenHeight = 512
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// Preload images
	img, _, err := image.Decode(bytes.NewReader(rplatformer.Right_png))
	if err != nil {
		panic(err)
	}
	rightSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Left_png))
	if err != nil {
		panic(err)
	}
	leftSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.MainChar_png))
	if err != nil {
		panic(err)
	}
	idleSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Background_png))
	if err != nil {
		panic(err)
	}
	backgroundImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type Char struct {
	X                 float64
	Y                 float64
	sprite            *ebiten.Image
	isKeyLeftPressed  bool
	isKeyRightPressed bool

	// used by server only
	lastUpdatedTimer time.Time
}

func (c *Char) RemoveFromSimulation() {
	// Unordered remove
	for i, char := range chars {
		if char == c {
			chars[i] = chars[len(chars)-1] // Replace it with the last one.
			chars = chars[:len(chars)-1]   // delete last element
			return
		}
	}
}

var (
	you *Char = &Char{
		X: 50,
		Y: 380,
	}
	chars []*Char = make([]*Char, 0, 256)
)

func update(screen *ebiten.Image) error {
	// Read/write network information
	if server != nil {
		server.Update()
	}
	if client != nil {
		client.Update()
	}

	// Controls
	if you != nil {
		you.isKeyLeftPressed = false
		you.isKeyRightPressed = false
		if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
			you.isKeyLeftPressed = true
		} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
			you.isKeyRightPressed = true
		}
	}

	// Simulate
	for _, char := range chars {
		char.sprite = idleSprite
		if char.isKeyLeftPressed {
			// Selects preloaded sprite
			char.sprite = leftSprite
			// Moves character 3px right
			char.X -= 3
		} else if char.isKeyRightPressed {
			// Selects preloaded sprite
			char.sprite = rightSprite
			// Moves character 3px left
			char.X += 3
		}
	}

	if ebiten.IsRunningSlowly() {
		return nil
	}

	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)

	// Draws selected sprite image
	for _, char := range chars {
		if char.sprite == nil {
			continue
		}
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(char.X, char.Y)
		screen.DrawImage(char.sprite, op)
	}

	// FPS counter
	fps := fmt.Sprintf("FPS: %f", ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, fps)

	return nil
}

func main() {
	// Setup network
	isServer := false
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if firstArg == "--server" {
			isServer = true
		}
	}
	if isServer {
		server = NewServer()
		go server.Listen()
	} else {
		client = NewClient()
		err := client.Dial("localhost:8080")
		if err != nil {
			panic(err)
		}
		go client.Listen()
	}

	// This is required so the server can run when the window isn't focused.
	ebiten.SetRunnableInBackground(true)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Platformer (Ebiten Demo)"); err != nil {
		panic(err)
	}
}
