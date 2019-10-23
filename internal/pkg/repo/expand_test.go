package repo

import (
	"testing"
)

func TestRepoExpander_Explode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 string
		want2 string
		want3 string
	}{
		{
			name:  "github-default",
			input: "org/name",
			want1: "",
			want2: "org",
			want3: "name",
		},
		{
			name:  "github",
			input: "g/org/name",
			want1: "g",
			want2: "org",
			want3: "name",
		},
		{
			name:  "bitbucket",
			input: "bb/org/name",
			want1: "bb",
			want2: "org",
			want3: "name",
		},
		{
			name:  "ext",
			input: "bb/org/name.git",
			want1: "bb",
			want2: "org",
			want3: "name",
		},
		{
			name:  "gitscheme",
			input: "github.com:org/name.git",
			want1: "github.com",
			want2: "org",
			want3: "name",
		},
		{
			name:  "full-git",
			input: "git@github.com:org/name.git",
			want1: "github.com",
			want2: "org",
			want3: "name",
		},
		{
			name:  "full-https",
			input: "https://github.com/org/name.git",
			want1: "github.com",
			want2: "org",
			want3: "name",
		},
		{
			name:  "git-inc-dot",
			input: "git@github.com:org/name.some",
			want1: "github.com",
			want2: "org",
			want3: "name.some",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRepoExpander()
			got1, got2, got3 := rs.Explode(tt.input)
			if got1 != tt.want1 {
				t.Errorf("RepoExpander.Explode() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("RepoExpander.Explode() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("RepoExpander.Explode() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestRepoExpander_ExpandTransport(t *testing.T) {
	type args struct {
		transport string
		platfrom  string
		org       string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "github-git",
			args: args{
				transport: "git",
				platfrom:  "github.com",
				org:       "orgname",
				name:      "reponame",
			},
			want:    "git@github.com:orgname/reponame.git",
			wantErr: false,
		},
		{
			name: "github-https",
			args: args{
				transport: "https",
				platfrom:  "github.com",
				org:       "orgname",
				name:      "reponame",
			},
			want:    "https://github.com/orgname/reponame.git",
			wantErr: false,
		},
		{
			name: "unknown transport",
			args: args{
				transport: "thisshouldnotexist",
				platfrom:  "github.com",
				org:       "orgname",
				name:      "reponame",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unknown platform",
			args: args{
				transport: "https",
				platfrom:  "thisshouldnotexist",
				org:       "orgname",
				name:      "reponame",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRepoExpander()
			got, err := rs.ExpandTransport(tt.args.transport, tt.args.platfrom, tt.args.org, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepoExpander.ExpandTransport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RepoExpander.ExpandTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepoExpander_ExpandPlatform(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "github",
			input: "g",
			want:  "github.com",
		},
		{
			name:  "default",
			input: "",
			want:  "github.com",
		},
		{
			name:  "bitbucket",
			input: "bb",
			want:  "bitbucket.org",
		},
		{
			name:  "fallback",
			input: "alkjsdhflkashdflakjhsdf",
			want:  "alkjsdhflkashdflakjhsdf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRepoExpander()
			if got := rs.ExpandPlatform(tt.input); got != tt.want {
				t.Errorf("RepoExpander.ExpandPlatform() = %v, want %v", got, tt.want)
			}
		})
	}
}
