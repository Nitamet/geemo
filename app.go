package main

import (
	"context"
	"fmt"
	"github.com/Nitamet/geemo/backend"
	"github.com/Nitamet/geemo/backend/lcu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
)

// App struct
type App struct {
	ctx      context.Context
	LCU      *lcu.Client
	Settings backend.Settings
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{Settings: backend.InitializeSettings()}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetLCUState returns the current state of the LCU
func (a *App) GetLCUState() string {
	if a.LCU != nil {
		state := a.LCU.UpdateState()

		// If we got "StateNotLaunched" state while we have a LCU instance, it means that the league client was closed
		if state != lcu.StateNotLaunched {
			return state
		}

		// So we can get rid of the old instance and try to get a new one
		a.LCU = nil
	}

	instance := lcu.TryToGetLCU()
	if instance == nil {
		return lcu.StateNotLaunched
	}

	a.LCU = instance

	return a.LCU.UpdateState()
}

func (a *App) GetSummoner() lcu.Summoner {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	return a.LCU.CurrentSummoner()
}

func (a *App) GetCurrentChampion() int {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	championId, _ := a.LCU.GetCurrentChampion()
	return championId
}

func (a *App) ApplyRunes(runes lcu.RunePage) {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	a.LCU.ApplyRunes(runes)
}

func (a *App) ApplySummonerSpells(firstSpellId int, secondSpellId int) {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	a.LCU.ApplySummonerSpells(firstSpellId, secondSpellId)
}

func (a *App) ApplyItemSet(itemSet lcu.ItemSet) {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	a.LCU.ApplyItemSet(itemSet)
}

func (a *App) GetGameMode() []string {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	gameMode, gameModeAsString := a.LCU.GetCurrentGameMode()

	return []string{gameMode, gameModeAsString}
}

func (a *App) GetAssignedRole() string {
	if a.LCU == nil {
		log.Panic("LCU not found")
	}

	position, _ := a.LCU.GetAssignedRole()

	lcuRoleToAppRole := map[string]string{
		"top":     "top",
		"jungle":  "jungle",
		"middle":  "mid",
		"bottom":  "adc",
		"utility": "support",
		"":        "",
	}

	return lcuRoleToAppRole[position]
}

func (a *App) Minimize() {
	if runtime.WindowIsMinimised(a.ctx) {
		runtime.WindowUnminimise(a.ctx)
		return
	}

	runtime.WindowMinimise(a.ctx)
}

func (a *App) Maximize() {
	if runtime.WindowIsMaximised(a.ctx) {
		runtime.WindowUnmaximise(a.ctx)
		return
	}

	runtime.WindowMaximise(a.ctx)
}

func (a *App) Close() {
	runtime.Quit(a.ctx)
}

func (a *App) GetAutoImportSetting() bool {
	return a.Settings.AutoImport
}

func (a *App) SetAutoImportSetting(value bool) {
	a.Settings.AutoImport = value
	a.Settings.Save()
}

func (a *App) GetShowNativeTitleBarSetting() bool {
	return a.Settings.ShowNativeTitleBar
}

func (a *App) SetShowNativeTitleBarSetting(value bool) {
	a.Settings.ShowNativeTitleBar = value
	a.Settings.Save()
}
