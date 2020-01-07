import withPassport, { passport } from '../../../lib/withPassport';

const Provider = (req, res) => {
  const { provider } = req.query;
  if (!provider) {
    return { statusCode: 404 };
  }

  passport.authenticate(provider)(req, res, (...args) => {
    console.log('passport authenticated', args);
  });
};

export default withPassport(Provider);
