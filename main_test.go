package main

import "testing"

func TestGet(t *testing.T) {
	type args struct {
		id   string
		name string
	}

	testCases := map[string]struct {
		args args
		want *Item
	}{
		"success": {
			args: args{
				id:   "100001",
				name: "test",
			},
			want: &Item{
				ID:   "100001",
				Name: "test",
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := GetItem(tc.args.id, tc.args.name)
			if got != tc.want {
				t.Fatalf("GetItem() = %v, want = %v", got, tc.want)
			}
		})
	}
}