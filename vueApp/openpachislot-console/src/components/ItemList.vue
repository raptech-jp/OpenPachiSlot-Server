<template>
  <div>
    <h1>Items List</h1>
    <table v-if="items.length">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Count</th>
          <th>Card ID</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>{{ item.id }}</td>
          <td>{{ item.name }}</td>
          <td>{{ item.count }}</td>
          <td>{{ item.card_id }}</td>
        </tr>
      </tbody>
    </table>
    <p v-else>No items found.</p>
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
  }
};
</script>

<style>
/* 必要に応じてスタイルを追加 */
table {
  width: 80%;
  margin: auto;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  text-align: left;
  background-color: #f2f2f2;
}
</style>