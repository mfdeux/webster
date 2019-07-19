package store

import (
	"fmt"
	"testing"
	"time"
)

func TestMakeDateIndex(t *testing.T) {
	indexName := "twitter"
	dateIndexName := MakeDateIndex(indexName)
	currentTime := time.Now().Format("01-02-2006")
	dateIndexNameTest := fmt.Sprintf("%s-%s", indexName, currentTime)
	if dateIndexName != "twitter-07-19-2019" {
		t.Errorf("Got %s, expected %s", dateIndexName, dateIndexNameTest)
		return
	}
}
