<template>
    <div class="p-4">
        <div class="flex mb-4">
            <div class="flex-1 mr-2">
                <label for="from" class="block text-sm font-medium text-gray-700">From:</label>
                <input type="text" id="from" v-model="searchCriteria.from"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm">
            </div>
            <div class="flex-1 mx-2">
                <label for="to" class="block text-sm font-medium text-gray-700">To:</label>
                <input type="text" id="to" v-model="searchCriteria.to"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm">
            </div>
            <div class="flex-1 ml-2">
                <label for="subject" class="block text-sm font-medium text-gray-700">Subject:</label>
                <input type="text" id="subject" v-model="searchCriteria.subject"
                    class="mt-1 block w-full border-gray-300 rounded-md shadow-sm">
            </div>
        </div>
        <button @click="searchEmails"
            class="mb-4 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">
            Search
        </button>
        <div v-if="emails && emails.length > 0" class="p-4">
            <table class="min-w-full table-auto border-collapse bg-white">
                <thead>
                    <tr>
                        <th class="border px-4 py-2 text-sm font-medium text-gray-700">From</th>
                        <th class="border px-4 py-2 text-sm font-medium text-gray-700">To</th>
                        <th class="border px-4 py-2 text-sm font-medium text-gray-700">Subject</th>
                        <th class="border px-4 py-2 text-sm font-medium text-gray-700">Actions</th>
                    </tr>
                </thead>
                <tbody class="text-gray-700">
                    <tr v-for="email in emails" :key="email.MessageID">
                        <td class="border px-4 py-2">{{ email.From }}</td>
                        <td class="border px-4 py-2">{{ email.To }}</td>
                        <td class="border px-4 py-2">{{ email.Subject }}</td>
                        <td class="border px-4 py-2">
                            <button @click="$emit('email-selected', email)"
                                class="text-blue-600 hover:text-blue-800 visited:text-purple-600">
                                View
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div v-else class="p-4">
            <p>No emails to display.</p>
        </div>
        <!-- <div class="mt-4">
        <ul>
          <li v-for="email in emails" :key="email.MessageID" @click="$emit('email-selected', email)" class="cursor-pointer hover:bg-gray-200">
            {{ email.Subject }} ({{ email.From }})
          </li>
        </ul>
      </div> -->
    </div>
</template>
  
<script>
import apiService from '@/services/emailService'; // Asegúrate de que la ruta al archivo apiService.js sea correcta

export default {
    data() {
        return {
            searchCriteria: {
                from: '',
                to: '',
                subject: ''
            },
            emails: []
        };
    },
    methods: {
        async searchEmails() {
            this.loading = true;
            this.error = null;

            try {
                // Construye los parámetros que serán enviados en la petición
                const params = {
                    username: 'admin', // Esto debería ser manejado de forma más segura, como con tokens
                    password: 'Complexpass#123', // Nunca expongas contraseñas así en una aplicación real
                };

                // Añade los parámetros de búsqueda si están presentes
                if (this.searchCriteria.from) params.from = this.searchCriteria.from;
                if (this.searchCriteria.to) params.to = this.searchCriteria.to;
                if (this.searchCriteria.subject) params.subject = this.searchCriteria.subject;

                // Hace la llamada a la API y espera por la respuesta
                const response = await apiService.searchEmails(params);
                this.emails = response.data; // Asume que la respuesta de la API es un array de correos
            } catch (error) {
                this.error = 'Error al buscar correos: ' + error.message;
            } finally {
                this.loading = false;
            }
        }

    }
};
</script>
  