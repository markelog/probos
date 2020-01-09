const passport = require('passport');

const github = require('../passport/strategies/github');
const jwt = require('../passport/strategies/jwt');

passport.use(github);
passport.use(jwt);

passport.serializeUser((user, done) => {
  done(null, { username: user.username });
});
passport.deserializeUser((serializedUser, done) => {
  if (!serializedUser) {
    return done(new Error(`User not found: ${serializedUser}`));
  }

  done(null, serializedUser);
});

module.exports = passport;
