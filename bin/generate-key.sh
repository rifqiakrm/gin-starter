#!/bin/bash
set -e

ENV_FILE=".env"

echo "Generating Ed25519 keys..."

cat << 'EOF' > /tmp/gen.go
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JWT_PUBLIC_KEY=\"%s\"\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("JWT_PRIVATE_KEY=\"%s\"\n", base64.StdEncoding.EncodeToString(priv))
}
EOF

OUTPUT=$(go run /tmp/gen.go)
rm /tmp/gen.go

echo "$OUTPUT"
echo "" >> "$ENV_FILE"
echo "# Generated on $(date)" >> "$ENV_FILE"
echo "$OUTPUT" >> "$ENV_FILE"

echo "✅ Keys appended to $ENV_FILE"
