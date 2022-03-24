package helper

import (
	"log"
	"time"
)

const layout = "2006-01-02"

func TimeConverter(stringTime string) string {
	t, err := time.Parse(layout, stringTime)
	if err != nil {
		log.Fatal(err)
	}
	return t.Format(time.RFC3339)
}
