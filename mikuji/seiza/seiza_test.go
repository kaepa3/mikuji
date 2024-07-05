package seiza

import (
	"testing"
)

func TestSeizaInfo_GetCurrent(t *testing.T) {
	type fields struct {
		nowValue Seiza
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"初期化チェック", fields{nowValue: Aries}, "おひつじ座"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SeizaInfo{
				nowValue: tt.fields.nowValue,
			}
			if got := s.GetCurrent(); got != tt.want {
				t.Errorf("SeizaInfo.GetCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeizaInfo_Next(t *testing.T) {
	type fields struct {
		nowValue Seiza
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"初期化チェック", fields{nowValue: Aries}, "おうし座"},
		{"初期化チェック", fields{nowValue: Pisces}, "おひつじ座"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SeizaInfo{
				nowValue: tt.fields.nowValue,
			}
			s.Next()
			if got := s.GetCurrent(); got != tt.want {
				t.Errorf("SeizaInfo.GetCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
