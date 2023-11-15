package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadConfig() map[string]string {
	log.Println("Read config")
	bytes, err := os.ReadFile("app.config")

	if err != nil {
		log.Fatalln(err.Error())
	}

	content := string(bytes)
	splitContent := strings.Split(content, "\r\n")
	configDictionary := map[string]string{}
	for _, item := range splitContent {
		kvPair := strings.Split(item, "=")
		if len(kvPair) == 1 { // prevent empty entries
			continue
		}

		configDictionary[kvPair[0]] = kvPair[1]
	}

	return configDictionary
}

func GetAddress(config map[string]string) string {
	host := "localhost"
	port := "3000"

	configuredHost, doesConfiguredHostExist := config["host"]

	if doesConfiguredHostExist {
		host = configuredHost
	}

	configuredPort, doesConfiguredPortExist := config["port"]

	if doesConfiguredPortExist {
		port = configuredPort
	}

	return fmt.Sprintf("%v:%v", host, port)
}
