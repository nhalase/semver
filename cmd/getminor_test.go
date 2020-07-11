package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func Test_runGetMinorCmd(t *testing.T) {
	type args struct {
		in0  *cobra.Command
		args []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"2 is extracted from 1.2.0-SNAPSHOT+abc.1", args{args: []string{"1.2.0-SNAPSHOT+abc.1"}}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runGetMinorCmd(tt.args.in0, tt.args.args)
			if actual != tt.expected {
				t.Errorf("runGetMinorCmd() expected = %v, actual %v", tt.expected, actual)
			}
		})
	}
}

func init() {
	printf = capturePrintfToActual
}
