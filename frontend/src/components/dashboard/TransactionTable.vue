<template>
<div class="bg-white rounded-lg shadow">
    <div class="px-6 py-5">
    <h3 class="text-lg font-semibold mb-4">Transactions</h3>
    
    <div v-if="loading" class="flex justify-center">
        <LoadingSpinner />
    </div>

    <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
        <thead>
            <tr class="border-b">
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Category</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Type</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Description</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Amount</th>
            </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="item in sortedCategories" :key="item.category + item.type" class="border-b">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ formatCategory(item.category) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-600" :class="item.type === 'income' ? 'text-green-600' : 'text-red-600'">
                {{ item.type.charAt(0).toUpperCase() + item.type.slice(1) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-600">
                {{ item.description }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-600" :class="item.type === 'income' ? 'text-green-600' : 'text-red-600'">
                ₺{{ item.amount.toFixed(2) }}
              </td>
            </tr>
        </tbody>
        </table>
    </div>
    </div>
</div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import LoadingSpinner from '../common/LoadingSpinner.vue';
import type { Expense, Income } from '@/types/transaction';


const props = defineProps<{
loading?: boolean;
incomes: Income[];
expenses: Expense[];
}>();


interface transaction {
  id: string;
  amount: number;
  description: string;
  date: string;
  type: string;
  category: string;
}

const sortedCategories = computed(() => {
  const allCategories : transaction[] = [];
  
    // Add income categories
    if (props.incomes) {
      Object.entries(props.incomes).forEach(([_, income]) => {
        allCategories.push({
            category: income.source, 
            type: 'income', 
            amount: income.amount, 
            description: income.description,
            id: income.id,
            date: income.date
        });
      });
    }
    
    // Add expense categories
    if (props.expenses) {
      Object.entries(props.expenses).forEach(([_, expense]) => {
        allCategories.push({ 
            category: expense.category, 
            type: 'expense', 
            amount: expense.amount, 
            description: expense.description,
            id: expense.id,
            date: expense.date
         });
      });
    }

  return allCategories.sort((a, b) => b.date - a.date);
});

const formatCategory = (category: string) => {
  return category.charAt(0).toUpperCase() + category.slice(1);
};


</script>