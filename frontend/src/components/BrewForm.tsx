import React, { useState } from "react";

export type BrewFormValues = {
  teaType: string;
  temp: number;
  time: number;
};

type BrewFormProps = {
  onSimulate: (values: BrewFormValues) => void;
  loading?: boolean;
};

export const BrewForm: React.FC<BrewFormProps> = ({ onSimulate, loading }) => {
  const [teaType, setTeaType] = useState("緑茶");
  const [temp, setTemp] = useState(80);
  const [time, setTime] = useState(60);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSimulate({ teaType, temp, time });
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-full max-w-md">
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          茶葉の種類
        </label>
        <select
          value={teaType}
          onChange={(e) => setTeaType(e.target.value)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
        >
          <option value="緑茶">緑茶</option>
          <option value="紅茶">紅茶</option>
          <option value="ウーロン茶">ウーロン茶</option>
          <option value="白茶">白茶</option>
          <option value="プーアル茶">プーアル茶</option>
        </select>
      </div>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          温度: {temp}°C
        </label>
        <input
          type="range"
          min="60"
          max="100"
          value={temp}
          onChange={(e) => setTemp(Number(e.target.value))}
          className="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700 slider"
        />
      </div>

      <div className="mb-6">
        <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          時間 (秒)
        </label>
        <input
          type="number"
          min="10"
          max="300"
          value={time}
          onChange={(e) => setTime(Number(e.target.value))}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white"
        />
      </div>

      <button
        type="submit"
        disabled={loading}
        className="w-full bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-medium py-2 px-4 rounded-md transition duration-200"
      >
        {loading ? "シミュレーション中..." : "シミュレーション"}
      </button>
    </form>
  );
};
