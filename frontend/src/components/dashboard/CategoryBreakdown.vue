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
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Category
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Amount
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="[category, amount] in sortedCategories" :key="category">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ formatCategory(category) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-right text-red-600">
                ${{ amount.toFixed(2) }}
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

const props = defineProps<{
  categoryTotals?: Record<string, number>;
  loading?: boolean;
}>();

const sortedCategories = computed(() => {
  if (!props.categoryTotals) return [];
  return Object.entries(props.categoryTotals)
    .sort((a, b) => b[1] - a[1]);
});

const formatCategory = (category: string) => {
  return category.charAt(0).toUpperCase() + category.slice(1);
};</script>