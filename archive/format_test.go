package archive

import "testing"

func TestDetectFormat(t *testing.T) {

	cases := []struct {
		filename     string
		expected     Format
		expectingErr bool
	}{
		{"archive.tar", FormatTar, false},
		{"archive.gz", FormatTar, false},
		{"archive.bz2", FormatTar, false},
		{"archive.xz", FormatTar, false},
		{"archive.zst", FormatTar, false},
		{"archive.tgz", FormatTar, false},
		{"archive.tbz2", FormatTar, false},
		{"archive.txz", FormatTar, false},
		{"archive.tzst", FormatTar, false},
		{"archive.zip", FormatZip, false},
		{"archive.7z", Format7z, false},
		{"archive.rar", "", true},
		{"archive.unknown", "", true},
	}

	for _, c := range cases {
		t.Run(c.filename, func(t *testing.T) {
			format, err := detectFormat(c.filename)
			if c.expectingErr {
				if err == nil {
					t.Errorf("expected an error for %s, but got none", c.filename)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error for %s, but got: %v", c.filename, err)
				}
				if format != c.expected {
					t.Errorf("expected format %s for %s, but got %s", c.expected, c.filename, format)
				}
			}
		})
	}
}
