package commands

import (
	"fmt"
	"mao/src/libs" // Assuming this path is correct
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:          "(tagall)", // Clear and consistent naming
		As:            []string{"tagall"},
		Tags:          "group",
		IsPrefix:      true,
		// IsOwner:       true, // Consider whether ownership is necessary
		// IsQuerry:      false, // Remove irrelevant flags
		IsWaitt:       true,
		Exec:          tagAllMembers,
		Usage:         ".tagall", // Example usage for users
		Description:   "Tag all members in the current group",
		MinParameters: 0,
		MaxParameters: 0,
	})
}

func tagAllMembers(client *libs.NewClientImpl, m *libs.IMessage) {
	// Retrieve participant information
	participants, err := client.WA.GetGroupParticipants(m.From)
	if err != nil {
		m.Reply("Gagal mendapatkan informasi anggota grup.")
		return
	}

	// Send individual messages to tag each participant
	for _, participant := range participants {
		// Assuming participant.JID is the correct identifier for tagging
		client.WA.SendMessage(m.From, fmt.Sprintf("@%s", participant.JID))
	}

	// Send a final confirmation message
	m.Reply("Semua anggota grup telah di-tag.")
}
