package archiver

// Copied from https://github.com/mholt/archiver

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kyma-project/kyma/components/helm-broker/platform/deferutil"
)

// TarGz is for TarGz format
var TarGz tarGzFormat

func init() {
	RegisterFormat("TarGz", TarGz)
}

type tarGzFormat struct{}

func (tarGzFormat) Match(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".tar.gz") ||
		strings.HasSuffix(strings.ToLower(filename), ".tgz") ||
		isTarGz(filename)
}

// isTarGz checks the file has the gzip compressed Tar format header by reading
// its beginning block.
func isTarGz(targzPath string) bool {
	f, err := os.Open(targzPath)
	if err != nil {
		return false
	}
	defer deferutil.CheckFn(f.Close, "while closing %s file", targzPath)

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return false
	}
	defer deferutil.CheckFn(gzr.Close, "while closing %s file gzip reader", targzPath)

	buf := make([]byte, tarBlockSize)
	n, err := gzr.Read(buf)
	if err != nil || n < tarBlockSize {
		return false
	}

	return hasTarHeader(buf)
}

// Write outputs a .tar.gz file to a Writer containing
// the contents of files listed in filePaths. It works
// the same way Tar does, but with gzip compression.
func (tarGzFormat) Write(output io.Writer, filePaths []string) error {
	return writeTarGz(filePaths, output, "")
}

// Make creates a .tar.gz file at targzPath containing
// the contents of files listed in filePaths. It works
// the same way Tar does, but with gzip compression.
func (tarGzFormat) Make(targzPath string, filePaths []string) error {
	out, err := os.Create(targzPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %v", targzPath, err)
	}
	defer deferutil.CheckFn(out.Close, "while closing %s file", targzPath)

	return writeTarGz(filePaths, out, targzPath)
}

func writeTarGz(filePaths []string, output io.Writer, dest string) error {
	gzw := gzip.NewWriter(output)
	defer deferutil.CheckFn(gzw.Close, "while closing gzip writer")

	return writeTar(filePaths, gzw, dest)
}

// Read untars a .tar.gz file read from a Reader and decompresses
// the contents into destination.
func (tarGzFormat) Read(input io.Reader, destination string) error {
	gzr, err := gzip.NewReader(input)
	if err != nil {
		return fmt.Errorf("error decompressing: %v", err)
	}
	defer deferutil.CheckFn(gzr.Close, "while closing gzip reader")

	return Tar.Read(gzr, destination)
}

// Open untars source and decompresses the contents into destination.
func (tarGzFormat) Open(source, destination string) error {
	f, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("%s: failed to open archive: %v", source, err)
	}
	defer deferutil.CheckFn(f.Close, "while closing %s file", source)

	return TarGz.Read(f, destination)
}
