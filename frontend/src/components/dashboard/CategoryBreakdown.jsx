import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import LoadingSpinner from '../common/LoadingSpinner';

interface CategoryBreakdownProps {
  categoryTotals: Record<string, number> | undefined;
  isLoading?: boolean;
}

const CategoryBreakdown: React.FC<CategoryBreakdownProps> = ({ 
  categoryTotals,
  isLoading 
}) => {
  if (isLoading) {
    return (
      <Card>
        <CardContent className="p-6">
          <LoadingSpinner />
        </CardContent>
      </Card>
    );
  }

  const sortedCategories = Object.entries(categoryTotals || {})
    .sort((a, b) => b[1] - a[1]);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Category Breakdown</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="overflow-x-auto">
          <table className="w-full border-collapse">
            <thead>
              <tr className="border-b">
                <th className="text-left p-2 font-semibold">Category</th>
                <th className="text-right p-2 font-semibold">Amount</th>
              </tr>
            </thead>
            <tbody>
              {sortedCategories.map(([category, amount]) => (
                <tr key={category} className="border-b">
                  <td className="p-2 font-medium">
                    {category.charAt(0).toUpperCase() + category.slice(1)}
                  </td>
                  <td className="text-right p-2 text-red-600">
                    ${amount.toFixed(2)}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </CardContent>
    </Card>
  );
};

export default CategoryBreakdown;