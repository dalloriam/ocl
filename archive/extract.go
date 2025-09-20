package archive

import (
	"errors"

	"github.com/dalloriam/ocl/osx"
)

var ErrSourceDoesNotExist = errors.New("source file does not exist")

// Extract extracts the archive file located at srcFile into the directory dstDir.
func Extract(srcFile, dstDir string) error {
	if !osx.FileExists(srcFile) {
		return ErrSourceDoesNotExist
	}

	if err := osx.EnsureDirectoryExists(dstDir); err != nil {
		return err
	}

	format, err := detectFormat(srcFile)
	if err != nil {
		return err
	}

	extractor, err := newExtractor(format)
	if err != nil {
		return err
	}

	return extractor.Extract(srcFile, dstDir)
}
