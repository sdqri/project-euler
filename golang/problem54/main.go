package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Suit string

var (
	Clubs    Suit = "C"
	Diamonds Suit = "D"
	Hearts   Suit = "H"
	Spades   Suit = "S"
)

var SuitMap = map[string]Suit{
	string(Clubs):    Clubs,
	string(Diamonds): Diamonds,
	string(Hearts):   Hearts,
	string(Spades):   Spades,
}

type Rank string

var (
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "T"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
	Ace   Rank = "A"
)

var RankMap = map[string]Rank{
	string(Two):   Two,
	string(Three): Three,
	string(Four):  Four,
	string(Five):  Five,
	string(Six):   Six,
	string(Seven): Seven,
	string(Eight): Eight,
	string(Nine):  Nine,
	string(Ten):   Ten,
	string(Jack):  Jack,
	string(Queen): Queen,
	string(King):  King,
	string(Ace):   Ace,
}

var RankScoreMap = map[Rank]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   0xA,
	Jack:  0xB,
	Queen: 0xC,
	King:  0xD,
	Ace:   0xE,
}

type Card struct {
	Rank
	Suit
}

func (card Card) String() string {
	return fmt.Sprintf("%s%s", string(card.Rank), string(card.Suit))
}

var (
	ErrImposibleCardLength error = errors.New("card length isn't equal to two")
	ErrInvalidRank         error = errors.New("rank value is invalid")
	ErrInvalidSuit         error = errors.New("suit value is invalid")
)

func ParseCard(cardStr string) (Card, error) {
	newCard := Card{}

	if len(cardStr) != 2 {
		return newCard, ErrImposibleCardLength
	}

	// Parse Rank
	rankStr := string(cardStr[0])
	rank, ok := RankMap[rankStr]
	if !ok {
		return newCard, ErrInvalidRank
	}
	newCard.Rank = rank

	suitStr := string(cardStr[1])
	suit, ok := SuitMap[suitStr]
	if !ok {
		return newCard, ErrInvalidSuit
	}
	newCard.Suit = suit

	return newCard, nil
}

var _ sort.Interface = &Hand{}

type Hand [5]Card

func (h *Hand) Len() int {
	return len(h)
}

func (h *Hand) Less(i, j int) bool {
	return RankScoreMap[h[i].Rank] > RankScoreMap[h[j].Rank]
}

func (h *Hand) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

type HandRank int

var (
	HighCard      HandRank = 0
	OnePair       HandRank = 1000
	TwoPairs      HandRank = 2000
	ThreeOfAKind  HandRank = 3000
	Straight      HandRank = 4000
	Flush         HandRank = 5000
	FullHouse     HandRank = 6000
	FourOfAKind   HandRank = 7000
	StraightFlush HandRank = 8000
	RoyalFlush    HandRank = 9000
)

func (h Hand) IsStraight() bool {
	ranksOrder := "A23456789TJQKA"
	for _, card := range h {
		ranksOrder = strings.ReplaceAll(ranksOrder, string(card.Rank), "-")
	}

	ordered := 0
	for _, r := range ranksOrder {
		if r == '-' {
			ordered++
		} else {
			ordered = 0
		}

		if ordered == 5 {
			return true
		}
	}

	return false
}

func (h Hand) GetHandRank() (HandRank, int) { // HandRank, Score
	sort.Sort(&h)
	ranksMap := make(map[Rank]int)
	suitsMap := make(map[Suit]int)
	for _, card := range h {
		s := suitsMap[card.Suit]
		suitsMap[card.Suit] = s + 1
		r := ranksMap[card.Rank]
		ranksMap[card.Rank] = r + 1
	}

	_, hasAce := ranksMap[Ace]
	_, hasKing := ranksMap[King]
	if h.IsStraight() && len(suitsMap) == 1 && hasAce && hasKing {
		return RoyalFlush, int(RoyalFlush)
	}

	if h.IsStraight() && len(suitsMap) == 1 {
		return StraightFlush, int(StraightFlush) + RankScoreMap[h[1].Rank]
	}

	for rank, count := range ranksMap {
		if count == 4 {
			return FourOfAKind, int(FourOfAKind) + RankScoreMap[rank]
		} else if count == 3 && len(ranksMap) == 2 {
			return FullHouse, int(FullHouse) + RankScoreMap[rank]
		}
	}

	if len(suitsMap) == 1 {
		return Flush, int(Flush)
	}

	if h.IsStraight() {
		return Straight, int(StraightFlush) + RankScoreMap[h[1].Rank]
	}

	for rank, count := range ranksMap {
		if count == 3 {
			return ThreeOfAKind, int(ThreeOfAKind) + RankScoreMap[rank]
		} else if count == 2 && len(ranksMap) == 3 {
			i := 0
			multiplyer := 14
			score := int(TwoPairs)
			for i < len(h) {
				count := ranksMap[h[i].Rank]
				if count == 2 {
					score += multiplyer * RankScoreMap[h[i].Rank]
					multiplyer /= 14
					i += 2
				}
				i++
			}
			return TwoPairs, score
		} else if count == 2 && len(ranksMap) == 4 {
			return OnePair, int(OnePair) + RankScoreMap[rank]
		}
	}

	return HighCard, int(HighCard)
}

type Result string

var (
	Win  Result = "win"
	Lose Result = "lose"
	Tie  Result = "tie"
)

func expose(hi, hj Hand) Result {
	sort.Sort(&hi)
	sort.Sort(&hj)
	_, scorei := hi.GetHandRank()
	_, scorej := hj.GetHandRank()
	if scorei > scorej {
		return Win
	} else if scorei < scorej {
		return Lose
	}

	for idx := 0; idx < 5; idx++ {
		if RankScoreMap[hi[idx].Rank] > RankScoreMap[hj[idx].Rank] {
			return Win
		} else if RankScoreMap[hi[idx].Rank] < RankScoreMap[hj[idx].Rank] {
			return Lose
		}
	}

	return Tie
}

func main() {
	start := time.Now()

	file, err := os.OpenFile("./0054_poker.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	player1Wins := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		p1c1, _ := ParseCard(line[0:2])
		p1c2, _ := ParseCard(line[3:5])
		p1c3, _ := ParseCard(line[6:8])
		p1c4, _ := ParseCard(line[9:11])
		p1c5, _ := ParseCard(line[12:14])

		h1 := Hand{p1c1, p1c2, p1c3, p1c4, p1c5}

		p2c1, _ := ParseCard(line[15:17])
		p2c2, _ := ParseCard(line[18:20])
		p2c3, _ := ParseCard(line[21:23])
		p2c4, _ := ParseCard(line[24:26])
		p2c5, _ := ParseCard(line[27:29])

		h2 := Hand{p2c1, p2c2, p2c3, p2c4, p2c5}

		if expose(h1, h2) == Win {
			player1Wins++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("number of times player 1 wins : %v (elapsed=%v)\n", player1Wins, elapsed)
}
