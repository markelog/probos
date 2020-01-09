const passport = require('../api/auth/passport');

module.exports = function jwt(req, res, next) {
  passport.authenticate('jwt', (err, user, info) => {
    if (err) {
      return next(err);
    }
    req.user = user;
    next();
  })(req, res, next);
};
