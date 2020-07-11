package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func Test_runGetPreCmd(t *testing.T) {
	type args struct {
		in0  *cobra.Command
		args []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{"SNAPSHOT is extracted from 1.2.0-SNAPSHOT+abc.1", args{args: []string{"1.2.0-SNAPSHOT+abc.1"}}, "SNAPSHOT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runGetPreCmd(tt.args.in0, tt.args.args)
			if actual != tt.expected {
				t.Errorf("runGetPreCmd() expected = %v, actual %v", tt.expected, actual)
			}
		})
	}
}

func init() {
	printf = capturePrintfToActual
}
