package golog

import (
	"os"
	"fmt"
	"log"
	"sort"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"

	"github.com/MrAndreID/gohelpers"
)

const (
	Reset = "\033[0m"

	PrimaryRed = "\033[1;41m"
	PrimaryGreen = "\033[1;42m"
	PrimaryYellow = "\033[1;43m"
	PrimaryCyan = "\033[1;46m"

	SecondaryRed = "\033[0;91m"
	SecondaryGreen = "\033[0;92m"
	SecondaryYellow = "\033[0;93m"
	SecondaryCyan = "\033[0;96m"
)

var (
	Limit int
	LogLevel, LastUpdate string
	CurrentDate = time.Now().Format("2006-01-02")
	CurrentDatetime = time.Now().Format("2006-01-02 15:04:05")
)

func Init(logLevel string, limit int) {
	Limit = limit
	LogLevel = logLevel
	newLine := gohelpers.GetNewLine()

	ignoreFile, _ := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer ignoreFile.Close()

	contentIgnoreFile, _ := ioutil.ReadFile(".gitignore")
	contents := strings.Split(string(contentIgnoreFile), newLine)

	var availabilityLogFolder bool = false
	for i := 0; i < len(contents); i++ {
		if contents[i] == "logs/" {
			availabilityLogFolder = true
		}
	}

	if availabilityLogFolder == false {
		ignoreFile.Write([]byte("logs/" + newLine))
	}

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
}

func ExecutionLimit() {
	if LastUpdate != CurrentDate {
		if logFiles, _ := filepath.Glob("logs/*"); len(logFiles) > Limit + 1 {
			sort.Strings(logFiles)

			os.Remove(logFiles[0])

			LastUpdate = CurrentDate
		}
	}
}

func Error(message string) {
	if LogLevel == "all" || LogLevel == "error" {
		logFile, _ := os.OpenFile("logs/" + CurrentDate + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		errorLog := log.New(logFile, "[ ERROR ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
		errorLog.Println(message)

		ExecutionLimit()
	}

	fmt.Println(CurrentDatetime + " " + PrimaryRed + "[ ERROR ]" + SecondaryRed + " " + message + Reset)
}

func Success(message string) {
	if LogLevel == "all" || LogLevel == "success" {
		logFile, _ := os.OpenFile("logs/" + CurrentDate + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		successLog := log.New(logFile, "[ SUCCESS ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
		successLog.Println(message)

		ExecutionLimit()
	}

	fmt.Println(CurrentDatetime + " " + PrimaryGreen + "[ SUCCESS ]" + SecondaryGreen + " " + message + Reset)
}

func Warning(message string) {
	if LogLevel == "all" || LogLevel == "warning" {
		logFile, _ := os.OpenFile("logs/" + CurrentDate + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		warningLog := log.New(logFile, "[ WARNING ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
		warningLog.Println(message)

		ExecutionLimit()
	}

	fmt.Println(CurrentDatetime + " " + PrimaryYellow + "[ WARNING ]" + SecondaryYellow + " " + message + Reset)
}

func Info(message string) {
	if LogLevel == "all" || LogLevel == "info" {
		logFile, _ := os.OpenFile("logs/" + CurrentDate + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer logFile.Close()

		infoLog := log.New(logFile, "[ INFO ] ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile)
		infoLog.Println(message)

		ExecutionLimit()
	}

	fmt.Println(CurrentDatetime + " " + PrimaryCyan + "[ INFO ]" + SecondaryCyan + " " + message + Reset)
}
