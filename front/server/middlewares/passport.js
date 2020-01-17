const passport = require('../api/auth/passport');

module.exports = (server) => {
  server.use(passport.initialize());
  server.use(passport.session());
};
