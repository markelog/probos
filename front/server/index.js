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
    if (req.user === undefined) {
      app.render(req, res, '/landing/index');
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

  server.get('/repos/github.com/:org/:name', jwtAuth, (req, res) => {
    const { org, name } = req.params;

    return app.render(req, res, '/repo', {
      repository: `github.com/${org}/${name}`,
      user: req.user
    });
  });

  server.get(
    '/repos/github.com/:org/:name/branch/:branch',
    jwtAuth,
    (req, res) => {
      const { org, name, branch } = req.params;
      return app.render(req, res, '/repo', {
        repository: `github.com/${org}/${name}`,
        branch,
        user: req.user
      });
    }
  );

  server.all('*', (req, res) => {
    return handle(req, res);
  });

  server.listen(port, err => {
    if (err) {
      throw err;
    }

    console.log(`> Ready on http://localhost:${port}`);
  });
});
