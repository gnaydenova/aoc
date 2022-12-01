package flags

import "flag"

var (
	fileName string
	part int
)

func File() string {
	initFlags()

	return fileName
}

func Part() int {
	initFlags()

	return part
}

func initFlags() {
	if flag.Parsed() {
		return 
	}

	flag.StringVar(&fileName, "file", "input", "input file name")
	flag.IntVar(&part, "part", 1, "part")
	flag.Parse()
}
