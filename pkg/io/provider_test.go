package io

import (
	"fmt"
	"testing"

	"github.com/ereOn/organize/pkg/model"
	"github.com/stretchr/testify/assert"
)

type testProvider struct {
	DocumentDataMapping map[model.DocumentID][]byte

	name string
}

func (p *testProvider) Name() string {
	return p.name
}

func (p *testProvider) ReadDocumentData(document *model.Document) error {
	if p.name != document.ProviderName {
		return fmt.Errorf("mismatching provider name: %s != %s", document.ProviderName, p.name)
	}

	if data, ok := p.DocumentDataMapping[document.ID]; ok {
		document.Data = data
		return nil
	}

	return fmt.Errorf("no document data found in the mapping for document %s", document.ID)
}

func newTestProvider(name string) *testProvider {
	return &testProvider{
		name: name,
	}
}

func TestProviders(t *testing.T) {
	providerA := newTestProvider("a")
	providerB := newTestProvider("b")
	providers := Providers{providerA, providerB}

	t.Run("exists", func(t *testing.T) {
		provider, err := providers.Get("a")

		assert.NoError(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("does-not-exist", func(t *testing.T) {
		provider, err := providers.Get("z")

		assert.Error(t, err)
		assert.Nil(t, provider)
	})
}
