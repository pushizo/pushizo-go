package herosignal

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestIdentityTokenImzaVeTalepler(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, _ := x509.MarshalPKCS8PrivateKey(key)
	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
	token, err := IdentityToken(IdentityTokenOptions{PrivateKeyPEM: pemBytes, Kid: "main", AppID: "app_1", ExternalID: "kullanici-1", TTL: 10 * time.Minute, Now: time.Unix(1_800_000_000, 0)})
	if err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("üç parça beklenirdi: %s", token)
	}
	enc := base64.RawURLEncoding
	var header map[string]string
	h, _ := enc.DecodeString(parts[0])
	_ = json.Unmarshal(h, &header)
	if header["alg"] != "ES256" || header["kid"] != "main" {
		t.Fatalf("başlık: %v", header)
	}
	var payload map[string]any
	p, _ := enc.DecodeString(parts[1])
	_ = json.Unmarshal(p, &payload)
	if payload["sub"] != "kullanici-1" || payload["aud"] != "app_1" || payload["exp"].(float64) != 1_800_000_600 {
		t.Fatalf("talepler: %v", payload)
	}
	sig, _ := enc.DecodeString(parts[2])
	if len(sig) != 64 {
		t.Fatalf("imza 64 bayt olmalı: %d", len(sig))
	}
	digest := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if !ecdsa.Verify(&key.PublicKey, digest[:], new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])) {
		t.Fatal("imza doğrulanamadı")
	}
}

func TestIdentityTokenOmurSiniri(t *testing.T) {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	der, _ := x509.MarshalECPrivateKey(key)
	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: der})
	if _, err := IdentityToken(IdentityTokenOptions{PrivateKeyPEM: pemBytes, Kid: "k", AppID: "a", ExternalID: "e", TTL: 25 * time.Hour}); err == nil {
		t.Fatal("24 saatten uzun ömür reddedilmeli")
	}
	if _, err := IdentityToken(IdentityTokenOptions{PrivateKeyPEM: pemBytes, Kid: "k", AppID: "a", ExternalID: "e"}); err != nil {
		t.Fatalf("SEC1 anahtarı kabul edilmeli: %v", err)
	}
}
