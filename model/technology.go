package model

// Technology represents a detected technology on a target page.
// It includes a name and optional version information.

type Technology struct {
    // Name of the technology (e.g., "React", "nginx").
    Name string `json:"name"`
    // Version is optional and may be empty if unknown.
    Version string `json:"version,omitempty"`
    // Category groups technologies (e.g., "framework", "webserver").
    Category TechnologyCategory `json:"category,omitempty"`
}

// TechnologyCategory defines a limited set of technology categories.
type TechnologyCategory string

const (
    CategoryFramework TechnologyCategory = "framework"
    CategoryWebServer TechnologyCategory = "webserver"
    CategoryDatabase  TechnologyCategory = "database"
    CategoryLanguage  TechnologyCategory = "language"
    CategoryOther     TechnologyCategory = "other"
)
