// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package precess_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/elementequinox"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/precess"
)

func ExampleApproxAnnualPrecession() {
	// Example 21.a, p. 132.
	eq := &coord.Equatorial{
		base.NewRA(10, 8, 22.3).Rad(),
		base.NewAngle(false, 11, 58, 2).Rad(),
	}
	epochFrom := 2000.0
	epochTo := 1978.0
	Δα, Δδ := precess.ApproxAnnualPrecession(eq, epochFrom, epochTo)
	fmt.Printf("%+.3d\n", base.NewFmtHourAngle(Δα.Rad()))
	fmt.Printf("%+.2d\n", base.NewFmtAngle(Δδ.Rad()))
	// Output:
	// +3ˢ.207
	// -17″.71
}

func ExampleApproxPosition() {
	// Example 21.a, p. 132.
	eq := &coord.Equatorial{
		base.NewRA(10, 8, 22.3).Rad(),
		base.NewAngle(false, 11, 58, 2).Rad(),
	}
	epochFrom := 2000.0
	epochTo := 1978.0
	mα := base.NewHourAngle(true, 0, 0, 0.0169)
	mδ := base.NewAngle(false, 0, 0, 0.006)
	precess.ApproxPosition(eq, eq, epochFrom, epochTo, mα, mδ)
	fmt.Printf("%0.1d\n", base.NewFmtRA(eq.RA))
	fmt.Printf("%+0d\n", base.NewFmtAngle(eq.Dec))
	// Output:
	// 10ʰ07ᵐ12ˢ.1
	// +12°04′32″
}

// test example epochs on p. 133 that are not constants in meeus/julian.go
func TestEpoch(t *testing.T) {
	if math.Abs(base.BesselianYearToJDE(1950)-2433282.4235) > 1e-4 {
		t.Fatal("B1950")
	}
	if math.Abs((base.JulianYearToJDE(2050)-2469807.5)/2469807.5) > 1e-15 {
		t.Fatal("J2050")
	}
}

func ExamplePosition() {
	// Example 21.b, p. 135.
	eq := &coord.Equatorial{
		base.NewRA(2, 44, 11.986).Rad(),
		base.NewAngle(false, 49, 13, 42.48).Rad(),
	}
	epochFrom := 2000.0
	jdTo := julian.CalendarGregorianToJD(2028, 11, 13.19)
	epochTo := base.JDEToJulianYear(jdTo)
	precess.Position(eq, eq, epochFrom, epochTo,
		base.NewHourAngle(false, 0, 0, 0.03425),
		base.NewAngle(true, 0, 0, 0.0895))
	fmt.Printf("%0.3d\n", base.NewFmtRA(eq.RA))
	fmt.Printf("%+0.2d\n", base.NewFmtAngle(eq.Dec))
	// Output:
	// 2ʰ46ᵐ11ˢ.331
	// +49°20′54″.54
}

// Exercise, p. 136.
func TestPosition(t *testing.T) {
	eqFrom := &coord.Equatorial{
		base.NewRA(2, 31, 48.704).Rad(),
		base.NewAngle(false, 89, 15, 50.72).Rad(),
	}
	eqTo := &coord.Equatorial{}
	mα := base.NewHourAngle(false, 0, 0, 0.19877)
	mδ := base.NewAngle(true, 0, 0, 0.0152)
	for _, tc := range []struct {
		α, δ string
		jde  float64
	}{
		{"1 22 33.90", "88 46 26.18", base.BesselianYearToJDE(1900)},
		{"3 48 16.43", "89 27 15.38", base.JulianYearToJDE(2050)},
		{"5 53 29.17", "89 32 22.18", base.JulianYearToJDE(2100)},
	} {
		epochTo := base.JDEToJulianYear(tc.jde)
		precess.Position(eqFrom, eqTo, 2000.0, epochTo, mα, mδ)
		αStr := fmt.Sprintf("%.2x", base.NewFmtRA(eqTo.RA))
		δStr := fmt.Sprintf("%.2x", base.NewFmtAngle(eqTo.Dec))
		if αStr != tc.α {
			t.Fatal("got:", αStr, "want:", tc.α)
		}
		if δStr != tc.δ {
			t.Fatal(δStr)
		}
	}
}

func TestPrecessor_Precess(t *testing.T) {
	// Exercise, p. 136.
	eqFrom := &coord.Equatorial{
		RA:  base.NewRA(2, 31, 48.704).Rad(),
		Dec: base.NewAngle(false, 89, 15, 50.72).Rad(),
	}
	mα := base.NewHourAngle(false, 0, 0, .19877)
	mδ := base.NewAngle(false, 0, 0, -.0152)
	epochs := []float64{
		base.JDEToJulianYear(base.B1900),
		2050,
		2100,
	}
	answer := []string{
		"α = 1ʰ22ᵐ33ˢ.90   δ = +88°46′26″.18",
		"α = 3ʰ48ᵐ16ˢ.43   δ = +89°27′15″.38",
		"α = 5ʰ53ᵐ29ˢ.17   δ = +89°32′22″.18",
	}
	eqTo := &coord.Equatorial{}
	for i, epochTo := range epochs {
		precess.Position(eqFrom, eqTo, 2000, epochTo, mα, mδ)
		if answer[i] != fmt.Sprintf("α = %0.2d   δ = %+0.2d",
			base.NewFmtRA(eqTo.RA), base.NewFmtAngle(eqTo.Dec)) {
			t.Fatal(i)
		}
	}
}

func ExampleEclipticPosition() {
	// Example 21.c, p. 137.
	eclFrom := &coord.Ecliptic{
		Lat: 1.76549 * math.Pi / 180,
		Lon: 149.48194 * math.Pi / 180,
	}
	eclTo := &coord.Ecliptic{}
	epochFrom := 2000.0
	epochTo := base.JDEToJulianYear(julian.CalendarJulianToJD(-214, 6, 30))
	precess.EclipticPosition(eclFrom, eclTo, epochFrom, epochTo, 0, 0)
	fmt.Printf("%.3f\n", eclTo.Lon*180/math.Pi)
	fmt.Printf("%+.3f\n", eclTo.Lat*180/math.Pi)
	// Output:
	// 118.704
	// +1.615
}

func ExampleProperMotion3D() {
	// Example 21.d, p. 141.
	eqFrom := &coord.Equatorial{
		RA:  base.NewRA(6, 45, 8.871).Rad(),
		Dec: base.NewAngle(true, 16, 42, 57.99).Rad(),
	}
	mra := base.NewHourAngle(false, 0, 0, -0.03847)
	mdec := base.NewAngle(false, 0, 0, -1.2053)
	r := 2.64           // given in correct unit
	mr := -7.6 / 977792 // magic conversion factor
	eqTo := &coord.Equatorial{}
	fmt.Printf("Δr = %.9f, Δα = %.10f, Δδ = %.10f\n", mr, mra, mdec)
	for _, epoch := range []float64{1000, 0, -1000, -2000, -10000} {
		precess.ProperMotion3D(eqFrom, eqTo, 2000, epoch, r, mr, mra, mdec)
		fmt.Printf("%8.1f  %0.2d  %-0.1d\n", epoch,
			base.NewFmtRA(eqTo.RA), base.NewFmtAngle(eqTo.Dec))
	}
	// Output:
	// Δr = -0.000007773, Δα = -0.0000027976, Δδ = -0.0000058435
	//   1000.0  6ʰ45ᵐ47ˢ.16  -16°22′56″.0
	//      0.0  6ʰ46ᵐ25ˢ.09  -16°03′00″.8
	//  -1000.0  6ʰ47ᵐ02ˢ.67  -15°43′12″.3
	//  -2000.0  6ʰ47ᵐ39ˢ.91  -15°23′30″.6
	// -10000.0  6ʰ52ᵐ25ˢ.72  -12°50′06″.7
}

func ExampleEclipticPrecessor_ReduceElements() {
	// Example 24.a, p. 160.
	ele := &elementequinox.Elements{
		Inc:  47.122 * math.Pi / 180,
		Peri: 151.4486 * math.Pi / 180,
		Node: 45.7481 * math.Pi / 180,
	}
	JFrom := base.JDEToJulianYear(base.BesselianYearToJDE(1744))
	JTo := base.JDEToJulianYear(base.BesselianYearToJDE(1950))
	p := precess.NewEclipticPrecessor(JFrom, JTo)
	p.ReduceElements(ele, ele)
	fmt.Printf("i = %.4f\n", ele.Inc*180/math.Pi)
	fmt.Printf("Ω = %.4f\n", ele.Node*180/math.Pi)
	fmt.Printf("ω = %.4f\n", ele.Peri*180/math.Pi)
	// Output:
	// i = 47.1380
	// Ω = 48.6037
	// ω = 151.4782
}
