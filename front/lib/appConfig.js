const isDevelopment = process.env.NODE_ENV !== 'production';
const hostingURL = process.env.HOSTING_URL || 'http://localhost:3000';

const appConfig = {
  isDevelopment,
  hostingURL,
  github: {
    passReqToCallback: false,
    clientID: process.env.GITHUB_CLIENT_ID,
    clientSecret: process.env.GITHUB_CLIENT_SECRET,
    callbackURL: process.env.CALLBACK_URL,
    scope: 'user:email'
  }
};

export default appConfig;
