package gomplate

import (
	"github.com/readystock/golog"
	"go/ast"
	"path/filepath"
)

type ReflectionInfo struct {
	PackageName   string
	TypeName      string
	LowerTypeName string
}

func GenerateTemplate(templateName string, typeName string, output string, args []string) {
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

	values := make([]Value, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		file.values = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			values = append(values, file.values...)
		}
	}
	golog.Infof("working in dir: %s", dir)
}
