package io

import (
	"fmt"

	"github.com/ereOn/organize/pkg/model"
)

// Provider implements data fetching for documents.
type Provider interface {
	// Returns the name of the provider.
	Name() string

	// Read a document's data.
	//
	// If the specified document already has its data, nothing is done and the
	// operation succeeds.
	ReadDocumentData(document *model.Document) error
}

// Providers is a list of providers.
type Providers []Provider

// Get a provider from its name
func (p Providers) Get(providerName string) (Provider, error) {
	for _, provider := range p {
		if provider.Name() == providerName {
			return provider, nil
		}
	}

	return nil, fmt.Errorf("could not find a provider named `%s`", providerName)
}
