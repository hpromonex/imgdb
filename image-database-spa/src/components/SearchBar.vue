<!-- SearchBar.vue -->
<template>
  <div class="input-group search-bar">
    <input
      type="text"
      class="form-control"
      placeholder="Search images..."
      v-model="searchQuery"
      @keyup.enter="submitSearch"
    />
    <div class="input-group-append">
      <button class="btn btn-light border-0 search-btn" type="button" @click="submitSearch">
        <font-awesome-icon icon="search" />
      </button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
    };
  },
  methods: {
    submitSearch() {
      this.$emit('search', this.searchQuery);
    },
  },
};
</script>

<style scoped>
.search-bar {
  width: 100%;
  max-width: 400px;
}

.search-bar .form-control {
  border-top-left-radius: 50px;
  border-bottom-left-radius: 50px;
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.search-bar .input-group-append {
  margin-left: -1px;
}

.search-bar .search-btn {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  border-top-right-radius: 50px;
  border-bottom-right-radius: 50px;
  transition: background-color 0.3s;
}

.search-bar .search-btn:hover {
  background-color: #343a40;
  color: white;
}
</style>
