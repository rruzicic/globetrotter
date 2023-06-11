package repos

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var neo4jDriver neo4j.Driver

func Connect() {
	driver, err := neo4j.NewDriver("bolt://neo4j:7687", neo4j.BasicAuth("neo4j", "Auth1234", ""))
	if err != nil {
		log.Panic("Could not connect to neo4j. Error: ", err.Error())
	}

	neo4jDriver = driver
}

func Disconnect() {
	if err := neo4jDriver.Close(); err != nil {
		log.Panic("Could not close driver for neo4j. Error", err.Error())
	}
}
