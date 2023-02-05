package day_25_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_25"
	"github.com/go-playground/assert/v2"
	"testing"
)

//go:embed testdata/test_input.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestConvertDecimalSumToSnafu(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&testInput)
		sum := day_25.ConvertToDecimalSum(arr)
		assert.Equal(t, 4890, sum)
		assert.Equal(t, "2=-1=0", day_25.ConvertToSnafu(sum))
	})

	t.Run("real input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&input)
		sum := day_25.ConvertToDecimalSum(arr)
		assert.Equal(t, 32969743607087, sum)
		t.Logf("SNAFU: %s", day_25.ConvertToSnafu(sum))
	})
}

func TestConvertToDecimal(t *testing.T) {
	tc := []struct {
		name  string
		snafu []byte
		want  int
	}{
		{
			name:  "test 1=-0-2",
			snafu: []byte("1=-0-2"),
			want:  1747,
		},
		{
			name:  "test 12111",
			snafu: []byte("12111"),
			want:  906,
		},
		{
			name:  "test 2=0=",
			snafu: []byte("2=0="),
			want:  198,
		},
		{
			name:  "test 21",
			snafu: []byte("21"),
			want:  11,
		},
		{
			name:  "test 2=01",
			snafu: []byte("2=01"),
			want:  201,
		},
		{
			name:  "test 111",
			snafu: []byte("111"),
			want:  31,
		},
		{
			name:  "test 20012",
			snafu: []byte("20012"),
			want:  1257,
		},
		{
			name:  "test 112",
			snafu: []byte("112"),
			want:  32,
		},
		{
			name:  "test 1=-1=",
			snafu: []byte("1=-1="),
			want:  353,
		},
		{
			name:  "test 1-12",
			snafu: []byte("1-12"),
			want:  107,
		},
		{
			name:  "test 12",
			snafu: []byte("12"),
			want:  7,
		},
		{
			name:  "test 1=",
			snafu: []byte("1="),
			want:  3,
		},
		{
			name:  "test 122",
			snafu: []byte("122"),
			want:  37,
		},
	}
	for _, c := range tc {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := day_25.ConvertToDecimal(&c.snafu)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestConvertToSnafu(t *testing.T) {
	tc := []struct {
		name string
		sum  int
		want string
	}{
		{
			name: "test 1",
			sum:  1,
			want: "1",
		},
		{
			name: "test 2",
			sum:  2,
			want: "2",
		},
		{
			name: "test 3",
			sum:  3,
			want: "1=",
		},
		{
			name: "test 4",
			sum:  4,
			want: "1-",
		},
		{
			name: "test 5",
			sum:  5,
			want: "10",
		},
		{
			name: "test 6",
			sum:  6,
			want: "11",
		},
		{
			name: "test 7",
			sum:  7,
			want: "12",
		},
		{
			name: "test 8",
			sum:  8,
			want: "2=",
		},
		{
			name: "test 9",
			sum:  9,
			want: "2-",
		},
		{
			name: "test 10",
			sum:  10,
			want: "20",
		},
		{
			name: "test 15",
			sum:  15,
			want: "1=0",
		},
		{
			name: "test 20",
			sum:  20,
			want: "1-0",
		},
		{
			name: "test 2022",
			sum:  2022,
			want: "1=11-2",
		},
		{
			name: "test 12345",
			sum:  12345,
			want: "1-0---0",
		},
		{
			name: "test 314159265",
			sum:  314159265,
			want: "1121-1110-1=0",
		},
	}
	for _, c := range tc {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := day_25.ConvertToSnafu(c.sum)
			assert.Equal(t, c.want, got)
		})
	}
}
