const GitHubStrategy = require('passport-github').Strategy;
const passport = require('passport');

const JWT = require('passport-jwt');
const JWTStrategy = JWT.Strategy;
const ExtractJWT = JWT.ExtractJwt;

const cookieExtractor = function(req) {
  var token = null;
  if (req && req.cookies) {
    token = req.cookies['jwt'];
  }
  return token;
};

const strategy = new JWTStrategy(
  {
    jwtFromRequest: cookieExtractor,
    secretOrKey: process.env.JWT_SECRET
  },
  function(user, done) {
    return done(null, user);
  }
);

module.exports = strategy;
