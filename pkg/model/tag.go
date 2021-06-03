package model

// Tag represent a tag on a document or document group.
type Tag string

// Tags represent a list of tags.
type Tags []Tag

// Contains checks whether a list of tags contains the specified tag.
func (t Tags) Contains(tag Tag) bool {
	for _, v := range t {
		if v == tag {
			return true
		}
	}

	return false
}
