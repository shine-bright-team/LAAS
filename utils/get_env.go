package utils

import (
	"log"
	"os"
)

func GetEnv(name string) string {
	value, isfound := os.LookupEnv(name)
	if !isfound {
		log.Fatalf("Couldn't find %s in env", name)
	}
	return value
}
