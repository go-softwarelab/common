package main

import (
	"fmt"
	"go/build"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/princjef/gomarkdoc"
	"github.com/princjef/gomarkdoc/lang"
	"github.com/princjef/gomarkdoc/logger"
)

func main() {
	log := logger.New(logger.DebugLevel)

	// Create a renderer to output data
	out, err := gomarkdoc.NewRenderer(
		gomarkdoc.WithTemplateOverride("package", packageTpl),
		gomarkdoc.WithTemplateOverride("func", funcTpl),
		gomarkdoc.WithTemplateOverride("example", exampleTpl),
		gomarkdoc.WithTemplateFunc("accordionHeader", accordionHeader),
		gomarkdoc.WithTemplateFunc("accordionTerminator", accordionTerminator),
	)
	if err != nil {
		panic(err)
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}

	root := path.Dir(path.Dir(path.Dir(currentFile)))
	docsRoot := path.Join(root, "docs", "docs")
	packageRoot := path.Join(root, "pkg")

	// Replace the incomplete line with:
	err = filepath.WalkDir(packageRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the current directory
		if path == packageRoot {
			return nil
		}

		// Only process directories
		if !d.IsDir() {
			return nil
		}
		if d.Name()[0] == '.' || d.Name() == "internal" {
			return filepath.SkipDir
		}

		// Try to import the directory as a package
		buildPkg, err := build.ImportDir(path, build.ImportComment)
		if err != nil {
			// Skip directories that aren't Go packages
			return nil
		}

		// Create a documentation package
		pkg, err := lang.NewPackageFromBuild(log, buildPkg, lang.PackageWithRepositoryOverrides(&lang.Repo{DefaultBranch: "main"}))
		if err != nil {
			return fmt.Errorf("failed to create package from %s: %w", path, err)
		}

		// get difference between package root and current package
		relPath, err := filepath.Rel(packageRoot, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		docsDir := filepath.Join(docsRoot, relPath)
		if err := os.MkdirAll(docsDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", docsDir, err)
		}

		// Create output file in the package directory
		outputFile := filepath.Join(docsDir, "README.md")
		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outputFile, err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(fmt.Errorf("failed to close file %s: %w", outputFile, err))
			}
		}(file)

		// Generate documentation and write to file
		docs, err := out.Package(pkg)
		if err != nil {
			return fmt.Errorf("failed to generate docs for %s: %w", path, err)
		}

		if _, err = file.WriteString(docs); err != nil {
			return fmt.Errorf("failed to write docs to %s: %w", outputFile, err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

}

var packageTpl = `{{- if eq .Name "main" -}}
	{{- header .Level .Dirname -}}
{{- else -}}
	{{- header .Level .Name -}}
{{- end -}}
{{- spacer -}}

{{- template "import" . -}}
{{- spacer -}}

{{- if len .Doc.Blocks -}}
	{{- template "doc" .Doc -}}
	{{- spacer -}}
{{- end -}}

{{- range (iter .Examples) -}}
	{{- template "example" .Entry -}}
	{{- spacer -}}
{{- end -}}


{{- if len .Consts -}}
	{{- spacer -}}

	{{- header (add .Level 1) "Constants" -}}
	{{- spacer -}}

	{{- range (iter .Consts) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}

{{- end -}}

{{- if len .Vars -}}
	{{- spacer -}}

	{{- header (add .Level 1) "Variables" -}}
	{{- spacer -}}

	{{- range (iter .Vars) -}}
		{{- template "value" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}

{{- end -}}

{{- if len .Funcs -}}
	{{- spacer -}}

	{{- range (iter .Funcs) -}}
		{{- template "func" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

{{- if len .Types -}}
	{{- spacer -}}

	{{- range (iter .Types) -}}
		{{- template "type" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}
`

var funcTpl = `{{- if .Receiver -}}
	{{- rawAnchorHeader .Level (codeHref .Location | link (printf "%s.%s" .Receiver .Name)) .Anchor -}}
{{- else -}}
	{{- rawAnchorHeader .Level (codeHref .Location | link (escape .Name) | printf "%s") .Anchor -}}
{{- end -}}
{{- spacer -}}

{{- codeBlock "go" .Signature -}}
{{- spacer -}}

{{- template "doc" .Doc -}}

{{- if len .Examples -}}
	{{- spacer -}}

	{{- range (iter .Examples) -}}
		{{- template "example" .Entry -}}
		{{- if (not .Last) -}}{{- spacer -}}{{- end -}}
	{{- end -}}
{{- end -}}

`

var exampleTpl = `{{- accordionHeader .Title -}}
{{- spacer -}}

{{- template "doc" .Doc -}}
{{- spacer -}}

{{- codeBlock "go" .Code -}}
{{- spacer -}}

{{- if .HasOutput -}}

	{{- bold "Output" -}}
	{{- spacer -}}

	{{- codeBlock "" .Output -}}
	{{- spacer -}}
    
{{- end -}}

{{- accordionTerminator -}}

`

func accordionHeader(title string) string {
	return fmt.Sprintf("<details>\n<summary>%s</summary>\n", title)
}

func accordionTerminator() string {
	return "\n</details>"
}
