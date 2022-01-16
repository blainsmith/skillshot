package skillshot_test

import (
	"github.com/blainsmith/skillshot"
	"math"
	"testing"

	"github.com/frankban/quicktest"
)

func TestV(t *testing.T) {
	qt := quicktest.New(t)

	qt.Assert(skillshot.V(1, 2), quicktest.Equals, 1.5251352044082924)
	qt.Assert(skillshot.V(0, 2), quicktest.Equals, 2.3732157475120528)
	qt.Assert(skillshot.V(0, -1), quicktest.Equals, 0.2875999734906994)
	qt.Assert(skillshot.V(0, 10), quicktest.Equals, 10.0)
}

func TestW(t *testing.T) {
	qt := quicktest.New(t)

	qt.Assert(skillshot.W(1, 2), quicktest.Equals, 0.8009021873172315)
	qt.Assert(skillshot.W(0, 2), quicktest.Equals, 0.8857214892150859)
	qt.Assert(skillshot.W(0, -1), quicktest.Equals, 0.3703137182425503)
	qt.Assert(skillshot.W(0, 10), quicktest.Equals, 0.0)
	qt.Assert(skillshot.W(-1, 10), quicktest.Equals, 1.0)
}

func TestVT(t *testing.T) {
	qt := quicktest.New(t)

	qt.Assert(skillshot.VT(-1000, -100), quicktest.Equals, 1100.0)
	qt.Assert(skillshot.VT(1000, -100), quicktest.Equals, -1100.0)
	qt.Assert(skillshot.VT(-1000, 1000), quicktest.Equals, 0.7978845368663289)
	qt.Assert(skillshot.VT(0, 1000), quicktest.Equals, 0.0)
}

func TestWT(t *testing.T) {
	qt := quicktest.New(t)

	qt.Assert(skillshot.WT(1, 2), quicktest.Equals, 0.38385826878672835)
	qt.Assert(skillshot.WT(0, 2), quicktest.Equals, 0.22625869547437663)
	qt.Assert(skillshot.WT(0, -1), quicktest.Equals, 1.0)
	qt.Assert(skillshot.WT(0, 0), quicktest.Equals, 1.0)
	qt.Assert(math.Round(skillshot.WT(0, 10)), quicktest.Equals, 0.0)
}
