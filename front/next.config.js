const dotenv = require("dotenv");
dotenv.config();

module.exports = {
  env: {
    API: process.env.API,

    CALLBACK_URL:
      process.env.CALLBACK_URL || "http://localhost:3000/api/auth/callback",
    REDIRECT_URI:
      process.env.REDIRECT_URI || "http://localhost:3000/api/callback",
    POST_LOGOUT_REDIRECT_URI:
      process.env.POST_LOGOUT_REDIRECT_URI || "http://localhost:3000/",
    SESSION_COOKIE_SECRET: process.env.SESSION_COOKIE_SECRET,
    SESSION_COOKIE_LIFETIME: 7200 // 2 hours
  }
};
