<template>
    <div class="flex items-center px-3 py-2 rounded-lg bg-gray-50">
        <input v-model="name" type="text" placeholder="Enter name" class="block mx-4 p-2.5 w-full text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
        <button @click="sendData">Register</button>
        <p v-if="showWarning" class="text-red-500">名前を入力してください。</p>
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
            showWarning: false, // 警告を表示するためのフラグ
        };
    },
    methods: {
        sendData() {
            if (this.name.trim() === '') {
                // テキストボックスが空の場合、警告を表示して送信しない
                this.showWarning = true;
                return;
            }

            axios.post('/api/user', { name: this.name })
                .then(response => {
                    this.cardId = response.data.cardId;
                    this.registrationSuccessful = true;
                    this.$emit('registration-success', { cardId: this.cardId, name: this.name});
                    this.name = ''; // 成功後にテキストボックスをクリア
                    this.showWarning = false;
                })
                .catch(error => {
                    this.responseMessage = error.response.data.error;
                });
        },
    }
};
</script>
