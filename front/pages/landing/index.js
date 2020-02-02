import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';

import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

import Layout from '../../components/layout';
import Repos from '../../components/repos';
import Top from '../../components/top';

import Background from './background.jpg';

const useStyles = makeStyles(theme => ({
  main: {
    backgroundImage: `url(${Background})`,
    height: '100%'
  },
  link: {
    margin: theme.spacing(1)
  }
}));

function Index({ user }) {
  const classes = useStyles();

  return (
    <main className={classes.main}>
      <Top />
      <Layout>
        <h1>Hello world</h1>
        <style>{`
          html,
          body,
          #__next {
            height: 100%;
          }
        `}</style>
      </Layout>
    </main>
  );
}

export default Index;
