import React from "react";
import {
  Radar, RadarChart, PolarGrid, PolarAngleAxis, PolarRadiusAxis, ResponsiveContainer
} from "recharts";

type ResultChartProps = {
  astringency: number;
  sweetness: number;
  aroma: number;
};

export const ResultChart: React.FC<ResultChartProps> = ({ astringency, sweetness, aroma }) => {
  const data = [
    { name: "渋み", value: astringency },
    { name: "甘み", value: sweetness },
    { name: "香り", value: aroma },
  ];

  return (
    <div className="mx-auto">
      <ResponsiveContainer width={256} height={256}>
        <RadarChart data={data}>
          <PolarGrid />
          <PolarAngleAxis dataKey="name" />
          <PolarRadiusAxis angle={30} domain={[0, 1]} tickCount={6} />
          <Radar 
            name="スコア" 
            dataKey="value" 
            stroke="#16a34a" 
            fill="#16a34a" 
            fillOpacity={0.5} 
          />
        </RadarChart>
      </ResponsiveContainer>
    </div>
  );
};
