package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Error_200(t *testing.T) {
	r := Response{
		Body: nil,
		Code: 200,
	}

	err := r.Error()

	// then: no error returned
	assert.NoError(t, err)
}

func Test_Error_Body_Nil(t *testing.T) {
	r := Response{
		Body: nil}

	err := r.Error()

	// then: error returned
	assert.Error(t, err)
}

func Test_Error_UnmarshallingError(t *testing.T) {
	r := Response{
		Body: []byte(``),
	}

	err := r.Error()

	// then: error returned
	assert.Error(t, err)
}
func Test_Error_ErrorResponse(t *testing.T) {
	r := Response{
		Body: []byte(`{"error":"some example error"}`),
	}

	err := r.Error()

	// then: error returned
	assert.Error(t, err)
}
