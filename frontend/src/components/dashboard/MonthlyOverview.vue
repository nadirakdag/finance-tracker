<template>
  <div class="bg-white rounded-lg shadow">
    <div class="px-6 py-5">
      <h3 class="text-lg font-semibold mb-4">Monthly Overview</h3>
      
      <div v-if="loading" class="flex justify-center">
        <LoadingSpinner />
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead>
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Month
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Income
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Expenses
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Balance
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="[month, data] in sortedMonths" :key="month">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ month }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-right text-green-600">
                ${{ data.income.toFixed(2) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-right text-red-600">
                ${{ data.expenses.toFixed(2) }}
              </td>
              <td 
                class="px-6 py-4 whitespace-nowrap text-sm text-right"
                :class="data.balance >= 0 ? 'text-green-600' : 'text-red-600'"
              >
                ${{ data.balance.toFixed(2) }}
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
import type { MonthData } from '@/types/transaction';
import LoadingSpinner from '../common/LoadingSpinner.vue';

const props = defineProps<{
  monthlyBreakdown?: Record<string, MonthData>;
  loading?: boolean;
}>();

const sortedMonths = computed(() => {
  if (!props.monthlyBreakdown) return [];
  return Object.entries(props.monthlyBreakdown)
    .sort((a, b) => b[0].localeCompare(a[0]));
});</script>