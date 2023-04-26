import { createStore } from 'vuex';

export default createStore({
  state: {
    images: [],
    currentPage: 1,
  },
  mutations: {
    SET_IMAGES(state, images) {
      state.images = images;
    },
    SET_CURRENT_PAGE(state, page) {
      state.currentPage = page;
    },
  },
  actions: {
    // Add actions here
  },
  getters: {
    // Add getters here
  },
});
