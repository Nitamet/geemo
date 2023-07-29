package backend

import (
	"github.com/natefinch/lumberjack"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const logFileName = "log.log"

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

func LogPanic() {
	if r := recover(); r != nil {
		log.Printf("Panic: %v", r)

		// Save stack trace
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		log.Println("Stack Trace: " + string(buf[0:stackSize]))

		// Emit error to front-end
		EmitError()
	}
}
