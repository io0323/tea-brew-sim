/*
TeaBrewSim テストファイル
各種機能のユニットテスト
*/
package main

import (
	"math"
	"testing"
	"time"
)

// TestCalcScores スコア計算のテスト
func TestCalcScores(t *testing.T) {
	tests := []struct {
		name     string
		request  SimulateRequest
		expected struct {
			astringency float64
			sweetness   float64
			aroma       float64
		}
	}{
		{
			name: "煎茶の理想的な条件",
			request: SimulateRequest{
				TeaType: "煎茶",
				Temp:    75.0,
				Time:    60,
			},
			expected: struct {
				astringency float64
				sweetness   float64
				aroma       float64
			}{
				astringency: 0.5,
				sweetness:   0.8,
				aroma:       0.6,
			},
		},
		{
			name: "玉露の理想的な条件",
			request: SimulateRequest{
				TeaType: "玉露",
				Temp:    55.0,
				Time:    120,
			},
			expected: struct {
				astringency float64
				sweetness   float64
				aroma       float64
			}{
				astringency: 0.5,
				sweetness:   0.8,
				aroma:       0.6,
			},
		},
		{
			name: "高温すぎる煎茶",
			request: SimulateRequest{
				TeaType: "煎茶",
				Temp:    85.0,
				Time:    60,
			},
			expected: struct {
				astringency float64
				sweetness   float64
				aroma       float64
			}{
				astringency: 0.8, // 高温で苦味が強くなる
				sweetness:   0.6, // 甘味が減少
				aroma:       0.7, // 香りが増加
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			astringency, sweetness, aroma := calcScores(tt.request)
			
			if astringency != tt.expected.astringency {
				t.Errorf("astringency = %v, expected %v", astringency, tt.expected.astringency)
			}
			if math.Abs(sweetness-tt.expected.sweetness) > 0.0001 {
				t.Errorf("sweetness = %v, expected %v", sweetness, tt.expected.sweetness)
			}
			if aroma != tt.expected.aroma {
				t.Errorf("aroma = %v, expected %v", aroma, tt.expected.aroma)
			}
		})
	}
}

// TestMakeComment コメント生成のテスト
func TestMakeComment(t *testing.T) {
	tests := []struct {
		name     string
		request  SimulateRequest
		expected string
	}{
		{
			name: "理想的な煎茶",
			request: SimulateRequest{
				TeaType: "煎茶",
				Temp:    75.0,
				Time:    60,
			},
			expected: "理想的な抽出です！煎茶は75℃・60秒がおすすめ。",
		},
		{
			name: "高温すぎる煎茶",
			request: SimulateRequest{
				TeaType: "煎茶",
				Temp:    85.0,
				Time:    60,
			},
			expected: "温度が高めです。渋みが強く出るかもしれません。やや低めの温度がおすすめ。",
		},
		{
			name: "時間が短すぎる玉露",
			request: SimulateRequest{
				TeaType: "玉露",
				Temp:    55.0,
				Time:    100,
			},
			expected: "抽出時間が短めです。味が薄くなるかもしれません。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := makeComment(tt.request)
			if comment != tt.expected {
				t.Errorf("comment = %v, expected %v", comment, tt.expected)
			}
		})
	}
}

// TestTeaIdeal お茶の理想条件のテスト
func TestTeaIdeal(t *testing.T) {
	// 煎茶の理想条件をテスト
	senchaIdeal, exists := teaIdeal["煎茶"]
	if !exists {
		t.Error("煎茶の理想条件が定義されていません")
	}
	if senchaIdeal.Temp != 75.0 {
		t.Errorf("煎茶の理想温度 = %v, expected 75.0", senchaIdeal.Temp)
	}
	if senchaIdeal.Time != 60 {
		t.Errorf("煎茶の理想時間 = %v, expected 60", senchaIdeal.Time)
	}

	// 玉露の理想条件をテスト
	gyokuroIdeal, exists := teaIdeal["玉露"]
	if !exists {
		t.Error("玉露の理想条件が定義されていません")
	}
	if gyokuroIdeal.Temp != 55.0 {
		t.Errorf("玉露の理想温度 = %v, expected 55.0", gyokuroIdeal.Temp)
	}
	if gyokuroIdeal.Time != 120 {
		t.Errorf("玉露の理想時間 = %v, expected 120", gyokuroIdeal.Time)
	}
}

// TestSimulateRequest リクエスト構造体のテスト
func TestSimulateRequest(t *testing.T) {
	req := SimulateRequest{
		TeaType: "煎茶",
		Temp:    75.0,
		Time:    60,
	}

	if req.TeaType != "煎茶" {
		t.Errorf("TeaType = %v, expected 煎茶", req.TeaType)
	}
	if req.Temp != 75.0 {
		t.Errorf("Temp = %v, expected 75.0", req.Temp)
	}
	if req.Time != 60 {
		t.Errorf("Time = %v, expected 60", req.Time)
	}
}

// TestSimulateResponse レスポンス構造体のテスト
func TestSimulateResponse(t *testing.T) {
	res := SimulateResponse{
		Astringency: 0.3,
		Sweetness:   0.7,
		Aroma:       0.8,
		Comment:     "理想的な抽出条件です！",
	}

	if res.Astringency != 0.3 {
		t.Errorf("Astringency = %v, expected 0.3", res.Astringency)
	}
	if res.Sweetness != 0.7 {
		t.Errorf("Sweetness = %v, expected 0.7", res.Sweetness)
	}
	if res.Aroma != 0.8 {
		t.Errorf("Aroma = %v, expected 0.8", res.Aroma)
	}
	if res.Comment != "理想的な抽出条件です！" {
		t.Errorf("Comment = %v, expected 理想的な抽出条件です！", res.Comment)
	}
}

// TestBrewingRecord 抽出記録のテスト
func TestBrewingRecord(t *testing.T) {
	now := time.Now()
	record := &BrewingRecord{
		ID:          1,
		TeaType:     "煎茶",
		Temperature: 75.0,
		Time:        60,
		Rating:      4.5,
		Comment:     "美味しかった",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if record.ID != 1 {
		t.Errorf("ID = %v, expected 1", record.ID)
	}
	if record.TeaType != "煎茶" {
		t.Errorf("TeaType = %v, expected 煎茶", record.TeaType)
	}
	if record.Temperature != 75.0 {
		t.Errorf("Temperature = %v, expected 75.0", record.Temperature)
	}
	if record.Time != 60 {
		t.Errorf("Time = %v, expected 60", record.Time)
	}
	if record.Rating != 4.5 {
		t.Errorf("Rating = %v, expected 4.5", record.Rating)
	}
	if record.Comment != "美味しかった" {
		t.Errorf("Comment = %v, expected 美味しかった", record.Comment)
	}
}

// TestConfig 設定のテスト
func TestConfig(t *testing.T) {
	config := &Config{
		Server: ServerConfig{
			Port:         "8080",
			Host:         "localhost",
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "teabrewsim",
			Password: "password",
			DBName:   "teabrewsim",
			SSLMode:  "disable",
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
	}

	if config.Server.Port != "8080" {
		t.Errorf("Server.Port = %v, expected 8080", config.Server.Port)
	}
	if config.Database.Port != 5432 {
		t.Errorf("Database.Port = %v, expected 5432", config.Database.Port)
	}
	if config.Logging.Level != "info" {
		t.Errorf("Logging.Level = %v, expected info", config.Logging.Level)
	}
}

// BenchmarkCalcScores スコア計算のベンチマーク
func BenchmarkCalcScores(b *testing.B) {
	req := SimulateRequest{
		TeaType: "煎茶",
		Temp:    75.0,
		Time:    60,
	}

	for i := 0; i < b.N; i++ {
		calcScores(req)
	}
}

// BenchmarkMakeComment コメント生成のベンチマーク
func BenchmarkMakeComment(b *testing.B) {
	req := SimulateRequest{
		TeaType: "煎茶",
		Temp:    75.0,
		Time:    60,
	}

	for i := 0; i < b.N; i++ {
		makeComment(req)
	}
}
