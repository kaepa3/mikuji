package uranai

import (
	"testing"

	"github.com/kaepa3/mikuji/mikuji/seiza"
)

var text = `
{"horoscope":{"2022/12/01":[{"content":"素直になれないとき。片思いの相手とは、できるだけ友達も交えたコミュニケーションを取り、聞き役に徹するのが◎。","item":"厚手のカーテン","money":2,"total":2,"job":2,"color":"グリーン","day":1,"love":2,"rank":11,"sign":"牡羊座"},{"content":"年下の異性との出会いが、あなたに転機を与えてくれそうです。面倒くさがらずに、いいところを探すよう心掛けましょう。","item":"卵焼き","money":5,"total":4,"job":4,"color":"パープル","day":"","love":5,"rank":3,"sign":"牡牛座"},{"content":"お金に関する運気が急上昇。株の勉強などをするのにも良い日です。貯金箱を新調すると、さらに嬉しい出来事の予感！","item":"クッキー","money":5,"total":5,"job":5,"color":"ゴールド","day":"","love":5,"rank":1,"sign":"双子座"},{"content":"努力が実る日。一緒に頑張る人から温かい言葉をかけられて、勇気付けてもらえそうです。感謝の気持ちを素直に表現して。","item":"赤いサインペン","money":3,"total":3,"job":4,"color":"ホワイト","day":"","love":3,"rank":7,"sign":"蟹座"},{"content":"過去に使っていた家具など、古く歴史があるものにツキがある日です。部屋の掃除をすると金運アップ♪面倒くさがらないで。","item":"みかん","money":5,"total":5,"job":4,"color":"ブルー","day":"","love":5,"rank":2,"sign":"獅子座"},{"content":"運気が停滞中。新しいことをするよりも、これまでの自分を振り返りましょう。昔のアルバムを見ると元気になれます。","item":"カイロ","money":4,"total":4,"job":3,"color":"シルバー","day":"","love":4,"rank":5,"sign":"乙女座"},{"content":"新しいプロジェクトの計画は率先して関わって。人任せにすると評判が下がりそう。献身的な気持ちを大事にしましょう。","item":"パイプ椅子","money":4,"total":3,"job":3,"color":"イエロー","day":"","love":4,"rank":6,"sign":"天秤座"},{"content":"身近な人から良い知らせが舞い込んできそうです。お祝い事は派手にして吉。ただし、浪費しすぎないように注意しましょう。","item":"携帯電話","money":3,"total":3,"job":3,"color":"グレー","day":"","love":3,"rank":8,"sign":"蠍座"},{"content":"今日はできるだけいつもと違う行動をしてみましょう。ずっと前に失くした物が返ってきたり、見つかりそうな予感です。","item":"ノートパソコン","money":3,"total":3,"job":2,"color":"オレンジ","day":"","love":3,"rank":9,"sign":"射手座"},{"content":"今日はいつもより清楚なイメージ作りが吉。好きな人の前では、できるだけ言葉遣いや仕草などに、気を配りましょう。","item":"恋愛小説","money":2,"total":2,"job":3,"color":"レッド","day":"","love":2,"rank":10,"sign":"山羊座"},{"content":"うっかりミスが増えそうな日。何事も後片付けは積極的に。いつも丁寧な気持ちを持っていれば、大きな問題にはなりません。","item":"カフェラテ","money":2,"total":1,"job":1,"color":"ブラック","day":"","love":2,"rank":12,"sign":"水瓶座"},{"content":"今日はのんびりモードがGOOD。気ままにショッピングを楽しむのも、良い気分転換に。ランチはのんびり時間をかけて♪","item":"毛糸の靴下","money":4,"total":4,"job":5,"color":"ピンク","day":"","love":4,"rank":4,"sign":"魚座"}]}}
`

func TestGetData(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
	}{
		{"hoge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getData([]byte(text))
			if err != nil || len(got.HoroScope) == 0 {
				t.Fail()
			}
			rec := got.GetRecord(seiza.Aquarius)
			if rec == nil {
				t.Fail()
			}
		})
	}
}
