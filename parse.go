package sics

import (
	"errors"
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"strings"
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
	var (
		skin             Skin
		skinBatsmanIndex int
	)
	rows, err := oversTable.Search(".//tr")
	if err != nil {
		return
	}
	for _, row := range rows {
		// Figure out type of row and handle
		cells, err := row.Search(".//td")
		if err != nil {
			return i, err
		}
		if len(cells) == 0 {
			continue
		}
		if len(cells) > 1 && cells[0].Attr("class") == "TeamHeader" {
			i.Team = strings.TrimSpace(cells[1].Content())
		} else if len(cells) > 2 &&
			cells[2].Attr("class") == "Bwl" {
			skin = Skin{}
			skinBatsmanIndex = 0
			for _, cell := range cells {
				if cell.Attr("class") == "Bwl" {
					skin.Overs = append(skin.Overs, Over{
						Bowler: strings.TrimSpace(cell.Content()),
					})
				}
			}
		} else if skinBatsmanIndex <= 1 && len(skin.Overs) > 0 &&
			cells[0].Attr("class") == "BatsmanCell" {
			skin.Batsmen[skinBatsmanIndex] = strings.TrimSpace(
				cells[0].Content())
			overIndex := 0
			ballIndex := 0
			for _, cell := range cells {
				if (cell.Attr("class") == "BallCell" ||
					cell.Attr("class") == "extraBall") &&
					overIndex < len(skin.Overs) {
					var ball *Ball
					ballRaw := strings.ToLower(strings.TrimSpace(
						cell.Content()))
					fmt.Printf("%#v\n", ballRaw)
					if skinBatsmanIndex == 0 {
						ball = &Ball{}
						skin.Overs[overIndex].Balls = append(
							skin.Overs[overIndex].Balls, *ball)
					} else {
						ball = &skin.Overs[overIndex].Balls[ballIndex]
					}
					ball.Bowler = skin.Overs[overIndex].Bowler
					if ballRaw != "" {
						ball.Batsman = skin.Batsmen[skinBatsmanIndex]
						ball.Kind = ballRaw
					}
					// skin.Overs[overIndex].Balls = append(
					// 	skin.Overs[overIndex].Balls, Ball{
					// 		Bowler:  skin.Overs[overIndex],
					// 		Batsman: skin.Batsmen[skinBatsmanIndex],
					// 	})
				} else if cell.Attr("class") == "OverTotalCell rightAligned" {
					overIndex++
					ballIndex = 0
				}
			}
			if skinBatsmanIndex == 1 {
				i.Skins = append(i.Skins, skin)
			}
			skinBatsmanIndex++
		}
	}
	fmt.Printf("%#v\n", i)
	return
}
