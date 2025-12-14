package logger

import (
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{
			name: "should return trace",
			l:    LEVEL_TRACE,
			want: "trace",
		},
		{
			name: "should return debug",
			l:    LEVEL_DEBUG,
			want: "debug",
		},
		{
			name: "should return info",
			l:    LEVEL_INFO,
			want: "info",
		},
		{
			name: "should return warn",
			l:    LEVEL_WARN,
			want: "warn",
		},
		{
			name: "should return error",
			l:    LEVEL_ERROR,
			want: "error",
		},
		{
			name: "should return fatal",
			l:    LEVEL_FATAL,
			want: "fatal",
		},
		{
			name: "should return unknown",
			l:    0,
			want: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("Level.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_IsValid(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want bool
	}{
		{
			name: "should return true",
			l:    LEVEL_TRACE,
			want: true,
		},
		{
			name: "should return true",
			l:    LEVEL_DEBUG,
			want: true,
		},
		{
			name: "should return true",
			l:    LEVEL_INFO,
			want: true,
		},
		{
			name: "should return true",
			l:    LEVEL_WARN,
			want: true,
		},
		{
			name: "should return true",
			l:    LEVEL_ERROR,
			want: true,
		},
		{
			name: "should return true",
			l:    LEVEL_FATAL,
			want: true,
		},
		{
			name: "should return false",
			l:    -1,
			want: false,
		},
		{
			name: "should return false",
			l:    0,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsValid(); got != tt.want {
				t.Errorf("Level.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevel_NewLevelFromString(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		want        Level
		shouldError bool
	}{
		{
			name:        "should return trace",
			s:           "trace",
			want:        LEVEL_TRACE,
			shouldError: false,
		},
		{
			name:        "should return debug",
			s:           "debug",
			want:        LEVEL_DEBUG,
			shouldError: false,
		},
		{
			name:        "should return info",
			s:           "info",
			want:        LEVEL_INFO,
			shouldError: false,
		},
		{
			name:        "should return warn",
			s:           "warn",
			want:        LEVEL_WARN,
			shouldError: false,
		},
		{
			name:        "should return error",
			s:           "error",
			want:        LEVEL_ERROR,
			shouldError: false,
		},
		{
			name:        "should return fatal",
			s:           "fatal",
			want:        LEVEL_FATAL,
			shouldError: false,
		},
		{
			name:        "should return error",
			s:           "something_else",
			want:        0,
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLevelFromString(tt.s)
			if tt.shouldError {
				if err == nil {
					t.Errorf("NewLevelFromString() error = %v, wantErr %v", err, tt.shouldError)
					return
				}
			} else {
				if err != nil {
					t.Errorf("NewLevelFromString() error = %v, wantErr %v", err, tt.shouldError)
					return
				}

				if got != tt.want {
					t.Errorf("NewLevelFromString() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
