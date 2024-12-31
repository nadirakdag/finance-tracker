<template>
  <div class="bg-white rounded-lg shadow">
    <div class="px-6 py-5">
      <h3 class="text-lg font-semibold mb-4">Category Breakdown</h3>
      
      <div v-if="loading" class="flex justify-center">
        <LoadingSpinner />
      </div>

      <div v-else class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead>
            <tr class="border-b">
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Category</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Type</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Amount</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="[category, data] in sortedCategories" :key="category" class="border-b">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ formatCategory(category) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-600" :class="data.type === 'income' ? 'text-green-600' : 'text-red-600'">
                {{ data.type.charAt(0).toUpperCase() + data.type.slice(1) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-600" :class="data.type === 'income' ? 'text-green-600' : 'text-red-600'">
                ₺{{ data.amount.toFixed(2) }}
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
import type { CategoryData } from '@/types/transaction';

const props = defineProps<{
  categoryBreakdown?: Record<string, CategoryData>;
  loading?: boolean;
}>();

const sortedCategories = computed(() => {
  if (!props.categoryBreakdown) return [];
  return  Object.entries(props.categoryBreakdown)
  .sort((a, b) => b[0].localeCompare(a[0]));
});

const formatCategory = (category: string) => {
  return category.charAt(0).toUpperCase() + category.slice(1);
};

</script>