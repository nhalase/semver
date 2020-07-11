package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func Test_runGetReleaseCmd(t *testing.T) {
	type args struct {
		in0  *cobra.Command
		args []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"1.2.0 is extracted from 1.2.0-SNAPSHOT+abc.1", args{args: []string{"1.2.0-SNAPSHOT+abc.1"}}, "1.2.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runGetReleaseCmd(tt.args.in0, tt.args.args)
			if actual != tt.expected {
				t.Errorf("runGetReleaseCmd() expected = %v, actual %v", tt.expected, actual)
			}
		})
	}
}

func init() {
	printf = capturePrintfToActual
}
