<template>
    <div>
        <input v-model="name" type="text" placeholder="Enter name">
        <button @click="sendData">Register</button>
    </div>
</template>
  
<script>
import axios from 'axios';
export default {
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
                    this.cardId = response.data.cardId;
                    this.registrationSuccessful = true;
                    this.$emit('registration-success', { cardId: this.cardId, name: this.name});
                    this.name = ''; // 成功後にテキストボックスをクリア
                })
                .catch(error => {
                    this.responseMessage = error.response.data.error;
                });
        },
    }
};
</script>