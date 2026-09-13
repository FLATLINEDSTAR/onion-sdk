package model

// SearchQuery defines the query parameters for searching observations.
// It includes optional filters such as target IDs, technology names, and time ranges.

type SearchQuery struct {
    // TargetIDs filters observations to specific targets (optional).
    TargetIDs []string `json:"target_ids,omitempty"`
    // TechnologyNames filters observations containing given technologies (optional).
    TechnologyNames []string `json:"technology_names,omitempty"`
    // Since filters observations after this RFC3339 timestamp (optional).
    Since string `json:"since,omitempty"`
    // Until filters observations before this RFC3339 timestamp (optional).
    Until string `json:"until,omitempty"`
    // Status filters by ObservationStatus (optional).
    Status []ObservationStatus `json:"status,omitempty"`
}
