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
 * google.go
 * Contains abstraction of the Google web search api results
 * google api link: https://ajax.googleapis.com/ajax/services/search/web
 * Note: the api is deprecated
 */

package google

type GoogleApiDataType struct {
	ResponseData ResponseDataType
	ResponseDetails string
	ResponseStatus int
}

type ResponseDataType struct {
	Results []ResultsType
	Cursor CursorType
}

type ResultsType struct {
	GsearchResultClass string
	UnescapedUrl string
	Url string
	VisibleUrl string
	CacheUrl string
	Title string
	TitleNoFormatting string
	Content string
}

type CursorType struct {
	ResultCount string
	Pages []PagesType
	EstimatedResultCount string
	CurrentPageIndex int
	MoreResultsUrl string
	SearchResultTime string
}

type PagesType struct {
	Start string
	Label int
}
