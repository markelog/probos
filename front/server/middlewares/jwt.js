const fetch = require('isomorphic-unfetch');

const passport = require('../api/auth/passport');

const { API } = process.env;

module.exports = function jwt(req, res, next) {
  passport.authenticate('jwt', (err, user, info) => {
    if (err) {
      return next(err);
    }

    if (user === false) {
      return next();
    }

    fetch(`${API}/users/${user.user}`)
      .then(response => response.json())
      .then(response => {
        const { payload } = response;
        if (Object.keys(payload).length === 0) {
          req.user = null;
        } else {
          req.user = response.payload;
        }
        next();
      })
      .catch(next);
  })(req, res, next);
};
