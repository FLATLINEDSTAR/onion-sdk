package model

import (
    "encoding/json"
    "testing"
)

func TestOnionTargetJSONRoundTrip(t *testing.T) {
    original := OnionTarget{ID: "t1", URL: "https://example.com", Description: "test target"}
    data, err := json.Marshal(original)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    var decoded OnionTarget
    if err := json.Unmarshal(data, &decoded); err != nil {
        t.Fatalf("unmarshal failed: %v", err)
    }
    if decoded != original {
        t.Fatalf("decoded %+v does not match original %+v", decoded, original)
    }
}
