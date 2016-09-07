package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/*
	 * 读取和加载class文件
	 */
	readClass(className string) ([]byte, Entry, error)
	/*
	 * toString
	 */
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}

	if strings.HasSuffix(strings.ToLower(path), "jar") || strings.HasSuffix(strings.ToLower(path), "zip") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
