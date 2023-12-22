package mylogger

import "log"

func LogInfo(input string) {
	log.Println("INFO: " + input)
}

func LogError(input string) {
	log.Println("ERROR: " + input)
}

func LogWarning(input string) {
	log.Println("WARNING: " + input)
}
