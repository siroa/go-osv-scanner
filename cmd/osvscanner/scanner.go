/*

go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE

*/

package main

import (
	"github.com/siroa/go-osv-scanner/cmd/osvscanner/cmd"
)

func main() {
	cmd.Execute()
}
