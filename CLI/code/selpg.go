package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	inFileName string
	printDest  string
	pageLen    int
	pageType   bool
}

var progname string /* program name, for error messages */

func main() {
	var args selpgArgs
	getArgs(&args)
	checkArgs(&args)
	excute(&args)
}

func execError(err error, content string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n[Error]%s:", content)
		os.Exit(10)
	}
}

func getArgs(args *selpgArgs) {
	flag.IntVarP(&(args.startPage), "startPage", "s", -1, "start page")
	flag.IntVarP(&(args.endPage), "endPage", "e", -1, "end page")
	flag.IntVarP(&(args.pageLen), "pageLength", "l", 1, "page length")
	flag.StringVarP(&(args.printDest), "printDest", "d", "", "print dest")
	flag.BoolVarP(&(args.pageType), "pageType", "f", false, "page type")
	flag.Parse()

	argLeft := flag.Args()
	if len(argLeft) > 0 {
		args.inFileName = string(argLeft[0])
	} else {
		args.inFileName = ""
	}
}

func checkArgs(sa *selpgArgs) {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "\n%s: not enough arguments\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1] != "-s" {
		fmt.Fprintf(os.Stderr, "\n%s: 1st arg should be -s startPage\n", progname)
		flag.Usage()
		os.Exit(2)

	}

	INT_MAX := 1<<32 - 1

	if sa.startPage < 1 || sa.startPage > INT_MAX {
		fmt.Fprintf(os.Stderr, "\n%s: invalid start page %s\n", progname, os.Args[2])
		flag.Usage()
		os.Exit(3)

	}

	if os.Args[3] != "-e" {
		fmt.Fprintf(os.Stderr, "\n%s: 2nd arg should be -e end_page\n", progname)
		flag.Usage()
		os.Exit(4)
	}

	if sa.endPage < sa.startPage || sa.endPage > INT_MAX {
		fmt.Fprintf(os.Stderr, "\n%s: invalid end page %d\n", progname, sa.endPage)
		flag.Usage()
		os.Exit(5)
	}

	if sa.pageLen < 1 || sa.pageLen > (INT_MAX-1) {
		fmt.Fprintf(os.Stderr, "\n%s: invalid page length %d\n", progname, sa.pageLen)
		flag.Usage()
		os.Exit(6)
	}

	if len(flag.Args()) == 1 {
		_, err := os.Stat(flag.Args()[0])
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "\n%s: input file \"%s\" does not exist\n", progname, flag.Args()[0])
			os.Exit(7)
		}
		sa.inFileName = flag.Args()[0]
	}

	fmt.Printf("\n[ArgsStart]\n")
}

func checkInputFile(filename string) {
	_, errFileExits := os.Stat(filename)
	if os.IsNotExist(errFileExits) {
		fmt.Fprintf(os.Stderr, "\n[Error]: the input file \"%s\" does not exist\n", filename)
		os.Exit(8)
	}
}

func excute(sa *selpgArgs) {
	var fin *os.File
	if sa.inFileName == "" {
		fin = os.Stdin
	} else {
		checkInputFile(sa.inFileName)
		var err error
		fin, err = os.Open(sa.inFileName)
		execError(err, "encounter an error when open input file")
	}

	if len(sa.printDest) == 0 {
		printToDes(os.Stdout, fin, sa.startPage, sa.endPage, sa.pageLen, sa.pageType)
	} else {
		printToDes(createPipe(sa.printDest), fin, sa.startPage, sa.endPage, sa.pageLen, sa.pageType)
	}
}

func createPipe(printDest string) io.WriteCloser {
	cmd := exec.Command("lp", "-d"+printDest)
	fout, err := cmd.StdinPipe()
	execError(err, "StdinPipe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errStart := cmd.Run()
	execError(errStart, "run cmd")
	return fout
}

func printToDes(fout interface{}, fin *os.File, pageStart int, pageEnd int, pageLen int, pageType bool) {
	lineCount := 0
	finalPage := 1
	buf := bufio.NewReader(fin)
	for true {
		var line string
		var err error
		if pageType {
			line, err = buf.ReadString('\f')
			finalPage++
		} else {
			line, err = buf.ReadString('\n')
			lineCount++
			if lineCount > pageLen {
				finalPage++
				lineCount = 1
			}
		}

		if err == io.EOF {
			break
		}

		execError(err, "encounter an error when read in file")

		if (finalPage >= pageStart) && (finalPage <= pageEnd) {
			var outputErr error
			if stdOutput, ok := fout.(*os.File); ok {
				_, outputErr = fmt.Fprintf(stdOutput, "%s", line)
			} else if pipeOutput, ok := fout.(io.WriteCloser); ok {
				_, outputErr = pipeOutput.Write([]byte(line))
			} else {
				fmt.Fprintf(os.Stderr, "\n[Error]:fout type error. ")
				os.Exit(8)
			}
			execError(outputErr, "Error happend when output the pages.")
		}
	}

	if finalPage < pageStart {
		fmt.Fprintf(os.Stderr, "\n[Error]: startPage (%d) greater than total pages (%d), no output written\n", pageStart, finalPage)
		os.Exit(11)
	} else if finalPage < pageEnd {
		fmt.Fprintf(os.Stderr, "\n[Error]: endPage (%d) greater than total pages (%d), less output than expected\n", pageEnd, finalPage)
		os.Exit(12)
	}
}
