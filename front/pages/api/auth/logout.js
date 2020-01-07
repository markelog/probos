import withPassport from '../../../lib/withPassport';
const redirect = require('micro-redirect');

const Provider = (req, res) => {
  req.logout();
  redirect(res, 302, '/');
};

export default withPassport(Provider);
