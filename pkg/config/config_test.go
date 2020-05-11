package config

import (
	"iceforg/pkg/utils"
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
		//{
		//	name: "loadPropertiesFile_case",
		//	args: args{
		//		opts: []opt{SetConfigFile(
		//			"F:/goproject/iceforg/resource/config-Files/multilingual_zh.properties")},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *config
			if got = loadConfig(tt.args.opts...); got == nil {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}
			utils.PrettyJsonPrint(got)
		})
	}
}
