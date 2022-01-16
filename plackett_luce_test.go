package skillshot_test

import (
	"testing"

	"github.com/blainsmith/skillshot"

	"github.com/frankban/quicktest"
)

func TestPlackettLuce(t *testing.T) {
	qt := quicktest.New(t)

	qt.Run("solo game does not change rating", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		expteams := []skillshot.Team{expteam1}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("2p FFA", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.63523138347365, Sigma: 8.065506316323548}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 22.36476861652635, Sigma: 8.065506316323548}},
			Rank:    1,
		}
		expteams := []skillshot.Team{expteam1, expteam2}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("3p FFA", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2, team3}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.868876552746237, Sigma: 8.204837030780652}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.717219138186557, Sigma: 8.057829747583874}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 21.413904309067206, Sigma: 8.057829747583874}},
			Rank:    2,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("4p FFA", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team4 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2, team3, team4}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.795084971874736, Sigma: 8.263160757613477}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 26.552824984374855, Sigma: 8.179213704945203}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 24.68943500312503, Sigma: 8.083731307186588}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 20.96265504062538, Sigma: 8.083731307186588}},
			Rank:    3,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("5p FFA", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team4 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team5 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team1, team2, team3, team4, team5}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.666666666666668, Sigma: 8.290556877154474}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 26.833333333333332, Sigma: 8.240145629781066}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.72222222222222, Sigma: 8.179996679645559}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 24.055555555555557, Sigma: 8.111796013701358}},
			Rank:    3,
		}
		expteam5 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 20.72222222222222, Sigma: 8.111796013701358}},
			Rank:    4,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4, expteam5}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("3 teams different sized players", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating},
		}
		team3 := skillshot.Team{
			Ratings: []skillshot.Rating{skillshot.DefaultRating, skillshot.DefaultRating, skillshot.DefaultRating},
		}
		teams := []skillshot.Team{team3, team1, team2}

		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 25.939870821784513, Sigma: 8.247641552260456},
				{Mu: 25.939870821784513, Sigma: 8.247641552260456},
				{Mu: 25.939870821784513, Sigma: 8.247641552260456},
			},
			Rank: 0,
		}
		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 27.21366020491262, Sigma: 8.274321317985242},
			},
			Rank: 1,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 21.84646897330287, Sigma: 8.213058173195341},
				{Mu: 21.84646897330287, Sigma: 8.213058173195341},
			},
			Rank: 2,
		}
		expteams := []skillshot.Team{expteam3, expteam1, expteam2}

		skillshot.PlackettLuce(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})

	qt.Run("series", func(qt *quicktest.C) {
		p00 := skillshot.DefaultRating
		p10 := skillshot.DefaultRating
		p20 := skillshot.DefaultRating
		p30 := skillshot.DefaultRating
		p40 := skillshot.DefaultRating

		teams0 := []skillshot.Team{
			{Ratings: []skillshot.Rating{p00}, Rank: 0},
			{Ratings: []skillshot.Rating{p10}, Rank: 1},
			{Ratings: []skillshot.Rating{p20}, Rank: 1},
			{Ratings: []skillshot.Rating{p30}, Rank: 3},
			{Ratings: []skillshot.Rating{p40}, Rank: 3},
		}

		skillshot.PlackettLuce(teams0, skillshot.NormalizeByRank)

		p01 := teams0[0].Ratings[0]
		p11 := teams0[1].Ratings[0]
		p21 := teams0[2].Ratings[0]
		p31 := teams0[3].Ratings[0]
		p41 := teams0[4].Ratings[0]

		p02 := p01
		p32 := p31

		teams1 := []skillshot.Team{
			{Ratings: []skillshot.Rating{p41}, Rank: 0},
			{Ratings: []skillshot.Rating{p21}, Rank: 1},
			{Ratings: []skillshot.Rating{p11}, Rank: 1},
		}

		skillshot.PlackettLuce(teams1, skillshot.NormalizeByRank)

		p42 := teams1[0].Ratings[0]
		p22 := teams1[1].Ratings[0]
		p12 := teams1[2].Ratings[0]

		p43 := p42

		teams2 := []skillshot.Team{
			{Ratings: []skillshot.Rating{p32}, Rank: 0},
			{Ratings: []skillshot.Rating{p12}, Rank: 0},
			{Ratings: []skillshot.Rating{p22}, Rank: 2},
			{Ratings: []skillshot.Rating{p02}, Rank: 2},
		}

		skillshot.PlackettLuce(teams2, skillshot.NormalizeByRank)

		p33 := teams2[0].Ratings[0]
		p13 := teams2[1].Ratings[0]
		p23 := teams2[2].Ratings[0]
		p03 := teams2[3].Ratings[0]

		qt.Assert(p03, quicktest.DeepEquals, skillshot.Rating{Mu: 26.3537611030628, Sigma: 8.111027060497456})
		qt.Assert(p13, quicktest.DeepEquals, skillshot.Rating{Mu: 24.618479788611904, Sigma: 7.905335509558602})
		qt.Assert(p23, quicktest.DeepEquals, skillshot.Rating{Mu: 23.065819512381218, Sigma: 7.822005594848687})
		qt.Assert(p33, quicktest.DeepEquals, skillshot.Rating{Mu: 24.47633240286842, Sigma: 8.106111471280611})
		qt.Assert(p43, quicktest.DeepEquals, skillshot.Rating{Mu: 26.385499684561076, Sigma: 8.05409080928062})
	})
}
