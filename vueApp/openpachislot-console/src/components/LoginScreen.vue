<template>
    <div class="container mx-auto px-4">
        <h1 class="text-2xl font-bold mb-4 text-center">Admin Page</h1>
        <div class="w-full w-[300px] mx-auto">             
            <form @submit.prevent="login" class="flex flex-col items-center gap-4">
                <input type="text" v-model="userId" placeholder="User ID" class="p-2.5 w-4/5 text-sm text-gray-700 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
                <input type="password" v-model="password" placeholder="Password" class="p-2.5 w-4/5 text-sm text-gray-700 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
                <button type="submit" class="mb-2 bg-transparent hover:bg-gray-500 text-gray-700 font-semibold hover:text-white py-2 px-4 border border-gray-500 hover:border-transparent rounded">Login</button>
                <p v-if="loginError" class="text-red-500">Login failed. Please try again.</p>
            </form>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import { mapActions } from 'vuex';

export default {
    data() {
        return {
            userId: '',
            password: '',
            loginError: false
        };
    },
    methods: {
        ...mapActions(['setUser']),
        async login() {
            try {
                const response = await axios.post('/api/login', {
                    id: this.userId,
                    password: this.password
                });
                this.setUser({ token: response.data.token, userId: this.userId });
                this.loginError = false;
                this.userId = '';
                this.password = '';
            } catch (error) {
                this.loginError = true;
                this.password = '';
            }
        }
    }
};
</script>
