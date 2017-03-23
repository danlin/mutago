package parser

import (
	"io"

	"github.com/dhowden/tag"
)

func filterTags(input map[string]interface{}) map[string]string {
	output := make(map[string]string)
	for key, value := range input {
		value, ok := value.(string)
		if ok {
			output[key] = value
		}
	}
	return output
}

func tags(reader io.ReadSeeker) (map[string]string, error) {
	tags, err := tag.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return filterTags(tags.Raw()), err
}
