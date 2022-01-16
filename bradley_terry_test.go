package skillshot_test

import (
	"testing"

	"github.com/blainsmith/skillshot"

	"github.com/frankban/quicktest"
)

func TestBradleyTerry(t *testing.T) {
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

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 30.2704627669473, Sigma: 7.788474807872566}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.0, Sigma: 7.788474807872566}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 19.7295372330527, Sigma: 7.788474807872566}},
			Rank:    2,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3}

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 32.90569415042095, Sigma: 7.5012190693964005}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.63523138347365, Sigma: 7.5012190693964005}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 22.36476861652635, Sigma: 7.5012190693964005}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 17.09430584957905, Sigma: 7.5012190693964005}},
			Rank:    3,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4}

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 35.5409255338946, Sigma: 7.202515895247076}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 30.2704627669473, Sigma: 7.202515895247076}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.0, Sigma: 7.202515895247076}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 19.729537233052703, Sigma: 7.202515895247076}},
			Rank:    3,
		}
		expteam5 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 14.4590744661054, Sigma: 7.202515895247076}},
			Rank:    4,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4, expteam5}

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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
				{Mu: 25.992743915179297, Sigma: 8.19709997489984},
				{Mu: 25.992743915179297, Sigma: 8.19709997489984},
				{Mu: 25.992743915179297, Sigma: 8.19709997489984},
			},
			Rank: 0,
		}
		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 28.48909130001799, Sigma: 8.220848339985736},
			},
			Rank: 1,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 20.518164784802718, Sigma: 8.127515465304823},
				{Mu: 20.518164784802718, Sigma: 8.127515465304823},
			},
			Rank: 2,
		}
		expteams := []skillshot.Team{expteam3, expteam1, expteam2}

		skillshot.BradleyTerry(teams, skillshot.NormalizeByRank)

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

		skillshot.BradleyTerry(teams0, skillshot.NormalizeByRank)

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

		skillshot.BradleyTerry(teams1, skillshot.NormalizeByRank)

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

		skillshot.BradleyTerry(teams2, skillshot.NormalizeByRank)

		p33 := teams2[0].Ratings[0]
		p13 := teams2[1].Ratings[0]
		p23 := teams2[2].Ratings[0]
		p03 := teams2[3].Ratings[0]

		qt.Assert(p03, quicktest.DeepEquals, skillshot.Rating{Mu: 27.643471362460662, Sigma: 6.716636757697851})
		qt.Assert(p13, quicktest.DeepEquals, skillshot.Rating{Mu: 28.844979380406198, Sigma: 6.310098391579701})
		qt.Assert(p23, quicktest.DeepEquals, skillshot.Rating{Mu: 20.705805440763985, Sigma: 6.310098391579701})
		qt.Assert(p33, quicktest.DeepEquals, skillshot.Rating{Mu: 24.387640206683077, Sigma: 6.6687559074968545})
		qt.Assert(p43, quicktest.DeepEquals, skillshot.Rating{Mu: 23.354955778428952, Sigma: 6.854096854822289})
	})
}
