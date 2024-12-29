<template>
  <div class="min-h-screen bg-gray-50">
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">
          Finance Tracker
        </h1>
      </div>
    </header>

    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0 space-y-6">
          <div v-if="error" class="rounded-md bg-red-50 p-4">
            <div class="text-sm text-red-700">{{ error }}</div>
          </div>

          <SummaryCards 
            :summary="summary"
            :loading="loading"
          />

          <TransactionForm 
            :loading="loading"
            :income-categories="incomeCategories"
            :expense-categories="expenseCategories"
            @submit="handleTransactionSubmit"
          />

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <MonthlyOverview 
              :monthly-breakdown="summary?.monthlyBreakdown"
              :loading="loading"
            />

            <CategoryBreakdown 
              :incomes="incomes"
              :expenses="expenses"
              :loading="loading"
            />
          </div>

          <TransactionTable :loading="false" :incomes="incomes" :expenses="expenses" />
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useFinanceStore } from '@/stores/finance';
import type { TransactionType } from '@/types/transaction';
import SummaryCards from '@/components/dashboard/SummaryCards.vue';
import TransactionForm from '@/components/dashboard/TransactionForm.vue';
import MonthlyOverview from '@/components/dashboard/MonthlyOverview.vue';
import CategoryBreakdown from '@/components/dashboard/CategoryBreakdown.vue';
import TransactionTable from '@/components/dashboard/TransactionTable.vue';

const store = useFinanceStore();
const { summary, incomes, expenses, loading, error, incomeCategories, expenseCategories } = storeToRefs(store);

onMounted(() => {
  store.fetchData();
});

const handleTransactionSubmit = async (type: TransactionType, data: any) => {
  try {
    await store.addTransaction(type, data);
  } catch (error) {
    console.error('Failed to add transaction:', error);
  }
};</script>