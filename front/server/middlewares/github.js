const passport = require('../api/auth/passport');

module.exports = passport.authenticate('github', {
  failureRedirect: '/'
});
