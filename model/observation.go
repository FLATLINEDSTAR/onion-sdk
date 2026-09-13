package model

// OnionObservation captures a generic observation made about a target.
// It links a target to the observed data and includes a status field.

type OnionObservation struct {
    // TargetID references the OnionTarget.ID this observation belongs to.
    TargetID string `json:"target_id"`
    // Timestamp of when the observation was recorded (RFC3339 format).
    Timestamp string `json:"timestamp"`
    // Status indicates the health or result of the observation.
    Status ObservationStatus `json:"status"`
    // Details holds optional free‑form text describing the observation.
    Details string `json:"details,omitempty"`
}

// ObservationStatus describes the outcome of an observation.
type ObservationStatus string

const (
    ObservationStatusSuccess ObservationStatus = "success"
    ObservationStatusFailure ObservationStatus = "failure"
    ObservationStatusUnknown ObservationStatus = "unknown"
)
