import React from "react";

type RecommendCommentProps = {
  comment: string;
};

export const RecommendComment: React.FC<RecommendCommentProps> = ({ comment }) => (
  <div className="bg-green-50 dark:bg-green-900 border border-green-200 dark:border-green-700 rounded-lg p-4">
    <h3 className="text-lg font-semibold text-green-800 dark:text-green-200 mb-2">
      おすすめの抽出方法
    </h3>
    <p className="text-green-700 dark:text-green-300">{comment}</p>
  </div>
);
