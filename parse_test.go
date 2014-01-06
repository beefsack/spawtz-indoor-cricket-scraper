package sics

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := ioutil.ReadFile(path.Join("test_files", "match.html"))
	if err != nil {
		t.Fatalf("Failed reading test file: %s", err.Error())
	}
	match, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse match from file: %s", err.Error())
	}
	if len(match.Teams) != 2 {
		t.Fatal("Two teams weren't parsed")
	}
	expectedTeam1 := "Kalathumpian Bullants"
	if match.Teams[0].Name != expectedTeam1 {
		t.Fatalf("Team name was not '%s', got '%s' instead", expectedTeam1,
			match.Teams[0].Name)
	}
	expectedTeam2 := "Devon Patrol"
	if match.Teams[1].Name != expectedTeam2 {
		t.Fatalf("Team name was not '%s', got '%s' instead", expectedTeam2,
			match.Teams[1].Name)
	}
	if len(match.Innings[0].Skins) != 4 {
		t.Fatalf("Expected there to be 4 innings for team 1, got %d",
			len(match.Innings[0].Skins))
	}
}
