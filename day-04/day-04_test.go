package main

import "testing"

func TestOverlap(t *testing.T) {
	tt := []struct {
		name     string
		spaces   []Space
		expected bool
	}{
		{
			name: "s1 start and s2 end are same",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 2, end: 4,
				},
			},
			expected: true,
		},
		{
			name: "s1 start, s2 start and s2 end are same",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 4, end: 4,
				},
			},
			expected: true,
		},
		{
			name: "s1 end, s2 start  are same",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 6, end: 8,
				},
			},
			expected: true,
		},
		{
			name: "s1 end, s2 start and s2 end  are same",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 6, end: 6,
				},
			},
			expected: true,
		},
		{
			name: "s1 inside s2",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 2, end: 8,
				},
			},
			expected: true,
		},
		{
			name: "s2 inside s1",
			spaces: []Space{
				{
					start: 2, end: 8,
				}, {
					start: 4, end: 6,
				},
			},
			expected: true,
		},
		{
			name: "s1 inside s2 same end",
			spaces: []Space{
				{
					start: 4, end: 6,
				}, {
					start: 2, end: 6,
				},
			},
			expected: true,
		},
		{
			name: "s2 inside s1 same end",
			spaces: []Space{
				{
					start: 2, end: 8,
				}, {
					start: 4, end: 8,
				},
			},
			expected: true,
		},
		{
			name: "s1 inside s2 same start",
			spaces: []Space{
				{
					start: 2, end: 4,
				}, {
					start: 2, end: 6,
				},
			},
			expected: true,
		},
		{
			name: "s2 inside s1 same start",
			spaces: []Space{
				{
					start: 2, end: 8,
				}, {
					start: 2, end: 6,
				},
			},
			expected: true,
		},
	}
	t.Parallel()
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			result := isOverlap(tc.spaces)
			if tc.expected != result {
				t.Errorf("%t expected, got %t", tc.expected, result)
			}
		})
	}
}
