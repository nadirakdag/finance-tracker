import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { MonthData } from '../../types/transaction';
import LoadingSpinner from '../common/LoadingSpinner';

interface MonthlyOverviewProps {
  monthlyBreakdown: Record<string, MonthData> | undefined;
  isLoading?: boolean;
}

const MonthlyOverview: React.FC<MonthlyOverviewProps> = ({ 
  monthlyBreakdown,
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

  const sortedMonths = Object.entries(monthlyBreakdown || {})
    .sort((a, b) => b[0].localeCompare(a[0]));

  return (
    <Card>
      <CardHeader>
        <CardTitle>Monthly Overview</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="overflow-x-auto">
          <table className="w-full border-collapse">
            <thead>
              <tr className="border-b">
                <th className="text-left p-2 font-semibold">Month</th>
                <th className="text-right p-2 font-semibold">Income</th>
                <th className="text-right p-2 font-semibold">Expenses</th>
                <th className="text-right p-2 font-semibold">Balance</th>
              </tr>
            </thead>
            <tbody>
              {sortedMonths.map(([month, data]) => (
                <tr key={month} className="border-b">
                  <td className="p-2 font-medium">{month}</td>
                  <td className="text-right p-2 text-green-600">
                    ${data.income.toFixed(2)}
                  </td>
                  <td className="text-right p-2 text-red-600">
                    ${data.expenses.toFixed(2)}
                  </td>
                  <td className={`text-right p-2 ${data.balance >= 0 ? 'text-green-600' : 'text-red-600'}`}>
                    ${data.balance.toFixed(2)}
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

export default MonthlyOverview;