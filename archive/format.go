package archive

import (
	"path/filepath"
)

type Format string

const (
	FormatTar Format = "tar" // Anything extracted with `tar` command
	FormatZip Format = "zip"
	Format7z  Format = "7z"
)

type ErrUnknownFormat struct {
	Ext string
}

func (e *ErrUnknownFormat) Error() string {
	return "unknown archive format: " + e.Ext
}

func (e *ErrUnknownFormat) Is(target error) bool {
	_, ok := target.(*ErrUnknownFormat)
	return ok
}

func detectFormat(filename string) (Format, error) {
	ext := filepath.Ext(filename)

	switch ext {
	case ".tar", ".gz", ".bz2", ".xz", ".zst", ".tgz", ".tbz2", ".txz", ".tzst":
		return FormatTar, nil
	case ".zip":
		return FormatZip, nil
	case ".7z":
		return Format7z, nil
	default:
		return "", &ErrUnknownFormat{Ext: ext}
	}
}
