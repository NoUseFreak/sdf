package utils

import (
	"reflect"
	"regexp"
	"testing"
)

func TestReSubMatchMap(t *testing.T) {
	type args struct {
		r   *regexp.Regexp
		str string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "success",
			args: args{
				r:   regexp.MustCompile("^((?P<short>[a-z]+)/)?(?P<org>[^/]+)/(?P<name>[^/]+)"),
				str: "g/theorg/thename",
			},
			want: map[string]string{
				"short": "g",
				"org":   "theorg",
				"name":  "thename",
			},
		},
		{
			name: "failure",
			args: args{
				r:   regexp.MustCompile("^((?P<short>[a-z]+)/)?(?P<org>[^/]+)/(?P<name>[^/]+)"),
				str: "not-match",
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReSubMatchMap(tt.args.r, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReSubMatchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
