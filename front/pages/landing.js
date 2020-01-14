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

  return (
    <main>
      <Top />
      <Layout>
        <h1>Hello world</h1>
      </Layout>
    </main>
  );
}

export default Index;
