package main

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	tt := []struct {
		name        string
		startStacks [][]string
		move        movement
		endStacks   [][]string
	}{
		{
			name: "example move 1",
			startStacks: [][]string{
				[]string{"Z", "N"},
				[]string{"M", "C", "D"},
				[]string{"P"},
			},
			move: movement{from: 1, to: 0, amount: 1},
			endStacks: [][]string{
				[]string{"Z", "N", "D"},
				[]string{"M", "C"},
				[]string{"P"},
			},
		},
	}
	t.Parallel()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			result := moveStacks(tc.startStacks, []movement{tc.move})
			fmt.Println(result)
		})
	}
}

func TestSlice(t *testing.T) {
	t.Run("length", func(t *testing.T) {
		old := []string{"a", "b", "c"}
		old = old[:len(old)-1]

		if len(old) != 2 {
			t.Errorf("length is not 2")
		}
	})
}
