const jwt = require('jsonwebtoken');

const passport = require('./passport');

module.exports = server => {
  server.get(
    '/api/auth/callback/github',
    passport.authenticate('github', {
      session: false,
      failureRedirect: '/'
    }),
    (req, res) => {
      jwt.sign(
        {
          user: req.user
        },
        process.env.JWT_SECRET,
        (err, token) => {
          // Send Set-Cookie header
          res.cookie('jwt', token, {
            httpOnly: false,
            signed: false,
            maxAge: 24 * 60 * 60 * 1000 // 24 hours
          });

          res.redirect(302, '/');
        }
      );
    }
  );
};
