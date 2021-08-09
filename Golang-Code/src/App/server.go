package main

import (
	"github.com/hewiefreeman/GopherGameServer"
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
	"module/"
)
func main() {

	settings := gopher.ServerSettings{
		ServerName:     "!s!",
		MaxConnections: 0,
		
		// This should work for localhost. But setting hostname and alias help for WAN
		HostName:  "", // Enter a Hostname
		HostAlias: "", // Enter Alias
		IP:        "0.0.0.0",
		Port:      443,

		TLS: false,
		CertFile: "", // Cert (optional)
		PrivKeyFile: "", // Enter a PrivKey (optional)

		AdminLogin: "admin",
		AdminPassword: "admin",
		//OriginOnly: true,
	}
	// Make a Room type and set broadcasts and callbacks
	createCustomActions() // Creates all custom actions required for client
	createServerCallbacks()

	mainchat := createChatRoomType()
	lobby := createLobbyRoomType()

	//chat := createChat()
	log.Println(mainchat)
	log.Println(lobby)
	//createPokerTable() // if button click
	createPokerRoomTypes()

	openSecondTestTable() // FOR DEBUGGING!
	openTestTable() // FOR DEBUGGING!
	gopher.Start(&settings)
}

func openRoom(rName string, rtype string) *core.Room{
	// Open a Room
	room, roomErr := core.NewRoom(rName, rtype, false, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return room
	}
}

func openPrivateRoom(rName string, rtype string) *core.Room{ // TODO IDK IF THIS IS NEEDED
	// Open a Room
	room, roomErr := core.NewRoom(rName, rtype, true, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
		return nil
	} else {
		return room
	}
}




