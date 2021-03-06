package fs

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrNoDriver is the error for a nonfunctional file system.
var ErrNoDriver = errors.New("fs: nil driver")

// FileSystem is a concrete file system that implements a VFS contract.
type FileSystem struct {
	drv Driver
}

// Driver is the internal file system implementation.
type Driver interface {
	MkdirAll(dirname string) error
	ReadDir(dirname string) ([]FileInfo, error)
	ReadFile(filename string) ([]byte, error)
	Stat(filename string) (FileInfo, error)
	Symlink(oldname, newname string) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

// New creates a new FileSystem with drv as its engine.
func New(drv Driver) FileSystem {
	return FileSystem{drv}
}

// MkdirAll creates directories and their parents, if needed.
func (fs FileSystem) MkdirAll(dirname string) error {
	fs.testDriver()
	dirname = filepath.FromSlash(dirname)
	return fs.drv.MkdirAll(dirname)
}

// ReadDir lists names of files from dirname.
func (fs FileSystem) ReadDir(dirname string) ([]FileInfo, error) {
	fs.testDriver()
	dirname = filepath.FromSlash(dirname)
	return fs.drv.ReadDir(dirname)
}

// ReadFile returns the content of filename.
func (fs FileSystem) ReadFile(filename string) ([]byte, error) {
	fs.testDriver()
	filename = filepath.FromSlash(filename)
	return fs.drv.ReadFile(filename)
}

// Stat returns information about a file.
func (fs FileSystem) Stat(filename string) (FileInfo, error) {
	fs.testDriver()
	filename = filepath.FromSlash(filename)
	return fs.drv.Stat(filename)
}

// Symlink creates a symlink of oldname as newname.
func (fs FileSystem) Symlink(oldname, newname string) error {
	fs.testDriver()
	oldname = filepath.FromSlash(oldname)
	newname = filepath.FromSlash(newname)
	return fs.drv.Symlink(oldname, newname)
}

// WriteFile writes data to filename with permission perm.
func (fs FileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	fs.testDriver()
	filename = filepath.FromSlash(filename)
	return fs.drv.WriteFile(filename, data, perm)
}

func (fs FileSystem) testDriver() {
	if fs.drv == nil {
		panic(ErrNoDriver)
	}
}
