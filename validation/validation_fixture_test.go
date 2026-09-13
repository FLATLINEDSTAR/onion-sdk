package validation

import (
    "encoding/json"
    "os"
    "path/filepath"
    "testing"

    "github.com/FLATLINEDSTAR/onion-sdk/model"
)

func loadFixture(t *testing.T, relPath string, v interface{}) {
    t.Helper()
    // Path relative to the repository root (one level up from the validation package)
    path := filepath.Join("..", "testdata", "fixtures", relPath)
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("failed to read fixture %s: %v", relPath, err)
    }
    if err := json.Unmarshal(data, v); err != nil {
        t.Fatalf("failed to unmarshal fixture %s: %v", relPath, err)
    }
}

func TestValidateTargetFixtures(t *testing.T) {
    var valid model.OnionTarget
    loadFixture(t, "onion_target_valid.json", &valid)
    if err := ValidateTarget(valid); err != nil {
        t.Fatalf("valid fixture should pass validation, got error: %v", err)
    }
    var invalid model.OnionTarget
    loadFixture(t, "onion_target_invalid_missing_id.json", &invalid)
    if err := ValidateTarget(invalid); err == nil {
        t.Fatalf("invalid fixture (missing ID) should fail validation, got nil")
    }
}

func TestValidateObservationFixtures(t *testing.T) {
    var valid model.OnionObservation
    loadFixture(t, "onion_observation_valid.json", &valid)
    if err := ValidateObservation(valid); err != nil {
        t.Fatalf("valid observation fixture should pass validation, got error: %v", err)
    }
    var invalid model.OnionObservation
    loadFixture(t, "onion_observation_invalid_missing_timestamp.json", &invalid)
    if err := ValidateObservation(invalid); err == nil {
        t.Fatalf("invalid observation fixture (missing timestamp) should fail validation, got nil")
    }
}

func TestValidateTechnologyFixtures(t *testing.T) {
    var valid model.Technology
    loadFixture(t, "technology_valid.json", &valid)
    if err := ValidateTechnology(valid); err != nil {
        t.Fatalf("valid technology fixture should pass validation, got error: %v", err)
    }
    var invalid model.Technology
    loadFixture(t, "technology_invalid_missing_name.json", &invalid)
    if err := ValidateTechnology(invalid); err == nil {
        t.Fatalf("invalid technology (missing name) should fail validation, got nil")
    }
    loadFixture(t, "technology_invalid_missing_category.json", &invalid)
    if err := ValidateTechnology(invalid); err == nil {
        t.Fatalf("invalid technology (missing category) should fail validation, got nil")
    }
}

func TestValidateEntityFixtures(t *testing.T) {
    var valid model.Entity
    loadFixture(t, "entity_valid.json", &valid)
    if err := ValidateEntity(valid); err != nil {
        t.Fatalf("valid entity fixture should pass validation, got error: %v", err)
    }
    var invalid model.Entity
    loadFixture(t, "entity_invalid_missing_id.json", &invalid)
    if err := ValidateEntity(invalid); err == nil {
        t.Fatalf("invalid entity (missing ID) should fail validation, got nil")
    }
}

func TestValidateEvidenceFixtures(t *testing.T) {
    var valid model.Evidence
    loadFixture(t, "evidence_valid.json", &valid)
    if err := ValidateEvidence(valid); err != nil {
        t.Fatalf("valid evidence fixture should pass validation, got error: %v", err)
    }
    var invalid model.Evidence
    loadFixture(t, "evidence_invalid_missing_observation_id.json", &invalid)
    if err := ValidateEvidence(invalid); err == nil {
        t.Fatalf("invalid evidence (missing observation_id) should fail validation, got nil")
    }
}

func TestValidateChangeEventFixtures(t *testing.T) {
    var valid model.ChangeEvent
    loadFixture(t, "change_event_valid.json", &valid)
    if err := ValidateChangeEvent(valid); err != nil {
        t.Fatalf("valid change event fixture should pass validation, got error: %v", err)
    }
    var invalid model.ChangeEvent
    loadFixture(t, "change_event_invalid_missing_changed_fields.json", &invalid)
    if err := ValidateChangeEvent(invalid); err == nil {
        t.Fatalf("invalid change event (missing changed_fields) should fail validation, got nil")
    }
}

func TestValidateSearchQueryFixtures(t *testing.T) {
    var valid model.SearchQuery
    loadFixture(t, "search_query_valid.json", &valid)
    if err := ValidateSearchQuery(valid); err != nil {
        t.Fatalf("valid search query fixture should pass validation, got error: %v", err)
    }
    // Create an invalid fixture on the fly with bad timestamp
    invalid := model.SearchQuery{Since: "not-rfc", Until: "2023-01-31T23:59:59Z"}
    if err := ValidateSearchQuery(invalid); err == nil {
        t.Fatalf("invalid search query (bad Since) should fail validation, got nil")
    }
    invalid = model.SearchQuery{Since: "2023-01-01T00:00:00Z", Until: "bad"}
    if err := ValidateSearchQuery(invalid); err == nil {
        t.Fatalf("invalid search query (bad Until) should fail validation, got nil")
    }
}
