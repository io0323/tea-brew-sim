/*
  シミュレーションロジック
  茶葉ごとの理想温度・時間、スコア計算、コメント生成
*/
package main

import (
  "fmt"
  "math"
)

// 茶葉ごとの理想値定義
// 茶葉名: 理想温度(℃), 理想時間(秒)
var teaIdeal = map[string]struct {
  Temp float64
  Time int
}{
  "煎茶":   {Temp: 75, Time: 60},
  "ほうじ茶": {Temp: 90, Time: 30},
  "玉露":   {Temp: 55, Time: 120},
  "烏龍茶": {Temp: 95, Time: 60},
}

/*
  スコア計算関数
  温度・時間が理想から外れるほど渋みUP/甘みDOWNなどの傾向を反映
*/
func calcScores(req SimulateRequest) (astringency, sweetness, aroma float64) {
  ideal, ok := teaIdeal[req.TeaType]
  if !ok {
    // 未知の茶葉はデフォルト
    ideal = teaIdeal["煎茶"]
  }
  // 温度・時間のズレ
  tempDiff := req.Temp - ideal.Temp
  timeDiff := float64(req.Time - ideal.Time)

  // 渋み: 温度・時間が高いほど増加
  astringency = 0.5 + 0.03*tempDiff + 0.01*timeDiff
  // 甘み: 理想から外れるほど減少
  sweetness = 0.8 - 0.02*math.Abs(tempDiff) - 0.01*math.Abs(timeDiff)
  // 香り: 温度が高いほどやや増加
  aroma = 0.6 + 0.01*tempDiff

  // 0〜1にクリップ
  astringency = math.Max(0, math.Min(1, astringency))
  sweetness = math.Max(0, math.Min(1, sweetness))
  aroma = math.Max(0, math.Min(1, aroma))
  return
}

/*
  おすすめコメント生成関数
  条件に応じて3パターン程度に分類
*/
func makeComment(req SimulateRequest) string {
  ideal, ok := teaIdeal[req.TeaType]
  if !ok {
    ideal = teaIdeal["煎茶"]
  }
  tempDiff := req.Temp - ideal.Temp
  timeDiff := req.Time - ideal.Time

  if math.Abs(tempDiff) <= 3 && math.Abs(float64(timeDiff)) <= 10 {
    return fmt.Sprintf("理想的な抽出です！%sは%.0f℃・%d秒がおすすめ。", req.TeaType, ideal.Temp, ideal.Time)
  }
  if tempDiff > 5 {
    return "温度が高めです。渋みが強く出るかもしれません。やや低めの温度がおすすめ。"
  }
  if tempDiff < -5 {
    return "温度が低めです。甘みは出ますが、香りが弱くなるかもしれません。"
  }
  if timeDiff > 15 {
    return "抽出時間が長めです。渋みが強くなりやすいので注意。"
  }
  if timeDiff < -15 {
    return "抽出時間が短めです。味が薄くなるかもしれません。"
  }
  return fmt.Sprintf("%sのおすすめは%.0f℃・%d秒です。", req.TeaType, ideal.Temp, ideal.Time)
} 