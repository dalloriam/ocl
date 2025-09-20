package archive

import "os/exec"

type extractor interface {
	// Extract extracts the archive file located at srcFile into the directory dstDir.
	// It returns an error if the extraction fails.
	// `srcFile` is guaranteed to exist and be a regular file.
	// `dstDir` is guaranteed to exist and be a directory.
	Extract(srcFile, dstDir string) error
}

type tarExtractor struct{}

func (e *tarExtractor) Extract(srcFile, dstDir string) error {
	cmd := exec.Command("tar", "-xf", srcFile, "-C", dstDir)
	return cmd.Run()
}

type zipExtractor struct{}

func (e *zipExtractor) Extract(srcFile, dstDir string) error {
	cmd := exec.Command("unzip", "-q", srcFile, "-d", dstDir)
	return cmd.Run()
}

type sevenZExtractor struct{}

func (e *sevenZExtractor) Extract(srcFile, dstDir string) error {
	cmd := exec.Command("7z", "x", "-y", "-o"+dstDir, srcFile)
	return cmd.Run()
}

func newExtractor(format Format) (extractor, error) {
	switch format {
	case FormatTar:
		return &tarExtractor{}, nil
	case FormatZip:
		return &zipExtractor{}, nil
	case Format7z:
		return &sevenZExtractor{}, nil
	default:
		return nil, &ErrUnknownFormat{}
	}
}
