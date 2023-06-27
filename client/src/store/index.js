import { createStore } from 'vuex';

const store = createStore({
  state() {
    return {
      counter: 0,
    };
  },
  mutations: {
    increment(state) {
      state.counter++;
    },
    decrement(state) {
      state.counter--;
    },
  },
  actions: {
    incrementAsync(context) {
      setTimeout(() => {
        context.commit('increment');
      }, 1000);
    },
  },
  getters: {
    counterSquared(state) {
      return state.counter ** 2;
    },
  },
});

export default store;