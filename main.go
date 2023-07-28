package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	rectangles   = map[rl.Vector2]rl.Rectangle{}
	camera       = rl.Camera2D{}
	ghost        = rl.NewRectangle(0, 0, 50, 50)
	screenWidth  = int32(1920)
	screenHeight = int32(1080)
)

func (g *Game) Update() {
	cameraController()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if g.state == Placing {
			v := rl.Vector2{
				X: ghost.X,
				Y: ghost.Y,
			}

			if _, ok := rectangles[v]; !ok {
				r := rl.NewRectangle(ghost.X, ghost.Y, 50, 50)
				rectangles[rl.Vector2{X: ghost.X, Y: ghost.Y}] = r
				g.state = Selection
			} else {
				fmt.Println("Could not place")
			}
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		if g.state == Selection {
			g.state = Placing
			fmt.Println(g.state.String())
		} else {
			g.state = Selection
			fmt.Println(g.state.String())
		}
	}

	ghostController()
	g.ui.updateUIPanel(g)

}

func (g *Game) Render() {
	renderPlacementObject(ghost, g.state)

	for _, r := range rectangles {
		rl.DrawRectangleRec(r, rl.Red)
	}

}

func (g *Game) RenderUI() {
	g.ui.drawUIPanel()
}

func snapToGridVector(x, y float64) rl.Vector2 {
	gridSize := 50.0
	xPrime := math.Round(x/gridSize) * gridSize
	yPrime := math.Round(y/gridSize) * gridSize
	return rl.NewVector2(float32(xPrime), float32(yPrime))
}

func getWorldMousePos(c rl.Camera2D) rl.Vector2 {
	return rl.GetScreenToWorld2D(rl.GetMousePosition(), c)
}

func renderPlacementObject(placementObject rl.Rectangle, state State) {
	if state != Placing {
		return
	}

	v := snapToGridVector(float64(placementObject.X), float64(placementObject.Y))

	if _, ok := rectangles[v]; ok {
		rl.DrawRectangleRec(placementObject, rl.Red)
	} else {
		rl.DrawRectangleRec(placementObject, rl.Blue)
	}
}

func (u *UI) updateUIPanel(g *Game) {
	if rl.GetMouseY() >= screenHeight-u.rectHeight && g.state != Interface {
		u.previousState = g.state
		g.state = Interface
		fmt.Println(g.state.String())
	}

	if rl.GetMouseY() <= screenHeight-u.rectHeight && g.state == Interface {
		g.state = u.previousState
		fmt.Println(g.state.String())
	}
}

func (u *UI) drawUIPanel() {
	x := 0
	y := screenHeight - u.rectHeight

	rl.DrawRectangle(int32(x), int32(y), int32(u.rectWidth), int32(u.rectHeight), rl.Yellow)

	rl.DrawRectangle(10, y+10, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+50, y+10, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+10+50+50, y+10, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+10+10+50+50+50, y+10, 50, 50, rl.Green)

	rl.DrawRectangle(10, y+50+10+10, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+50, y+10+10+50, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+10+50+50, y+10+10+50, 50, 50, rl.Green)
	rl.DrawRectangle(10+10+10+10+50+50+50, y+10+10+50, 50, 50, rl.Green)
}

func cameraController() {
	camera.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

	if camera.Zoom > 3.0 {
		camera.Zoom = 3.0
	} else if camera.Zoom < 0.1 {
		camera.Zoom = 0.1
	}

	if rl.IsKeyPressed(rl.KeyR) {
		camera.Zoom = 1.0
	}
	if rl.IsKeyDown(rl.KeyW) {
		camera.Target.Y -= 10
	}

	if rl.IsKeyDown(rl.KeyS) {
		camera.Target.Y += 10
	}

	if rl.IsKeyDown(rl.KeyA) {
		camera.Target.X -= 10
	}

	if rl.IsKeyDown(rl.KeyD) {
		camera.Target.X += 10
	}
}

func ghostController() {
	p := snapToGridVector(float64(getWorldMousePos(camera).X), float64(getWorldMousePos(camera).Y))
	ghost.X = p.X - ghost.Width/2
	ghost.Y = p.Y - ghost.Height/2
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 2d camera")

	camera.Target = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	camera.Zoom = 1.0

	rl.SetTargetFPS(60)

	g := &Game{
		ui: &UI{
			rectWidth:  screenWidth,
			rectHeight: int32(screenHeight) / int32(3),
		},
	}

	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		g.Render()

		rl.EndMode2D()

		g.RenderUI()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
