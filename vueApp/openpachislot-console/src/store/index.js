import { createStore } from 'vuex';

export default createStore({
    state: {
    user: {
      token: null,
      userId: null
    }
  },
  mutations: {
    setUser(state, payload) {
      state.user.token = payload.token;
      state.user.userId = payload.userId;
    }
  },
  actions: {
    setUser({ commit }, payload) {
      commit('setUser', payload);
    }
  },
  getters: {
    isAuthenticated(state) {
      return !!state.user.token;
    }
  }
});

