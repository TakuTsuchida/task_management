export default {
  setUser(state, userData) {
    state.userId = userData.userId;
    state.password = userData.password;
    state.token = userData.token;
  },
};
