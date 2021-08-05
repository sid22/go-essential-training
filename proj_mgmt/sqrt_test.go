package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
		{4.0, 2.0},
		// {4.0, 1.5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f \n", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatalf("error in calculation - %s", err)
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}

func TestFile(t *testing.T) {
	filep := "dt.csv"
	f, err := os.Open(filep)
	if err != nil {
		t.Fatalf("error in opening - %s", err)
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	testCases := []testCase{}
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		tc := testCase{}
		if tc.value, err = strconv.ParseFloat(row[0], 64); err != nil {
			t.Fatalf("error in parsing - %s", err)
		}
		if tc.expected, err = strconv.ParseFloat(row[1], 64); err != nil {
			t.Fatalf("error in parsing - %s", err)
		}
		testCases = append(testCases, tc)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f \n", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatalf("error in calculation - %s", err)
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
