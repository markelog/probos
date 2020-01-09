const auth = require('./auth');

module.exports = server => {
  auth(server);
};
