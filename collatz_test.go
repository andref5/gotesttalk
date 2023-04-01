package collatz

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollatz(t *testing.T) {
	// TODO: Change to table driven tests
	hailstone, err := Collatz(10)
	if err != nil {
		t.Error(err)
	}
	want := []int{10, 5, 16, 8, 4, 2, 1}
	if reflect.DeepEqual(hailstone, want) == false {
		t.Errorf("want %v, got %v", want, hailstone)
	}
}

func TestCollatzGoRoutine(t *testing.T) {
	// TODO: Change to table driven tests
	hailstone, err := CollatzGoRoutine(10)
	if err != nil {
		t.Error(err)
	}
	want := 7
	if len(hailstone) != want {
		t.Errorf("want %v, got %v", want, len(hailstone))
	}
}

func TestCollatzRaceDetector(t *testing.T) {
	// TODO: Change to table driven tests
	hailstone, err := CollatzWithDataRace(10)
	if err != nil {
		t.Error(err)
	}
	want := 7
	if len(hailstone) != want {
		t.Errorf("want %v, got %v", want, len(hailstone))
	}
}

func BenchmarkCollatz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Collatz(1000000)
	}
}

func BenchmarkCollatzGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = CollatzGoRoutine(1000000)
	}
}
func TestEven(t *testing.T) {
	tests := map[string]struct {
		x       int
		want    int
		wantErr string
	}{
		"even": {
			x:    10,
			want: 5,
		},
		"zero": {
			x:       0,
			wantErr: "0 is not positive",
		},
		"negative": {
			x:       -10,
			wantErr: "-10 is not positive",
		},
		"odd": {
			x:       11,
			wantErr: "11 is not even",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Even(tc.x)
			if tc.wantErr != "" {
				require.EqualError(t, err, tc.wantErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	// TODO: Change to table driven tests
	got, err := Odd(11)
	if err != nil {
		t.Error(err)
	}
	want := 34
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
