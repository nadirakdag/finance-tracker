// src/components/dashboard/SummaryCards.tsx

import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Summary } from '../../types/transaction';
import LoadingSpinner from '../common/LoadingSpinner';
import ErrorBoundary from '../common/ErrorBoundary';

interface SummaryCardsProps {
  summary: Summary | null;
  isLoading?: boolean;
}

const SummaryCards: React.FC<SummaryCardsProps> = ({ summary, isLoading }) => {
  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {[1, 2, 3].map((i) => (
          <Card key={i}>
            <CardContent className="p-6">
              <LoadingSpinner size="small" />
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  if (!summary) {
    return null;
  }

  const cards = [
    {
      title: 'Total Income',
      amount: summary.totalIncome,
      color: 'text-green-600'
    },
    {
      title: 'Total Expenses',
      amount: summary.totalExpenses,
      color: 'text-red-600'
    },
    {
      title: 'Balance',
      amount: summary.remainingAmount,
      color: summary.remainingAmount >= 0 ? 'text-green-600' : 'text-red-600'
    }
  ];

  return (
    <ErrorBoundary>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {cards.map((card) => (
          <Card key={card.title}>
            <CardHeader>
              <CardTitle>{card.title}</CardTitle>
            </CardHeader>
            <CardContent>
              <p className={`text-2xl font-bold ${card.color}`}>
                ${card.amount.toFixed(2)}
              </p>
            </CardContent>
          </Card>
        ))}
      </div>
    </ErrorBoundary>
  );
};

export default SummaryCards;