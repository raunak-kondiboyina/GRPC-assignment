package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsvalid(t *testing.T) {
	actualresult := Isvalid("Raunak")
	//assert.NotNil(t, nil, "bool should not be nil")
	assert.NotNil(t, actualresult, "bool should not be nil")
	assert.Equal(t, true, actualresult, "expecting true")

	actualresult1 := Isvalid("987s")
	assert.NotNil(t, actualresult1, "bool should not be nil")
	assert.Equal(t, true, actualresult1, "expecting true")

	actualresult2 := Isvalid("^%$")
	assert.NotNil(t, actualresult2, "bool should not be nil")
	assert.Equal(t, false, actualresult2, "expecting true")

}
