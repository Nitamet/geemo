package main

import (
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"log"
)

const version = "v0.4.1"
const updaterURL = "https://geemo.app/"

func setupUpdater() *selfupdate.Updater {
	updater := &selfupdate.Updater{
		CurrentVersion: version, // the current version of your app used to determine if an update is necessary
		ApiURL:         updaterURL,
		BinURL:         updaterURL,
		DiffURL:        updaterURL,
		Dir:            "geemoupdate/",
		CmdName:        "update",
		CheckTime:      1,
		OnSuccessfulUpdate: func() {
			RestartSelf()
		},
	}

	return updater
}

func autoUpdate(updater *selfupdate.Updater) {
	go func() {
		err := updater.BackgroundRun()
		if err != nil {
			log.Println("Error occurred while updating: ", err)
		}
	}()
}
