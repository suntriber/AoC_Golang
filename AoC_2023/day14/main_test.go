package main

import "testing"

func TestTiltString(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "cubes last", input: "OO.O.O..##", want: "OOOO....##"},
		{name: "no cubes", input: "...OO....O", want: "OOO......."},
		{name: "cube middle", input: ".O...#O..O", want: "O....#OO.."},
		{name: "cube up top", input: "#.#O..#.##", want: "#.#O..#.##"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tiltString(tc.input)
			if got != tc.want {
				t.Errorf("%s error, got %s want %s", tc.name, got, tc.want)
			}
		})
	}
}
