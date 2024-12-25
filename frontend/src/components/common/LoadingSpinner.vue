<template>
  <div 
    :class="[
      'flex items-center justify-center',
      fullScreen ? 'fixed inset-0 bg-white/80' : 'p-4'
    ]"
  >
    <div 
      :class="[
        'animate-spin rounded-full border-t-blue-500 border-r-blue-500 border-b-transparent border-l-transparent',
        sizeClasses[size]
      ]"
    />
  </div>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
  size?: 'small' | 'medium' | 'large'
  fullScreen?: boolean
}>(), {
  size: 'medium',
  fullScreen: false
});

const sizeClasses = {
  small: 'h-4 w-4 border-2',
  medium: 'h-8 w-8 border-2',
  large: 'h-12 w-12 border-3',
};</script>