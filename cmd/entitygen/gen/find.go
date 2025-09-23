package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindMigration finds the migration file for given schema + table
func FindMigration(root, schema, table string) (string, error) {
	var found string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ensure it's inside the schema folder
		if strings.Contains(path, string(filepath.Separator)+schema+string(filepath.Separator)) &&
			strings.Contains(path, "_"+table) &&
			strings.HasSuffix(path, ".up.sql") {
			found = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if found == "" {
		return "", fmt.Errorf("no migration for %s.%s", schema, table)
	}
	return found, nil
}
