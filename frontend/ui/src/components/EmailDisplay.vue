<template>
  <div v-if="email" class="p-4">
    <div class="flex mb-2">
      <div class="flex-1 mr-2">
        <label class="block text-sm font-medium text-gray-700">From:</label>
        <div class="mt-1 block w-full bg-gray-100 rounded-md shadow-sm">{{ email.From }}</div>
      </div>
      <div class="flex-1 mx-2">
        <label class="block text-sm font-medium text-gray-700">To:</label>
        <div class="mt-1 block w-full bg-gray-100 rounded-md shadow-sm">{{ email.To }}</div>
      </div>
      <div class="flex-1 ml-2">
        <label class="block text-sm font-medium text-gray-700">Subject:</label>
        <div class="mt-1 block w-full bg-gray-100 rounded-md shadow-sm">{{ email.Subject }}</div>
      </div>
    </div>
    <div class="mt-4 bg-gray-100 rounded-md shadow-sm p-2">
      <p>{{ email.Body }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    email: Object
  }
};
</script>
