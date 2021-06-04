package io

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ereOn/organize/pkg/model"
)

// DocumentReader can read documents.
type DocumentReader struct {
	// A flag that indicates whether files should be read with their data.
	WithData bool

	// A list of supported providers.
	//
	// If nil, the default providers are used.
	Providers Providers
}

// ReadDocumentFile reads a document from a file-name.
func (dr DocumentReader) ReadDocumentFile(fileName string) (document *model.Document, err error) {
	r, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("failed to open document file: %w", err)
	}

	defer r.Close()

	return dr.ReadDocument(r)
}

// ReadDocument reads a document from an io.Reader.
func (dr DocumentReader) ReadDocument(r io.Reader) (document *model.Document, err error) {
	document = &model.Document{}

	if err = json.NewDecoder(r).Decode(document); err != nil {
		return nil, fmt.Errorf("failed to decode document metadata: %w", err)
	}

	if dr.WithData {
		err = dr.ReadDocumentData(document)
	}

	return
}

// ReadDocumentData completes a document and read its data from its associated
// provider.
//
// If the document already has its data, the operation does nothing and
// succeeds.
func (dr DocumentReader) ReadDocumentData(document *model.Document) error {
	if document == nil {
		panic("document cannot be nil")
	}

	if document.ProviderName == "" {
		return fmt.Errorf("document %s has no defined provider", document.ID)
	}

	provider, err := dr.Providers.Get(document.ProviderName)

	if err != nil {
		return fmt.Errorf("failed to read document %s data: %w", document.ID, err)
	}

	return provider.ReadDocumentData(document)
}
