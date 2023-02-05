package day_2

import (
	"fmt"
	a "github.com/akolybelnikov/advent-of-code"
)

const (
	Lose = 0
	Draw = 3
	Win  = 6

	Rock     = 1
	Paper    = 2
	Scissors = 3

	ARock     = 65
	BPaper    = 66
	CScissors = 67
	XRock     = 88
	YPaper    = 89
	ZScissors = 90

	XLose = 88
	YDraw = 89
	ZWin  = 90
)

func FindMyStrategyTotal(data []byte) (int, error) {
	return a.HandleBytes(data, myGameScore)
}

func FindElfStrategyTotal(data []byte) (int, error) {
	return a.HandleBytes(data, elfGameScore)
}

func myHand(b byte) (int, error) {
	switch b {
	case XRock:
		return Rock, nil
	case YPaper:
		return Paper, nil
	case ZScissors:
		return Scissors, nil
	default:
		return 0, fmt.Errorf("unknown my hand %s", string(b))
	}
}

func opponentHand(b byte) (int, error) {
	switch b {
	case ARock:
		return ARock, nil
	case BPaper:
		return BPaper, nil
	case CScissors:
		return CScissors, nil
	default:
		return 0, fmt.Errorf("unknown opponent's hand %s", string(b))
	}
}

func myStrategy(b byte) (int, error) {
	switch b {
	case XLose:
		return Lose, nil
	case YDraw:
		return Draw, nil
	case ZWin:
		return Win, nil
	default:
		return 0, fmt.Errorf("unknown my strategy %s", string(b))
	}
}

func myGameScore(b []byte) (int, error) {
	score, err := myHand(b[2])
	if err != nil {
		return 0, err
	}

	opponent, err := opponentHand(b[0])
	if err != nil {
		return 0, err
	}

	me := b[2]

	switch {
	case opponent == ARock && me == YPaper:
		score += Win
	case opponent == BPaper && me == ZScissors:
		score += Win
	case opponent == CScissors && me == XRock:
		score += Win
	case opponent == ARock && me == XRock:
		score += Draw
	case opponent == BPaper && me == YPaper:
		score += Draw
	case opponent == CScissors && me == ZScissors:
		score += Draw
	}

	return score, nil
}

func elfGameScore(b []byte) (int, error) {
	score, err := myStrategy(b[2])
	if err != nil {
		return 0, err
	}

	opponent, err := opponentHand(b[0])
	if err != nil {
		return 0, err
	}

	me := b[2]

	switch {
	case opponent == ARock && me == XLose:
		score += Scissors
	case opponent == ARock && me == YDraw:
		score += Rock
	case opponent == ARock && me == ZWin:
		score += Paper
	case opponent == BPaper && me == XLose:
		score += Rock
	case opponent == BPaper && me == YDraw:
		score += Paper
	case opponent == BPaper && me == ZWin:
		score += Scissors
	case opponent == CScissors && me == XLose:
		score += Paper
	case opponent == CScissors && me == YDraw:
		score += Scissors
	case opponent == CScissors && me == ZWin:
		score += Rock
	}

	return score, nil
}
