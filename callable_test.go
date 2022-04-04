package gc

import (
	"testing"
	"time"
)

func TestSetCallableOK(t *testing.T) {
	SetCallable("test", func() (v interface{}, expAt *time.Time, err error) {
		return "vitego", nil, nil
	})

	v, _, err := callableItems["test"]()

	if v != "vitego" || err != nil {
		t.Error("wrong result or error not nil")
	}
}

func TestSetCallableWithGc(t *testing.T) {
	SetCallable("testWithGc", func() (v interface{}, expAt *time.Time, err error) {
		return "vitego", nil, nil
	})

	if Get("testWithGc").(string) != "vitego" {
		t.Error("wrong result")
	}
}
