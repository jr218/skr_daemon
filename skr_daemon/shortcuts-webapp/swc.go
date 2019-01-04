/*
SWC stands for 'shortcuts webapp controller. Its goal is to provide a 
command line interface for managing my webapp. With this, the user will 
be able to turn the web server on and off, and provide a the option to 
turn it on by default.
*/

package main

import "os"

func main() {

	argc := os.Args
	argv := os.Args[1:]

	if(argc != 0) {
		switch(argv) {
			case "start":
				break
			case "stop":
				if(is_stopped() == true) {
				}
				break
			case "restart":
				break
			case "addlink":
				addlink()
				break	
		}
	}
}

func addlink() {

	//A function that edits a text file (links.txt) for use in
	//list_shortcuts.js.
}

func is_stopped() bool {

	//Test to see if the server is running.

		if(server_running) {

			return false
		}

		return true
}

func restart_server() {

	if(is_stopped == true) {

		return
	} else {

		//Stop the server and start it up again.
	}
}




















