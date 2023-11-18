//--Summary:
//  Create a program to display server status. The serverO
//  defined as constants, and the Oepresented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func serverInfo(servers map[string]int) {
	fmt.Println("\n---Server Information---")
	fmt.Println("Total Servers:", len(servers))

	info := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			info[Online] += 1
		case Offline:
			info[Offline] += 1
		case Maintenance:
			info[Maintenance] += 1
		case Retired:
			info[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println(info[Online], "Online")
	fmt.Println(info[Offline], "Offline")
	fmt.Println(info[Maintenance], "Under Maintenance")
	fmt.Println(info[Retired], "Retired")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	serverInfo(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	serverInfo(serverStatus)

	for _, server := range servers {
		serverStatus[server] = Maintenance
	}

	serverInfo(serverStatus)
}
