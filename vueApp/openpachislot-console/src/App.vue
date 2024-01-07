<template>
  <div id="app">
    <input v-model="name" type="text" placeholder="Enter name">
    <button @click="sendData">Register</button>
    <generateQRCode v-if="cardId" :uuid="cardId"></generateQRCode> <!-- UUID を渡し、cardId がある場合にのみ表示 -->
    <ItemList ref="itemList" />
  </div>
</template>

<script>
import axios from 'axios';
import ItemList from './components/ItemList.vue'; // Importing the new component
import generateQRCode from './components/GenerateQRCode.vue'; // Importing the new component

export default {
  name: 'App',
  components: {
    ItemList, // Registering the new component
    generateQRCode
  },
  data() {
    return {
      name: '',
      cardId: '',
      responseMessage: '',
      registrationSuccessful: false,
    };
  },
  methods: {
    sendData() {
      axios.post('/api/register', { name: this.name })
        .then(response => {
          this.cardId = response.data.cardId; // API からのレスポンスで cardId を取得
          this.responseMessage = `Registered! Card ID: ${this.cardId}`;
          this.registrationSuccessful = true;
          this.$refs.itemList.fetchItems();
        })
        .catch(error => {
          this.responseMessage = error.response.data.error;
        });
    },
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