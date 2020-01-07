import passport from 'passport';
import cookieSession from 'cookie-session';
import url from 'url';
import redirect from 'micro-redirect';
import { github } from './passport';
import { useIdentity } from './withIdentity';
import { Profile } from 'passport-github';
export { default as passport } from 'passport';

passport.use(github);
passport.serializeUser((user, done) => {
  done(null, { username: user.username });
});
passport.deserializeUser((serializedUser, done) => {
  if (!serializedUser) {
    return done(new Error(`User not found: ${serializedUser}`));
  }

  done(null, serializedUser);
});

// export middleware to wrap api/auth handlers
export default fn => (req, res) => {
  if (!res.redirect) {
    // passport.js needs res.redirect:
    // https://github.com/jaredhanson/passport/blob/1c8ede/lib/middleware/authenticate.js#L261
    // Monkey-patch res.redirect to emulate express.js's res.redirect,
    // since it doesn't exist in micro. default redirect status is 302
    // as it is in express. https://expressjs.com/en/api.html#res.redirect
    res.redirect = location => redirect(res, 302, location);
  }

  // Initialize Passport and restore authentication state, if any, from the
  // session. This nesting of middleware handlers basically does what app.use(passport.initialize())
  // does in express.
  cookieSession({
    name: 'auth',
    signed: false,
    domain: url.parse(req.url).host,
    maxAge: 24 * 60 * 60 * 1000 // 24 hours
  })(req, res, () =>
    passport.initialize()(req, res, () =>
      passport.session()(req, res, () => {
        // call wrapped api route as innermost handler
        fn(req, res);
      })
    )
  );
};
