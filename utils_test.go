package skillshot_test

import (
	"testing"

	"github.com/blainsmith/skillshot"

	"github.com/frankban/quicktest"
)

func TestScore(t *testing.T) {
	qt := quicktest.New(t)

	tests := []struct {
		name string
		q    float64
		i    float64
		s    float64
	}{
		{
			name: "returns 1.0 on wins",
			q:    2,
			i:    1,
			s:    1.0,
		},
		{
			name: "returns 1.0 on big wins",
			q:    34,
			i:    6,
			s:    1.0,
		},
		{
			name: "returns 0.0 on losses",
			q:    1,
			i:    2,
			s:    0.0,
		},
		{
			name: "returns 0.0 on big loss",
			q:    3,
			i:    58,
			s:    0.0,
		},
		{
			name: "returns 0.5 on tie",
			q:    1,
			i:    1,
			s:    0.5,
		},
	}

	for _, test := range tests {
		s := skillshot.Score(test.q, test.i)
		qt.Assert(s, quicktest.DeepEquals, test.s)
	}
}

func TestNormalizeByRank(t *testing.T) {
	qt := quicktest.New(t)

	tests := []struct {
		name string
		in   []skillshot.Team
		out  []skillshot.Team
	}{
		{
			name: "no rank",
			in: []skillshot.Team{
				{Rank: 0, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 0, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 0, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 0, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 2, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 3, Ratings: []skillshot.Rating{{5, 5}}},
			},
		},
		{
			name: "basic order",
			in: []skillshot.Team{
				{Rank: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 5, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 1, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 2, Ratings: []skillshot.Rating{{5, 5}}},
				{Rank: 3, Ratings: []skillshot.Rating{{3, 3}}},
			},
		},
		{
			name: "ties",
			in: []skillshot.Team{
				{Rank: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 3, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 1, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 1, Ratings: []skillshot.Rating{{5, 5}}},
				{Rank: 3, Ratings: []skillshot.Rating{{3, 3}}},
			},
		},
	}

	for _, test := range tests {
		skillshot.NormalizeByRank(test.in)
		qt.Assert(test.in, quicktest.DeepEquals, test.out)
	}
}

func TestNormalizeByScore(t *testing.T) {
	qt := quicktest.New(t)

	tests := []struct {
		name string
		in   []skillshot.Team
		out  []skillshot.Team
	}{
		{
			name: "no score",
			in: []skillshot.Team{
				{Score: 0, Ratings: []skillshot.Rating{{3, 3}}},
				{Score: 0, Ratings: []skillshot.Rating{{1, 1}}},
				{Score: 0, Ratings: []skillshot.Rating{{10, 10}}},
				{Score: 0, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Score: 0, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Score: 0, Ratings: []skillshot.Rating{{1, 1}}},
				{Rank: 2, Score: 0, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 3, Score: 0, Ratings: []skillshot.Rating{{5, 5}}},
			},
		},
		{
			name: "basic order",
			in: []skillshot.Team{
				{Score: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Score: 1, Ratings: []skillshot.Rating{{1, 1}}},
				{Score: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Score: 5, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Score: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Score: 5, Ratings: []skillshot.Rating{{5, 5}}},
				{Rank: 2, Score: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 3, Score: 1, Ratings: []skillshot.Rating{{1, 1}}},
			},
		},
		{
			name: "ties",
			in: []skillshot.Team{
				{Score: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Score: 1, Ratings: []skillshot.Rating{{1, 1}}},
				{Score: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Score: 3, Ratings: []skillshot.Rating{{5, 5}}},
			},
			out: []skillshot.Team{
				{Rank: 0, Score: 10, Ratings: []skillshot.Rating{{3, 3}}},
				{Rank: 1, Score: 3, Ratings: []skillshot.Rating{{10, 10}}},
				{Rank: 1, Score: 3, Ratings: []skillshot.Rating{{5, 5}}},
				{Rank: 3, Score: 1, Ratings: []skillshot.Rating{{1, 1}}},
			},
		},
	}

	for _, test := range tests {
		skillshot.NormalizeByScore(test.in)
		qt.Assert(test.in, quicktest.DeepEquals, test.out, quicktest.Commentf(test.name))
	}
}

func TestC(t *testing.T) {
	qt := quicktest.New(t)

	qt.Run("computes default", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2}
		skillshot.NormalizeByRank(teams)

		c := skillshot.C(teams)
		qt.Assert(skillshot.Precision(c, 4), quicktest.Equals, 15.5902)
	})

	qt.Run("computes 5v5", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
			},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
			},
		}
		teams := []skillshot.Team{team1, team2}
		skillshot.NormalizeByRank(teams)

		c := skillshot.C(teams)
		qt.Assert(skillshot.Precision(c, 4), quicktest.Equals, 27.0031)
	})
}

func TestSumQ(t *testing.T) {
	qt := quicktest.New(t)

	qt.Run("computes default", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2}
		skillshot.NormalizeByRank(teams)

		c := skillshot.C(teams)
		sumQ := skillshot.SumQ(teams, c)
		qt.Assert(skillshot.Precision(sumQ[0], 4), quicktest.Equals, 29.6789)
		qt.Assert(skillshot.Precision(sumQ[1], 4), quicktest.Equals, 24.7082)
	})

	qt.Run("computes 5v5", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
			},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
				skillshot.DefaultRating,
			},
		}
		teams := []skillshot.Team{team1, team2}
		skillshot.NormalizeByRank(teams)

		c := skillshot.C(teams)
		sumQ := skillshot.SumQ(teams, c)
		qt.Assert(skillshot.Precision(sumQ[0], 4), quicktest.Equals, 204.8438)
		qt.Assert(skillshot.Precision(sumQ[1], 4), quicktest.Equals, 102.4219)
	})
}

func TestA(t *testing.T) {
	qt := quicktest.New(t)

	qt.Run("computes default", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2}
		skillshot.NormalizeByRank(teams)

		a := skillshot.A(teams)
		qt.Assert(a[0], quicktest.Equals, 1)
		qt.Assert(a[1], quicktest.Equals, 1)
	})

	qt.Run("computes 1 per rank", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		team4 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2, team3, team4}
		skillshot.NormalizeByRank(teams)

		a := skillshot.A(teams)
		qt.Assert(a[0], quicktest.Equals, 1)
		qt.Assert(a[1], quicktest.Equals, 1)
		qt.Assert(a[2], quicktest.Equals, 1)
		qt.Assert(a[3], quicktest.Equals, 1)
	})

	qt.Run("computes number per shared rank", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
			Rank:    1,
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
			Rank:    1,
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
			Rank:    1,
		}
		team4 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
			Rank:    4,
		}
		teams := []skillshot.Team{team1, team2, team3, team4}
		skillshot.NormalizeByRank(teams)

		a := skillshot.A(teams)
		qt.Assert(a[0], quicktest.Equals, 3)
		qt.Assert(a[1], quicktest.Equals, 3)
		qt.Assert(a[2], quicktest.Equals, 3)
		qt.Assert(a[3], quicktest.Equals, 1)
	})
}
