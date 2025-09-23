package gen

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ParseSQL parses a CREATE TABLE SQL file into a Table struct.
// It is robust to multi-line SQL, BEGIN/COMMIT wrappers, quoted identifiers,
// and types containing parentheses (e.g. VARCHAR(128), numeric(10,2)).
func ParseSQL(path string) (*Table, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(data)

	// find "create table" (case-insensitive)
	lower := strings.ToLower(s)
	ctIdx := strings.Index(lower, "create table")
	if ctIdx == -1 {
		return nil, fmt.Errorf("could not find CREATE TABLE in file")
	}

	// Find '(' that starts the column block after the CREATE TABLE
	openIdx := strings.Index(s[ctIdx:], "(")
	if openIdx == -1 {
		return nil, fmt.Errorf("could not find opening '(' after CREATE TABLE")
	}
	openIdx += ctIdx // make absolute index

	// Find matching closing ')' by scanning and counting parentheses
	level := 0
	closeIdx := -1
	for i := openIdx; i < len(s); i++ {
		ch := s[i]
		if ch == '(' {
			level++
		} else if ch == ')' {
			level--
			if level == 0 {
				closeIdx = i
				break
			}
		}
	}
	if closeIdx == -1 {
		return nil, fmt.Errorf("could not find matching ')' for CREATE TABLE")
	}

	// header contains everything between "CREATE TABLE" and '('
	header := s[ctIdx:openIdx]
	body := s[openIdx+1 : closeIdx] // column definitions and constraints

	// Extract name token from header:
	// Remove "CREATE TABLE" and optional "IF NOT EXISTS"
	headerLower := strings.ToLower(header)
	after := header
	if idx := strings.Index(headerLower, "create table"); idx != -1 {
		after = header[idx+len("create table"):]
	}
	after = strings.TrimSpace(after)
	if strings.HasPrefix(strings.ToLower(after), "if not exists") {
		after = strings.TrimSpace(after[len("if not exists"):])
	}

	// header name should now be something like:
	// - auth.users
	// - "auth"."users"
	// - users
	// It may contain whitespace/newlines; take the first token (until whitespace)
	after = strings.TrimSpace(strings.ReplaceAll(after, "\n", " "))
	tokens := strings.Fields(after)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("could not parse table name from header: %q", header)
	}
	nameToken := tokens[0]

	// handle quoted identifiers and schema.table
	var schemaName, tableName string
	if strings.Contains(nameToken, ".") {
		parts := strings.SplitN(nameToken, ".", 2)
		schemaName = trimQuotes(parts[0])
		tableName = trimQuotes(parts[1])
	} else {
		tableName = trimQuotes(nameToken)
		// schema is not specified in header; we'll leave schema empty
	}

	// split body into column/constraint lines safely (ignore commas inside parentheses)
	colLines := splitColumns(body)

	var cols []Column
	for _, raw := range colLines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)

		// skip constraint lines
		if strings.HasPrefix(upper, "PRIMARY KEY") ||
			strings.HasPrefix(upper, "CONSTRAINT") ||
			strings.HasPrefix(upper, "UNIQUE") ||
			strings.HasPrefix(upper, "FOREIGN KEY") ||
			strings.HasPrefix(upper, "CHECK") {
			continue
		}

		// parse column name
		colName, rest, err := splitNameAndRest(line)
		if err != nil {
			// if can't parse, skip (safe)
			continue
		}

		// determine type: take everything until one of stop tokens (NOT NULL, DEFAULT, PRIMARY, UNIQUE, REFERENCES, CHECK, CONSTRAINT)
		stopTokens := []string{"not null", "null", "default", "primary", "unique", "references", "check", "constraint"}
		idx := indexOfAny(strings.ToLower(rest), stopTokens)
		var typePart string
		if idx >= 0 {
			typePart = strings.TrimSpace(rest[:idx])
		} else {
			typePart = strings.TrimSpace(rest)
		}

		nullability := true
		if strings.Contains(strings.ToUpper(rest), "NOT NULL") {
			nullability = false
		}
		isPrimary := strings.Contains(strings.ToUpper(rest), "PRIMARY KEY")

		cols = append(cols, Column{
			Name:       colName,
			Type:       strings.ToUpper(typePart),
			Nullable:   nullability,
			PrimaryKey: isPrimary,
		})
	}

	// Table name conversions for generator (CamelCase, lower)
	entityName := toCamel(naiveSingular(tableName))

	return &Table{
		Schema:    schemaName,
		Name:      tableName,
		NameUpper: entityName,
		NameLower: strings.ToLower(tableName),
		Columns:   cols,
	}, nil
}

// trimQuotes removes surrounding double quotes or backticks (if any)
func trimQuotes(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && ((s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '`' && s[len(s)-1] == '`')) {
		return s[1 : len(s)-1]
	}
	return s
}

// splitColumns splits the column/constraint block by commas that are at parentheses level 0.
// e.g. handles numeric(10,2) or CHECK (...), etc.
func splitColumns(body string) []string {
	var out []string
	var sb strings.Builder
	level := 0
	for i := 0; i < len(body); i++ {
		c := body[i]
		if c == '(' {
			level++
			sb.WriteByte(c)
			continue
		}
		if c == ')' {
			level--
			sb.WriteByte(c)
			continue
		}
		if c == ',' && level == 0 {
			out = append(out, sb.String())
			sb.Reset()
			continue
		}
		sb.WriteByte(c)
	}
	if s := strings.TrimSpace(sb.String()); s != "" {
		out = append(out, s)
	}
	return out
}

// splitNameAndRest returns (name, rest-of-line-after-name)
// name may be quoted: "user_id" or unquoted: user_id
func splitNameAndRest(line string) (string, string, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return "", "", fmt.Errorf("empty line")
	}
	if line[0] == '"' || line[0] == '`' {
		quote := line[0]
		// find closing quote
		j := -1
		for i := 1; i < len(line); i++ {
			if line[i] == quote {
				j = i
				break
			}
		}
		if j == -1 {
			return "", "", fmt.Errorf("unclosed quoted identifier")
		}
		name := line[1:j]
		rest := strings.TrimSpace(line[j+1:])
		return name, rest, nil
	}

	// unquoted: first token is name
	i := 0
	for i < len(line) && !unicode.IsSpace(rune(line[i])) {
		i++
	}
	if i == 0 {
		return "", "", fmt.Errorf("cannot parse column name")
	}
	name := line[:i]
	rest := strings.TrimSpace(line[i:])
	return name, rest, nil
}

// indexOfAny returns the lowest index >=0 of any word in words inside s (case-insensitive), or -1
func indexOfAny(s string, words []string) int {
	low := -1
	ls := strings.ToLower(s)
	for _, w := range words {
		if idx := strings.Index(ls, w); idx >= 0 {
			if low == -1 || idx < low {
				low = idx
			}
		}
	}
	return low
}

// toCamel converts snake_case to CamelCase (basic)
func toCamel(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '_' || r == '-' || r == ' ' })
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToUpper(parts[i][:1]) + strings.ToLower(parts[i][1:])
	}
	return strings.Join(parts, "")
}

// naiveSingular removes trailing 's' for simple plural -> singular conversion
// (this is naive but works for many table names like "users" -> "auth")
func naiveSingular(s string) string {
	if len(s) > 1 && strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") {
		return s[:len(s)-1]
	}
	return s
}
