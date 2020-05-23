package chezmoi

var _ System = struct {
	nullReaderSystem
	nullWriterSystem
}{}
