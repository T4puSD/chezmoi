package chezmoi

import (
	"os"
	"os/exec"
)

type DiscardWritesSystem struct {
	nullWriterSystem
	s System
}

func NewDiscardWritesSystem(s System) System {
	return &DiscardWritesSystem{
		s: s,
	}
}

func (s *DiscardWritesSystem) Get(bucket, key []byte) ([]byte, error) { return s.s.Get(bucket, key) }
func (s *DiscardWritesSystem) Glob(pattern string) ([]string, error)  { return s.s.Glob(pattern) }
func (s *DiscardWritesSystem) IdempotentCmdOutput(cmd *exec.Cmd) ([]byte, error) {
	return s.s.IdempotentCmdOutput(cmd)
}
func (s *DiscardWritesSystem) Lstat(name string) (os.FileInfo, error) { return s.s.Lstat(name) }
func (s *DiscardWritesSystem) Stat(name string) (os.FileInfo, error)  { return s.s.Stat(name) }
func (s *DiscardWritesSystem) ReadDir(dirname string) ([]os.FileInfo, error) {
	return s.s.ReadDir(dirname)
}

func (s *DiscardWritesSystem) ReadFile(filename string) ([]byte, error) {
	return s.s.ReadFile(filename)
}
func (s *DiscardWritesSystem) Readlink(name string) (string, error) { return s.s.Readlink(name) }
