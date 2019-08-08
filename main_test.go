package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
				ID:    "100001",
				Name:  "test",
				token: "token",
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := GetItem(tc.args.id, tc.args.name)
			opt := cmp.AllowUnexported(*got)
			if diff := cmp.Diff(got, tc.want, opt); diff != "" {
				t.Fatalf("GetItem() = %v, want = %v", got, tc.want)
			}
		})
	}
}
