package model

// DocumentTypeID is a unique identifier for a document type.
type DocumentTypeID string

// DocumentType represents an unique type of document.
type DocumentType struct {
	// The document type unique identifier.
	ID DocumentTypeID `json:"id"`

	// A friendly name for the type.
	Name string `json:"name"`

	// FileNameTemplate is a template for the file names.
	FileNameTemplate string `json:"fileNamePrefix"`
}
