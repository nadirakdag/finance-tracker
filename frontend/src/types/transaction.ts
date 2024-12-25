
  export interface Expense  {
    id: string;
    amount: number;
    description: string;
    date: string;
    category: string;
  }
  
  export interface Income  {
    id: string;
    amount: number;
    description: string;
    date: string;
    source: string;
  }
  
  export interface MonthData {
    income: number;
    expenses: number;
    balance: number;
  }
  
  export interface Summary {
    totalIncome: number;
    totalExpenses: number;
    remainingAmount: number;
    categoryTotals: Record<string, number>;
    monthlyBreakdown: Record<string, MonthData>;
  }
  
  export type TransactionType = 'income' | 'expense';