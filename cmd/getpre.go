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

var getPreCmd = &cobra.Command{
	Use:   "pre",
	Short: "Extract the pre-release part of <version>.",
	Long: `Extract the pre-release part of <version>.

A pre-release version MAY be denoted by appending a hyphen and a series of dot 
separated identifiers immediately following the patch version. Identifiers 
MUST comprise only ASCII alphanumerics and hyphens [0-9A-Za-z-]. Identifiers 
MUST NOT be empty. Numeric identifiers MUST NOT include leading zeroes. 
Pre-release versions have a lower precedence than the associated normal 
version. A pre-release version indicates that the version is unstable and 
might not satisfy the intended compatibility requirements as denoted by its 
associated normal version. 

Examples: 1.0.0-alpha, 1.0.0-alpha.1, 1.0.0-0.3.7, 1.0.0-x.7.z.92, 
1.0.0-x-y-z.–.`,
	Args: ValidateGetCmd,
	Run:  runGetPreCmd,
}

func runGetPreCmd(_ *cobra.Command, args []string) {
	version, _ := semver.Make(args[0])
	var preStrings = make([]string, len(version.Pre))
	for i, pre := range version.Pre {
		preStrings[i] = pre.String()
	}
	_, _ = printf("%s", strings.Join(preStrings, "."))
}

func init() {
	getCmd.AddCommand(getPreCmd)
}
