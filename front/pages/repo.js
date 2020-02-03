import React, { useState, useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

import Layout from '../components/layout';
import Graphs from '../components/graphs';
import Top from '../components/Top';

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

const Repo = function(data) {
  const { user, repository } = data;
  const classes = useStyles();

  return (
    <>
      <Top user={user} />
      <Layout user={user}>
        <Graphs repository={repository} branch="master" />
      </Layout>
    </>
  );
};

Repo.getInitialProps = ({ query, asPath }) => {
  return {
    user: query.user,
    repository: asPath.replace('/repos/', '')
  };
};

export default Repo;
