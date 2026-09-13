package model

// Evidence captures proof supporting an entity or technology detection.
// It references the source observation and includes optional raw data.

type Evidence struct {
    // ObservationID links to the OnionObservation that produced this evidence.
    ObservationID string `json:"observation_id"`
    // Source describes where the evidence originated (e.g., "header", "script", "meta").
    Source string `json:"source"`
    // Value holds the raw evidence value (e.g., header string, script snippet).
    Value string `json:"value"`
    // Confidence indicates certainty of the evidence (0‑1).
    Confidence float64 `json:"confidence,omitempty"`
}
