export interface TransactionFormData {
    amount: string;
    description: string;
    category: string;
    source: string;
    date: string;
  }
  
  export interface Option {
    value: string;
    label: string;
  }

  export interface CategoryListData {
    category: string;
    type: string;
    amount: number;
  }