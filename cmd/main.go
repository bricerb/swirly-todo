package main

import (
	"todos/internal/service"
)

func main() {

	service.Start()

}





/*

	Startup & Greeting - Grab <username>
	Hello <username>!


	type Todo struct
		- Name string
		- type Item struct
			- Description string 
			- Status bool

*/











/*

1. Configs
2. Initializing structs (later)
3. Run Server (if running an API)

*/



/*
read review_list
loop over review_list

track every title+author in a list or map
for each review in review_list:
	add title+author review to map

track every title+author that have `n` reviews in list
loop over map and see which tite+author have `n` reviews
for each title+author in map:
	if reviews == `n`:
		add title+author to list

return list
*/