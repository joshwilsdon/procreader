package procreader

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	statContent    string
	statmContent   string
	statusContent  string
	cmdlineContent string
	environContent string
	expected       Proc
}

// NOTE: you can generate test cases using examples/proc_read_struct.go
var testCases = map[uint64]testCase{
	15220: {
		statContent:    "15220 (bash) S 15160 15220 15220 34817 29367 4219136 161706 6374605 11 796 28 33 25909 3879 20 0 1 0 131158 21934080 985 18446744073709551615 4194304 5173212 140736926389104 140736926387816 140716594644428 0 65536 3670020 1266777851 18446744071579277074 0 0 17 0 0 0 7 0 0 7273968 7310504 32763904 140736926396005 140736926396011 140736926396011 140736926396398 0\n",
		statmContent:   "5355 985 450 239 0 533 0\n",
		statusContent:  "Name:\tbash\nState:\tS (sleeping)\nTgid:\t15220\nNgid:\t0\nPid:\t15220\nPPid:\t15160\nTracerPid:\t0\nUid:\t0\t0\t0\t0\nGid:\t0\t0\t0\t0\nFDSize:\t256\nGroups:\t0 \nVmPeak:\t   21480 kB\nVmSize:\t   21420 kB\nVmLck:\t       0 kB\nVmPin:\t       0 kB\nVmHWM:\t    3964 kB\nVmRSS:\t    3940 kB\nVmData:\t    1996 kB\nVmStk:\t     136 kB\nVmExe:\t     956 kB\nVmLib:\t    2288 kB\nVmPTE:\t      60 kB\nVmSwap:\t       0 kB\nThreads:\t1\nSigQ:\t0/3838\nSigPnd:\t0000000000000000\nShdPnd:\t0000000000000000\nSigBlk:\t0000000000010000\nSigIgn:\t0000000000380004\nSigCgt:\t000000004b817efb\nCapInh:\t0000000000000000\nCapPrm:\t0000001fffffffff\nCapEff:\t0000001fffffffff\nCapBnd:\t0000001fffffffff\nSeccomp:\t0\nCpus_allowed:\t3\nCpus_allowed_list:\t0-1\nMems_allowed:\t00000000,00000001\nMems_allowed_list:\t0\nvoluntary_ctxt_switches:\t8367\nnonvoluntary_ctxt_switches:\t7268\n",
		cmdlineContent: "-bash\x00",
		environContent: "LANG=en_US.UTF-8\x00USER=root\x00LOGNAME=root\x00HOME=/root\x00PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games\x00MAIL=/var/mail/root\x00SHELL=/bin/bash\x00SSH_CLIENT=10.88.0.1 52420 22\x00SSH_CONNECTION=10.88.0.1 52420 10.88.0.151 22\x00SSH_TTY=/dev/pts/1\x00TERM=xterm-256color\x00XDG_SESSION_ID=7\x00XDG_RUNTIME_DIR=/run/user/0\x00SSH_AUTH_SOCK=/tmp/ssh-W5s8CB8xWd/agent.15160\x00",
		expected: Proc{
			Stat_t{Pid: 0x3b74, Tcomm: "bash", State: "S", Ppid: 15160, Pgrp: 15220, Sid: 15220, Tty_nr: 34817, Tty_pgrp: 29367, Flags: 0x406100, Min_flt: 0x277aa, Cmin_flt: 0x6144cd, Maj_flt: 0xb, Cmaj_flt: 0x31c, Utime: 0x1c, Stime: 0x21, Cutime: 0x6535, Cstime: 0xf27, Priority: 20, Nice: 0, Num_threads: 0x1, it_real_value: 0x0, Start_time: 0x20056, Vsize: 0x14eb000, Rss: 0x3d9, Rsslim: 0xffffffffffffffff, Start_code: 0x400000, End_code: 0x4eefdc, Start_stack: 0x7fffde811370, Esp: 0x7fffde810e68, Eip: 0x7ffb22a345cc, Pending: "0", Blocked: "65536", Sigign: "3670020", Sigcatch: "1266777851", Wchan: 0xffffffff81069712, placeholder1: 0x0, placeholder2: 0x0, Exit_signal: 0x11, Task_cpu: 0x0, Rt_priority: 0x0, Policy: 0x0, Blkio_ticks: 0x7, Gtime: 0x0, Cgtime: 0x0, Start_data: 0x6efdf0, End_data: 0x6f8ca8, Start_brk: 0x1f3f000, Arg_start: 0x7fffde812e65, Arg_end: 0x7fffde812e6b, Env_start: 0x7fffde812e6b, Env_end: 0x7fffde812fee, Exit_code: 0x0},
			Statm_t{Size: 0x14eb, Resident: 0x3d9, Shared: 0x1c2, Trs: 0xef, Lrs: 0x0, Drs: 0x215, Dt: 0x0},
			Status_t{Name: "bash", State: "S (sleeping)", Tgid: 0x3b74, Ngid: 0x0, Pid: 0x3b74, PPid: 0x3b38, TracerPid: 0x0, Uid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Gid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Effective_Gid: 0x0, Saved_Gid: 0x0, FS_Gid: 0x0, FDSize: 0x100, Groups: []uint64{0x0}, VmPeak: 0x53e8, VmSize: 0x53ac, VmLck: 0x0, VmPin: 0x0, VmHWM: 0xf7c, VmRSS: 0xf64, VmData: 0x7cc, VmStk: 0x88, VmExe: 0x3bc, VmLib: 0x8f0, VmPTE: 0x3c, VmSwap: 0x0, Threads: 0x1, SigQ: SigQVal{Num: 0x0, Max: 0xefe}, SigPnd: "0000000000000000", ShdPnd: "0000000000000000", SigBlk: "0000000000010000", SigIgn: "0000000000380004", SigCgt: "000000004b817efb", CapInh: "0000000000000000", CapPrm: "0000001fffffffff", CapEff: "0000001fffffffff", CapBnd: "0000001fffffffff", Seccomp: 0x0, Cpus_allowed: "3", Cpus_allowed_list: "0-1", Mems_allowed: "00000000,00000001", Mems_allowed_list: "0", Voluntary_ctxt_switches: 0x20af, Nonvoluntary_ctxt_switches: 0x1c64},
			[]string{"-bash"},
			[]string{"LANG=en_US.UTF-8", "USER=root", "LOGNAME=root", "HOME=/root", "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games", "MAIL=/var/mail/root", "SHELL=/bin/bash", "SSH_CLIENT=10.88.0.1 52420 22", "SSH_CONNECTION=10.88.0.1 52420 10.88.0.151 22", "SSH_TTY=/dev/pts/1", "TERM=xterm-256color", "XDG_SESSION_ID=7", "XDG_RUNTIME_DIR=/run/user/0", "SSH_AUTH_SOCK=/tmp/ssh-W5s8CB8xWd/agent.15160"},
		},
	},
	29821: {
		statContent:    "29821 (:-) 0 1 2 3 4 5) R 15220 29821 15220 34817 29852 4218880 823 0 1 0 3980 3 0 0 20 0 1 0 5829898 11390976 293 18446744073709551615 4194304 5173212 140734601257184 140734601255848 4541996 0 0 4 65536 0 0 0 17 0 0 0 13 0 0 7273968 7310504 21405696 140734601263196 140734601263256 140734601263256 140734601265094 0\n",
		statmContent:   "2781 293 244 239 0 65 0\n",
		statusContent:  "Name:\t:-) 0 1 2 3 4 5\nState:\tR (running)\nTgid:\t29821\nNgid:\t0\nPid:\t29821\nPPid:\t15220\nTracerPid:\t0\nUid:\t0\t0\t0\t0\nGid:\t0\t0\t0\t0\nFDSize:\t256\nGroups:\t0 \nVmPeak:\t   11124 kB\nVmSize:\t   11124 kB\nVmLck:\t       0 kB\nVmPin:\t       0 kB\nVmHWM:\t    1172 kB\nVmRSS:\t    1172 kB\nVmData:\t     124 kB\nVmStk:\t     136 kB\nVmExe:\t     956 kB\nVmLib:\t    2072 kB\nVmPTE:\t      40 kB\nVmSwap:\t       0 kB\nThreads:\t1\nSigQ:\t0/3838\nSigPnd:\t0000000000000000\nShdPnd:\t0000000000000000\nSigBlk:\t0000000000000000\nSigIgn:\t0000000000000004\nSigCgt:\t0000000000010000\nCapInh:\t0000000000000000\nCapPrm:\t0000001fffffffff\nCapEff:\t0000001fffffffff\nCapBnd:\t0000001fffffffff\nSeccomp:\t0\nCpus_allowed:\t3\nCpus_allowed_list:\t0-1\nMems_allowed:\t00000000,00000001\nMems_allowed_list:\t0\nvoluntary_ctxt_switches:\t2\nnonvoluntary_ctxt_switches:\t302\n",
		cmdlineContent: "/bin/bash\x00/root/gops/procreader/testdata/:-) 0 1 2 3 4 5 6 \x00",
		environContent: "XDG_SESSION_ID=7\x00SHELL=/bin/bash\x00TERM=xterm-256color\x00SSH_CLIENT=10.88.0.1 52420 22\x00SSH_TTY=/dev/pts/1\x00USER=root\x00LS_COLORS=rs=0:di=01;34:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.lz=01;31:*.xz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.cgm=01;35:*.emf=01;35:*.axv=01;35:*.anx=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:*.axa=00;36:*.oga=00;36:*.spx=00;36:*.xspf=00;36:\x00SSH_AUTH_SOCK=/tmp/ssh-W5s8CB8xWd/agent.15160\x00PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/usr/local/go/bin\x00MAIL=/var/mail/root\x00_=./execer\x00PWD=/root/gops/procreader/testdata\x00LANG=en_US.UTF-8\x00HOME=/root\x00SHLVL=1\x00LOGNAME=root\x00SSH_CONNECTION=10.88.0.1 52420 10.88.0.151 22\x00LESSOPEN=| /usr/bin/lesspipe %s\x00XDG_RUNTIME_DIR=/run/user/0\x00LESSCLOSE=/usr/bin/lesspipe %s %s\x00",
		expected: Proc{
			Stat_t{Pid: 0x747d, Tcomm: ":-) 0 1 2 3 4 5", State: "R", Ppid: 15220, Pgrp: 29821, Sid: 15220, Tty_nr: 34817, Tty_pgrp: 29852, Flags: 0x406000, Min_flt: 0x337, Cmin_flt: 0x0, Maj_flt: 0x1, Cmaj_flt: 0x0, Utime: 0xf8c, Stime: 0x3, Cutime: 0x0, Cstime: 0x0, Priority: 20, Nice: 0, Num_threads: 0x1, it_real_value: 0x0, Start_time: 0x58f50a, Vsize: 0xadd000, Rss: 0x125, Rsslim: 0xffffffffffffffff, Start_code: 0x400000, End_code: 0x4eefdc, Start_stack: 0x7fff53ea60e0, Esp: 0x7fff53ea5ba8, Eip: 0x454e2c, Pending: "0", Blocked: "0", Sigign: "4", Sigcatch: "65536", Wchan: 0x0, placeholder1: 0x0, placeholder2: 0x0, Exit_signal: 0x11, Task_cpu: 0x0, Rt_priority: 0x0, Policy: 0x0, Blkio_ticks: 0xd, Gtime: 0x0, Cgtime: 0x0, Start_data: 0x6efdf0, End_data: 0x6f8ca8, Start_brk: 0x146a000, Arg_start: 0x7fff53ea785c, Arg_end: 0x7fff53ea7898, Env_start: 0x7fff53ea7898, Env_end: 0x7fff53ea7fc6, Exit_code: 0x0},
			Statm_t{Size: 0xadd, Resident: 0x125, Shared: 0xf4, Trs: 0xef, Lrs: 0x0, Drs: 0x41, Dt: 0x0},
			Status_t{Name: ":-) 0 1 2 3 4 5", State: "R (running)", Tgid: 0x747d, Ngid: 0x0, Pid: 0x747d, PPid: 0x3b74, TracerPid: 0x0, Uid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Gid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Effective_Gid: 0x0, Saved_Gid: 0x0, FS_Gid: 0x0, FDSize: 0x100, Groups: []uint64{0x0}, VmPeak: 0x2b74, VmSize: 0x2b74, VmLck: 0x0, VmPin: 0x0, VmHWM: 0x494, VmRSS: 0x494, VmData: 0x7c, VmStk: 0x88, VmExe: 0x3bc, VmLib: 0x818, VmPTE: 0x28, VmSwap: 0x0, Threads: 0x1, SigQ: SigQVal{Num: 0x0, Max: 0xefe}, SigPnd: "0000000000000000", ShdPnd: "0000000000000000", SigBlk: "0000000000000000", SigIgn: "0000000000000004", SigCgt: "0000000000010000", CapInh: "0000000000000000", CapPrm: "0000001fffffffff", CapEff: "0000001fffffffff", CapBnd: "0000001fffffffff", Seccomp: 0x0, Cpus_allowed: "3", Cpus_allowed_list: "0-1", Mems_allowed: "00000000,00000001", Mems_allowed_list: "0", Voluntary_ctxt_switches: 0x2, Nonvoluntary_ctxt_switches: 0x12e},
			[]string{"/bin/bash", "/root/gops/procreader/testdata/:-) 0 1 2 3 4 5 6 "},
			[]string{"XDG_SESSION_ID=7", "SHELL=/bin/bash", "TERM=xterm-256color", "SSH_CLIENT=10.88.0.1 52420 22", "SSH_TTY=/dev/pts/1", "USER=root", "LS_COLORS=rs=0:di=01;34:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.lz=01;31:*.xz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.cgm=01;35:*.emf=01;35:*.axv=01;35:*.anx=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:*.axa=00;36:*.oga=00;36:*.spx=00;36:*.xspf=00;36:", "SSH_AUTH_SOCK=/tmp/ssh-W5s8CB8xWd/agent.15160", "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/usr/local/go/bin", "MAIL=/var/mail/root", "_=./execer", "PWD=/root/gops/procreader/testdata", "LANG=en_US.UTF-8", "HOME=/root", "SHLVL=1", "LOGNAME=root", "SSH_CONNECTION=10.88.0.1 52420 10.88.0.151 22", "LESSOPEN=| /usr/bin/lesspipe %s", "XDG_RUNTIME_DIR=/run/user/0", "LESSCLOSE=/usr/bin/lesspipe %s %s"},
		},
	},
	// This one came from 2.6.18 and has a different number of fields
	29167: {
		statContent:    "29167 (sshd) S 1 29167 29167 0 -1 4202816 34440643 2073340695 0 512 495 2615 147515 115358 15 0 1 0 53885311 50077696 300 18446744073709551615 93824992231424 93824992662604 140734328009440 18446744073709551615 47340894086243 0 0 4096 81925 0 0 0 17 0 0 0 0\n",
		statmContent:   "12226 300 171 106 0 138 0\n",
		statusContent:  "Name:\tsshd\nState:\tS (sleeping)\nSleepAVG:\t98%\nTgid:\t29167\nPid:\t29167\nPPid:\t1\nTracerPid:\t0\nUid:\t0\t0\t0\t0\nGid:\t0\t0\t0\t0\nFDSize:\t64\nGroups:\t\nVmPeak:\t   48908 kB\nVmSize:\t   48904 kB\nVmLck:\t       0 kB\nVmHWM:\t    1200 kB\nVmRSS:\t    1200 kB\nVmData:\t     468 kB\nVmStk:\t      84 kB\nVmExe:\t     424 kB\nVmLib:\t    4652 kB\nVmPTE:\t     112 kB\nThreads:\t1\nSigQ:\t0/2112\nSigPnd:\t0000000000000000\nShdPnd:\t0000000000000000\nSigBlk:\t0000000000000000\nSigIgn:\t0000000000001000\nSigCgt:\t0000000180014005\nCapInh:\t0000000000000000\nCapPrm:\t00000000fffffeff\nCapEff:\t00000000fffffeff\nCpus_allowed:\tffffffff\nMems_allowed:\t1\n",
		cmdlineContent: "/usr/sbin/sshd\x00",
		environContent: "SUDO_GID=1000\x00USER=root\x00MAIL=/var/mail/josh\x00HOME=/home/josh\x00SUDO_UID=1000\x00LOGNAME=root\x00USERNAME=root\x00TERM=xterm-color\x00PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/X11R6/bin:/usr/sbin:/sbin\x00SSHD_OOM_ADJUST=-17\x00LS_COLORS=no=00:fi=00:di=01;34:ln=01;36:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.svgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.lzma=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.rar=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:\x00SUDO_COMMAND=/etc/init.d/ssh restart\x00SHELL=/bin/bash\x00SUDO_USER=josh\x00PWD=/home/josh\x00",
		expected: Proc{
			Stat_t{Pid: 0x71ef, Tcomm: "sshd", State: "S", Ppid: 1, Pgrp: 29167, Sid: 29167, Tty_nr: 0, Tty_pgrp: -1, Flags: 0x402140, Min_flt: 0x20d85c3, Cmin_flt: 0x7b94ab17, Maj_flt: 0x0, Cmaj_flt: 0x200, Utime: 0x1ef, Stime: 0xa37, Cutime: 0x2403b, Cstime: 0x1c29e, Priority: 15, Nice: 0, Num_threads: 0x1, it_real_value: 0x0, Start_time: 0x336397f, Vsize: 0x2fc2000, Rss: 0x12c, Rsslim: 0xffffffffffffffff, Start_code: 0x555555554000, End_code: 0x5555555bd44c, Start_stack: 0x7fff43a0f2e0, Esp: 0xffffffffffffffff, Eip: 0x2b0e692ce463, Pending: "0", Blocked: "0", Sigign: "4096", Sigcatch: "81925", Wchan: 0x0, placeholder1: 0x0, placeholder2: 0x0, Exit_signal: 0x11, Task_cpu: 0x0, Rt_priority: 0x0, Policy: 0x0, Blkio_ticks: 0x0, Gtime: 0x0, Cgtime: 0x0, Start_data: 0x0, End_data: 0x0, Start_brk: 0x0, Arg_start: 0x0, Arg_end: 0x0, Env_start: 0x0, Env_end: 0x0, Exit_code: 0x0},
			Statm_t{Size: 0x2fc2, Resident: 0x12c, Shared: 0xab, Trs: 0x6a, Lrs: 0x0, Drs: 0x8a, Dt: 0x0},
			Status_t{Name: "sshd", State: "S (sleeping)", Tgid: 0x71ef, Ngid: 0x0, Pid: 0x71ef, PPid: 0x1, TracerPid: 0x0, Uid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Gid: Ids{Real: 0x0, Effective: 0x0, Saved: 0x0, FS: 0x0}, Effective_Gid: 0x0, Saved_Gid: 0x0, FS_Gid: 0x0, FDSize: 0x40, Groups: []uint64(nil), VmPeak: 0xbf0c, VmSize: 0xbf08, VmLck: 0x0, VmPin: 0x0, VmHWM: 0x4b0, VmRSS: 0x4b0, VmData: 0x1d4, VmStk: 0x54, VmExe: 0x1a8, VmLib: 0x122c, VmPTE: 0x70, VmSwap: 0x0, Threads: 0x1, SigQ: SigQVal{Num: 0x0, Max: 0x840}, SigPnd: "0000000000000000", ShdPnd: "0000000000000000", SigBlk: "0000000000000000", SigIgn: "0000000000001000", SigCgt: "0000000180014005", CapInh: "0000000000000000", CapPrm: "00000000fffffeff", CapEff: "00000000fffffeff", CapBnd: "", Seccomp: 0x0, Cpus_allowed: "ffffffff", Cpus_allowed_list: "", Mems_allowed: "1", Mems_allowed_list: "", Voluntary_ctxt_switches: 0x0, Nonvoluntary_ctxt_switches: 0x0},
			[]string{"/usr/sbin/sshd"},
			[]string{"SUDO_GID=1000", "USER=root", "MAIL=/var/mail/josh", "HOME=/home/josh", "SUDO_UID=1000", "LOGNAME=root", "USERNAME=root", "TERM=xterm-color", "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/X11R6/bin:/usr/sbin:/sbin", "SSHD_OOM_ADJUST=-17", "LS_COLORS=no=00:fi=00:di=01;34:ln=01;36:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:su=37;41:sg=30;43:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.svgz=01;31:*.arj=01;31:*.taz=01;31:*.lzh=01;31:*.lzma=01;31:*.zip=01;31:*.z=01;31:*.Z=01;31:*.dz=01;31:*.gz=01;31:*.bz2=01;31:*.bz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.rar=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.jpg=01;35:*.jpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.aac=00;36:*.au=00;36:*.flac=00;36:*.mid=00;36:*.midi=00;36:*.mka=00;36:*.mp3=00;36:*.mpc=00;36:*.ogg=00;36:*.ra=00;36:*.wav=00;36:", "SUDO_COMMAND=/etc/init.d/ssh restart", "SHELL=/bin/bash", "SUDO_USER=josh", "PWD=/home/josh"},
		},
	},
}

func TestReadProc(t *testing.T) {
	var cfg procConfig
	var pid uint64
	var tc testCase

	cfg.basepath = "/proc"

	for pid, tc = range testCases {
		contents := map[string]string{
			"stat":    tc.statContent,
			"statm":   tc.statmContent,
			"status":  tc.statusContent,
			"cmdline": tc.cmdlineContent,
			"environ": tc.environContent,
		}
		cfg.contents = contents

		actual, err := readProc(&cfg, pid)
		if err != nil {
			t.Errorf("readProc: %s\n%s\n", err.Error(), err.(*ProcErr).Stack)
		}

		if !reflect.DeepEqual(actual.Stat, testCases[pid].expected.Stat) {
			t.Errorf("<%d> stat: actual != expected\n", pid)
		} else {
			fmt.Printf("ok <%d> stat matches\n", pid)
		}
		if !reflect.DeepEqual(actual.Statm, testCases[pid].expected.Statm) {
			t.Errorf("<%d> statm: actual != expected\n", pid)
		} else {
			fmt.Printf("ok <%d> statm matches\n", pid)
		}
		if !reflect.DeepEqual(actual.Status, testCases[pid].expected.Status) {
			t.Errorf("<%d> status: actual != expected\n", pid)
		} else {
			fmt.Printf("ok <%d> status matches\n", pid)
		}
		if !reflect.DeepEqual(actual.Cmdline, testCases[pid].expected.Cmdline) {
			t.Errorf("<%d> actual != expected\n", pid)
		} else {
			fmt.Printf("ok <%d> cmdline matches\n", pid)
		}
		if !reflect.DeepEqual(actual.Environ, testCases[pid].expected.Environ) {
			t.Errorf("<%d> environ: actual != expected\n", pid)
		} else {
			fmt.Printf("ok <%d> environ matches\n", pid)
		}
	}
}

func TestBadProcDir(t *testing.T) {
	var cfg procConfig
	var pid uint64

	pid = 1

	cfg.basepath = "/nonexistent/path"
	_, err := readProc(&cfg, pid)
	if err == nil {
		t.Errorf("readProc: %s\n%s\n", err.Error(), err.(*ProcErr).Stack)
	} else {
		fmt.Printf("ok invalid path == %s\n", err.Error())
	}
}

func TestIncompleteStat(t *testing.T) {
	var cfg procConfig
	var pid uint64

	pid = 1

	cfg.basepath = "/proc"

	contents := map[string]string{
		"stat":    "",
		"statm":   "",
		"status":  "",
		"cmdline": "",
		"environ": "",
	}
	cfg.contents = contents

	_, err := readProc(&cfg, pid)
	if err == nil {
		t.Errorf("readProc: %s\n%s\n", err.Error(), err.(*ProcErr).Stack)
	} else {
		fmt.Printf("ok fails: %s\n", err.Error())
	}
}
