<!-- Fullscreen.vue -->
<template>
  <div class="fullscreen-container" @click="closeFullscreen">
    <div class="image-content" @click.stop>
      <img :src="image.url" :alt="image.title" />
      <div class="image-info">
        <h2>{{ image.title }}</h2>
        <p><strong>Tags:</strong> {{ image.tags.join(', ') }}</p>
        <p><strong>Date:</strong> {{ image.date }}</p>
        <p><strong>Rating:</strong> {{ image.rating }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    image: {
      type: Object,
      required: true,
    },
  },
  methods: {
    closeFullscreen() {
      this.$emit('close-fullscreen');
    },
  },
};
</script>

<style scoped>
.fullscreen-container {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.image-content {
  position: relative;
  background-color: #222;
  padding: 16px;
  border-radius: 4px;
  max-width: 90%;
  max-height: 90%;
  overflow-y: auto;
}

img {
  max-width: 100%;
  max-height: 60vh;
  object-fit: contain;
}

.image-info {
  color: white;
  margin-top: 16px;
}

.image-info h2 {
  margin: 0 0 16px;
  font-size: 24px;
}

.image-info p {
  margin: 8px 0;
}
</style>
