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
      axios.post('/api/register', { name: this.name, card_id: this.cardId })
        .then(response => {
          this.responseMessage = response.data.message; // 登録成功時のメッセージ
          this.registrationSuccessful = true; // 入力フィールドを非表示にする
          this.$refs.itemList.fetchItems();
        })
        .catch(error => {
          this.responseMessage = error.response.data.error; // エラーメッセージ
        });
    },
  },
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
