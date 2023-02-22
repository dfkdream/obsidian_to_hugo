package config

import (
	"reflect"
	"testing"
)

func TestFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			"FromFile",
			args{filename: "../test_site/config.yml"},
			Config{
				BaseURL: "https://blog.dfkdream.dev",
				IgnoreFiles: []string{
					"templates",
				},
				Permalinks: map[string]string{
					"posts": "/:2006/:01/:02/:filename",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
