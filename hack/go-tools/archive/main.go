package main

import (
	"compress/gzip"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var inPath, outPath string
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:        "archive",
				Usage:       "Convert json file to gzip",
				UsageText:   "archive <in> <out>",
				Description: "Arguments:\n  in   input json file\n  out  output gzip file",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:        "in",
						Destination: &inPath,
						UsageText:   "input json file",
					},
					&cli.StringArg{
						Name:        "out",
						Destination: &outPath,
						UsageText:   "output binary file",
					},
				},
				Action: func(context.Context, *cli.Command) error {
					checksumMatch, err := Archive(inPath, outPath)
					if err != nil {
						return err
					}

					if checksumMatch {
						fmt.Println("Checksum matched...")
					}

					return nil
				},
			},
			{
				Name:        "unarchive",
				Usage:       "Unarchive input file",
				UsageText:   "unarchive <in> <out>",
				Description: "Arguments:\n  in   input archive file\n  out  output file to unarchive",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:        "in",
						Destination: &inPath,
						UsageText:   "input archive file",
					},
					&cli.StringArg{
						Name:        "out",
						Destination: &outPath,
						UsageText:   "output file to unarchive",
					},
				},
				Action: func(context.Context, *cli.Command) error {
					err := Unarchive(inPath, outPath)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func Archive(inPath, outPath string) (bool, error) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return false, fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	outMd5File, err := os.OpenFile(fmt.Sprintf("%s.md5", outPath), os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return false, fmt.Errorf("failed to open md5 file: %w", err)
	}
	defer outMd5File.Close()

	hasher := md5.New()
	match, err := checksumMatch(inFile, outMd5File, hasher)
	if err != nil {
		return false, fmt.Errorf("checksum match failed: %w", err)
	}
	if match {
		return true, nil
	}

	// return input file to beginning
	_, err = inFile.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}

	// clear hash file
	_, err = outMd5File.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}
	err = outMd5File.Truncate(0)
	if err != nil {
		return false, err
	}

	outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return false, fmt.Errorf("failed to open out file: %w", err)
	}
	defer outFile.Close()

	compressWriter := gzip.NewWriter(outFile)
	defer compressWriter.Close()
	multiWriter := io.MultiWriter(compressWriter, hasher)
	_, err = io.Copy(multiWriter, inFile)
	if err != nil {
		return false, fmt.Errorf("multi write failed: %w", err)
	}
	if err := compressWriter.Close(); err != nil {
		return false, fmt.Errorf("compressed file did not finalize: %w", err)
	}

	sum := hex.EncodeToString(hasher.Sum(nil))
	_, err = outMd5File.WriteString(sum)
	if err != nil {
		return false, fmt.Errorf("hash write failed: %w", err)
	}

	if err := os.Remove(inFile.Name()); err != nil {
		return false, fmt.Errorf("failed to remove source file: %w", err)
	}

	return false, nil
}

func checksumMatch(inFile, hashFile *os.File, hasher hash.Hash) (bool, error) {
	_, err := io.Copy(hasher, inFile)
	if err != nil {
		return false, err
	}

	newHash := hex.EncodeToString(hasher.Sum(nil))
	// do not leave polluted hasher, if there is no hash match it will be reused
	hasher.Reset()

	hashBytes, err := io.ReadAll(hashFile)
	if err != nil {
		return false, err
	}

	if newHash == string(hashBytes) {
		return true, nil
	}

	return false, nil
}

func Unarchive(inPath, outPath string) error {
	inFile, err := os.OpenFile(inPath, os.O_RDONLY, 0o644)
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer inFile.Close()

	zipReader, err := gzip.NewReader(inFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	outFile, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("could not open output file: %w", err)
	}

	_, err = io.Copy(outFile, zipReader)
	if err != nil {
		return fmt.Errorf("could not unzip file: %w", err)
	}
	if err := outFile.Close(); err != nil {
		return fmt.Errorf("could not close output file and flush it to disk: %w", err)
	}

	return nil
}
