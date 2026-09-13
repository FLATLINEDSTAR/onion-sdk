package validation

import (
    "testing"
    "github.com/FLATLINEDSTAR/onion-sdk/model"
)

func TestValidateTarget(t *testing.T) {
    valid := model.OnionTarget{ID: "t1", URL: "https://example.com"}
    if err := ValidateTarget(valid); err != nil {
        t.Fatalf("expected valid target, got error: %v", err)
    }
    // missing ID
    missingID := model.OnionTarget{URL: "https://example.com"}
    if err := ValidateTarget(missingID); err == nil {
        t.Fatalf("expected error for missing ID, got nil")
    }
    // invalid URL
    badURL := model.OnionTarget{ID: "t2", URL: "not-a-url"}
    if err := ValidateTarget(badURL); err == nil {
        t.Fatalf("expected error for invalid URL, got nil")
    }
}

func TestValidateObservation(t *testing.T) {
    valid := model.OnionObservation{TargetID: "t1", Timestamp: "2023-01-01T00:00:00Z", Status: model.ObservationStatusSuccess}
    if err := ValidateObservation(valid); err != nil {
        t.Fatalf("expected valid observation, got error: %v", err)
    }
    // missing Timestamp
    missingTS := model.OnionObservation{TargetID: "t1", Status: model.ObservationStatusSuccess}
    if err := ValidateObservation(missingTS); err == nil {
        t.Fatalf("expected error for missing Timestamp, got nil")
    }
    // bad timestamp format
    badTS := model.OnionObservation{TargetID: "t1", Timestamp: "not-rfc3339", Status: model.ObservationStatusSuccess}
    if err := ValidateObservation(badTS); err == nil {
        t.Fatalf("expected error for malformed Timestamp, got nil")
    }
}

func TestValidateTechnology(t *testing.T) {
    valid := model.Technology{Name: "React"}
    if err := ValidateTechnology(valid); err != nil {
        t.Fatalf("expected valid technology, got error: %v", err)
    }
    missingName := model.Technology{}
    if err := ValidateTechnology(missingName); err == nil {
        t.Fatalf("expected error for missing Name, got nil")
    }
}

func TestValidateEntity(t *testing.T) {
    valid := model.Entity{ID: "e1", Name: "service", Type: model.EntityTypeService, Confidence: 0.9}
    if err := ValidateEntity(valid); err != nil {
        t.Fatalf("expected valid entity, got error: %v", err)
    }
    // invalid confidence
    badConf := model.Entity{ID: "e2", Name: "svc", Type: model.EntityTypeService, Confidence: 1.5}
    if err := ValidateEntity(badConf); err == nil {
        t.Fatalf("expected error for confidence out of range, got nil")
    }
    // missing fields
    missing := model.Entity{Confidence: 0.5}
    if err := ValidateEntity(missing); err == nil {
        t.Fatalf("expected error for missing required fields, got nil")
    }
}

func TestValidateEvidence(t *testing.T) {
    valid := model.Evidence{ObservationID: "obs1", Source: "header", Value: "X-Test", Confidence: 0.8}
    if err := ValidateEvidence(valid); err != nil {
        t.Fatalf("expected valid evidence, got error: %v", err)
    }
    // missing ObservationID
    missingID := model.Evidence{Source: "header", Value: "X-Test", Confidence: 0.8}
    if err := ValidateEvidence(missingID); err == nil {
        t.Fatalf("expected error for missing ObservationID, got nil")
    }
    // confidence out of range
    badConf := model.Evidence{ObservationID: "obs2", Source: "meta", Value: "val", Confidence: -0.1}
    if err := ValidateEvidence(badConf); err == nil {
        t.Fatalf("expected error for confidence out of range, got nil")
    }
}

func TestValidateChangeEvent(t *testing.T) {
    ev := model.Evidence{ObservationID: "obs1", Source: "header", Value: "v", Confidence: 0.9}
    valid := model.ChangeEvent{TargetID: "t1", BeforeObservationID: "obs1", AfterObservationID: "obs2", ChangedFields: []string{"status"}, Timestamp: "2023-01-02T00:00:00Z", Evidence: []model.Evidence{ev}}
    if err := ValidateChangeEvent(valid); err != nil {
        t.Fatalf("expected valid change event, got error: %v", err)
    }
    // missing ChangedFields
    missingFields := valid
    missingFields.ChangedFields = nil
    if err := ValidateChangeEvent(missingFields); err == nil {
        t.Fatalf("expected error for empty ChangedFields, got nil")
    }
    // bad timestamp
    badTS := valid
    badTS.Timestamp = "not-rfc"
    if err := ValidateChangeEvent(badTS); err == nil {
        t.Fatalf("expected error for malformed Timestamp, got nil")
    }
}

func TestValidateSearchQuery(t *testing.T) {
    valid := model.SearchQuery{Since: "2023-01-01T00:00:00Z", Until: "2023-01-31T23:59:59Z"}
    if err := ValidateSearchQuery(valid); err != nil {
        t.Fatalf("expected valid search query, got error: %v", err)
    }
    // bad Since format
    badSince := model.SearchQuery{Since: "bad"}
    if err := ValidateSearchQuery(badSince); err == nil {
        t.Fatalf("expected error for bad Since, got nil")
    }
    // bad Until format
    badUntil := model.SearchQuery{Until: "bad"}
    if err := ValidateSearchQuery(badUntil); err == nil {
        t.Fatalf("expected error for bad Until, got nil")
    }
}
