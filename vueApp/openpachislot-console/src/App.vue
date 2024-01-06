<template>
  <div id="app">
    <input v-model="name" type="text" placeholder="Enter name">
    <input v-model="cardId" type="text" placeholder="Enter card ID">
    <button @click="sendData">Send</button>
    <div v-if="responseMessage">
      <h2>Response:</h2>
      <p>{{ responseMessage }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      name: '',
      cardId: '',
      responseMessage: ''
    };
  },
  methods: {
    async sendData() {
      try {
        const payload = {
          name: this.name,
          card_id: this.cardId
        };
        const response = await axios.post(`/api/register`, payload);
        this.responseMessage = response.data;
        // Process the API response here
      } catch (error) {
        console.error('API request failed', error);
      }
    }
  }
};
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
