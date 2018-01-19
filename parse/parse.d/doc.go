// Dataflow kit - main
//
// Copyright © 2017-2018 Slotix s.r.o. <dm@slotix.sk>
//
//
// All rights reserved. Use of this source code is governed
// by the BSD 3-Clause License license.

/*
Parse service of the Dataflow kit parses html content from web pages following the rules described in Payload JSON file.

Here is a simple example for requesting Parse endpoint:
		curl -XPOST  127.0.0.1:8001/parse -d '
		{
		"name":"collection",
		"request":{
			"url":"https://example.com"
		},
		"fields":[
			{
				"name":"Title",
				"selector":".product-container a",
				"extractor":{
					"type":"link",
            		"filters":[  
                		"trim",
                		"lowerCase"
             		],
					"params":{
					"includeIfEmpty":false
					}
				}
			},
			{
				"name":"Image",
				"selector":"#product-container img",
				"extractor":{
					"type":"image",
					"filters":[  
						"trim",
						"upperCase"
					]
				}
			},
			{
				"name":"Buyinfo",
				"selector":".buy-info",
				"extractor":{
					"type":"text",
					"params":{
					"includeIfEmpty":false
					}
				}
			}
		],
		"paginator":{
			"selector":".next",
			"attr":"href",
			"maxPages":3
		},
		"format":"json",
		"paginateResults":false
		}'


Name

Collection name

Request

Request parameters are passed to Fetch Endpoint for downloading html pages.
URL holds the URL address of the web page to be downloaded. URL is required. All other fields including Params, Cookies, Func are optional.

Params is a string value for passing formdata parameters.
For example it may be used for processing pages which require authentication.
	"auth_key=880ea6a14ea49e853634fbdc5015a024&referer=http%3A%2F%2Fexample.com%2F&ips_username=user&ips_password=userpassword&rememberMe=1"

Cookies contain cookies to be added to request  before sending it to browser.
It may be used for processing pages after initial authentication. In the first step formdata with auth info is passed to a web page.
Response object headers may contain an Object like
	name: "Set-Cookie"
	value: "session_id=29d7b97879209ca89316181ed14eb01f; path=/; httponly"
These cookies should be passed to the next pages on the same domain.
	"session_id", "29d7b97879209ca89316181ed14eb01f", "/", domain="example.com"

Fields

A set of fields used to extract data from a web page.
A Field represents a given chunk of data to be extracted from every block in each page of a scrape.

Field name is required, and will be used to aggregate results.

Selector represents a CSS selector within the given block to process.  Pass in "." to use the root block's selector.

Extractor contains the logic on how to extract some results from the selector that is provided to this Field.

Paginator

Paginator is used to scrape multiple pages.
If there is no paginator in Payload, then no pagination is performed and it is assumed that the initial URL is the only page.
Paginator extracts the next page from a document by querying a given CSS selector and extracting the given HTML attribute from the resulting element.

Selector represents corresponding CSS selector for the next page along with
Attr defining HTML attribute for the next page.
MaxPages sets upper bound to maximum number of pages to scrape. The scrape will proceed until either this number of pages have been scraped, or until the paginator returns no further URLs.
Default value is 1.
Set maxPages value to 0 to indicate an unlimited number of pages to be scraped.

Format

The following Output formats are available: CSV, JSON, XML

paginateResults

Paginated results are returned if paginateResults is true.
Single list of combined results from every block on all pages is returned by default.
Paginated results are applicable for JSON and XML output formats.
Combined list of results is always returned for CSV format.
*/
//
package main

// EOF
