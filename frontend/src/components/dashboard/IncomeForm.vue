<template>
    <form @submit.prevent="handleSubmit" class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">
            Amount
          </label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <span class="text-gray-500 sm:text-sm">₺</span>
            </div>
            <input
              v-model="form.amount"
              type="number"
              step="0.01"
              :class="[
                'block w-full pl-7 pr-12 py-3 rounded-lg border transition-shadow duration-200',
                loading ? 'bg-gray-50' : 'bg-white',
                'focus:ring-2 focus:ring-green-500 focus:border-green-500',
                'border-gray-300 shadow-sm'
              ]"
              :disabled="loading"
            >
          </div>
        </div>
  
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">
            Description
          </label>
          <input
            v-model="form.description"
            type="text"
            placeholder="Brief description"
            :class="[
              'block w-full px-4 py-3 rounded-lg border transition-shadow duration-200',
              loading ? 'bg-gray-50' : 'bg-white',
              'focus:ring-2 focus:ring-green-500 focus:border-green-500',
              'border-gray-300 shadow-sm'
            ]"
            :disabled="loading"
          >
        </div>
  
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">
            Source
          </label>
          <div class="relative">
            <select
              v-model="form.source"
              :class="[
                'block w-full px-4 py-3 rounded-lg border appearance-none transition-shadow duration-200',
                loading ? 'bg-gray-50' : 'bg-white',
                'focus:ring-2 focus:ring-green-500 focus:border-green-500',
                'border-gray-300 shadow-sm',
                'cursor-pointer'
              ]"
              :disabled="loading"
            >
              <option value="" disabled>Select source</option>
              <option v-for="source in sources" :key="source.value" :value="source.value">
                {{ source.label }}
              </option>
            </select>
            <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-500">
              <svg class="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                <path d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" />
              </svg>
            </div>
          </div>
        </div>
  
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-700">
            Date
          </label>
          <input
            v-model="form.date"
            type="date"
            :class="[
              'block w-full px-4 py-3 rounded-lg border transition-shadow duration-200',
              loading ? 'bg-gray-50' : 'bg-white',
              'focus:ring-2 focus:ring-green-500 focus:border-green-500',
              'border-gray-300 shadow-sm'
            ]"
            :disabled="loading"
          >
        </div>
      </div>
  
      <div class="flex items-center justify-end space-x-4 pt-2">
        <button
          type="button"
          class="px-4 py-2 text-sm font-medium text-gray-700 hover:text-gray-800 focus:outline-none"
          @click="resetForm"
        >
          Clear
        </button>
        <button
          type="submit"
          :disabled="loading || !isValid"
          :class="[
            'px-6 py-2.5 rounded-lg text-sm font-medium transition-all duration-200',
            'focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500',
            loading || !isValid 
              ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
              : 'bg-green-600 text-white hover:bg-green-700 shadow-sm hover:shadow'
          ]"
        >
          <div class="flex items-center gap-2">
            <LoadingSpinner v-if="loading" size="small" />
            <span>Add Income</span>
          </div>
        </button>
      </div>
    </form>
  </template>
  
<script setup lang="ts">
import { ref, computed } from 'vue';
import type { TransactionFormData } from '@/types/forms';
import LoadingSpinner from '../common/LoadingSpinner.vue';

const props = defineProps<{
  loading?: boolean;
}>();

const emit = defineEmits<{
  (e: 'submit', data: TransactionFormData): void;
}>();

const sources = [
  { value: 'salary', label: 'Salary' },
  { value: 'freelance', label: 'Freelance' },
  { value: 'investments', label: 'Investments' },
  { value: 'business', label: 'Business' },
  { value: 'other', label: 'Other' }
];

const form = ref({
  amount: '',
  description: '',
  source: '',
  date: new Date().toISOString().split('T')[0]
});

const isValid = computed(() => {
  return (
    form.value.amount &&
    form.value.description &&
    form.value.source &&
    form.value.date
  );
});

const handleSubmit = () => {
  if (!isValid.value) return;
  
  // Create a date with time set to 00:00:00
  const dateWithTime = new Date(form.value.date);
  dateWithTime.setHours(0, 0, 0, 0);

  emit('submit', {
    ...form.value,
    date: dateWithTime.toISOString() // This will include time as 00:00:00
  });

  form.value = {
    amount: '',
    description: '',
    source: '',
    date: new Date().toISOString().split('T')[0]
  };
};</script>