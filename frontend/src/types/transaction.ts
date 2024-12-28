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
    monthlyBreakdown: Record<string, MonthData>;
  }
  
  export type TransactionType = 'income' | 'expense';

  export interface Category {
    id: string;
    name: string;
    type: string;
  }