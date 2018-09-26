package group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Group(t *testing.T) {
	input := []Score{
		Score{1, "Tom", "One", 60},
		Score{2, "Jerry", "One", 70},
		Score{3, "Woof", "Two", 90},
	}

	groupedScores := GroupByClass(input)
	assert.NotNil(t, groupedScores)
	assert.Equal(t, 2, len(groupedScores))
	classOneScores := groupedScores["One"]
	assert.Equal(t, 2, len(classOneScores))
	classTwoScores := groupedScores["Two"]
	assert.Equal(t, 1, len(classTwoScores))
}
