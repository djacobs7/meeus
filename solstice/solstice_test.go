package solstice_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solstice"
)

func ExampleJune() {
	// Example 27.a, p. 180
	fmt.Printf("%.5f\n", solstice.June(1962))
	// Output:
	// 2437837.39245
}

type eq struct {
	y    int
	d    float64
	h, m int
	s    float64
}

var (
	mar = []eq{
		{1996, 20, 8, 4, 7},
		{1997, 20, 13, 55, 42},
		{1998, 20, 19, 55, 35},
		{1999, 21, 1, 46, 53},
		{2000, 20, 7, 36, 19},

		{2001, 20, 13, 31, 47},
		{2002, 20, 19, 17, 13},
		{2003, 21, 1, 0, 50},
		{2004, 20, 6, 49, 42},
		{2005, 20, 12, 34, 29},
	}
	jun = []eq{
		{1996, 21, 2, 24, 46},
		{1997, 21, 8, 20, 59},
		{1998, 21, 14, 3, 38},
		{1999, 21, 19, 50, 11},
		{2000, 21, 1, 48, 46},

		{2001, 21, 7, 38, 48},
		{2002, 21, 13, 25, 29},
		{2003, 21, 19, 11, 32},
		{2004, 21, 0, 57, 57},
		{2005, 21, 6, 47, 12},
	}
	sep = []eq{
		{1996, 22, 18, 1, 8},
		{1997, 22, 23, 56, 49},
		{1998, 23, 5, 38, 15},
		{1999, 23, 11, 32, 34},
		{2000, 22, 17, 28, 40},

		{2001, 22, 23, 5, 32},
		{2002, 23, 4, 56, 28},
		{2003, 23, 10, 47, 53},
		{2004, 22, 16, 30, 54},
		{2005, 22, 22, 24, 14},
	}
	dec = []eq{
		{1996, 21, 14, 6, 56},
		{1997, 21, 20, 8, 5},
		{1998, 22, 1, 57, 31},
		{1999, 22, 7, 44, 52},
		{2000, 21, 13, 38, 30},

		{2001, 21, 19, 22, 34},
		{2002, 22, 1, 15, 26},
		{2003, 22, 7, 4, 53},
		{2004, 21, 12, 42, 40},
		{2005, 21, 18, 36, 1},
	}
)

func Test2000(t *testing.T) {
	for i := range mar {
		e := &mar[i]
		approx := solstice.March(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 3, e.d) +
			base.NewTime(false, e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("mar %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
	for i := range jun {
		e := &jun[i]
		approx := solstice.June(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 6, e.d) +
			base.NewTime(false, e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("jun %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
	for i := range sep {
		e := &sep[i]
		approx := solstice.September(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 9, e.d) +
			base.NewTime(false, e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("sep %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f day error", math.Abs(vsop87-approx))
		}
	}
	for i := range dec {
		e := &dec[i]
		approx := solstice.December(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 12, e.d) +
			base.NewTime(false, e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("dec %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
}

type seasons struct {
	y              int
	sp, su, au, wi float64
}

var years = []seasons{
	{-4000, 93.55, 89.18, 89.07, 93.44},
	{-3500, 93.83, 89.53, 88.82, 93.07},
	{-3000, 94.04, 89.92, 88.61, 92.67},
	{-2500, 94.20, 90.33, 88.47, 92.25},
	{-2000, 94.28, 90.76, 88.38, 91.81},
	{-1500, 94.30, 91.20, 88.38, 91.37},
	{-1000, 94.25, 91.63, 88.42, 90.94},
	{-500, 94.14, 92.05, 88.53, 90.52},
	{0, 93.96, 92.45, 88.69, 90.13},
	{500, 93.73, 92.82, 88.90, 89.78},
	{1000, 93.44, 93.15, 89.18, 89.47},
	{1500, 93.12, 93.42, 89.50, 89.20},
	{2000, 92.76, 93.65, 89.84, 88.99},
	{2500, 92.37, 93.81, 90.22, 88.84},
	{3000, 91.97, 93.92, 90.61, 88.74},
	{3500, 91.57, 93.96, 91.01, 88.71},
	{4000, 91.17, 93.93, 91.40, 88.73},
	{4500, 90.79, 93.84, 91.79, 88.82},
	{5000, 90.44, 93.70, 92.15, 88.96},
	{5500, 90.11, 93.50, 92.49, 89.15},
	{6000, 89.82, 93.25, 92.79, 89.38},
	{6500, 89.58, 92.96, 93.04, 89.66},
}

func ExampleJune2() {
	// Example 27.b, p. 180.
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	j := solstice.June2(1962, e)
	t := j - 2437836.5 // 0h 1962 June 21
	// result is VSOP87 result given in example 27.a, p. 180
	fmt.Println(base.NewFmtTime(t * 24 * 60 * 60))
	// Output:
	// 21ʰ24ᵐ42ˢ
}

/*
Commented out because results cannot be accurately determined.  The idea was
to use table 27.F, p. 182 to test functions over a wider range than the ten
years of Test2000.  The low accuracy functions of this package would only
agree with table 27.F to .2 day, not too surprising since the stated range
of those functions is only years -1000 to +3000.  The high accuracy functions
though, only agreed to .02 day, not the given precision of .01 day.  I suspect
the reason for this is Meeus using his truncated VSOP87 rather than full VSOP87
to construct the table but I have no way of knowing.

func Test10000(t *testing.T) {
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	for i := range years {
		y := &years[i]
		s0 := solstice.March2(y.y, e)
		s1 := solstice.June2(y.y, e)
		sp := s1 - s0
		if math.Abs(y.sp-sp) > .02 {
			t.Errorf("spring %d got %.2f expected %.2f", y.y, sp, y.sp)
		}
		s2 := solstice.September2(y.y, e)
		su := s2 - s1
		if math.Abs(y.su-su) > .02 {
			t.Errorf("summer %d got %.2f expected %.2f", y.y, su, y.su)
		}
		s3 := solstice.December2(y.y, e)
		au := s3 - s2
		if math.Abs(y.au-au) > .02 {
			t.Errorf("autumn %d got %.2f expected %.2f", y.y, au, y.au)
		}
		s4 := solstice.March2(y.y+1, e)
		wi := s4 - s3
		if math.Abs(y.wi-wi) > .02 {
			t.Errorf("winter %d got %.2f expected %.2f", y.y, wi, y.wi)
		}
	}
}
*/
