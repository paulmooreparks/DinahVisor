// Command dinahvisor is the desktop GUI for Dinah. The Wails shell, the
// pages and the command log are not built yet; this file exists so the
// module has a package, which is what the CI workflow needs in order to
// build, vet and test anything at all.
package main

import "fmt"

func main() {
	fmt.Println(greeting())
}

// greeting returns the line the unbuilt binary prints. It exists as a
// separate function so there is something for a test to call, which keeps
// the test job honest rather than vacuous while the real program is written.
func greeting() string {
	return "DinahVisor is not built yet."
}
