package lobby

import (
	"github.com/hewiefreeman/GopherGameServer/core"
)
func createLobbyRoomType() *core.Room{
	// Make a Room type and set broadcasts and callbacks
	lobbyRoomType := core.NewRoomType("Lobby", true)// TODO FIX NEEDING TO LOGIN FIRST!
	lobbyRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterLobby).SetUserLeaveCallback(onLeaveLobby)

	// Create lobby room on startup so people auto join in
	lobbyroom := openRoom("mainLobby", "Lobby") // TODO THIS WONT WORK FULLY, NOT CUSTOM
	return lobbyroom
}

