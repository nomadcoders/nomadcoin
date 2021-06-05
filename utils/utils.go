package utils

import "log"

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
