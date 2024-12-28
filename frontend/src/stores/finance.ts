import { defineStore } from 'pinia';
import { ApiService } from '@/services/api';
import type { Category, Expense, Income, Summary, TransactionType } from '@/types/transaction';

interface FinanceState {
  summary: Summary | null;
  incomes:  Income[] | [];
  expenses: Expense[] | [];
  incomeCategories: Category[] | [];
  expenseCategories: Category[] | [];
  loading: boolean;
  error: string | null;
}

export const useFinanceStore = defineStore('finance', {
  state: (): FinanceState => ({
    summary: null,
    incomes: [],
    expenses: [],
    incomeCategories: [],
    expenseCategories: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchData() {
      this.loading = true;
      this.error = null;
      try {
        const { expenses, incomes, summary, incomeCategories, expenseCategories } = await ApiService.fetchAllData();
        this.summary = summary;
        this.incomes = incomes;
        this.expenses = expenses;
        this.incomeCategories = incomeCategories;
        this.expenseCategories = expenseCategories;
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'An error occurred';
      } finally {
        this.loading = false;
      }
    },

    async addTransaction(type: TransactionType, data: any) {
      this.loading = true;
      this.error = null;
      try {
        await ApiService.addTransaction(type, data);
        await this.fetchData();
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'An error occurred';
        throw error;
      } finally {
        this.loading = false;
      }
    }
  }
});