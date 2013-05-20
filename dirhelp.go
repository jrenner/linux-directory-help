package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

var INFO_SOURCE string = "Information source: http://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard"

var VERSION string = "1.0"
var flagAllHelp = flag.Bool("a", false, "print info for all directories")
var flagVersion = flag.Bool("v", false, "show version number")
var flagPrintHelp = flag.Bool("h", false, "print usage info")

// grabs a help string from that mess at the bottom of this file
var testRE, _ = regexp.Compile("(/.*)\n(.*)")

var CURRENT_PATH, _ = os.Getwd()

func init() {
	flag.Parse()
}

func printUsage() {
	fmt.Println("Execute dirhelp with no arguments to get help on the current directory.")
	flag.Usage()
}

func handleError(err error) {
	if err != nil {
		fatalError(err)
	}
}

func fatalError(err error) {
	fmt.Println("FATAL ERROR: ", err)
	os.Exit(1)
}

func main() {
	if *flagPrintHelp {
		printUsage()
		os.Exit(0)
	}
	if *flagVersion {
		fmt.Printf("dirhelp version: %s\nby Jon Renner\nrennerjc@gmail.com\n", VERSION)
		fmt.Println(INFO_SOURCE)
		os.Exit(0)
	}
	var argDir string = CURRENT_PATH
	args := flag.Args()
	if len(args) > 0 {
		argDir = args[0]
		if argDir[0] != '/' {
			argDir = "/" + argDir
		}
		if argDir[len(argDir)-1] == '/' {
			argDir = argDir[:len(argDir)-1]
		}
	}
	foundInfo := false
	results := testRE.FindAllStringSubmatch(string(info), -1)
	for _, v := range results {
		dir := v[1]
		help := v[2]
		if dir == argDir || *flagAllHelp {
			fmt.Printf("[%v] %v\n", dir, help)
			foundInfo = true
		}
	}
	if !foundInfo {
		fmt.Printf("No information for directory: '%s'\n", argDir)
		printUsage()
	}
}

// source: http://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard
var info string = `/
Primary hierarchy root and root directory of the entire file system hierarchy.
/bin
Essential command binaries that need to be available in single user mode; for all users, e.g., cat, ls, cp.
/boot
Boot loader files, e.g., kernels, initrd.
/dev
Essential devices, e.g., /dev/null.
/etc
Host-specific system-wide configuration files. There has been controversy over the meaning of the name itself. In early versions of the UNIX Implementation Document from Bell labs, /etc is referred to as the etcetera directory, as this directory historically held everything that did not belong elsewhere (however, the FHS restricts /etc to static configuration files and may not contain binaries). Since the publication of early documentation, the directory name has been re-designated in various ways. Recent interpretations include backronyms such as "Editable Text Configuration" or "Extended Tool Chest".
/etc/opt
Configuration files for /opt/.
/etc/sgml
Configuration files for SGML.
/etc/X11
Configuration files for the X Window System, version 11.
/etc/xml
Configuration files for XML.
/home
Users' home directories, containing saved files, personal settings, etc.
/lib
Libraries essential for the binaries in /bin/ and /sbin/.
/media
Mount points for removable media such as CD-ROMs (appeared in FHS-2.3).
/mnt
Temporarily mounted filesystems.
/opt
Optional application software packages.
/proc
Virtual filesystem providing information about processes and kernel information as files. In Linux, corresponds to a procfs mount.
/root
Home directory for the root user.
/sbin
Essential system binaries, e.g., init, ip, mount.
/srv
Site-specific data which are served by the system.
/tmp
Temporary files (see also /var/tmp). Often not preserved between system reboots.
/usr
Secondary hierarchy for read-only user data; contains the majority of (multi-)user utilities and applications.
/usr/bin
Non-essential command binaries (not needed in single user mode); for all users.
/usr/include
Standard include files.
/usr/lib
Libraries for the binaries in /usr/bin/ and /usr/sbin/.
/usr/local
Tertiary hierarchy for local data, specific to this host. Typically has further subdirectories, e.g., bin/, lib/, share/.
/usr/sbin
Non-essential system binaries, e.g., daemons for various network-services.
/usr/share
Architecture-independent (shared) data.
/usr/src
Source code, e.g., the kernel source code with its header files.
/usr/X11R6
X Window System, Version 11, Release 6.
/var
Variable files—files whose content is expected to continually change during normal operation of the system—such as logs, spool files, and temporary e-mail files.
/var/cache
Application cache data. Such data are locally generated as a result of time-consuming I/O or calculation. The application must be able to regenerate or restore the data. The cached files can be deleted without loss of data.
/var/lib
State information. Persistent data modified by programs as they run, e.g., databases, packaging system metadata, etc.
/var/lock
Lock files. Files keeping track of resources currently in use.
/var/log
Log files. Various logs.
/var/mail
Users' mailboxes.
/var/run
Information about the running system since last boot, e.g., currently logged-in users and running daemons.
/var/spool
Spool for tasks waiting to be processed, e.g., print queues and unread mail.
/var/spool/mail
Deprecated location for users' mailboxes.
/var/tmp
Temporary files to be preserved between reboots.`
