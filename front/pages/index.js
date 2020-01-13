import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';

import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

import Layout from '../components/layout';
import Repos from '../components/repos';
import Top from '../components/top';

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

function Index({ user }) {
  const classes = useStyles();
  const router = useRouter();

  useEffect(() => {
    const { page } = router.query;

    if (page !== '') {
      return;
    }

    // Always do navigations after the first render
    router.push('/?page=1', '/?page=1', { shallow: true });
  }, []);

  return (
    <main>
      <Top user={user} />
      <Layout>
        <Repos username={user.username} page={1} />
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
