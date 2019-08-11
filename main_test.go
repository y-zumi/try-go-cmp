package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGetAllowUnexported(t *testing.T) {
	type args struct {
		id   string
		name string
	}

	testCases := map[string]struct {
		args args
		want Item
	}{
		"success": {
			args: args{
				id:   "100001",
				name: "test",
			},
			want: Item{
				ID:         "100001",
				Name:       "test",
				secretCode: int64(12345),
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := GetItem(tc.args.id, tc.args.name)
			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tc.want, opt); diff != "" {
				t.Fatalf("GetItem() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestGetIgnoreUnexported(t *testing.T) {
	type args struct {
		id   string
		name string
	}

	testCases := map[string]struct {
		args args
		want Item
	}{
		"success": {
			args: args{
				id:   "100001",
				name: "test",
			},
			want: Item{
				ID:         "100001",
				Name:       "test",
				secretCode: int64(12345),
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := GetItem(tc.args.id, tc.args.name)
			opt := cmpopts.IgnoreUnexported(got)
			if !cmp.Equal(got, tc.want, opt) {
				t.Fatalf("GetItem() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestGetDeepEqual(t *testing.T) {
	type args struct {
		id   string
		name string
	}

	testCases := map[string]struct {
		args args
		want Item
	}{
		"success": {
			args: args{
				id:   "100001",
				name: "test",
			},
			want: Item{
				ID:         "100001",
				Name:       "test",
				secretCode: int64(12345),
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := GetItem(tc.args.id, tc.args.name)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("GetItem() = %v, want = %v", got, tc.want)
			}
		})
	}
}
