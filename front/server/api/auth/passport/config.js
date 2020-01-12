const dotenv = require('dotenv');

dotenv.config();

module.exports = {
  github: {
    passReqToCallback: false,
    clientID: process.env.GITHUB_CLIENT_ID,
    clientSecret: process.env.GITHUB_CLIENT_SECRET,
    callbackURL: process.env.CALLBACK_URL
  }
};
