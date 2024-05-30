package fileutils

import "strings"

func SwapFsPathToRoutePath(fsPath string) string {
	return strings.Join(strings.Split(fsPath, "\\"), "/")
}
