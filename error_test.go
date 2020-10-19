package handle

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestHandle
func TestHandle(t *testing.T) {
	t.Run("handles panic", func(t *testing.T) {
		var err error
		result := errors.New("handled")

		defer func() {
			require.EqualError(t, result, err.Error())
		}()
		defer Handle(&err, func(err error) error {
			return errors.New("handled")
		})

		panic(errors.New("unhandled"))
	})

	t.Run("handles error", func(t *testing.T) {
		var err error
		result := errors.New("handled")

		defer func() {
			require.EqualError(t, result, err.Error())
		}()
		defer Handle(&err, func(err error) error {
			return errors.New("handled")
		})

		err = errors.New("unhandled")
	})

	t.Run("handles nil", func(t *testing.T) {
		var err error
		defer func() { require.Nil(t, err) }()

		defer Handle(&err, func(err error) error {
			return errors.New("handled")
		})

		require.Nil(t, err)
	})
}

// TestCheck
func TestCheck(t *testing.T) {
	t.Run("triggers panic", func(t *testing.T) {
		err := errors.New("panic")
		defer func() {
			require.EqualError(t, err, recover().(error).Error())
		}()
		Check(errors.New("panic"))
	})

	t.Run("accepts nil", func(t *testing.T) {
		defer func() {
			require.Nil(t, recover())
		}()
		Check(nil)
	})
}
