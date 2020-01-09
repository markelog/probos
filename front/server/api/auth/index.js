const githubAuth = require('../../middlewares/github');
const callbacks = require('./callbacks');

module.exports = server => {
  server.get('/api/auth/github', githubAuth);
  callbacks(server);
};
