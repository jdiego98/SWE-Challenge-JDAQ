<template>
  <div class="container mx-auto my-10 p-4">
    <div class="flex gap-10">
      <!-- Componente de Búsqueda de Correos -->
      <email-search class="w-" @email-selected="handleEmailSelected"></email-search>
      <!-- Componente de Visualización de Correos -->
      <email-display :email="selectedEmail"></email-display>
    </div>
  </div>
</template>

<script>
import EmailSearch from './components/EmailSearch.vue';
import EmailDisplay from './components/EmailDisplay.vue';

export default {
  components: {
    EmailSearch,
    EmailDisplay
  },
  data() {
    return {
      selectedEmail: null
    };
  },
  methods: {
    handleEmailSelected(email) {
      this.selectedEmail = email;
    }
  }
};
</script>
