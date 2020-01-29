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

const List = function(data) {
  const { user, repository } = data;
  const classes = useStyles();

  return (
    <>
      <Link href="/api/logout" color="inherit" className={classes.link}>
        logout
      </Link>
      <Layout user={user}>
        <Graphs repository={repository} branch="master" />
      </Layout>
    </>
  );
};

List.getInitialProps = ({ asPath }) => {
  return {
    repository: asPath.replace('/repos/', '')
  };
};

export default List;
