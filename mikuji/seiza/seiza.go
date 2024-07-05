package seiza

type Seiza int

const (
	Aries Seiza = iota + 1
	Taurus
	Gemini
	Cancer
	Leo
	Virgo
	Libra
	Scorpio
	Sagittarius
	Capricorn
	Aquarius
	Pisces
)
const (
	AriesText       = "おひつじ座"
	TaurusText      = "おうし座"
	GeminiText      = "ふたご座"
	CancerText      = "かに座"
	LeoText         = "しし座"
	VirgoText       = "おとめ座"
	LibraText       = "てんびん座"
	ScorpioText     = "さそり座"
	SagittariusText = "いて座"
	CapricornText   = "やぎ座"
	AquariusText    = "みずがめ座"
	PiscesText      = "うお座"
)

type SeizaInfo struct {
	nowValue Seiza
}

func NewSeizaInfo() *SeizaInfo {
	obj := SeizaInfo{nowValue: Aries}
	return &obj
}

func (s *SeizaInfo) GetCurrent() string {
	var name string
	switch s.nowValue {
	case Aries:
		name = AriesText
	case Taurus:
		name = TaurusText
	case Gemini:
		name = GeminiText
	case Cancer:
		name = CancerText
	case Leo:
		name = LeoText
	case Virgo:
		name = VirgoText
	case Libra:
		name = LibraText
	case Scorpio:
		name = ScorpioText
	case Sagittarius:
		name = SagittariusText
	case Capricorn:
		name = CapricornText
	case Aquarius:
		name = AquariusText
	case Pisces:
		name = PiscesText
	}
	return name
}

func (s *SeizaInfo) Next() {
	s.nowValue += 1
	if s.nowValue > Pisces {
		s.nowValue = Aries
	}
}

func (s *SeizaInfo) Before() {
	s.nowValue -= 1
	if s.nowValue < Aries {
		s.nowValue = Pisces
	}
}
