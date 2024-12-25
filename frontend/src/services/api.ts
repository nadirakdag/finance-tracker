import type { Expense, Income, Summary, TransactionType } from '@/types/transaction';

const API_BASE_URL = 'http://localhost:8080/api/v1';

export class ApiService {
  private static async fetchWithError(url: string, options?: RequestInit): Promise<any> {
    const response = await fetch(url, options);
    
    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new Error(error.message || 'An error occurred');
    }
    
    return response.json();
  }

  static async getExpenses(): Promise<Expense[]> {
    return this.fetchWithError(`${API_BASE_URL}/expenses`);
  }

  static async getIncomes(): Promise<Income[]> {
    return this.fetchWithError(`${API_BASE_URL}/incomes`);
  }

  static async getSummary(): Promise<Summary> {
    return this.fetchWithError(`${API_BASE_URL}/summary`);
  }

  static async addTransaction(type: TransactionType, data: any): Promise<Expense | Income> {
    return this.fetchWithError(`${API_BASE_URL}/${type}s`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
  }

  static async fetchAllData() {
    const [expenses, incomes, summary] = await Promise.all([
      this.getExpenses(),
      this.getIncomes(),
      this.getSummary(),
    ]);

    return { expenses, incomes, summary };
  }
}