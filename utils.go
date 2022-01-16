package skillshot

import (
	"math"
	"sort"
)

func Precision(x float64, prec int) float64 {
	p := math.Pow10(prec)
	return math.Round(x*p) / p
}

func Score(q float64, i float64) float64 {
	if q < i {
		return 0.0
	}
	if q > i {
		return 1.0
	}
	return 0.5
}

type NormalizeRankFunc func([]Team)

// NormalizeByRank orders and normalizes the Team.Rank to be 0-based
// and accounts for ties
func NormalizeByRank(teams []Team) {
	var ranked bool
	for i := range teams {
		if teams[i].Rank > 0 {
			ranked = true
		}
	}
	if !ranked {
		for i := range teams {
			teams[i].Rank = i
		}
		return
	}

	sort.Slice(
		teams,
		func(i int, j int) bool {
			return teams[i].Rank < teams[j].Rank
		},
	)

	ranks := make([]int, len(teams))

	var s int
	for j := range teams {
		if j > 0 && teams[j-1].Rank < teams[j].Rank {
			s = j
		}
		ranks[j] = s
	}

	for idx := range teams {
		teams[idx].Rank = ranks[idx]
	}
}

// NormalizeByScore orders and normalizes the Team.Score to be 0-based
// and accounts for ties
func NormalizeByScore(teams []Team) {
	var scored bool
	for i := range teams {
		if teams[i].Score > 0 {
			scored = true
		}
	}
	if !scored {
		for i := range teams {
			teams[i].Rank = i
		}
		return
	}

	sort.Slice(
		teams,
		func(i int, j int) bool {
			return teams[i].Score > teams[j].Score
		},
	)

	ranks := make([]int, len(teams))

	var s int
	for j := range teams {
		if j > 0 && teams[j-1].Score > teams[j].Score {
			s = j
		}
		ranks[j] = s
	}

	for idx := range teams {
		teams[idx].Rank = ranks[idx]
	}
}

func C(teams []Team) float64 {
	var c float64
	for _, team := range teams {
		c += (team.Sigma() + BetaSquared)
	}
	return math.Sqrt(c)
}

func SumQ(teams []Team, c float64) []float64 {
	sumQ := make([]float64, len(teams))

	filteredTeams := make(map[int][]Team)
	for q := range teams {
		for i := range teams {
			if teams[i].Rank >= teams[q].Rank {
				filteredTeams[q] = append(filteredTeams[q], teams[i])
			}
		}
	}

	for idx, teams := range filteredTeams {
		for i := range teams {
			sumQ[idx] += math.Exp(teams[i].Mu() / c)
		}
	}

	return sumQ
}

func A(teams []Team) []int {
	a := make([]int, len(teams))

	ranksCount := make(map[int]int)
	for q := range teams {
		for i := range teams {
			if teams[i].Rank == teams[q].Rank {
				ranksCount[q]++
			}
		}
	}

	for idx, count := range ranksCount {
		a[idx] = count
	}

	return a
}

func Gamma(sigma float64, c float64) float64 {
	return math.Sqrt(sigma) / c
}
