/*
/api/simulate 用のリクエスト・レスポンス型定義
*/
package main

// シミュレーションリクエストの構造体
type SimulateRequest struct {
	TeaType string  `json:"teaType"` // 茶葉の種類
	Temp    float64 `json:"temp"`    // 抽出温度（℃）
	Time    int     `json:"time"`    // 抽出時間（秒）
}

// シミュレーションレスポンスの構造体
type SimulateResponse struct {
	Astringency float64 `json:"astringency"` // 渋みスコア
	Sweetness   float64 `json:"sweetness"`   // 甘みスコア
	Aroma       float64 `json:"aroma"`       // 香りスコア
	Comment     string  `json:"comment"`     // おすすめコメント
}
