package sics

import (
	"errors"
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

func Parse(input []byte) (m Match, err error) {
	h, err := gokogiri.ParseHtml(input)
	if err != nil {
		return
	}
	// Find score tables and extract innings out of each
	oversTables, err := h.Search("//table[@class='OversTable']")
	if err != nil {
		return
	}
	if len(oversTables) < 2 {
		err = errors.New(fmt.Sprintf(
			"Could not parse out two innings tables, got %d", len(oversTables)))
		return
	}
	m.Innings = [2]Innings{}
	m.Teams = [2]Team{}
	for i, _ := range m.Innings {
		m.Innings[i], err = ParseInnings(oversTables[i])
		if err != nil {
			err = errors.New(fmt.Sprintf(
				"Could not parse innings: %s", err.Error()))
			return
		}
		m.Teams[i] = Team{
			Name: m.Innings[i].Team,
		}
	}
	return
}

func ParseInnings(oversTable xml.Node) (i Innings, err error) {
	teamHeaders, err := oversTable.Search("//td[@class='TeamHeader']")
	if err != nil {
		return
	}
	if len(teamHeaders) == 0 {
		err = errors.New("Couldn't parse team name")
		return
	}
	i.Team = teamHeaders[len(teamHeaders)-1].Content()
	return
}
