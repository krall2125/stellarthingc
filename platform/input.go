package platform

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/hellory4n/stellarthing/core"
)

var keymapma map[string][]Key

func init() {
	keymapma = make(map[string][]Key)
}

func AddKeymap(keymap string, key Key) {
	keymapma[keymap] = append(keymapma[keymap], key)
}

func IsKeyJustPressed(key Key) bool {
	return rl.IsKeyPressed(int32(key))
}

func IsKeyJustReleased(key Key) bool {
	return rl.IsKeyReleased(int32(key))
}

func IsKeyHeld(key Key) bool {
	return rl.IsKeyDown(int32(key))
}

func IsKeyNotPressed(key Key) bool {
	return rl.IsKeyUp(int32(key))
}

func IsMouseButtonJustPressed(btn MouseButton) bool {
	return rl.IsMouseButtonPressed(rl.MouseButton(btn))
}

func IsMouseButtonJustReleased(btn MouseButton) bool {
	return rl.IsMouseButtonReleased(rl.MouseButton(btn))
}

func IsMouseButtonHeld(btn MouseButton) bool {
	return rl.IsMouseButtonDown(rl.MouseButton(btn))
}

func IsMouseButtonNotPressed(btn MouseButton) bool {
	return rl.IsMouseButtonUp(rl.MouseButton(btn))
}

func IsKeymapJustPressed(keymap string) bool {
	// input fields haha
	if core.InternalInputFieldFocus {
		return false
	}

	keymapdeez, ok := keymapma[keymap]

	// i'm not ok :(
	if !ok {
		return false
	}

	for _, v := range keymapdeez {
		if IsKeyJustPressed(v) {
			return true
		}
	}
	return false
}

func IsKeymapJustReleased(keymap string) bool {
	// input fields haha
	if core.InternalInputFieldFocus {
		return false
	}

	keymapdeez, ok := keymapma[keymap]

	// i'm not ok :(
	if !ok {
		return false
	}

	for _, v := range keymapdeez {
		if IsKeyJustReleased(v) {
			return true
		}
	}
	return false
}

func IsKeymapHeld(keymap string) bool {
	// input fields haha
	if core.InternalInputFieldFocus {
		return false
	}

	keymapdeez, ok := keymapma[keymap]

	// i'm not ok :(
	if !ok {
		return false
	}

	for _, v := range keymapdeez {
		if IsKeyHeld(v) {
			return true
		}
	}
	return false
}

func IsKeymapNotPressed(keymap string) bool {
	// input fields haha
	if core.InternalInputFieldFocus {
		return false
	}

	keymapdeez, ok := keymapma[keymap]

	// i'm not ok :(
	if !ok {
		return false
	}

	for _, v := range keymapdeez {
		if IsKeyNotPressed(v) {
			return true
		}
	}
	return false
}

// returns the mouse position in the fake resolution thingy
func MousePosition() core.Vec2 {
	var realMouse rl.Vector2 = rl.GetMousePosition()
	var virtMouse core.Vec2 = core.Vec2{0, 0}

	scale := math.Min(
		float64(rl.GetRenderWidth())/float64(core.RenderSize.X),
		float64(rl.GetRenderHeight())/float64(core.RenderSize.Y),
	)
	// :(
	virtMouse.X = core.Clamp((float64(realMouse.X)-(float64(rl.GetRenderWidth())-(float64(core.RenderSize.X)*scale))*0.5)/scale, 0, core.RenderSize.ToVec2().X)
	virtMouse.Y = core.Clamp((float64(realMouse.Y)-(float64(rl.GetRenderHeight())-(float64(core.RenderSize.Y)*scale))*0.5)/scale, 0, core.RenderSize.ToVec2().Y)

	return virtMouse
}
