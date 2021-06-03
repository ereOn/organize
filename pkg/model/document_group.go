package model

import "time"

// DocumentGroupID is a unique document group identifier.
type DocumentGroupID string

// DocumentGroup is a list of documents.
type DocumentGroup struct {
	// ID is the unique document group identifier.
	//
	// This ID is generated at the creation of the document group and will never
	// change during its lifetime.
	ID DocumentGroupID `json:"id"`

	// The title of the document.
	Title string `json:"title"`

	// The creation date of the document or the date at which it was received.
	CreationDate time.Time `json:"creationDate"`

	// The list of sub-documents that this document group contains.
	//
	// A given document may belong to different groups at once.
	DocumentIDs DocumentIDs `json:"documentsIds,omitempty"`

	// The tags on the document group.
	Tags Tags `json:"tags"`
}
