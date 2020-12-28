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
	"fmt"
	"github.com/jszwedko/go-circleci"
	"io"
	"net/http"
	"os"
	"path"
)

// In this file contains all the circleci shit

type BuildOptions struct {
	Os string
	Processor int16
	BuildNum int32
}

func GetLatestBuildNum(client circleci.Client) (int32, error) {
	builds, err := client.ListRecentBuildsForProject("SunProxy", "sun", "master", "", 1, 0)
	if err != nil {
		// this will never happen.....
		return 0, err
	}
	return int32(builds[0].BuildNum), err
}

func GenerateDownloadLink(options BuildOptions) string {
	var ext = ""
	if options.Os == "windows" {
		ext = ".exe"
	}
	return fmt.Sprintf("https://%v-321156895-gh.circle-artifacts.com/0/tmp/artifacts/sun_%s%v%s", options.BuildNum, options.Os, options.Processor, ext)
}

func DownloadBuild(link string, dir string) error {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}
	_ = os.Mkdir(dir, 0644)
	f, err := os.Create(dir + "/" + path.Base(req.URL.Path))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	//Close http request after the copying is done
	defer resp.Body.Close()
	//Close file after the copying is done
	defer f.Close()
	//Copy file buffer from the resp into the file.
	_, err = io.Copy(f, resp.Body)
	return err
}