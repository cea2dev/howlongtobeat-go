package hltb

import "testing"

func TestGetSteamProfileIdNonZero(t *testing.T) {
	expected := int64(1)

	actual := getSteamProfileId(expected)

	if actual == nil {
		t.Error("got nil int64 ptr")
	}

	if *actual != expected {
		t.Errorf("expected: %d; actual: %d", expected, *actual)
	}
}

func TestGetSteamProfileIdPositive(t *testing.T) {
	actual := getSteamProfileId(0)

	if actual != nil {
		t.Error("expected nil pointer")
	}
}
