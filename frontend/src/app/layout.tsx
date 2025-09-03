import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "TeaBrewSim - お茶抽出シミュレーター",
  description: "お茶の抽出条件をシミュレーションして最適な味を見つけよう",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
