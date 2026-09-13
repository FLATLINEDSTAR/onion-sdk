package model

// Entity represents a discovered entity within a target observation.
// It includes a unique ID, a human‑readable name, and a type indicating the kind of entity.

type Entity struct {
    // ID is the unique identifier for the entity.
    ID string `json:"id"`
    // Name is a short descriptive name.
    Name string `json:"name"`
    // Type classifies the entity (e.g., "service", "host", "application").
    Type EntityType `json:"type"`
    // Confidence indicates how certain the detection is (0‑1).
    Confidence float64 `json:"confidence,omitempty"`
}

// EntityType defines allowed classifications for an Entity.
type EntityType string

const (
    EntityTypeService     EntityType = "service"
    EntityTypeHost        EntityType = "host"
    EntityTypeApplication EntityType = "application"
    EntityTypeOther       EntityType = "other"
)
