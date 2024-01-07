<template>
  <div id="app">
    <input v-model="name" type="text" placeholder="Enter name">
    <input v-model="cardId" type="text" placeholder="Enter card ID">
    <button @click="sendData">Send</button>
    <div v-if="responseMessage">
      <h2>Response:</h2>
      <p>{{ responseMessage }}</p>
    </div>
    <ItemList ref="itemList" /> <!-- New component for displaying items -->
  </div>
</template>

<script>
import axios from 'axios';
import ItemList from './components/ItemList.vue'; // Importing the new component

export default {
  name: 'App',
  components: {
    ItemList // Registering the new component
  },
  methods: {
    async sendData() {
      try {
        const payload = {
          name: this.name,
          card_id: this.cardId
        };
        await axios.post(`/api/register`, payload);
        this.$refs.itemList.fetchItems();
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
