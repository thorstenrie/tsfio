// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public Licence v3.0
// that can be found in the LICENSE file.
package tsfio

var (
	goldenDir      Directory = "testdata/"
	goldenFileType string    = ".golden"
)

func goldenPath(tc string) Filename {
	return Filename(goldenDir) + Filename(tc+goldenFileType)
}

func SetGoldenDir(dir Directory) error {
	return nil
}
