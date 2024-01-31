package cliprdr

type Control struct {
	hwnd uintptr
}

func ClipWatcher(c *CliprdrClient) {}

const (
	FILE_ATTRIBUTE_DIRECTORY = 0x00000010
	CFSTR_FILEDESCRIPTORW    = "FileGroupDescriptorW"
)

func (c *Control) withOpenClipboard(f func()) {}
func (c *Control) SendCliprdrMessage()        {}

func EmptyClipboard() bool {
	return true
}

func SetClipboardData(format uint32, mem uintptr) bool {
	return true
}

func GetFileInfo(sys interface{}) (uint32, []byte, uint32, uint32) {
	return 0, nil, 0, 0
}

func GetFileNames() []string {
	return nil
}

func GetClipboardData(formatId uint32) string {
	return ""
}

func RegisterClipboardFormat(format string) uint32 {
	return 0
}

func GetFormatList(hwnd uintptr) []CliprdrFormat {
	return nil
}
