package logger

import (
	"context"
	"errors"
	"testing"

	testutils "github.com/iolave/go-logger/internal/test_utils"
)

func Test_newEntry(t *testing.T) {
	t.Run("should return an entry with an empty info when info is nil", func(t *testing.T) {
		e := newEntry(
			context.Background(),
			"", "",
			0,
			"msg",
			nil,
			nil,
		)

		if e.Info == nil {
			t.Errorf("expected info to be empty, got %v", e.Info)
		}
	})
}

func TestEntry_serialize(t *testing.T) {
	t.Run("should destroy circular references", func(t *testing.T) {
		w := `{"timestamp":0,"level":"unknown","name":"","version":"","trace":{},"info":{"data":"my data"},"msg":""}`
		info := map[string]any{
			"data": "my data",
		}
		info["circular"] = &info

		e := newEntry(
			context.Background(),
			"",
			"",
			0,
			"",
			info,
			nil,
		)
		e.Timestamp = 0

		s := e.serialize()
		if s != w {
			testutils.CmpString(t, w, s)
		}

	})

	t.Run("should set the info to an empty map if it is not json marshalable", func(t *testing.T) {
		w := `{"timestamp":0,"level":"unknown","name":"","version":"","trace":{},"error":{"name":"error","message":"my error","original":{}},"info":{},"msg":""}`
		info := map[string]any{
			"func": func() {},
		}

		e := newEntry(
			context.Background(),
			"",
			"",
			0,
			"",
			info,
			errors.New("my error"),
		)
		e.Timestamp = 0

		s := e.serialize()
		if s != w {
			testutils.CmpString(t, w, s)
		}
	})
}
