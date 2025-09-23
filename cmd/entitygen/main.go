package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gin-starter/cmd/entitygen/gen"
)

func main() {
	table := flag.String("table", "", "Table name (e.g. users)")
	schema := flag.String("schema", "public", "Schema name (default=public)")
	migrations := flag.String("migrations", "./db/migrations", "Path to migrations folder")
	outDir := flag.String("out", "./entity", "Output dir for entities")
	flag.Parse()

	if *table == "" {
		log.Fatal("missing --table")
	}

	// find the SQL migration
	sqlFile, err := gen.FindMigration(*migrations, *schema, *table)
	if err != nil {
		log.Fatalf("migration not found: %v", err)
	}

	// parse SQL
	tbl, err := gen.ParseSQL(sqlFile)
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	// generate entity
	code, err := gen.GenerateEntity(tbl)
	if err != nil {
		log.Fatalf("generate error: %v", err)
	}

	// write entity file
	filename := fmt.Sprintf("%s.entity.go", tbl.NameLower)
	outPath := filepath.Join(*outDir, filename)
	if err := os.WriteFile(outPath, []byte(code), 0644); err != nil {
		log.Fatalf("write error: %v", err)
	}

	fmt.Printf("✅ Generated entity: %s\n", outPath)
}
