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
	AriesText       = "牡羊座"
	TaurusText      = "牡牛座"
	GeminiText      = "双子座"
	CancerText      = "蟹座"
	LeoText         = "獅子座"
	VirgoText       = "乙女座"
	LibraText       = "天秤座"
	ScorpioText     = "蠍座"
	SagittariusText = "射手座"
	CapricornText   = "山羊座"
	AquariusText    = "水瓶座"
	PiscesText      = "魚座"
)

type SeizaInfo struct {
	nowValue Seiza
}

func NewSeizaInfo() *SeizaInfo {
	obj := SeizaInfo{nowValue: Aries}
	return &obj
}

func (s *SeizaInfo) GetCurrentValue() Seiza {
	return s.nowValue
}

func (s *SeizaInfo) GetCurrent() string {
	return IndexToString(s.nowValue)
}

func IndexToString(sz Seiza) string {
	var name string
	switch sz {
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
