package backend

import (
	"github.com/natefinch/lumberjack"
	"log"
	"os"
	"path/filepath"
)

const logFileName = "geemo.log"

func InitializeLog() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Panic(err)
	}

	logPath := filepath.FromSlash(configDir + "/geemo/" + logFileName)

	log.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     15,   // days
		Compress:   true, // disabled by default
	})
}
