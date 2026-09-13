package model

// ChangeEvent represents a change detected between two observations of the same target.
// It records what changed, when, and optional evidence supporting the change.

type ChangeEvent struct {
    // TargetID references the OnionTarget affected.
    TargetID string `json:"target_id"`
    // BeforeObservationID references the earlier observation.
    BeforeObservationID string `json:"before_observation_id"`
    // AfterObservationID references the later observation.
    AfterObservationID string `json:"after_observation_id"`
    // ChangedFields lists the names of fields that changed (e.g., "technologies", "status").
    ChangedFields []string `json:"changed_fields"`
    // Evidence provides supporting evidence for the change (optional).
    Evidence []Evidence `json:"evidence,omitempty"`
    // Timestamp of when the change was recorded.
    Timestamp string `json:"timestamp"`
}
