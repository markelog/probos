import React, { useContext } from 'react';
import nextCookie from 'next-cookies';
import redirect from './redirect';
import NextApp, { AppInitialProps, AppContext } from 'next/app';

const IdentityContext = React.createContext(null);

const rootPage = '/';

export const redirectToLogin = ctx => {
  if (
    (ctx && ctx.pathname === rootPage) ||
    (typeof window !== 'undefined' && window.location.pathname === rootPage)
  ) {
    return;
  }

  redirect(ctx, rootPage);
};

// any is needed to use as JSX element
const withIdentity = App => {
  return class IdentityProvider extends React.Component {
    static displayName = `IdentityProvider(MyApp)`;
    static async getInitialProps(ctx) {
      // Get inner app's props
      let appProps;
      if (NextApp.getInitialProps) {
        appProps = await NextApp.getInitialProps(ctx);
      } else {
        appProps = { pageProps: {} };
      }

      const { auth } = nextCookie(ctx.ctx);

      // Redirect to login if page is protected but no session exists
      if (!auth) {
        redirectToLogin(ctx.ctx);
        return Promise.resolve({
          pageProps: null,
          session: null
        });
      }

      const serializedCookie = Buffer.from(auth, 'base64').toString();

      const {
        passport: { user }
      } = JSON.parse(serializedCookie);

      // redirect to login if cookie exists but is empty
      if (!user) {
        redirectToLogin(ctx.ctx);
      }

      const session = user;

      return {
        ...appProps,
        session
      };
    }

    render() {
      const { session, ...appProps } = this.props;

      return (
        <IdentityContext.Provider value={session}>
          <App {...appProps} />
        </IdentityContext.Provider>
      );
    }
  };
};

export default withIdentity;

export const useIdentity = () => useContext(IdentityContext);
