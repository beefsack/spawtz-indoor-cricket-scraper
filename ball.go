package sics

import (
// "errors"
// "strings"
)

const (
	BALL_FAIR              = "f"
	BALL_NO_BALL           = "n"
	BALL_WIDE              = "w"
	BALL_LEG_SIDE          = "ls"
	BALL_WICKET_BOWLED     = "b"
	BALL_WICKET_CAUGHT     = "c"
	BALL_WICKET_RUN_OUT    = "r"
	BALL_WICKET_STUMPED    = "s"
	BALL_WICKET_LBW        = "l"
	BALL_WICKET_MANKAD     = "m"
	BALL_WICKET_HIT_WICKET = "h"
	BALL_PENALTY           = "p"
)

var BallKindScores = map[string]int{
	BALL_FAIR:              0,
	BALL_NO_BALL:           2,
	BALL_WIDE:              2,
	BALL_LEG_SIDE:          2,
	BALL_WICKET_BOWLED:     -5,
	BALL_WICKET_CAUGHT:     -5,
	BALL_WICKET_RUN_OUT:    -5,
	BALL_WICKET_STUMPED:    -5,
	BALL_WICKET_LBW:        -5,
	BALL_WICKET_MANKAD:     -5,
	BALL_WICKET_HIT_WICKET: -5,
	BALL_PENALTY:           -5,
}

type Ball struct {
	Bowler  string
	Batsman string
	Kind    string
	Score   int
}

func ParseBall(input string) (kind string, score int, err error) {
	// rawInput := strings.ToLower(strings.TrimSpace(input))

	return
}
