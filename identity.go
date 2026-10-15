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
	"errors"
	"time"
)

// IdentityTokenMaxLifetime, SDK JWT'sinin en uzun ömrüdür (ADR-0023 §9).
const IdentityTokenMaxLifetime = 24 * time.Hour

// IdentityTokenOptions, IdentityToken parametreleridir.
type IdentityTokenOptions struct {
	// PrivateKeyPEM: ES256 (P-256) özel anahtarı, PKCS#8 veya SEC1 PEM.
	PrivateKeyPEM []byte
	// Kid: panelde/API'de kayıtlı genel anahtarın kimliği.
	Kid        string
	AppID      string
	ExternalID string
	// TTL: belirteç ömrü (varsayılan 1 saat, en çok 24 saat).
	TTL time.Duration
	// Now: testler için sabit zaman.
	Now time.Time
}

// IdentityToken, istemci SDK'sına verilecek ES256 JWT'yi üretir: sub = external_id, aud = app_id. Uygulamanız
// belirteci kendi sunucusundan alır ve SDK'nın login(externalId, jwt) çağrısına verir.
func IdentityToken(opts IdentityTokenOptions) (string, error) {
	ttl := opts.TTL
	if ttl == 0 {
		ttl = time.Hour
	}
	if ttl < time.Second || ttl > IdentityTokenMaxLifetime {
		return "", errors.New("herosignal: TTL 1 sn … 24 saat aralığında olmalı")
	}
	key, err := parseECPrivateKey(opts.PrivateKeyPEM)
	if err != nil {
		return "", err
	}
	now := opts.Now
	if now.IsZero() {
		now = time.Now()
	}
	header, _ := json.Marshal(map[string]string{"alg": "ES256", "typ": "JWT", "kid": opts.Kid})
	payload, _ := json.Marshal(map[string]any{"sub": opts.ExternalID, "aud": opts.AppID, "iat": now.Unix(), "exp": now.Add(ttl).Unix()})
	enc := base64.RawURLEncoding
	signingInput := enc.EncodeToString(header) + "." + enc.EncodeToString(payload)
	digest := sha256.Sum256([]byte(signingInput))
	r, s, err := ecdsa.Sign(rand.Reader, key, digest[:])
	if err != nil {
		return "", err
	}
	// JWS ES256 imzası DER değil r||s (her biri 32 bayt).
	sig := make([]byte, 64)
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:])
	return signingInput + "." + enc.EncodeToString(sig), nil
}

func parseECPrivateKey(pemBytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("herosignal: PEM özel anahtarı okunamadı")
	}
	if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		ec, ok := k.(*ecdsa.PrivateKey)
		if !ok || ec.Curve != elliptic.P256() {
			return nil, errors.New("herosignal: ES256 için P-256 EC özel anahtarı gerekli")
		}
		return ec, nil
	}
	ec, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil || ec.Curve != elliptic.P256() {
		return nil, errors.New("herosignal: ES256 için P-256 EC özel anahtarı gerekli")
	}
	return ec, nil
}
