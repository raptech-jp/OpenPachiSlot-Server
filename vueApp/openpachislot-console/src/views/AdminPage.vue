<template>
  <div>
    <LoginScreen v-if="!isAuthenticated" />
    <ItemList v-else class="p-4" />
  </div>
</template>

<script>
import LoginScreen from '../components/LoginScreen.vue';
import ItemList from '../components/ItemList.vue';
import { mapGetters } from 'vuex';

export default {
  components: {
    LoginScreen,
    ItemList
  },
  computed: {
    ...mapGetters(['isAuthenticated'])
  }
};
</script>
  