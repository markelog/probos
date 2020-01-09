const passport = require('passport');

const github = require('../passport/strategies/github');
const jwt = require('../passport/strategies/jwt');

passport.use(github);
passport.use(jwt);

module.exports = passport;
