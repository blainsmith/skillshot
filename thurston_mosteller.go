package skillshot

import "math"

func ThurstonMosteller(teams []Team, nrf NormalizeRankFunc) {
	nrf(teams)

	cachedTeamRatings := make([]Rating, 0)
	for i := range teams {
		cachedTeamRatings = append(cachedTeamRatings, Rating{
			Mu:    teams[i].Mu(),
			Sigma: teams[i].Sigma(),
		})
	}

	for i := range teams {
		iMu := cachedTeamRatings[i].Mu
		iSigma := cachedTeamRatings[i].Sigma

		var omegaSum float64
		var deltaSum float64
		for q := range teams {
			if q != i {
				qMu := cachedTeamRatings[q].Mu
				qSigma := cachedTeamRatings[q].Sigma

				ciq := math.Sqrt(iSigma + qSigma + DoubleBetaSquared)
				deltaMu := (iMu - qMu) / ciq
				sigSqToCiq := iSigma / ciq
				iGamma := Gamma(iSigma, ciq)

				if teams[q].Rank == teams[i].Rank {
					omegaSum += sigSqToCiq * VT(deltaMu, Epsilon/ciq)
					deltaSum += ((iGamma * sigSqToCiq) / ciq) * WT(deltaMu, Epsilon/ciq)
				} else {
					sign := -1.0
					if teams[q].Rank > teams[i].Rank {
						sign = 1.0
					}

					omegaSum += sign * sigSqToCiq * V(sign*deltaMu, Epsilon/ciq)
					deltaSum += ((iGamma * sigSqToCiq) / ciq) * W(sign*deltaMu, Epsilon/ciq)
				}
			}
		}

		for r := range teams[i].Ratings {
			ratingSigma := teams[i].Ratings[r].Sigma
			sigmaSquared := ratingSigma * ratingSigma

			teams[i].Ratings[r].Mu += (sigmaSquared / iSigma) * omegaSum
			teams[i].Ratings[r].Sigma = ratingSigma * math.Sqrt(math.Max(1-(sigmaSquared/iSigma)*deltaSum, Epsilon))
		}
	}
}

func ThurstonMostellerPartial(teams []Team) {}
