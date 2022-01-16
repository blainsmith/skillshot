package skillshot_test

import (
	"github.com/blainsmith/skillshot"
	"testing"

	"github.com/frankban/quicktest"
)

func TestElo(t *testing.T) {
	qt := quicktest.New(t)

	qt.Run("2p FFA", func(qt *quicktest.C) {
		team1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 12.4}},
		}
		team2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 20.0}},
		}
		teams := []skillshot.Team{team1, team2}

		expteam1 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 27.728066051289712, Sigma: 0.0}},
			Rank:    0,
		}
		expteam2 := skillshot.Team{
			Ratings: []skillshot.Rating{{Mu: 4.671933948710288, Sigma: 0.0}},
			Rank:    1,
		}
		expteams := []skillshot.Team{expteam1, expteam2}

		skillshot.Elo(teams, skillshot.NormalizeByRank)

		qt.Assert(teams, quicktest.DeepEquals, expteams)
	})
}
