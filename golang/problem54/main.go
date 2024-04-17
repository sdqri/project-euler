package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

func main() {
	file, err := os.OpenFile("./0054_poker.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
