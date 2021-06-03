package model

import (
	"crypto/sha256"
	"time"
)

// DocumentID is a unique document identifier.
type DocumentID string

// DocumentIDs is a list of document identifiers.
type DocumentIDs []DocumentID

// Document represents a stored document.
type Document struct {
	// ID is the unique document identifier.
	//
	// This ID is generated at the creation of the document and will never
	// change during its lifetime.
	ID DocumentID `json:"id"`

	// The type identifier for this document.
	TypeID DocumentTypeID `json:"typeId"`

	// A description for the document.
	Description string `json:"title"`

	// The creation date of the document or the date at which it was received.
	CreationDate time.Time `json:"creationDate"`

	// MediaType is the RFC 2045/RFC 2616 media type of the document.
	MediaType string `json:"mediaType,omitempty"`

	// The size of the document.
	Size int64 `json:"size"`

	// The document data.
	Data []byte `json:"data,omitempty"`

	// The sha256 hash of the document.
	Hash [sha256.Size]byte `json:"hash"`

	// The tags on the document.
	Tags Tags `json:"tags"`

	// The provider info.
	ProviderInfo []byte `json:"providerInfo"`
}
