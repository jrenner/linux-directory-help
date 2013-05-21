package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"dirinfo"
	"strings"
)

var (
	CURRENT_DIR = ""
	INFO_SOURCES = "Information sources:\n" +
		"\thttp://en.wikipedia.org/wiki/Filesystem_Hierarchy_Standard\n" +
		"\tcontents of 'man hier'"
	VERSION       = "1.1"
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
	if (err != nil) {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage of dirhelp:")
	fmt.Println("\t'dirhelp -a' print info for every directory")
	fmt.Println("\t'dirhelp -v' show version, URL and author info")
	fmt.Println("\t'dirhelp -h' show this help")
	fmt.Println("\t'dirhelp *' ")
}

func formatLookupDir(lookupDir string) {
	if lookupDir == "/" {
		return
	}
	// need to have the path begin with a slash for the regex
	if lookupDir[0] != '/' {
		lookupDir = "/" + lookupDir
	}
	// remove a last slash if there is one
	if lookupDir[len(lookupDir)-1] == '/' {
		lookupDir = lookupDir[:len(lookupDir)-1]
	}
	for strings.Contains(lookupDir, "//") {
		fmt.Println(lookupDir)
		lookupDir = strings.Replace(lookupDir, "//", "/", -1);
		fmt.Println("post: " + lookupDir)
	}
}

func printDirInfo(lookupDirList []string) {
	var foundInfoCount uint8 = 0
	var foundInfo bool = false;
	results := testRE.FindAllStringSubmatch(string(dirinfo.FHS_INFO), -1)
	infoToPrint := ""
	for _, lookupDir := range lookupDirList {
		foundInfo = false
		formatLookupDir(lookupDir)
		for _, regexResult := range results {
			dir := regexResult[1]
			help := regexResult[2]
			if  dir == lookupDir || *flagAllHelp {
				infoToPrint += fmt.Sprintf("[%v] %v\n", dir, help)
				foundInfo = true
				foundInfoCount++
			}
		}
		if (!foundInfo) {
			fmt.Printf("Did not find information for directory: '%s'\n", lookupDir)
		}
	}
	if foundInfoCount == 0 {
		printUsage()
	} else {
		fmt.Print(infoToPrint)
	}
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
		lookupDirList[i] = CURRENT_DIR + "/" + dir

	}
	if len(lookupDirList) == 0 {
		lookupDirList = append(lookupDirList, CURRENT_DIR)
	}
	printDirInfo(lookupDirList)
}
