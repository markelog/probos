import React, { useState, useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

import Layout from '../components/layout';

import Graphs from '../components/graphs';

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

function Index({ user }) {
  const classes = useStyles();
  const SignIn = () => (
    <p>
      <a href="/api/auth/github">Sign in with github</a>
    </p>
  );

  const User = () => {
    return (
      <>
        <h1>{user}</h1>
        <Link href="/api/auth/logout" color="inherit" className={classes.link}>
          logout
        </Link>
      </>
    );
  };

  return (
    <main>
      {user ? <User /> : <SignIn />}

      <Layout>
        <Graphs repository="github.com/markelog/adit" branch="master" />
      </Layout>
    </main>
  );
}

Index.getInitialProps = ({ query }) => {
  return {
    user: query.user
  };
};

export default Index;
