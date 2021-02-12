package golog

import (
	"fmt"
	"log"
	"bytes"
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
	buffer bytes.Buffer

	errorLog = log.New(&buffer, PrimaryRed + "[ Error ]" + Reset + " ", log.Ldate|log.Ltime|log.Lmsgprefix)
	successLog = log.New(&buffer, PrimaryGreen + "[ Success ]" + Reset + " ", log.Ldate|log.Ltime|log.Lmsgprefix)
	warningLog = log.New(&buffer, PrimaryYellow + "[ Warning ]" + Reset + " ", log.Ldate|log.Ltime|log.Lmsgprefix)
	infoLog = log.New(&buffer, PrimaryCyan + "[ Info ]" + Reset + " ", log.Ldate|log.Ltime|log.Lmsgprefix)
)

func Error(logType string, message string) {
	if logType == "primary" {
		errorLog.Println(SecondaryRed + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	} else {
		errorLog.Println(SecondaryRed + "-> " + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	}
}

func Success(logType string, message string) {
	if logType == "primary" {
		successLog.Println(SecondaryGreen + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	} else {
		successLog.Println(SecondaryGreen + "-> " + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	}
}

func Warning(logType string, message string) {
	if logType == "primary" {
		warningLog.Println(SecondaryYellow + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	} else {
		warningLog.Println(SecondaryYellow + "-> " + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	}
}

func Info(logType string, message string) {
	if logType == "primary" {
		infoLog.Println(SecondaryCyan + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	} else {
		infoLog.Println(SecondaryCyan + "-> " + message + "." + Reset)

		fmt.Print(&buffer)
		buffer.Reset()
	}
}
