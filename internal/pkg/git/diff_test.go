package git

import "testing"

func TestExistsIn(t *testing.T) {
	type args struct {
		repoURL    string
		devBranch  string
		mainBranch string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				repoURL:    "https://github.com/src-d/go-siva",
				devBranch:  "appveyor",
				mainBranch: "master",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "success",
			args: args{
				repoURL:    "https://github.com/src-d/go-siva",
				devBranch:  "master",
				mainBranch: "appveyor",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExistsIn(tt.args.repoURL, tt.args.devBranch, tt.args.mainBranch)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExistsIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExistsIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
