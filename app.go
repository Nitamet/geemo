package main

import (
	"changeme/internal/lcu"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	LCU *lcu.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
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
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetLCUState returns the current state of the LCU
func (a *App) GetLCUState() string {
	if a.LCU != nil {
		return a.LCU.UpdateState()
	}

	instance := lcu.TryToGetLCU()
	if instance == nil {
		return "NotLaunched"
	}

	a.LCU = instance

	return a.LCU.UpdateState()
}

func (a *App) GetSummoner() lcu.Summoner {
	if a.LCU == nil {
		return lcu.Summoner{}
	}

	return a.LCU.CurrentSummoner()
}

func (a *App) GetCurrentChampion() int {
	if a.LCU == nil {
		return 0
	}

	championId, _ := a.LCU.SelectedChampion()
	return championId
}
