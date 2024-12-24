import React, { useEffect, useState } from 'react';
import { Alert, AlertDescription } from "@/components/ui/alert";
import ErrorBoundary from '../common/ErrorBoundary';
import LoadingSpinner from '../common/LoadingSpinner';
import SummaryCards from '../dashboard/SummaryCards';
import TransactionForm from '../dashboard/TransactionForm';
import MonthlyOverview from '../dashboard/MonthlyOverview';
import CategoryBreakdown from '../dashboard/CategoryBreakdown';
import { ApiService } from '../../services/api';
import { Summary, TransactionFormData, TransactionType } from '../../types/transaction';

const DashboardLayout: React.FC = () => {
  const [summary, setSummary] = useState<Summary | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);

  const loadData = async () => {
    try {
      setLoading(true);
      const { summary } = await ApiService.fetchAllData();
      setSummary(summary);
    } catch (err) {
      setError('Failed to load data. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleTransactionSubmit = async (type: TransactionType, data: TransactionFormData) => {
    try {
      setIsSubmitting(true);
      await ApiService.addTransaction(type, data);
      await loadData();
    } catch (err) {
      setError('Failed to add transaction. Please try again.');
      throw err;
    } finally {
      setIsSubmitting(false);
    }
  };

  if (loading) {
    return <LoadingSpinner fullScreen />;
  }

  return (
    <ErrorBoundary>
      <div className="min-h-screen bg-gray-50">
        <header className="bg-white shadow">
          <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
            <h1 className="text-3xl font-bold text-gray-900">
              Finance Tracker
            </h1>
          </div>
        </header>

        <main>
          <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
            <div className="px-4 py-6 sm:px-0 space-y-6">
              {error && (
                <Alert variant="destructive">
                  <AlertDescription>{error}</AlertDescription>
                </Alert>
              )}

              <SummaryCards 
                summary={summary} 
                isLoading={loading} 
              />

              <TransactionForm 
                onSubmit={handleTransactionSubmit}
                isLoading={isSubmitting}
              />

              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <ErrorBoundary>
                  <MonthlyOverview 
                    monthlyBreakdown={summary?.monthlyBreakdown} 
                    isLoading={loading}
                  />
                </ErrorBoundary>

                <ErrorBoundary>
                  <CategoryBreakdown 
                    categoryTotals={summary?.categoryTotals}
                    isLoading={loading}
                  />
                </ErrorBoundary>
              </div>
            </div>
          </div>
        </main>
      </div>
    </ErrorBoundary>
  );
};

export default DashboardLayout;