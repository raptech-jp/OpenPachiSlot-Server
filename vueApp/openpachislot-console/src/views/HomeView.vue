<template>
    <div class="font-sans antialiased text-center text-gray-800 mt-15 user-select-none">
      <RegisterForms v-if="!qrGenerated" @registration-success="handleRegistration" class="p-4" />
      <generateQRCode v-if="cardId" :uuid="cardId" :name="name" @qr-generated="qrGenerated = $event" class="p-4 ">
      </generateQRCode>
    </div>
  </template>
  
  <script>
  import RegisterForms from '../components/RegisterForms.vue';
  import generateQRCode from '../components/GenerateQRCode.vue';
  
  export default {
    name: 'App',
    components: {
      RegisterForms,
      generateQRCode
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
      }
    }
  };
  </script>
  