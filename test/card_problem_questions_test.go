package test

import (
	"fmt"
	"github.com/LilMonk/go-tut/pkg/helper"
	"github.com/LilMonk/go-tut/pkg/probability"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

const SAMPLE_SIZE = 100000
var TOTAL_HANDS = probability.CreateHandsCombinations(probability.CreateDeck(), 5)


func convertToSet(hands *[][]probability.Card) mapset.Set[string] {
	result := mapset.NewSet[string]()
	for _, hand := range *hands {
		result.Add(fmt.Sprintf("%v", hand))
	}
	return result
}

// probability of being dealt a flush (5 cards of the same suit)
func TestProbabiltyOfFlush(t *testing.T) {
	hands := helper.Sample(TOTAL_HANDS, SAMPLE_SIZE)

	flush := make([][]probability.Card, 0)

	for _, hand := range hands {
		is_flush := false
		suit_count := make(map[string]int)
		for _, card := range hand {
			suit_count[card.Suit]++
		}
		for _, count := range suit_count {
			if (count == 5) && (!is_flush) {
				flush = append(flush, hand)
				is_flush = true
			}
		}
	}

	hands_set := convertToSet(&hands)
	flush_set := convertToSet(&flush)

	probability_of_flush := probability.CalcProb(flush_set, hands_set)
	t.Logf("The probability of being dealt a flush is %d / %d = %f", probability_of_flush.Numerator, probability_of_flush.Denominator, float64(probability_of_flush.Numerator)/float64(probability_of_flush.Denominator))

	// assert that the probability of being dealt a flush is approximately 0.001640
	if (float64(probability_of_flush.Numerator)/float64(probability_of_flush.Denominator) < 0.0015) || (float64(probability_of_flush.Numerator)/float64(probability_of_flush.Denominator) > 0.0022) {
		t.Errorf("The probability of being dealt a flush is not approximately 0.001640")
	}
}

// probability of four of a kind (4 cards of the same rank)
func TestProbabiltyOfFourOfAKind(t *testing.T) {
	hands := helper.Sample(TOTAL_HANDS, SAMPLE_SIZE)

	four_of_a_kind := make([][]probability.Card, 0)

	for _, hand := range hands {
		is_four_of_a_kind := false
		rank_count := make(map[string]int)
		for _, card := range hand {
			rank_count[card.Rank]++
		}
		for _, count := range rank_count {
			if (count == 4) && (!is_four_of_a_kind) {
				four_of_a_kind = append(four_of_a_kind, hand)
				is_four_of_a_kind = true
			}
		}
	}

	hands_set := convertToSet(&hands)
	four_of_a_kind_set := convertToSet(&four_of_a_kind)

	probability_of_four_of_a_kind := probability.CalcProb(four_of_a_kind_set, hands_set)
	t.Logf("The probability of being dealt four of a kind is %d / %d = %f", probability_of_four_of_a_kind.Numerator, probability_of_four_of_a_kind.Denominator, float64(probability_of_four_of_a_kind.Numerator)/float64(probability_of_four_of_a_kind.Denominator))

	// assert that the probability of being dealt four of a kind is approximately 0.000210
	if (float64(probability_of_four_of_a_kind.Numerator)/float64(probability_of_four_of_a_kind.Denominator) < 0.0001) || (float64(probability_of_four_of_a_kind.Numerator)/float64(probability_of_four_of_a_kind.Denominator) > 0.0003) {
		t.Errorf("The probability of being dealt four of a kind is not approximately 0.000210")
	}
}
