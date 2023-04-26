<!-- ImageGrid.vue -->
<template>
  <div class="image-grid">
    <div
      v-for="image in images"
      :key="image.id"
      class="image-thumbnail"
      @click="onImageClick(image)"
    >
      <img :src="image.thumbnailUrl" :alt="image.title" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    images: {
      type: Array,
      required: true,
    },
  },
  methods: {
    onImageClick(image) {
      this.$emit('image-click', image);
    },
  },
};
</script>

<style scoped>
.image-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  grid-gap: 16px;
  padding: 16px;
}

.image-thumbnail {
  cursor: pointer;
  overflow: hidden;
  position: relative;
}

.image-thumbnail img {
  width: 100%;
  height: auto;
  object-fit: cover;
  transition: transform 0.3s;
}

.image-thumbnail:hover img {
  transform: scale(1.1);
}
</style>
