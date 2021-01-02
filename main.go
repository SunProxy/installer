/**
      ___           ___           ___
     /  /\         /__/\         /__/\
    /  /:/_        \  \:\        \  \:\
   /  /:/ /\        \  \:\        \  \:\
  /  /:/ /::\   ___  \  \:\   _____\__\:\
 /__/:/ /:/\:\ /__/\  \__\:\ /__/::::::::\
 \  \:\/:/~/:/ \  \:\ /  /:/ \  \:\~~\~~\/
  \  \::/ /:/   \  \:\  /:/   \  \:\  ~~~
   \__\/ /:/     \  \:\/:/     \  \:\
     /__/:/       \  \::/       \  \:\
     \__\/         \__\/         \__\/

MIT License

Copyright (c) 2020 Jviguy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package main

import (
	"bufio"
	"fmt"
	"github.com/jszwedko/go-circleci"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	build, err := RequestOption("Please Select a build", "latest", "Custom Build Number")
	if err != nil {
		log.Fatal(err)
	}
	dos, err := RequestOption("Please Select a OS For the Build", runtime.GOOS)
	if err != nil {
		log.Fatal(err)
	}
	processor, err := RequestOption("Please Type the Bit size of the Said Build", "64", "32")
	if err != nil {
		log.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting the working directory some fucking how.")
	}
	dir, err := RequestOption("Please Type the Dir to install Sun too", cwd)
	if err != nil {
		log.Fatal(err)
	}
	var num	int32
	n, err := strconv.Atoi(build)
	num = int32(n)
	if err != nil {
		if build != "latest" {
			log.Print("Invalid Build Must be latest or a integer!")
			_, _ = fmt.Scanln()
			os.Exit(1)
		}
		n, err := GetLatestBuildNum(circleci.Client{})
		if err != nil {
			log.Fatal("error fetching latest build!")
		}
		num = n
	}
	proc, err := strconv.Atoi(processor)
	if err != nil {
		//leave msg for person to see.
		log.Print("Invalid Parameter Bit Size Must be one of a int32!")
		_, _ = fmt.Scanln()
		os.Exit(1)
	}
	link := GenerateDownloadLink(BuildOptions{BuildNum: num, Processor: int16(proc), Os: dos})
	fmt.Println("Downloading Sun Proxy with the artifact link: ", link)
	err = DownloadBuild(link, dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Downloaded Sun Proxy to ", dir, " if you encounter any bugs please report them on the repo or discord!")
	_, _ = fmt.Scanln("Press Any key to close this CLI!")
}

func RequestOption(msg string, defaults ...string) (string, error) {
	fmt.Print(fmt.Sprintf(msg	+ " [%s]: ", strings.Join(defaults, "/")))
	var selected string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	selected = scanner.Text()
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	if strings.TrimSpace(selected) == "" {
		selected = defaults[0]
	}
	return selected, scanner.Err()
}

