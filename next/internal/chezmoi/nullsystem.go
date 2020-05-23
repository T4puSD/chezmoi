package chezmoi

import (
	"os"
	"os/exec"
)

// A nullReaderSystem simulates an empty system.
type nullReaderSystem struct{}

// A nullWriterSystem is a silently ignores all modifications.
type nullWriterSystem struct{}

func (nullReaderSystem) Get(bucket, key []byte) ([]byte, error)            { return nil, nil }
func (nullReaderSystem) Glob(pattern string) ([]string, error)             { return nil, nil }
func (nullReaderSystem) IdempotentCmdOutput(cmd *exec.Cmd) ([]byte, error) { return cmd.Output() }
func (nullReaderSystem) Lstat(name string) (os.FileInfo, error)            { return nil, os.ErrNotExist }
func (nullReaderSystem) Stat(name string) (os.FileInfo, error)             { return nil, os.ErrNotExist }
func (nullReaderSystem) ReadDir(dirname string) ([]os.FileInfo, error)     { return nil, os.ErrNotExist }
func (nullReaderSystem) ReadFile(filename string) ([]byte, error)          { return nil, os.ErrNotExist }
func (nullReaderSystem) Readlink(name string) (string, error)              { return "", os.ErrNotExist }

func (nullWriterSystem) Chmod(name string, mode os.FileMode) error                      { return nil }
func (nullWriterSystem) Delete(bucket, key []byte) error                                { return nil }
func (nullWriterSystem) Mkdir(dirname string, perm os.FileMode) error                   { return nil }
func (nullWriterSystem) RemoveAll(name string) error                                    { return nil }
func (nullWriterSystem) Rename(oldpath, newpath string) error                           { return nil }
func (nullWriterSystem) RunScript(scriptname, dir string, data []byte) error            { return nil }
func (nullWriterSystem) Set(bucket, key, value []byte) error                            { return nil }
func (nullWriterSystem) WriteFile(filename string, data []byte, perm os.FileMode) error { return nil }
func (nullWriterSystem) WriteSymlink(oldname, newname string) error                     { return nil }
