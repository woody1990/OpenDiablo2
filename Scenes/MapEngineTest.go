package Scenes

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/Common"
	"github.com/OpenDiablo2/OpenDiablo2/Map"
	"github.com/OpenDiablo2/OpenDiablo2/Sound"
	"github.com/OpenDiablo2/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type MapEngineTest struct {
	uiManager     *UI.Manager
	soundManager  *Sound.Manager
	fileProvider  Common.FileProvider
	sceneProvider SceneProvider
	gameState     *Common.GameState
	mapEngine     *Map.Engine
}

func CreateMapEngineTest(
	fileProvider Common.FileProvider,
	sceneProvider SceneProvider,
	uiManager *UI.Manager,
	soundManager *Sound.Manager) *MapEngineTest {
	result := &MapEngineTest{
		fileProvider:  fileProvider,
		uiManager:     uiManager,
		soundManager:  soundManager,
		sceneProvider: sceneProvider,
	}
	result.gameState = Common.CreateGameState()
	return result
}

func (v *MapEngineTest) Load() []func() {
	// TODO: Game seed comes from the game state object

	v.soundManager.PlayBGM("")
	return []func(){
		func() {
			v.mapEngine = Map.CreateMapEngine(v.gameState, v.soundManager, v.fileProvider)

			v.mapEngine.GenerateAct1Overworld()
			//v.mapEngine.GenerateMap(Map.RegionAct1Tristram, 300)
			//v.mapEngine.GenerateMap(Map.RegionAct1Cathedral, 257)
			//v.mapEngine.GenerateMap(Map.RegionAct2Town, 301)
			//v.mapEngine.GenerateMap(Map.RegionAct2Harem, 353) // Crashes on dcc load
			//v.mapEngine.GenerateMap(Map.RegionAct3Town, 529)
			//v.mapEngine.GenerateMap(Map.RegionAct3Jungle, 574)
			//v.mapEngine.GenerateMap(Map.RegionAct4Town, 797) // Broken height of large objects
			//v.mapEngine.GenerateMap(Map.RegonAct5Town, 863) // Completely broken!!
			//v.mapEngine.GenerateMap(Map.RegionAct5IceCaves, 1038) // Completely broken!
			//v.mapEngine.GenerateMap(Map.RegionAct5Siege, 879) // Completely broken!
			//v.mapEngine.GenerateMap(Map.RegionAct5Lava, 1057) // Broken
			//v.mapEngine.GenerateMap(Map.RegionAct5Barricade, 880) // Broken

		},
	}
}

func (v *MapEngineTest) Unload() {

}

func (v *MapEngineTest) Render(screen *ebiten.Image) {
	v.mapEngine.Render(screen)
	actualX := float64(v.uiManager.CursorX) - v.mapEngine.OffsetX
	actualY := float64(v.uiManager.CursorY) - v.mapEngine.OffsetY
	tileX, tileY := Common.ScreenToIso(actualX, actualY)
	subtileX := int(math.Ceil(math.Mod((tileX*10), 10))) / 2
	subtileY := int(math.Ceil(math.Mod((tileY*10), 10))) / 2
	curRegion := v.mapEngine.GetRegionAt(int(tileX), int(tileY))
	if curRegion == nil {
		return
	}
	line := fmt.Sprintf("%d, %d (Tile %d.%d, %d.%d)",
		int(math.Ceil(actualX)),
		int(math.Ceil(actualY)),
		int(math.Ceil(tileX))-curRegion.Rect.Left,
		subtileX,
		int(math.Ceil(tileY))-curRegion.Rect.Top,
		subtileY,
	)
	ebitenutil.DebugPrintAt(screen, line, 5, 5)
	ebitenutil.DebugPrintAt(screen, "Map: "+curRegion.Region.LevelType.Name, 5, 17)
	ebitenutil.DebugPrintAt(screen, curRegion.Region.RegionPath, 5, 29)
}

func (v *MapEngineTest) Update(tickTime float64) {
	if v.uiManager.KeyPressed(ebiten.KeyDown) {
		v.mapEngine.OffsetY -= tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyUp) {
		v.mapEngine.OffsetY += tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyLeft) {
		v.mapEngine.OffsetX += tickTime * 800
	}
	if v.uiManager.KeyPressed(ebiten.KeyRight) {
		v.mapEngine.OffsetX -= tickTime * 800
	}
}
