package validation

import (
    "errors"
    "net/url"
    "time"

    "github.com/FLATLINEDSTAR/onion-sdk/model"
)

// ValidateTarget checks that an OnionTarget has required fields.
func ValidateTarget(t model.OnionTarget) error {
    if t.ID == "" {
        return errors.New("OnionTarget.ID must not be empty")
    }
    if t.URL == "" {
        return errors.New("OnionTarget.URL must not be empty")
    }
    if _, err := url.ParseRequestURI(t.URL); err != nil {
        return errors.New("OnionTarget.URL is not a valid URI: " + err.Error())
    }
    // Description is optional.
    return nil
}

// ValidateObservation checks required fields of OnionObservation.
func ValidateObservation(o model.OnionObservation) error {
    if o.TargetID == "" {
        return errors.New("OnionObservation.TargetID must not be empty")
    }
    if o.Timestamp == "" {
        return errors.New("OnionObservation.Timestamp must not be empty")
    }
    if _, err := time.Parse(time.RFC3339, o.Timestamp); err != nil {
        return errors.New("OnionObservation.Timestamp must be RFC3339 format: " + err.Error())
    }
    // Status should be one of the defined constants; we accept any string here.
    if o.Status == "" {
        return errors.New("OnionObservation.Status must not be empty")
    }
    return nil
}

// ValidatePageObservation validates PageObservation, including embedded OnionObservation.
func ValidatePageObservation(p model.PageObservation) error {
    // Validate the embedded observation first.
    if err := ValidateObservation(p.OnionObservation); err != nil {
        return err
    }
    // HTML is optional.
    // Validate each Technology in the list.
    for _, tech := range p.Technologies {
        if err := ValidateTechnology(tech); err != nil {
            return err
        }
    }
    return nil
}

// ValidateTechnology checks required fields of Technology.
func ValidateTechnology(t model.Technology) error {
    if t.Name == "" {
        return errors.New("Technology.Name must not be empty")
    }
    // Version optional.
    // Category optional; could be empty string.
    return nil
}

// ValidateEntity checks required fields of Entity.
func ValidateEntity(e model.Entity) error {
    if e.ID == "" {
        return errors.New("Entity.ID must not be empty")
    }
    if e.Name == "" {
        return errors.New("Entity.Name must not be empty")
    }
    if e.Type == "" {
        return errors.New("Entity.Type must not be empty")
    }
    if e.Confidence < 0.0 || e.Confidence > 1.0 {
        return errors.New("Entity.Confidence must be between 0 and 1")
    }
    return nil
}

// ValidateEvidence checks required fields of Evidence.
func ValidateEvidence(ev model.Evidence) error {
    if ev.ObservationID == "" {
        return errors.New("Evidence.ObservationID must not be empty")
    }
    if ev.Source == "" {
        return errors.New("Evidence.Source must not be empty")
    }
    if ev.Value == "" {
        return errors.New("Evidence.Value must not be empty")
    }
    if ev.Confidence < 0.0 || ev.Confidence > 1.0 {
        return errors.New("Evidence.Confidence must be between 0 and 1")
    }
    return nil
}

// ValidateChangeEvent checks required fields of ChangeEvent.
func ValidateChangeEvent(c model.ChangeEvent) error {
    if c.TargetID == "" {
        return errors.New("ChangeEvent.TargetID must not be empty")
    }
    if c.BeforeObservationID == "" || c.AfterObservationID == "" {
        return errors.New("ChangeEvent must have both BeforeObservationID and AfterObservationID")
    }
    if len(c.ChangedFields) == 0 {
        return errors.New("ChangeEvent.ChangedFields must contain at least one field")
    }
    if c.Timestamp == "" {
        return errors.New("ChangeEvent.Timestamp must not be empty")
    }
    if _, err := time.Parse(time.RFC3339, c.Timestamp); err != nil {
        return errors.New("ChangeEvent.Timestamp must be RFC3339 format: " + err.Error())
    }
    for _, ev := range c.Evidence {
        if err := ValidateEvidence(ev); err != nil {
            return err
        }
    }
    return nil
}

// ValidateSearchQuery checks optional fields of SearchQuery.
func ValidateSearchQuery(q model.SearchQuery) error {
    // Validate timestamps if provided.
    if q.Since != "" {
        if _, err := time.Parse(time.RFC3339, q.Since); err != nil {
            return errors.New("SearchQuery.Since must be RFC3339 format: " + err.Error())
        }
    }
    if q.Until != "" {
        if _, err := time.Parse(time.RFC3339, q.Until); err != nil {
            return errors.New("SearchQuery.Until must be RFC3339 format: " + err.Error())
        }
    }
    // Status slice can be empty; no further validation needed.
    return nil
}
