package sics

const (
	BALL_FAIR              = "f"
	BALL_NO_BALL           = "nb"
	BALL_WIDE              = "w"
	BALL_LEG_SIDE          = "ls"
	BALL_WICKET_BOWLED     = "b"
	BALL_WICKET_CAUGHT     = "c"
	BALL_WICKET_RUN_OUT    = "ro"
	BALL_WICKET_STUMPED    = "s"
	BALL_WICKET_LBW        = "l"
	BALL_WICKET_HIT_WICKET = "h"
)

var BallTypeScores = map[string]int{
	BALL_FAIR:              0,
	BALL_NO_BALL:           2,
	BALL_WIDE:              2,
	BALL_LEG_SIDE:          2,
	BALL_WICKET_BOWLED:     -5,
	BALL_WICKET_CAUGHT:     -5,
	BALL_WICKET_RUN_OUT:    -5,
	BALL_WICKET_STUMPED:    -5,
	BALL_WICKET_LBW:        -5,
	BALL_WICKET_HIT_WICKET: -5,
}

type Ball struct {
	Bowler  string
	Batsmen string
	Type    string
	Score   int
}
