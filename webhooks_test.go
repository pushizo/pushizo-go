package herosignal

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type webhookVector struct {
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
}

// Ortak vektörler: sdk/testdata/webhook-signature-vectors.json (Node, PHP, Python, Java, .NET, Ruby ile aynı).
func TestWebhookSignatureVectors(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "testdata", "webhook-signature-vectors.json"))
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Cases []webhookVector `json:"cases"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if len(file.Cases) == 0 {
		t.Fatal("vektör yok")
	}
	for _, c := range file.Cases {
		t.Run(c.Name, func(t *testing.T) {
			header := ""
			if c.Header != nil {
				header = *c.Header
			}
			opts := &WebhookVerifyOptions{Tolerance: time.Duration(c.ToleranceSeconds) * time.Second, Now: time.Unix(c.Now, 0)}
			res := VerifyWebhookSignature([]byte(c.Payload), header, c.Secrets, opts)
			if res.OK != c.Expected.OK {
				t.Fatalf("ok = %v, beklenen %v (%s)", res.OK, c.Expected.OK, res.Reason)
			}
			if c.Expected.OK && res.Timestamp != c.Expected.Timestamp {
				t.Fatalf("timestamp = %d", res.Timestamp)
			}
			if !c.Expected.OK && res.Reason != c.Expected.Reason {
				t.Fatalf("reason = %q, beklenen %q", res.Reason, c.Expected.Reason)
			}
			events, err := ConstructWebhookEvents([]byte(c.Payload), header, c.Secrets, opts)
			if c.Expected.OK {
				if err != nil {
					t.Fatal(err)
				}
				if c.Events != nil && len(events) != *c.Events {
					t.Fatalf("olay sayısı = %d, beklenen %d", len(events), *c.Events)
				}
			} else {
				var sigErr *WebhookSignatureError
				if !errors.As(err, &sigErr) || sigErr.Reason != c.Expected.Reason {
					t.Fatalf("beklenen WebhookSignatureError(%s), alınan %v", c.Expected.Reason, err)
				}
			}
		})
	}
}

func TestIdentityHashMatchesNode(t *testing.T) {
	// Node: createHmac('sha256', 'secret').update('user-1').digest('hex')
	if got := IdentityHash("secret", "user-1"); got != "1d16fd7e96e8a9681f283b8a251822ffdc71e2a5bda0c5267a044874fa21b82b" {
		t.Fatal(got)
	}
}
