/*
Copyright © 2020 Nicholas J Halase <njhalase@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"testing"
)

func TestValidateGetCmd(t *testing.T) {
	type args struct {
		in0  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"arguments are required", args{}, true},
		{"1.2.3 passes validation", args{args: []string{"1.2.3"}}, false},
		{"1.2 passes validation", args{args: []string{"1.2"}}, true},
		{"1.2.3--SNAPSHOT.ABN passes validation", args{args: []string{"1.2.3--SNAPSHOT.ABN"}}, false},
		{"1.2.3--SNAPSHOT.ABC+abc.123.-z1z passes validation", args{args: []string{"1.2.3--SNAPSHOT.ABC+abc.123.-z1z"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateGetCmd(tt.args.in0, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGetCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
