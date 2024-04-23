package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandRank(t *testing.T) {
	cases := []struct {
		name     string
		hand     Hand
		handRank HandRank
	}{
		{
			"high card",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Five, Suit: Diamonds},
				Card{Rank: Queen, Suit: Clubs},
				Card{Rank: Jack, Suit: Spades},
				Card{Rank: Ten, Suit: Hearts},
			},
			HighCard,
		},
		{
			"one pair",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Ace, Suit: Diamonds},
				Card{Rank: King, Suit: Clubs},
				Card{Rank: Queen, Suit: Spades},
				Card{Rank: Ten, Suit: Hearts},
			},
			OnePair,
		},
		{
			"two pairs",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Ace, Suit: Diamonds},
				Card{Rank: King, Suit: Clubs},
				Card{Rank: King, Suit: Spades},
				Card{Rank: Ten, Suit: Hearts},
			},
			TwoPairs,
		},
		{
			"three of a kind",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Ace, Suit: Diamonds},
				Card{Rank: Ace, Suit: Clubs},
				Card{Rank: King, Suit: Spades},
				Card{Rank: Ten, Suit: Hearts},
			},
			ThreeOfAKind,
		},
		{
			"straight",
			Hand{
				Card{Rank: Seven, Suit: Hearts},
				Card{Rank: Eight, Suit: Diamonds},
				Card{Rank: Nine, Suit: Clubs},
				Card{Rank: Ten, Suit: Spades},
				Card{Rank: Jack, Suit: Hearts},
			},
			Straight,
		},
		{
			"flush",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Three, Suit: Hearts},
				Card{Rank: Five, Suit: Hearts},
				Card{Rank: Seven, Suit: Hearts},
				Card{Rank: Ten, Suit: Hearts},
			},
			Flush,
		},
		{
			"full house",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Ace, Suit: Diamonds},
				Card{Rank: Ace, Suit: Clubs},
				Card{Rank: King, Suit: Spades},
				Card{Rank: King, Suit: Hearts},
			},
			FullHouse,
		},
		{
			"four of a kind",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: Ace, Suit: Diamonds},
				Card{Rank: Ace, Suit: Clubs},
				Card{Rank: Ace, Suit: Spades},
				Card{Rank: Ten, Suit: Hearts},
			},
			FourOfAKind,
		},
		{
			"straight flush",
			Hand{
				Card{Rank: Seven, Suit: Hearts},
				Card{Rank: Eight, Suit: Hearts},
				Card{Rank: Nine, Suit: Hearts},
				Card{Rank: Ten, Suit: Hearts},
				Card{Rank: Jack, Suit: Hearts},
			},
			StraightFlush,
		},
		{
			"royal flush",
			Hand{
				Card{Rank: Ace, Suit: Hearts},
				Card{Rank: King, Suit: Hearts},
				Card{Rank: Queen, Suit: Hearts},
				Card{Rank: Jack, Suit: Hearts},
				Card{Rank: Ten, Suit: Hearts},
			},
			RoyalFlush,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rank, _ := tc.hand.GetHandRank()
			assert.Equal(t, tc.handRank, rank)
		})
	}
}
