package gc

import (
	"testing"
	"time"
)

func TestSetOK(t *testing.T) {
	Set("test", 10, nil)

	if Get("test").(int) != 10 {
		t.Error("wrong result")
	}
}

func TestGetExpire(t *testing.T) {
	now := time.Now()
	now.Add(time.Duration(10) * time.Minute)
	Set("test2", 10, &now)

	if Get("test2") != nil {
		t.Error("test2 must be nil")
	}
}
