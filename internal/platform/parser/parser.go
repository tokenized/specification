package parser

import (
	"os"
	"path/filepath"
	"strings"
)

func FetchFiles(folder, version string) []string {

	dir := "./" + folder + "/" + version + "/"
	filenames := []string{}

	fn := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Do not go in to sub directories
		if filepath.Dir(path) != dir {
			return nil
		}

		filenames = append(filenames, path)

		return nil
	}

	if err := filepath.Walk(dir, fn); err != nil {
		panic(err)
	}

	return filenames
}
