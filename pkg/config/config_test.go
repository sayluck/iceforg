package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		opts []opt
	}
	tests := []struct {
		name string
		args args
		want *config
	}{
		{
			name: "loadYamlFile_case",
			args: args{
				opts: []opt{SetConfigFile(
					"F:/goproject/iceforg/resource/config-Files/config.yaml")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *config
			if got = loadConfig(tt.args.opts...); got == nil {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
			t.Logf("conf:%+v\n", got.DB.Mysql)
		})
	}
}
