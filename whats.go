/*
 * whats: A tool to quickly look up something
 *
 * Copyright (c) 2015 Arjun Sreedharan <arjun024@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
/*
 * whats.go
 * Entry source file
 */

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/arjun024/whats/whatslib/google"
)

const AUTHOR = "Arjun Sreedharan <arjun024@gmail.com>"
const VERSION = "0.0.1"

const DEBUG = false
const REFERER = "http://arjunsreedharan.org"
const GOOGLE_URI = "https://ajax.googleapis.com" +
	"/ajax/services/search/web?v=1.0&q="

func usage() {
	fmt.Printf("%s\n%s%s\n%s%s\n",
		"SYNTAX : whats <SOMETHING>",
		"AUTHOR : ", AUTHOR,
		"VERSION: ", VERSION)
	os.Exit(0)
}

func strip_html(str string) string {
	regexp_html := regexp.MustCompile("<[^>]*>")
	return html.UnescapeString(regexp_html.ReplaceAllString(str, ""))
}

/* From the top 4 results, let me guess which's best */
func guess(r []google.ResultsType) int {
	cues := []string{
		" is a ",
		" are a ",
		" was as ",
		" were a ",
		" defined as ",
		" developed as a ",
	}
	for i, result := range r {
		if strings.Contains(result.VisibleUrl, "wikipedia.org") {
			return i % 3
		}
	}
	for i, result := range r {
		for _, cue := range cues {
			if strings.Contains(result.Content, cue) {
				return i % 3
			}
		}
	}
	return 0
}

func output(g *google.GoogleApiDataType) {
	i := guess((*g).ResponseData.Results)
	content := strip_html((*g).ResponseData.Results[i].Content)
	fmt.Printf("\n%s\n\n", content)
}

func main() {
	var query string
	var gdata google.GoogleApiDataType

	if len(os.Args) == 1 {
		usage()
	}
	query = GOOGLE_URI + url.QueryEscape(strings.Join(os.Args[1:], " "))

	client := &http.Client{}
	req, err := http.NewRequest("GET", query, nil)
	req.Header.Set("Referer", REFERER)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in making HTTP request: %s\n",
			err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&gdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse json: %s\n",
			err.Error())
		os.Exit(1)
	}

	output(&gdata)
}
