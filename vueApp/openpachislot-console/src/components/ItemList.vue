<template>
  <div class="container mx-auto px-4">
    <h1 class="text-2xl font-bold mb-4">Items List</h1>
    <div class="mb-4">
      <button @click="setSort('id', 1)" class="mr-2 px-4 py-2 bg-blue-500 text-white rounded">Sort by ID (Asc)</button>
      <button @click="setSort('id', -1)" class="mr-2 px-4 py-2 bg-blue-500 text-white rounded">Sort by ID (Desc)</button>
      <button @click="setSort('count', 1)" class="mr-2 px-4 py-2 bg-green-500 text-white rounded">Sort by Count (Asc)</button>
      <button @click="setSort('count', -1)" class="mr-2 px-4 py-2 bg-green-500 text-white rounded">Sort by Count (Desc)</button>
    </div>

    <div v-if="sortedItems && sortedItems.length">
      <table class="w-full text-sm text-left rtl:text-right text-gray-700" >
        <thead class="text-xs text-gray-700 uppercase bg-gray-100">
          <tr>
            <th class="border px-4 py-2 bg-gray-100 text-left" @click="changeSort('id')">ID</th>
            <th class="border px-4 py-2 bg-gray-100 text-left">Name</th>
            <th class="border px-4 py-2 bg-gray-100 text-left" @click="changeSort('count')">Count</th>
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
      items: [],
      sortKey: 'id',
      sortOrder: 1
    };
  },
  computed: {
    sortedItems() {
      return this.items.slice().sort((a, b) => {
        let modifier = this.sortOrder;
        if (a[this.sortKey] < b[this.sortKey]) return -1 * modifier;
        if (a[this.sortKey] > b[this.sortKey]) return 1 * modifier;
        return 0;
      });
    }
  },
  created() {
    this.fetchItems();
  },
  methods: {
    async fetchItems() {
      try {
        const token = this.$store.state.user.token;
        const response = await axios.get('/api/all-items',{
          headers: {
            Authorization: token
          }
        });
        this.items = response.data;
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    },
    setSort(key, order) {
      this.sortKey = key;
      this.sortOrder = order;
    }
  }
};
</script>