<template>
  <div class="container mx-auto px-4">
    <h1 class="text-2xl font-bold mb-4">Items List</h1>
    <div v-if="sortedItems && sortedItems.length">
      <table class="w-full text-sm text-left rtl:text-right text-gray-700" >
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700">
          <tr>
            <th class="border px-4 py-2 bg-gray-100 text-left">ID</th>
            <th class="border px-4 py-2 bg-gray-100 text-left">Name</th>
            <th class="border px-4 py-2 bg-gray-100 text-left">Count</th>
            <th class="border px-4 py-2 bg-gray-100 text-left">Card ID</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in sortedItems" :key="item.id">
            <td class="border px-4 py-2">{{ item.id }}</td>
            <td class="border px-4 py-2">{{ item.name }}</td>
            <td class="border px-4 py-2">{{ item.count }}</td>
            <td class="border px-4 py-2">{{ item.card_id }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>


<script>
import axios from 'axios';
export default {
  data() {
    return {
      items: []
    };
  },
  computed: {
    sortedItems() {
      // IDを昇順にソートした新しい配列を返す
      return this.items.slice().sort((a, b) => a.id - b.id);
    }
  },
  created() {
    this.fetchItems();
  },
  methods: {
    async fetchItems() {
      try {
        const response = await axios.get('/api/all-items');
        this.items = response.data;
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    }
  },
  mounted() {
    this.fetchItems(); // コンポーネントのマウント時にアイテムを取得
  }
};
</script>