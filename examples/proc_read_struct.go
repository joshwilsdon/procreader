// Eg. 'go run examples/proc_read_struct.go $$'
//
// Usage: <proc_read_struct.go> <pid> [...]
//
// Generates a Go-structure version of the information from PIDs listed as
// arguments. These structures can be added to procreader_test.go as test-cases.
//

package main

import (
	"../../procreader"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var content map[string]string
	var err error
	var pid uint64
	var proc procreader.Proc

	flag.Parse()

	/* treat each arg as a PID */

	for _, arg := range flag.Args() {
		pid, err = strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Printf("'%s' is not an integer\n", arg)
			os.Exit(1)
		}
		proc, content, err = procreader.ReadProcData(pid)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\t%d: {\n", proc.Stat.Pid)
		fmt.Printf("\t\tstatContent: %#v,\n", content["stat"])
		fmt.Printf("\t\tstatmContent: %#v,\n", content["statm"])
		fmt.Printf("\t\tstatusContent: %#v,\n", content["status"])
		fmt.Printf("\t\tcmdlineContent: %#v,\n", content["cmdline"])
		fmt.Printf("\t\tenvironContent: %#v,\n", content["environ"])
		fmt.Printf("\t\texpected: Proc{\n")

		//
		// XXX WARNING:
		//
		// If any of stat/statm/status contain the string 'procreader.' this will
		// break. We have to do this gross Replace here becase %#v adds the
		// procreader. suffix, but in procreader_test.go, procreader is undefined.
		// Apparently we also can't easily get the current package from within that
		// package so we do this instead.
		//
		fmt.Printf("%s", strings.Replace(fmt.Sprintf("\t\t\t%#v,\n", proc.Stat),
			"procreader.", "", -1))
		fmt.Printf("%s", strings.Replace(fmt.Sprintf("\t\t\t%#v,\n", proc.Statm),
			"procreader.", "", -1))
		fmt.Printf("%s", strings.Replace(fmt.Sprintf("\t\t\t%#v,\n", proc.Status),
			"procreader.", "", -1))

		fmt.Printf("\t\t\t%#v,\n", proc.Cmdline)
		fmt.Printf("\t\t\t%#v,\n", proc.Environ)
		fmt.Printf("\t\t},\n")
		fmt.Printf("\t},\n")
	}
}
