<template>
  <div v-if="loading">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div v-for="i in 3" :key="i" class="bg-white rounded-lg shadow p-6">
        <LoadingSpinner size="small" />
      </div>
    </div>
  </div>

  <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <div
      v-for="card in cards"
      :key="card.title"
      class="bg-white rounded-lg shadow"
    >
      <div class="px-6 py-5">
        <h3 class="text-lg font-semibold">{{ card.title }}</h3>
        <p :class="['text-2xl font-bold mt-2', card.colorClass]">
          ${{ card.amount.toFixed(2) }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { Summary } from '@/types/transaction';
import LoadingSpinner from '../common/LoadingSpinner.vue';

const props = defineProps<{
  summary: Summary | null;
  loading?: boolean;
}>();

const cards = computed(() => {
  if (!props.summary) return [];

  return [
    {
      title: 'Total Income',
      amount: props.summary.totalIncome,
      colorClass: 'text-green-600'
    },
    {
      title: 'Total Expenses',
      amount: props.summary.totalExpenses,
      colorClass: 'text-red-600'
    },
    {
      title: 'Balance',
      amount: props.summary.remainingAmount,
      colorClass: props.summary.remainingAmount >= 0 ? 'text-green-600' : 'text-red-600'
    }
  ];
});</script>