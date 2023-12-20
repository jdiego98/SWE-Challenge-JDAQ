import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080'
});

const emailService = {
  searchEmails({ from, to, subject }) {
    // Inicializar los parámetros con las credenciales
    const params = {
      username: "admin",
      password: "Complexpass#123",
      // Agrega aquí los demás parámetros que siempre se incluyen
    };

    // Añadir condicionalmente los parámetros de búsqueda
    if (from) {
      params.From = from;
    }
    if (to) {
      params.To = to;
    }
    if (subject) {
      params.Subject = subject;
    }

    // Realizar la llamada a la API con los parámetros
    return apiClient.get('/api/search', { params });
  },
};

export default emailService;