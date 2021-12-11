package flags

import "flag"

func File () string {
	fileName := flag.String("file", "input", "input file name")
	flag.Parse()

	return *fileName
}
