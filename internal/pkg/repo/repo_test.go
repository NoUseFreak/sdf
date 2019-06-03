package repo

import "testing"

func Test_sfdRepo_TargetDir(t *testing.T) {
	type fields struct {
		Platform string
		Org      string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base",
			fields: fields{
				Platform: "platform",
				Org:      "org",
				Name:     "name",
			},
			want: "platform/org/name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sfdRepo{
				Platform: tt.fields.Platform,
				Org:      tt.fields.Org,
				Name:     tt.fields.Name,
			}
			if got := r.TargetDir(); got != tt.want {
				t.Errorf("sfdRepo.TargetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
