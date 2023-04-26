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
      <button class="btn btn-outline-secondary border-0 rounded-circle" type="button" @click="submitSearch">
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
  position: relative;
}

.search-bar .form-control {
  padding-right: 40px;
  border-radius: 50px;
}

.search-bar .input-group-append {
  position: absolute;
  right: 8px;
  top: 50%;
  transform: translateY(-50%);
}
</style>
