package ui

import (
	"log"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr"
	snr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	mn "github.com/sonr-io/core/pkg/menu"
	win "github.com/sonr-io/core/pkg/window"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Interface struct {
	// Client Properties
	node   *snr.Node
	assets ImageAssets

	// System Native Tray
	tLink      *systray.MenuItem
	tShare     *systray.MenuItem
	tCount     *systray.MenuItem
	tQuit      *systray.MenuItem
	tPeersList []*systray.MenuItem

	// Locale Info
	peerCount int32
	lobbySize int32
}

func Start() Interface {
	// Initialize Assets
	deviceBox := packr.NewBox("../assets/images/device")
	iconBox := packr.NewBox("../assets/images/icon")
	sysBox := packr.NewBox("../assets/images/system")

	// Default Interface Values
	ai := Interface{
		peerCount: 0,
		lobbySize: 1,
		assets: ImageAssets{
			Device: deviceBox,
			Icon:   iconBox,
			System: sysBox,
		},
	}

	// Set System Tray
	systray.SetTooltip("Sonr")

	// Check Platform
	if runtime.GOOS == "windows" {
		icon, err := iconBox.Find("tray.ico")
		if err != nil {
			log.Println(err)
		}
		systray.SetIcon(icon)
	} else {
		icon, err := iconBox.Find("tray.png")
		if err != nil {
			log.Println(err)
		}
		systray.SetTemplateIcon(icon, icon)
	}

	// Link Sonr Device
	ai.tLink = systray.AddMenuItem("Link Device", "Link a Device to Sonr")
	ai.tLink.SetTemplateIcon(ai.assets.SystemIcon(Link), ai.assets.SystemIcon(Link))
	ai.tShare = systray.AddMenuItem("Open Share", "Open Share Menu")

	// Quit Sonr
	ai.tQuit = systray.AddMenuItem("Quit", "Quit Sonr Desktop")
	ai.tQuit.SetTemplateIcon(ai.assets.SystemIcon(Close), ai.assets.SystemIcon(Close))
	systray.AddSeparator()

	// Pers Label
	ai.tCount = systray.AddMenuItem("Available Peers", "")
	ai.tCount.Disable()
	return ai
}

// ^ References Sonr Node and Handles Input ^ //
func (ai *Interface) Initialize(n *snr.Node) {
	// Set Node
	ai.node = n

	// Handle Menu Events
	go ai.HandleMenuInput()
}

// ^ Routine Handles Menu Input ^ //
func (ai *Interface) HandleMenuInput() {
	for {
		select {
		// @ File Item Clicked
		case <-ai.tQuit.ClickedCh:
			systray.Quit()

		case <-ai.tShare.ClickedCh:
			mn.OpenShareMenu()

			// @ Link Item Clicked
		case <-ai.tLink.ClickedCh:
			// Validate Node Set
			if ai.node != nil {
				// Create Link Request
				name := ShowNameDialog()
				linkRequest := ai.node.LinkRequest(name)

				// Get Node JSON
				jsonBytes, err := protojson.Marshal(linkRequest)
				if err != nil {
					log.Panicln(err)
				}

				// Display Window
				go win.OpenQRWindow(string(jsonBytes))
			} else {
				log.Println("Node not set.")
			}
		}
	}
}

// ^ Routine Handles Peer Item Input ^ //
func (ai *Interface) HandlePeerInput(fileItem *systray.MenuItem, linkItem *systray.MenuItem, peer *md.Peer) {
	for {
		select {
		// @ File Item Clicked
		case <-fileItem.ClickedCh:
			// Validate and Invite File
			if ai.node != nil {
				// Get File
				filename := ShowFileDialog()

				// Add Files
				files := make([]*md.InviteRequest_FileInfo, 0)
				files = append(files, &md.InviteRequest_FileInfo{
					Path: filename,
				})

				// Create Request
				procReq := md.InviteRequest{
					To:    peer,
					Files: files,
					Type:  md.InviteRequest_File,
				}

				// Convert to Bytes
				reqBytes, err := proto.Marshal(&procReq)
				if err != nil {
					log.Panicln(err)
				}

				ai.node.Invite(reqBytes)
			} else {
				log.Println("Node not set.")
			}

			// @ Link Item Clicked
		case <-linkItem.ClickedCh:
			// Validate and Invite URL
			if ai.node != nil {
				url := ShowURLDialog()

				// Create Request
				procReq := md.InviteRequest{
					To:   peer,
					Url:  url,
					Type: md.InviteRequest_File,
				}

				// Convert to Bytes
				reqBytes, err := proto.Marshal(&procReq)
				if err != nil {
					log.Panicln(err)
				}

				ai.node.Invite(reqBytes)
			} else {
				log.Println("Node not set.")
			}
		}
	}
}

// ^ Method to Rebuild Menu for Lobby Refresh ^ //
func (ai *Interface) RefreshPeers(newLob *md.Lobby, node *snr.Node) {
	// Set Node
	ai.node = node

	// Check if Lobby Updated
	if newLob.Size != ai.lobbySize {
		// Change Lobby
		ai.lobbySize = newLob.Size
		ai.peerCount = newLob.Size - 1

		// Reset Menu
		for _, mi := range ai.tPeersList {
			mi.Hide()
		}

		// Empty Peers
		ai.tPeersList = nil

		// Add Peers to Menu
		for _, p := range newLob.Peers {
			ai.SetPeerItem(p)
		}
	}
}

// ^ Method to Build Peer Item ^ //
func (ai *Interface) SetPeerItem(p *md.Peer) {
	// Add Peer to Menu
	peerItem := systray.AddMenuItem(p.Profile.FirstName, "")
	peerItem.SetTemplateIcon(ai.assets.DeviceIcon(p.Platform), ai.assets.DeviceIcon(p.Platform))

	// Add Peer Send Options
	urlItem := peerItem.AddSubMenuItem("Send URL", "Send a URL to "+p.Profile.FirstName)
	urlItem.SetTemplateIcon(ai.assets.SystemIcon(URL), ai.assets.SystemIcon(URL))
	fileItem := peerItem.AddSubMenuItem("Send File", "Send a File to "+p.Profile.FirstName)
	fileItem.SetTemplateIcon(ai.assets.SystemIcon(File), ai.assets.SystemIcon(File))

	// Spawn Routine to handle Peer Item Actions
	go ai.HandlePeerInput(fileItem, urlItem, p)

	// Add to Menu List
	ai.tPeersList = append(ai.tPeersList, peerItem)
}