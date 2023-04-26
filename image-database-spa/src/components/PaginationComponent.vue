<template>
  <div class="pagination-container">
    <button @click="previousPage" :disabled="currentPage === 1">Previous</button>
    <span>Page {{ currentPage }} of {{ totalPages }}</span>
    <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
  </div>
</template>

<script>
export default {
  props: {
    totalPages: {
      type: Number,
      required: true,
    },
    currentPage: {
      type: Number,
      required: true,
    },
  },
  methods: {
    previousPage() {
      if (this.currentPage > 1) {
        this.$emit('update-page', this.currentPage - 1);
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.$emit('update-page', this.currentPage + 1);
      }
    },
  },
};
</script>

<style scoped>
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 16px;
}

button {
  background-color: #4a4a4a;
  color: white;
  border: none;
  padding: 8px 16px;
  margin: 0 8px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #3a3a3a;
}

button:disabled {
  background-color: #2a2a2a;
  cursor: not-allowed;
}

span {
  color: white;
}
</style>
