<!-- TagSidePanel.vue -->
<template>
  <div class="tag-side-panel">
    <h3>Tags</h3>
    <ul class="tag-list">
      <li v-for="tag in tags" :key="tag" @click="onTagClick(tag)">
        {{ tag }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    tags: {
      type: Array,
      required: true,
    },
  },
  methods: {
    onTagClick(tag) {
      this.$emit('tag-click', tag);
    },
  },
};
</script>

<style scoped>
.tag-side-panel {
  background-color: #2a2a2a;
  color: white;
  padding: 16px;
  width: 200px;
  height: 100%;
  overflow-y: auto;
}

h3 {
  margin-top: 0;
  margin-bottom: 16px;
}

.tag-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.tag-list li {
  cursor: pointer;
  padding: 8px;
  transition: background-color 0.3s;
}

.tag-list li:hover {
  background-color: #1a1a1a;
}
</style>
