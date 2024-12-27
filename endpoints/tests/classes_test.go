package tests

import (
	"context"
	"testing"
)

func TestGetClasses(t *testing.T) {
	c := CreateClient(t)

	classes, err := c.GetClasses(context.Background(), "football")

	if err != nil {
		t.Fatal(err)
	}

	if len := len(classes); len == 0 {
		t.Fatal(len)
	}
}

func TestGetClassesByID(t *testing.T) {
	c := CreateClient(t)

	id := 121

	classes, err := c.GetClassesByID(context.Background(), "football", id)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(classes); len == 0 {
		t.Fatal(len)
	}

	if classes[0].ID != id {
		t.Fail()
	}
}

func TestGetClassesByAlpha(t *testing.T) {
	c := CreateClient(t)

	alpha := "NX"

	classes, err := c.GetClassesByAlpha(context.Background(), "football", alpha)

	if err != nil {
		t.Fatal(err)
	}

	if len := len(classes); len == 0 {
		t.Fatal(len)
	}

	if classes[0].Alpha != alpha {
		t.Fail()
	}
}
