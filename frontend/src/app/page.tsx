"use client";

import { BrewForm, BrewFormValues } from "../components/BrewForm";
import { ResultChart } from "../components/ResultChart";
import { RecommendComment } from "../components/RecommendComment";
import { useState } from "react";

export default function Home() {
  const [result, setResult] = useState<{
    astringency: number;
    sweetness: number;
    aroma: number;
    comment: string;
  } | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSimulate = async (values: BrewFormValues) => {
    setLoading(true);
    setError("");
    setResult(null);
    try {
      const res = await fetch("http://localhost:8080/api/simulate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
      });
      if (!res.ok) throw new Error("APIエラー");
      const data = await res.json();
      setResult(data);
    } catch (e) {
      setError("シミュレーションに失敗しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-100 via-white to-green-200 dark:from-gray-900 dark:via-gray-800 dark:to-green-900 flex flex-col items-center py-8 px-2">
      <h1 className="text-3xl sm:text-4xl font-bold mb-6 text-green-800 dark:text-green-200 text-center drop-shadow">
        TeaBrewSim<br />
        <span className="text-base font-normal">お茶抽出シミュレーター</span>
      </h1>
      <BrewForm onSimulate={handleSimulate} loading={loading} />
      {error && <div className="text-red-600 mt-4">{error}</div>}
      {result && (
        <div className="w-full flex flex-col items-center mt-6">
          <ResultChart 
            astringency={result.astringency} 
            sweetness={result.sweetness} 
            aroma={result.aroma} 
          />
          <div className="w-full flex justify-center">
            <div className="w-full max-w-md mt-2">
              <RecommendComment comment={result.comment} />
            </div>
          </div>
        </div>
      )}
      <footer className="mt-auto py-6 text-xs text-gray-500 dark:text-gray-400 text-center w-full">
        &copy; {new Date().getFullYear()} TeaBrewSim
      </footer>
    </div>
  );
}
