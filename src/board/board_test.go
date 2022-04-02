package board

import "testing"

func Test_coordsToPos(t *testing.T) {
	type args struct {
		coords string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a1 is zero", args{"a1"}, 0},
		{"b1 is one", args{"b1"}, 1},
		{"h1 is seven", args{"h1"}, 7},
		{"a2 is eight", args{"a2"}, 8},
		{"h8 is 63", args{"h8"}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coordsToPos(tt.args.coords); got != tt.want {
				t.Errorf("coordsToPos(%q) = %v, want %v", tt.args.coords, got, tt.want)
			}
		})
	}
}
