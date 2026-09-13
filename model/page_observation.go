package model

// PageObservation extends OnionObservation with page‑specific fields.
// It captures the HTML content snapshot and a list of extracted technologies.

type PageObservation struct {
    // Embedded generic observation fields.
    OnionObservation
    // HTML contains the raw page source (optional).
    HTML string `json:"html,omitempty"`
    // Technologies lists detected technologies on the page.
    Technologies []Technology `json:"technologies,omitempty"`
}
