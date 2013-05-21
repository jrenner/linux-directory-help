package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"dirinfo"
	"strings"
	"runtime"
)

var (
	USER_HOME_DIR string
	CURRENT_DIR string
	INFO_SOURCES = "Information sources:\n" +
		"    http://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard\n" +
		"    contents of 'man hier'"
	VERSION       = "1.2"
	flagAllHelp   = flag.Bool("a", false, "print info for all directories")
	flagVersion   = flag.Bool("v", false, "show version number")
	flagPrintHelp = flag.Bool("h", false, "print usage info")
	// grabs a help string from that mess at the bottom of this file
	testRE, _ = regexp.Compile("(/.*)\n(.*)")
)

func init() {
	flag.Parse()
	var err error
	CURRENT_DIR, err = os.Getwd()
	handleFatalError(err)

	// because of cross-compiling disabling cgo, this feature is not available
	// a solution is to compile on the respective platforms
	// a bit of a pain for such a small feature, though
	if (runtime.GOARCH == "amd64") {
		currentUser, err := user.Current()
		handleFatalError(err)
		USER_HOME_DIR = currentUser.HomeDir
	} else {
		// I hope there isn't a user called DISABLED_DUE_TO_CROSSCOMPILE_WOES!
		USER_HOME_DIR = "DISABLED_DUE_TO_CROSSCOMPILE_WOES"
	}
}

func handleFatalError(err error) {
	if (err != nil) {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage of dirhelp:")
	fmt.Println("    dirhelp          get info on current working directory")
	fmt.Println("    dirhelp <dir>    get info on directory - can be multiple dirs or wildcards")
	fmt.Println("    dirhelp -a       print info for every directory")
	fmt.Println("    dirhelp -v       show version, URL and author info")
	fmt.Println("    dirhelp -h       show this help")
}

func formatLookupDir(lookupDir *string) {
	tempDir := *lookupDir
	if tempDir == "/" {
		return
	}
	// need to have the path begin with a slash for the regex
	if tempDir[0] != '/' {
		tempDir = "/" + tempDir
	}
	// remove a last slash if there is one
	if tempDir[len(tempDir)-1] == '/' {
		tempDir = tempDir[:len(tempDir)-1]
	}
	for strings.Contains(tempDir, "//") {
		tempDir = strings.Replace(tempDir, "//", "/", -1);
	}
	*lookupDir = tempDir
}

func isDir(pth string) (bool, error) {
	fi, err := os.Stat(pth)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsDir(), nil
}

func isADirectory(path string) bool {
	f, err := os.Open(path)
	if (err != nil) {
		fmt.Println("ERROR: " + err.Error())
		return false
	}
	stat, err := f.Stat()
	if (err != nil) {
		fmt.Println("ERROR: " + err.Error())
		return false;
	}
	return stat.IsDir()
}

func printDirInfo(lookupDirList []string) {
	var foundInfo bool = false;
	didNotFind := make([]string, 0)
	results := testRE.FindAllStringSubmatch(string(dirinfo.FHS_INFO), -1)
	infoToPrint := ""
	for _, lookupDir := range lookupDirList {
		isDir, err := isDir(lookupDir)
		handleFatalError(err)
		if !isDir {
			continue
		}
		foundInfo = false
		if lookupDir == USER_HOME_DIR || *flagAllHelp {
			infoToPrint += fmt.Sprintf("[%v] %v\n", USER_HOME_DIR, dirinfo.HOME_DIR_INFO)
			foundInfo = true;
		}
		formatLookupDir(&lookupDir)
		for _, regexResult := range results {
			dir := regexResult[1]
			help := regexResult[2]
			if  dir == lookupDir || *flagAllHelp {
				infoToPrint += fmt.Sprintf("[%v] %v\n", dir, help)
				foundInfo = true
			}
		}
		if (!foundInfo) {
			didNotFind = append(didNotFind, lookupDir)
		}
	}
	if (len(didNotFind) > 0) {
		fmt.Println("Did not find information for the following directories:")
		for _, dir := range didNotFind {
			fmt.Println(dir)
		}
		printUsage()
	}
	fmt.Print(infoToPrint)
}

func main() {
	if *flagPrintHelp {
		printUsage()
		os.Exit(0)
	}
	if *flagVersion {
		fmt.Printf("dirhelp %s\nhttps://github.com/jrenner/linux-directory-help\nby Jon Renner\nrennerjc@gmail.com\n", VERSION)
		fmt.Println(INFO_SOURCES)
		os.Exit(0)
	}
	lookupDirList := flag.Args()
	for i, dir := range lookupDirList {
		if (lookupDirList[i][0] != '/') {
			lookupDirList[i] = CURRENT_DIR + "/" + dir
		}
	}
	if len(lookupDirList) == 0 {
		lookupDirList = append(lookupDirList, CURRENT_DIR)
	}
	printDirInfo(lookupDirList)
}
