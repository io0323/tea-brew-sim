# TeaBrewSim

お茶の抽出条件（茶葉の種類・温度・時間）を入力すると、最適な抽出方法と味の傾向（渋み・香り・甘み）をシミュレーションできるWebアプリです。

## 構成
- フロントエンド: Next.js (TypeScript) + Tailwind CSS
- バックエンド: Go (Gin)
- データベース不要、APIとUIのみのシンプルな構成

## ディレクトリ構成
```
frontend/   # Next.js + Tailwind CSS フロントエンド
*.go        # Go (Gin) バックエンドAPI（ルートディレクトリ）
go.mod      # Go モジュール定義
```

## 起動方法

### 1. バックエンドAPI（Go）
```sh
# 依存解決（初回のみ）
go mod tidy
# サーバー起動
go run .
```
- ポート: 8080

### 2. フロントエンド（Next.js）
```sh
cd frontend
# 依存解決（初回のみ）
npm install
# サーバー起動
npm run dev
```
- ポート: 3000

### 3. ブラウザでアクセス
[http://localhost:3000](http://localhost:3000)

## 使い方
1. 茶葉の種類・抽出温度・抽出時間を入力
2. 「シミュレーションする」ボタンを押す
3. 渋み・甘み・香りのレーダーチャートと、おすすめ抽出コメントが表示されます

## 主な技術
- Next.js 15 (App Router, TypeScript)
- Tailwind CSS
- recharts（レーダーチャート表示）
- Go 1.21+
- Gin（APIサーバー）
- CORS対応済み

## ライセンス
MIT
