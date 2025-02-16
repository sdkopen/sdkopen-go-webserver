package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonType(t *testing.T) {
	t.Run("Should error when scan with a nil value", func(t *testing.T) {
		var result JsonB
		err := result.Scan(nil)

		assert.Error(t, err, "type assertion to []byte failed")
		assert.Nil(t, result)
	})

	t.Run("Should error when scan with a invalid json", func(t *testing.T) {
		value := "{\"name\":\"teste\""

		var result JsonB
		err := result.Scan([]byte(value))

		assert.Error(t, err, "type assertion to []byte failed")
		assert.Nil(t, result)
	})

	t.Run("Should scan with a valid value", func(t *testing.T) {
		value := "{\"name\":\"teste\"}"

		var result JsonB
		err := result.Scan([]byte(value))

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "teste", result["name"])
	})

	t.Run("Should get value with a valid json", func(t *testing.T) {
		expected := "{\"name\":\"teste\"}"

		var model JsonB
		err := model.Scan([]byte(expected))
		assert.Nil(t, err)
		assert.NotNil(t, model)
		assert.Equal(t, "teste", model["name"])

		result, err := model.Value()
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, string(result.([]byte)))
	})
}
