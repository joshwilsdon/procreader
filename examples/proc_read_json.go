// Eg. 'go run examples/proc_read_json.go $$ | json -g'
//
// Usage: <proc_read_json.go> <pid> [...]
//
// returns JSON structures with the properties of the /proc info for the
// specified processes.
//

package main

import (
	"../../procreader"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

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
		proc, err = procreader.ReadProc(pid)
		if err != nil {
			panic(err)
		}
		out, err := json.Marshal(proc)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))
	}
}
