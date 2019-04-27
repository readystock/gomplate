package gomplate

import (
	"log"
	"path/filepath"
)

type ReflectionInfo struct {
	PackageName string
	TypeName    string
}

func GenerateTemplate(templateName string, typeNames []string, output string, args []string) {
	// Parse the package once.
	var dir string
	g := Generator{
		trimPrefix:  *trimprefix,
		lineComment: *linecomment,
	}
	// TODO(suzmue): accept other patterns for packages (directories, list of files, import paths, etc).
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}

	g.parsePackage(args, make([]string, 0))

}
