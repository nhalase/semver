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
	"github.com/blang/semver/v4"
	"strings"

	"github.com/spf13/cobra"
)

var getBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Extract the build metadata part of <version>.",
	Long: `Extract the build metadata part of <version>.

Build metadata MAY be denoted by appending a plus sign and a series of dot 
separated identifiers immediately following the patch or pre-release version. 
Identifiers MUST comprise only ASCII alphanumerics and hyphens [0-9A-Za-z-]. 
Identifiers MUST NOT be empty. Build metadata MUST be ignored when 
determining version precedence. Thus two versions that differ only in the 
build metadata, have the same precedence. 

Examples: 1.0.0-alpha+001, 1.0.0+20130313144700, 1.0.0-beta+exp.sha.5114f85, 
1.0.0+21AF26D3—-117B344092BD.`,
	Args: ValidateGetCmd,
	Run:  runGetBuildCmd,
}

func runGetBuildCmd(_ *cobra.Command, args []string) {
	version, _ := semver.Make(args[0])
	_, _ = printf("%s", strings.Join(version.Build, "."))
}

func init() {
	getCmd.AddCommand(getBuildCmd)
}
