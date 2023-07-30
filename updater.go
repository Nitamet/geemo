package main

import (
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"log"
)

const version = "v0.1.5"
const updaterURL = "https://geemo.app/"

func setupUpdater() {
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

	go func() {
		err := updater.BackgroundRun()
		if err != nil {
			log.Println("Error occurred while updating: ", err)
		}
	}()
}
