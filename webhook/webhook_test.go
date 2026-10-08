package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// vectorFile, sdk/testdata altındaki ortak webhook imza vektörleridir (Node/PHP/Python/Java/.NET/Ruby ile aynı).
type vectorFile struct {
	Cases []struct {
		Name             string   `json:"name"`
		ToleranceSeconds int64    `json:"toleranceSeconds"`
		Payload          string   `json:"payload"`
		Header           *string  `json:"header"`
		Secrets          []string `json:"secrets"`
		Now              int64    `json:"now"`
		Expected         struct {
			OK        bool   `json:"ok"`
			Timestamp int64  `json:"timestamp"`
			Reason    string `json:"reason"`
		} `json:"expected"`
		Events *int `json:"events"`
	} `json:"cases"`
}

func loadVectors(t *testing.T) vectorFile {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("..", "..", "testdata", "webhook-signature-vectors.json"))
	if err != nil {
		t.Fatalf("ortak vektörler okunamadı: %v", err)
	}
	var v vectorFile
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatal(err)
	}
	if len(v.Cases) == 0 {
		t.Fatal("vektör yok")
	}
	return v
}

func TestSharedVectors(t *testing.T) {
	for _, c := range loadVectors(t).Cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			header := ""
			if c.Header != nil {
				header = *c.Header
			}
			opts := &VerifyOptions{Tolerance: time.Duration(c.ToleranceSeconds) * time.Second, Now: time.Unix(c.Now, 0)}
			res := Verify([]byte(c.Payload), header, c.Secrets, opts)
			if res.OK != c.Expected.OK {
				t.Fatalf("ok: %v != %v (%s)", res.OK, c.Expected.OK, res.Reason)
			}
			if c.Expected.OK && res.Timestamp != c.Expected.Timestamp {
				t.Fatalf("timestamp: %d", res.Timestamp)
			}
			if !c.Expected.OK && string(res.Reason) != c.Expected.Reason {
				t.Fatalf("reason: %s != %s", res.Reason, c.Expected.Reason)
			}
			events, err := ConstructEvents([]byte(c.Payload), header, c.Secrets, opts)
			if c.Expected.OK {
				if err != nil {
					t.Fatal(err)
				}
				if c.Events != nil && len(events) != *c.Events {
					t.Fatalf("olay sayısı: %d != %d", len(events), *c.Events)
				}
				for _, e := range events {
					if e.Type == "" || e.ID == "" {
						t.Fatalf("olay ayrıştırılmadı: %+v", e)
					}
				}
			} else {
				var se *SignatureError
				if !errors.As(err, &se) || string(se.Reason) != c.Expected.Reason {
					t.Fatalf("SignatureError bekleniyordu: %v", err)
				}
			}
		})
	}
}

func TestMalformedBody(t *testing.T) {
	body := []byte("not json")
	mac := hmac.New(sha256.New, []byte("s"))
	mac.Write([]byte("1800000000."))
	mac.Write(body)
	header := "t=1800000000,v1=" + hex.EncodeToString(mac.Sum(nil))
	_, err := ConstructEvents(body, header, []string{"s"}, &VerifyOptions{Now: time.Unix(1800000000, 0)})
	var se *SignatureError
	if !errors.As(err, &se) || se.Reason != MalformedBody {
		t.Fatalf("malformed_body bekleniyordu: %v", err)
	}
}

func TestDefaultToleranceUsesCurrentTime(t *testing.T) {
	ts := time.Now().Unix()
	body := []byte(`{"id":"whe_1","type":"x"}`)
	sign := func(t int64) string {
		m := hmac.New(sha256.New, []byte("s"))
		m.Write([]byte(jsonInt(t) + "."))
		m.Write(body)
		return "t=" + jsonInt(t) + ",v1=" + hex.EncodeToString(m.Sum(nil))
	}
	if res := Verify(body, sign(ts), []string{"s"}, nil); !res.OK {
		t.Fatalf("güncel imza geçerli olmalı: %s", res.Reason)
	}
	if res := Verify(body, sign(ts-int64(DefaultTolerance/time.Second)-5), []string{"s"}, nil); res.Reason != TimestampOutOfTolerance {
		t.Fatalf("eski imza reddedilmeli: %+v", res)
	}
}

func jsonInt(n int64) string {
	b, _ := json.Marshal(n)
	return string(b)
}
