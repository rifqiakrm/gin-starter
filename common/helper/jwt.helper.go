//revive:disable:var-naming
package helper

import (
	"crypto/ed25519"
	"crypto/md5" // #nosec
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gin-starter/common/constant"
	"gin-starter/config"
)

const (
	daysInYear = 365
	dayInHour  = 24
)

// TokenClaims define available data in JWT Token
type TokenClaims struct {
	ExpiresAt int64  `json:"exp,omitempty"`
	ID        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   int64  `json:"sub,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	jwt.Claims
}

// getEd25519Public loads Ed25519 public key from base64 (env or config)
func getEd25519Public(b64 string) (ed25519.PublicKey, error) {
	if b64 == "" {
		return nil, fmt.Errorf("JWT public key not configured")
	}
	bytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 public key: %w", err)
	}
	return ed25519.PublicKey(bytes), nil
}

// getEd25519Private loads Ed25519 private key from base64 (env or config)
func getEd25519Private(b64 string) (ed25519.PrivateKey, error) {
	if b64 == "" {
		return nil, fmt.Errorf("JWT private key not configured")
	}
	bytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 private key: %w", err)
	}
	return ed25519.PrivateKey(bytes), nil
}

// JWTDecode decodes JWT to token claims using Ed25519 public key
func JWTDecode(cfg config.Config, t string) (*TokenClaims, error) {
	publicKey, err := getEd25519Public(cfg.JWTConfig.Public)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(t, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// JWTEncode encodes token claims to JWT using Ed25519 private key
func JWTEncode(cfg config.Config, id int64, iss string) (string, error) {
	privateKey, err := getEd25519Private(cfg.JWTConfig.Private)
	if err != nil {
		return "", err
	}

	// generate jti
	hashQuery := md5.New() // #nosec
	hashQuery.Write([]byte(fmt.Sprintf("secret123:%v", time.Now().Add(time.Hour*dayInHour*daysInYear).Unix())))
	jti := hex.EncodeToString(hashQuery.Sum(nil))

	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * constant.TwentyFourHour * constant.DaysInOneYear).Unix(),
		"jti": jti,
		"sub": id,
		"iss": iss,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	signed, signErr := token.SignedString(privateKey)
	if signErr != nil {
		return "", fmt.Errorf("failed to sign token: %w", signErr)
	}
	return signed, nil
}
