/*
filename extensions enhancement
Copyright © 2024 sunmoondevlab
*/
package fileextenh

import (
	"path/filepath"
	"strings"
)

/*
Get basename from filename.
*/
func Base(fname string, withExt bool) string {
	if withExt {
		return filepath.Base(fname)
	} else {
		x := Ext(fname)
		fb := strings.TrimSuffix(fname, x)
		return filepath.Base(fb)
	}
}

/*
Get extension from filename.
*/
func Ext(fname string) string {
	x := filepath.Ext(fname)
	// If there is a possibility of a double extension defined in this program (such as .gz or tar compressed file)
	if possibleDoubleExt(x) {
		des := getDblExts(x)
		// If the specified file has a double extension, rewrite it.
		for _, de := range des {
			if strings.HasSuffix(fname, de) {
				x = de
				break
			}
		}
	}
	return x
}

/*
Match file name extension  the specified extension?
*/
func MatchFileExt(fname string, ext string) bool {
	ax := Ext(fname)
	return ax == ext
}

func getDblExts(ext string) []string {
	de, e := doubleExtMap[ext]
	if e {
		return de
	} else {
		return []string{}
	}
}

func possibleDoubleExt(ext string) bool {
	_, e := doubleExtMap[ext]
	return e
}

var doubleExtMap = map[string][]string{
	".gz":   {".tar.gz"},
	".bz2":  {".tar.bz2"},
	".Z":    {".tar.Z"},
	".xz":   {".tar.xz"},
	".lz":   {".tar.lz"},
	".lzma": {".tar.lzma"},
	".lzo":  {".tar.lzo"},
	".zst":  {".tar.tzst"},
}
