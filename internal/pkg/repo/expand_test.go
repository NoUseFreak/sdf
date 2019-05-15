package repo

import "testing"

func TestRepoExpander_Expand(t *testing.T) {
	type args struct {
		shortRepo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "github-default",
			args: args{shortRepo: "org/name"},
			want: "github.com/org/name",
		},
		{
			name: "github",
			args: args{shortRepo: "g/org/name"},
			want: "github.com/org/name",
		},
		{
			name: "bitbucket",
			args: args{shortRepo: "bb/org/name"},
			want: "bitbucket.org/org/name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRepoExpander()
			if got := rs.Expand(tt.args.shortRepo); got != tt.want {
				t.Errorf("RepoExpander.Expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
