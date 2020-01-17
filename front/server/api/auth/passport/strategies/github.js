const GitHubStrategy = require('passport-github').Strategy;
const fetch = require('isomorphic-unfetch');
const Octokit = require('@octokit/rest');

const config = require('../config');

const { API } = process.env;

async function getRepos(token) {
  const perPage = 100;

  const octokit = new Octokit({
    auth() {
      return `token ${token}`;
    }
  });

  let result = [];

  async function get(page) {
    const repos = await octokit.repos.list({
      page,
      per_page: perPage
    });

    if (repos.status != 200) {
      return;
    }

    const { data } = repos;
    const names = data.map(repo => repo.full_name);

    result = result.concat(names);

    if (data.length === perPage) {
      await get(page + 1);
    }
  }

  await get(1);
  return result;
}

// STATICALLY configure the Github strategy for use by Passport.
//
// OAuth 2.0-based strategies require a `verify` function which receives the
// credential (`accessToken`) for accessing the Github API on the user's
// behalf, along with the user's profile.  The function must invoke `cb`
// with a user object, which will be exposed in the request as `req.user`
// in api handlers after authentication.
const strategy = new GitHubStrategy(
  config.github,
  async (accessToken, refreshToken, profile, cb) => {
    getRepos(accessToken).then((repositories) => {
      const data = {
        name: profile.displayName,
        username: profile.username,
        email: profile.emails[0].value,
        avatar: profile.photos[0].value,
        provider: profile.provider,
        repositories
      };

      console.log(accessToken);

      fetch(`${API}/users`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      })
        .then(response => response.json())
        .then((response) => {
          if (response.status === 'failed') {
            throw new Error(response.message);
          }

          cb(null, data.username);
        })
        .catch((err) => {
          console.error('Cannot create a user', err);
          cb(err, false);
        });
    });
  }
);

module.exports = strategy;
