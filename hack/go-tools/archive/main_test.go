package main_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	main "github.com/afterlook/poe-carlo/hack/go-tools/archive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {
	// given: in
	inDir, err := os.CreateTemp("", "in")
	require.NoError(t, err, "Input file not created")
	inFile, err := os.OpenFile(inDir.Name(), os.O_WRONLY, 0o644)
	require.NoError(t, err, "Must open file created for testing")

	testContent := []byte("test content")
	_, err = inFile.Write(testContent)
	require.NoError(t, err, "Must be able to write test content to input file")

	hasher := md5.New()
	_, err = hasher.Write(testContent)
	require.NoError(t, err, "Must be able to write md5 hash")
	want := make([]byte, hex.EncodedLen(hasher.Size()))
	_ = hex.Encode(want, hasher.Sum(nil))

	// given: out
	outDir, err := os.CreateTemp("", "in")
	require.NoError(t, err, "Output file not created")
	t.Cleanup(func() {
		_ = os.Remove(outDir.Name())
	})

	// when
	checksumMatched, err := main.Archive(inDir.Name(), outDir.Name())
	require.NoError(t, err, "Failed to run compression")

	// then
	md5FileName := fmt.Sprintf("%s.md5", outDir.Name())
	t.Cleanup(func() {
		_ = os.Remove(md5FileName)
	})

	outMd5File, err := os.Open(md5FileName)
	require.NoError(t, err, "After compression md5 file does not exist")

	got, err := io.ReadAll(outMd5File)
	require.NoError(t, err, "After compression could not read file")

	assert.Equal(t, string(want), string(got), "Hashes should be equal")
	assert.NoFileExists(t, inFile.Name(), "source file should not exist anymore")
	// first pass, fresh files
	assert.False(t, checksumMatched)

	// recreate input file
	inDir, err = os.CreateTemp("", "in")
	require.NoError(t, err, "Input file not created")
	inFile, err = os.OpenFile(inDir.Name(), os.O_WRONLY, 0o644)
	require.NoError(t, err, "Must open file created for testing")
	_, err = inFile.Write(testContent)
	require.NoError(t, err, "Must be able to write test content to input file")

	checksumMatched, err = main.Archive(inDir.Name(), outDir.Name())
	require.NoError(t, err, "Second run of zipping should not fail on the same files")

	// now we should just exit as the content was the same
	assert.True(t, checksumMatched)
}

func TestDecompression(t *testing.T) {
	// given: in
	inDir, err := os.CreateTemp("", "in")
	require.NoError(t, err, "Input file not created")
	inFile, err := os.OpenFile(inDir.Name(), os.O_WRONLY, 0o644)
	require.NoError(t, err, "Must open file created for testing")

	testContent := []byte("test content")
	_, err = inFile.Write(testContent)
	require.NoError(t, err, "Must be able to write test content to input file")

	// given: out
	outDir, err := os.CreateTemp("", "in")
	require.NoError(t, err, "Output file not created")
	t.Cleanup(func() {
		_ = os.Remove(outDir.Name())
	})

	// given: out
	outDecompressionDir, err := os.CreateTemp("", "in")
	require.NoError(t, err, "Output decompression file not created")
	t.Cleanup(func() {
		_ = os.Remove(outDecompressionDir.Name())
	})

	// given succesful archive
	_, err = main.Archive(inDir.Name(), outDir.Name())
	require.NoError(t, err, "Failed to run compression")

	// when
	err = main.Unarchive(outDir.Name(), outDecompressionDir.Name())
	require.NoError(t, err, "Failed to decompress file")

	content, err := io.ReadAll(outDecompressionDir)
	require.NoError(t, err, "Could not read decompressed output")

	assert.Equal(t, testContent, content)
}
