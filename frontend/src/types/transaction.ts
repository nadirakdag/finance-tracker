// src/types/transaction.ts

export interface Transaction {
    id: string;
    amount: number;
    description: string;
    date: string;
  }
  
  export interface Expense extends Transaction {
    category: string;
  }
  
  export interface Income extends Transaction {
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
  
  export interface TransactionFormData {
    amount: string;
    description: string;
    category: string;
    date: string;
  }
  
  export type TransactionType = 'income' | 'expense';