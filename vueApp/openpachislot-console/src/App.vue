<template>
  <div id="app" class="font-sans antialiased text-center text-gray-800 mt-15">
    <RegisterForms v-if="!qrGenerated" @registration-success="handleRegistration" class="p-4" />
    <generateQRCode v-if="cardId" :uuid="cardId" :name="name" @qr-generated="qrGenerated = $event" class="p-4">
    </generateQRCode>
    <ItemList ref="itemList" class="p-4" />
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
      cardId: null,
      qrGenerated: false
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