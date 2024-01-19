<template>
    <div class="absolute inset-0 flex items-center justify-center px-6 py-8 mx-auto lg:py-0">
        <div class="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
            <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
                    Sign in to your account
                </h1>
                <form @submit.prevent="login" class="space-y-4 md:space-y-6">
                    <div>
                        <label for="text" class="block mb-2 text-sm font-medium text-gray-900">Account</label>
                        <input type="text" v-model="userId" placeholder="Account"
                            class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5">
                    </div>
                    <div>
                        <label for="password" class="block mb-2 text-sm font-medium text-gray-900">Password</label>
                        <input type="password" v-model="password" placeholder="••••••••"
                            class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5">
                    </div>
                    <button type="submit"
                        class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center">Sign
                        in</button>

                    <p v-if="loginError" class="text-red-500">Login failed. Please try again.</p>
                </form>
            </div>
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
                const token = response.data.token;
                this.setUser({ token: token, userId: this.userId });

                localStorage.setItem('token', token);
                localStorage.setItem('userId', this.userId);

                this.loginError = false;
            } catch (error) {
                this.loginError = true;
                this.password = '';
            }
        }
    }
};
</script>
