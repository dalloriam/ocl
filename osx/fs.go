package osx

import "os"

// FileExists returns true if the file exists at the given path.
// It returns false if the path does not exist, if the path is a directory, or if there was an error checking.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists returns true if the directory exists at the given path.
// It returns false if the path does not exist, if the path is a file, or if there was an error checking.
func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// EnsureDirectoryExists creates the directory at the given path if it does not exist.
func EnsureDirectoryExists(path string) error {
	if DirectoryExists(path) {
		return nil
	}

	if FileExists(path) {
		return os.ErrExist
	}

	return os.MkdirAll(path, os.ModePerm)
}
