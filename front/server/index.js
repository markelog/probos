const express = require('express');
const next = require('next');
const cookie = require('cookie-parser');
const dotenv = require('dotenv');

const jwtAuth = require('./middlewares/jwt');
const passport = require('./middlewares/passport');

const api = require('./api');

const port = parseInt(process.env.PORT, 10) || 3000;
const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle = app.getRequestHandler();

app.prepare().then(() => {
  const server = express();

  server.use(cookie());
  server.use(express.urlencoded({ extended: true }));

  api(server);
  passport(server);

  server.get('/', jwtAuth, (req, res) => {
    if (req.user === null) {
      app.render(req, res, '/landing');
      return;
    }

    app.render(req, res, '/profile', {
      user: req.user
    });
  });

  server.get('/api/auth/logout', (req, res) => {
    req.logout();
    res.clearCookie('jwt');
    res.redirect(302, '/');
  });

  server.get(/repos\/(.*)/, jwtAuth, (req, res) => {
    return app.render(req, res, '/repo/list', {
      user: req.user
    });
  });

  server.get('/repo/add', jwtAuth, (req, res) => {
    return app.render(req, res, '/repo/add', {
      user: req.user
    });
  });

  server.all('*', (req, res) => {
    return handle(req, res);
  });

  server.listen(port, err => {
    if (err) throw err;
    console.log(`> Ready on http://localhost:${port}`);
  });
});
