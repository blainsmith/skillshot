package skillshot_test

import (
	"testing"

	"github.com/blainsmith/skillshot"

	"github.com/frankban/quicktest"
)

func TestThurstonMosteller(t *testing.T) {
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

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 29.20524620886059, Sigma: 7.632833464033909}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 20.79475379113941, Sigma: 7.632833464033909}},
			Rank:    1,
		}
		expteams := []skillshot.Team{expteam1, expteam2}

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 33.41049241772118, Sigma: 6.861184222487201}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.0, Sigma: 6.861184222487201}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 16.58950758227882, Sigma: 6.861184222487201}},
			Rank:    2,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3}

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 37.61573862658177, Sigma: 5.99095578185474}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 29.20524620886059, Sigma: 5.99095578185474}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 20.79475379113941, Sigma: 5.99095578185474}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 12.38426137341823, Sigma: 5.99095578185474}},
			Rank:    3,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4}

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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
			Ratings: []skillshot.Rating{{Mu: 41.82098483544236, Sigma: 4.970639136506507}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 33.41049241772118, Sigma: 4.970639136506507}},
			Rank:    1,
		}
		expteam3 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 25.0, Sigma: 4.970639136506507}},
			Rank:    2,
		}
		expteam4 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 16.58950758227882, Sigma: 4.970639136506507}},
			Rank:    3,
		}
		expteam5 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 8.17901516455764, Sigma: 4.970639136506507}},
			Rank:    4,
		}
		expteams := []skillshot.Team{expteam1, expteam2, expteam3, expteam4, expteam5}

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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
				{Mu: 25.72407717517514, Sigma: 8.154234192355084},
				{Mu: 25.72407717517514, Sigma: 8.154234192355084},
				{Mu: 25.72407717517514, Sigma: 8.154234192355084},
			},
			Rank: 0,
		}
		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 34.0010841338675, Sigma: 7.7579369709569805},
			},
			Rank: 1,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{
				{Mu: 15.274838690957358, Sigma: 7.373381474061001},
				{Mu: 15.274838690957358, Sigma: 7.373381474061001},
			},
			Rank: 2,
		}
		expteams := []skillshot.Team{expteam3, expteam1, expteam2}

		skillshot.ThurstonMosteller(teams, skillshot.NormalizeByRank)

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

		skillshot.ThurstonMosteller(teams0, skillshot.NormalizeByRank)

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

		skillshot.ThurstonMosteller(teams1, skillshot.NormalizeByRank)

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

		skillshot.ThurstonMosteller(teams2, skillshot.NormalizeByRank)

		p33 := teams2[0].Ratings[0]
		p13 := teams2[1].Ratings[0]
		p23 := teams2[2].Ratings[0]
		p03 := teams2[3].Ratings[0]

		qt.Assert(p03, quicktest.DeepEquals, skillshot.Rating{Mu: 18.712293144521784, Sigma: 3.3771652952566362})
		qt.Assert(p13, quicktest.DeepEquals, skillshot.Rating{Mu: 26.994278073655067, Sigma: 3.2558490006114136})
		qt.Assert(p23, quicktest.DeepEquals, skillshot.Rating{Mu: 22.850233064805856, Sigma: 3.2677270391160795})
		qt.Assert(p33, quicktest.DeepEquals, skillshot.Rating{Mu: 27.276887427210042, Sigma: 3.392300886397173})
		qt.Assert(p43, quicktest.DeepEquals, skillshot.Rating{Mu: 22.64501648101657, Sigma: 3.755436639423349})
	})
}
