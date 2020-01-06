import { Strategy as GithubStrategy } from "passport-github";
import fetch from "isomorphic-unfetch";

import appConfig from "../appConfig";

const API = process.env.API;

// STATICALLY configure the Github strategy for use by Passport.
//
// OAuth 2.0-based strategies require a `verify` function which receives the
// credential (`accessToken`) for accessing the Github API on the user's
// behalf, along with the user's profile.  The function must invoke `cb`
// with a user object, which will be exposed in the request as `req.user`
// in api handlers after authentication.
const strategy = new GithubStrategy(
  appConfig.github,
  (accessToken, refreshToken, profile, cb) => {
    const data = {
      name: profile.displayName,
      username: profile.username,
      email: profile.emails[0].value,
      avatar: profile.photos[0].value,
      provider: profile.provider
    };

    fetch(`${API}/users`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })
      .then(response => response.json())
      .then(response => {
        if (response.status === "failed") {
          throw new Error(response.message);
        }

        cb(null, data);
      })
      .catch(err => {
        console.error("Cannot create a user", err);
        cb(err, data);
      });
  }
);

export default strategy;
