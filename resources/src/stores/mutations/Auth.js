export default {
  setUser(state, userData) {
    state.userName = userData.userName;
    state.password = userData.password;
    state.token = userData.token;
  },
  removeUser(state){
    state.userName = null;
    state.password = null;
    state.token = null;
  },
};
