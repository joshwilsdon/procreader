// Eg. 'go run examples/proc_read_dockerfields.go $$ | json -g'
//
// Usage: <proc_read_dockerfields.go> <pid> [...]
//
// returns JSON structures with the properties of the /proc info for the
// specified processes.
//
// The fields in the output match those described in
//
// https://github.com/docker/docker/pull/9232
//
// and you can compare these to the output of:
//
// ps -p <pid>[,<pid>,...] -o c,comm,command,cputime,gid,lwp,nice,pcpu,pid,pgid,pmem,ppid,psr,rgid,rss,ruid,start_time,state,stat,tty,uid,vsz
//

package main

import (
	"../../procreader"
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type DockerTop struct {
	C           uint64
	Comm        string
	Command     string
	Cputime     float64
	Gid         uint64
	Lwp         uint64
	Nice        int32
	Pcpu        float32
	Pid         uint64
	Pgid        int64
	Pmem        float32
	Ppid        int64
	Psr         uint64
	Rgid        uint64
	Rss         uint64
	Ruid        uint64
	Start_time  float64
	State       string
	State_flags []string
	Tty         string
	Uid         uint64
	Vsz         uint64
}

type ttyDriver struct {
	name         string
	default_node string
	major        uint32
	minor_min    uint32
	minor_max    uint32
	tty_type     string
}

type ProcUptime struct {
	Uptime float64
	Idle   float64
}

func getHertz() uint64 {
	// In docker, we'd use libcontainer/system which has 'GetClockTicks'
	// to calculate. But that falls back to 100, so we just use 100 for this
	// example.
	return 100
}

func getUptime() (ProcUptime, error) {
	var contents string
	var uptime ProcUptime

	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return uptime, err
	}

	contents = string(data)
	n, err := fmt.Sscanf(contents, "%f %f", &uptime.Uptime, &uptime.Idle)
	if err != nil {
		return uptime, err
	}
	if n != 2 {
		return uptime, errors.New(fmt.Sprintf(
			"getUptime: expected 2 fields, got %d: '%s'", n, contents))
	}

	return uptime, nil
}

func getSystemStartTime() float64 {
	var now int64
	var secs int64
	var nsecs int64

	now = time.Now().UnixNano()

	secs = now / 1000000000
	nsecs = now % 1000000000

	return float64(secs) + (float64(nsecs) / 1000000000)
}

func getStartTime(proc procreader.Proc) float64 {
	return getSystemStartTime() - float64(proc.Stat.Start_time)
}

// Returns system total memory in kB
func getSystemMemory() (uint64, error) {
	var result uint64

	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmpstr := scanner.Text()
		if strings.HasPrefix(tmpstr, "MemTotal:") {
			cnt, err := fmt.Sscanf(tmpstr, "MemTotal:\t%d kB", &result)
			if err != nil {
				return result, err
			}
			if cnt != 1 {
				return result, errors.New(fmt.Sprintf(
					"getSystemMemory(): bad 'MemTotal:' line: %s", tmpstr))
			}
			return result, nil
		}
	}

	return result, errors.New("getSystemMemory(): 'MemTotal:' not found")
}

// https://gitorious.org/procps/procps/source/3a66fba1e934cbd830df572d8d03c05b4f4a5f1e:ps/output.c#L487
func calculateC(proc procreader.Proc) (uint64, error) {
	var c uint64
	pcpu, err := calculatePcpu(proc)
	if err != nil {
		return c, err
	}
	if pcpu > 99.9 {
		pcpu = 99.9
	}
	c = uint64(pcpu)

	return c, nil
}

func calculatePcpu(proc procreader.Proc) (float32, error) {
	var pcpu float32
	var seconds uint64    // seconds of process life
	var total_time uint64 // jiffies used by this process

	uptime, err := getUptime()
	if err != nil {
		return pcpu, err
	}
	Hertz := getHertz()

	if err != nil {
		return pcpu, err
	}

	total_time = proc.Stat.Utime + proc.Stat.Stime
	if uint64(uptime.Uptime) >= (proc.Stat.Start_time / Hertz) {
		seconds = uint64(uptime.Uptime) - (proc.Stat.Start_time / Hertz)
		pcpu = float32(total_time*100/Hertz) / float32(seconds)
	} else {
		seconds = 0
		pcpu = 0
	}

	return pcpu, nil
}

func calculatePmem(proc procreader.Proc) (float32, error) {
	var pmem float32

	kb_main_total, err := getSystemMemory()
	if err != nil {
		return pmem, err
	}

	pmem = (float32(proc.Status.VmRSS) / float32(kb_main_total)) * 100

	if pmem > 99.9 {
		pmem = 99.9
	}

	return pmem, nil
}

func getState(proc procreader.Proc) string {
	switch proc.Stat.State {
	case "R":
		return "running"
	case "S":
		return "sleeping"
	case "D":
		return "uninteruptible"
	case "Z":
		return "zombie"
	case "T":
		return "stopped"
	default:
		return "unknown"
	}
}

func getTtyMajorMinor() ([]ttyDriver, error) {
	var cnt int
	var err error
	var td *ttyDriver
	var result []ttyDriver
	var minors string

	file, err := os.Open("/proc/tty/drivers")
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		td = &ttyDriver{}

		tmpstr := scanner.Text()

		cnt, err = fmt.Sscanf(tmpstr, "%s\t%s\t%d\t%s\t%s",
			&td.name, &td.default_node, &td.major, &minors, &td.tty_type)
		if err != nil {
			return result, err
		}
		if cnt != 5 {
			return result, errors.New(fmt.Sprintf(
				"getTtyMajorMinor(): bad /dev/tty/drivers line: %s",
				tmpstr))
		}
		cnt, err = fmt.Sscanf(minors, "%d-%d", &td.minor_min, &td.minor_max)
		if cnt == 1 {
			td.minor_max = td.minor_min
		} else if cnt != 2 {
			return result, errors.New(fmt.Sprintf(
				"getTtyMajorMinor(): bad minor range line: %s", minors))
		}

		result = append(result, *td)
	}

	return result, nil
}

// based on https://gitorious.org/procps/procps/source/3a66fba1e934cbd830df572d8d03c05b4f4a5f1e:proc/devname.c#L286-321
func getTty(proc procreader.Proc) (string, error) {
	var d uint32
	var major uint32
	var minor uint32
	var name string

	// no_tty if == 0
	if proc.Stat.Tty_nr == 0 {
		return "?", nil
	}

	d = uint32(proc.Stat.Tty_nr)
	major = (d >> 8) & uint32(0xfff)
	minor = (d & uint32(0xff)) | (d&uint32(0xfff00000))>>12

	tty_drivers, err := getTtyMajorMinor()
	if err != nil {
		return "", err
	}

	for _, drv := range tty_drivers {
		if drv.major == major && drv.minor_min <= minor && drv.minor_max >= minor {
			if drv.minor_min == drv.minor_max {
				// these drivers are special since there's only one device
				name = drv.default_node
			} else if drv.default_node == "/dev/pts" || drv.default_node == "/dev/ptm" {
				name = fmt.Sprintf("%s/%d", drv.default_node, minor-drv.minor_min)
			} else {
				name = fmt.Sprintf("%s%d", drv.default_node, minor-drv.minor_min)
			}

			return strings.TrimPrefix(name, "/dev/"), nil
		}
	}

	// XXX unlike procps we only try based on major/minor. Eventually probably we
	// should compile devname.c and w/ cgo

	return "?", nil
}

// https://gitorious.org/procps/procps/source/3a66fba1e934cbd830df572d8d03c05b4f4a5f1e:ps/output.c#L775-794
func getStateFlags(proc procreader.Proc) []string {
	var result []string

	if proc.Stat.Nice > 0 {
		// N flag in 'ps -o stat'
		result = append(result, "low")
	}

	if proc.Stat.Nice < 0 {
		// < flag in 'ps -o stat'
		result = append(result, "high")
	}

	if proc.Status.VmLck > 0 {
		// L flag in 'ps -o stat'
		result = append(result, "lock")
	}

	if proc.Status.VmLck > 0 {
		// L flag in 'ps -o stat'
		result = append(result, "lock")
	}

	if int64(proc.Status.Tgid) == proc.Stat.Sid {
		// s flag in 'ps -o stat'
		result = append(result, "leader")
	}

	if proc.Status.Threads > 1 {
		// l flag in 'ps -o stat'
		result = append(result, "threads")
	}

	if proc.Stat.Pgrp == proc.Stat.Tty_pgrp {
		// + flag in 'ps -o stat'
		result = append(result, "foreground")
	}

	return result
}

func formatCommand(proc procreader.Proc) string {
	if len(proc.Cmdline) > 0 {
		return strings.Join(proc.Cmdline, " ")
	}
	return fmt.Sprintf("[%s]", proc.Stat.Tcomm)
}

func toDocker(proc procreader.Proc) (DockerTop, error) {
	var dt DockerTop
	var err error

	dt.C, err = calculateC(proc)
	if err != nil {
		return dt, err
	}
	dt.Comm = proc.Stat.Tcomm
	dt.Command = formatCommand(proc)
	dt.Cputime = float64(proc.Stat.Utime+proc.Stat.Stime) / float64(getHertz())
	dt.Gid = proc.Status.Gid.Effective
	dt.Lwp = proc.Stat.Pid // procps uses lwp=pid, so good enough for me
	dt.Nice = proc.Stat.Nice
	dt.Pcpu, err = calculatePcpu(proc)
	if err != nil {
		return dt, err
	}
	dt.Pid = proc.Stat.Pid
	dt.Pgid = proc.Stat.Pgrp
	dt.Pmem, err = calculatePmem(proc)
	if err != nil {
		return dt, err
	}
	dt.Ppid = proc.Stat.Ppid
	dt.Psr = proc.Stat.Task_cpu
	dt.Rgid = proc.Status.Gid.Real
	dt.Rss = proc.Status.VmRSS
	dt.Ruid = proc.Status.Uid.Real
	dt.Start_time = getStartTime(proc)
	dt.State = getState(proc)
	dt.State_flags = getStateFlags(proc)
	dt.Tty, err = getTty(proc)
	if err != nil {
		return dt, err
	}
	dt.Uid = proc.Status.Uid.Effective
	dt.Vsz = proc.Status.VmSize

	return dt, err
}

func main() {
	var dt DockerTop
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
		dt, err = toDocker(proc)
		if err != nil {
			panic(err)
		}
		out, err := json.Marshal(dt)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))
	}
}
