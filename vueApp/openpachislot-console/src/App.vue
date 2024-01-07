<template>
  <div id="app">
    <RegisterForms @registration-success="handleRegistration" />
    <generateQRCode v-if="cardId" :uuid="cardId" :name="name"></generateQRCode>
    <ItemList ref="itemList" />
  </div>
</template>

<script>
import RegisterForms from './components/RegisterForms.vue';
import ItemList from './components/ItemList.vue';
import generateQRCode from './components/GenerateQRCode.vue';

export default {
  name: 'App',
  components: {
    RegisterForms,
    generateQRCode,
    ItemList
  },
  data() {
    return {
      name: '',
      cardId: null
    };
  },
  methods: {
    handleRegistration(data) {
      this.name = data.name;
      this.cardId = data.cardId;
      this.refreshItems();
    },
    refreshItems() {
      if (this.$refs.itemList) {
        this.$refs.itemList.fetchItems();
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