import React, { useState, useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

import Layout from '../components/layout';
import Graphs from '../components/graphs';
import Top from '../components/top';

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

function Index({ user }) {
  const classes = useStyles();

  return (
    <main>
      <Top user={user} />
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
