<template>
  <div class="bg-white rounded-lg shadow-lg border border-gray-100">
    <div class="px-8 py-6">
      <h3 class="text-xl font-semibold mb-6 text-gray-800">Add Transaction</h3>

      <!-- Tab Design -->
      <div class="flex rounded-xl bg-gray-50 p-1.5 mb-8 border border-gray-200">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'flex-1 text-sm font-medium rounded-lg py-2.5 px-6',
            'flex items-center justify-center gap-2',
            activeTab === tab.id
              ? `bg-white text-${tab.color}-600 shadow-sm ring-1 ring-gray-200`
              : 'text-gray-600 hover:text-gray-800 hover:bg-gray-100'
          ]"
        >
          <component 
            :is="tab.icon" 
            class="w-4 h-4"
            :class="activeTab === tab.id ? `text-${tab.color}-600` : 'text-gray-500'"
          />
          {{ tab.name }}
        </button>
      </div>

      <!-- Form Content -->
      <div>
        <component
          :is="activeTab === 'income' ? IncomeForm : ExpenseForm"
          :loading="loading"
          :categories="activeTab === 'income' ? incomeCategories : expenseCategories"
          @submit="data => handleSubmit(activeTab as TransactionType, data)"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ArrowDownCircle, ArrowUpCircle } from 'lucide-vue-next';
import type { Category, TransactionType } from '@/types/transaction';
import type { TransactionFormData } from '@/types/forms';
import IncomeForm from './IncomeForm.vue';
import ExpenseForm from './ExpenseForm.vue';

const props = defineProps<{
  loading?: boolean;
  incomeCategories?: Category[];
  expenseCategories?: Category[];
}>();

const emit = defineEmits<{
  (e: 'submit', type: TransactionType, data: TransactionFormData): void;
}>();

const tabs = [
  { 
    id: 'income', 
    name: 'Income', 
    color: 'green',
    icon: ArrowUpCircle 
  },
  { 
    id: 'expense', 
    name: 'Expense', 
    color: 'red',
    icon: ArrowDownCircle 
  }
];

const activeTab = ref('income');

const handleSubmit = (type: TransactionType, data: TransactionFormData) => {
  emit('submit', type, data);
};
</script>