package model

// OnionTarget defines a target for observation within the Onion ecosystem.
// It represents an entity that can be observed, such as a website URL or an IP address.
//
// Fields are intentionally minimal to keep the contract stable across versions.
type OnionTarget struct {
    // ID is a globally unique identifier for the target.
    ID string `json:"id"`
    // URL is the canonical address of the target (e.g., https://example.com).
    URL string `json:"url"`
    // Description provides a short human‑readable explanation of the target.
    Description string `json:"description,omitempty"`
}
