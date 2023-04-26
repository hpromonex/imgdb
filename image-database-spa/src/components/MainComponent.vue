<template>
  <div class="main-component d-flex flex-column h-100">
    <header class="bg-dark text-white d-flex align-items-center justify-content-between p-3">
      <h1 class="title m-0">Image Database</h1>
      <SearchBar @search="performSearch" />
    </header>
    <div class="content d-flex flex-grow-1">
      <TagSidePanel :tags="tags" @tag-click="handleTagClick" />
      <div class="center-content d-flex flex-column align-items-center flex-grow-1">
        <ImageGrid :images="images" @image-click="handleImageClick" />
        <PaginationComponent :totalPages="totalPages" :currentPage="currentPage" @page-change="handlePageChange" />
        <FullscreenComponent v-if="showFullscreen" :image="fullscreenImage" @close-fullscreen="showFullscreen = false" />
      </div>
    </div>
  </div>
</template>

<script>
import TagSidePanel from './TagSidePanel.vue';
import ImageGrid from './ImageGrid.vue';
import PaginationComponent from './PaginationComponent.vue';
import FullscreenComponent from './FullscreenComponent.vue';
import SearchBar from './SearchBar.vue';

export default {
  components: {
    TagSidePanel,
    ImageGrid,
    PaginationComponent,
    FullscreenComponent,
    SearchBar,
  },
  data() {
    return {
      tags: [
        // Your list of tags goes here
      ],
      images: [
        // Your image objects go here, each with an id, thumbnailUrl, and title
      ],
      showFullscreen: false,
      fullscreenImage: null,
      totalPages: 10,
      currentPage: 1,
    };
  },
  methods: {
    handleTagClick(clickedTag) {
      console.log('Clicked tag:', clickedTag);
      // Perform your desired action, such as filtering the image grid or performing a new search
    },
    handleImageClick(clickedImage) {
      console.log('Clicked image:', clickedImage);
      this.fullscreenImage = clickedImage;
      this.showFullscreen = true;
    },
    handlePageChange(newPage) {
      console.log('Changed to page:', newPage);
      this.currentPage = newPage;
      // Fetch images for the new page
    },
    performSearch(searchQuery) {
      console.log('Search query:', searchQuery);
      // Perform search with the search query and update the images
    },
  },
};
</script>

<style scoped>
.main-component {
  display: flex;
  flex-direction: column;
  height: 100%;
}

header {
  background-color: #1a1a1a;
  color: white;
  padding: 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.title {
  font-size: 24px;
  margin: 0;
}

.content {
  display: flex;
  flex-grow: 1;
}

.center-content {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
