import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';

import { makeStyles } from '@material-ui/core/styles';

import Layout from '../../components/layout';
import Repos from '../../components/repos';
import Top from '../../components/top';

import Background from './background.jpg';

const useStyles = makeStyles(theme => ({
  main: {
    // backgroundImage: `url(${Background})`,
    height: '100%'
  },

  ship: {
    fontSize: '2rem',
    fontFamily:
      'SFMono-Regular,Menlo,Monaco,Consolas,"Liberation Mono","Courier New",monospace'
  },
  link: {
    margin: theme.spacing(1)
  }
}));

const ship = `
               |    |    |
             )_)  )_)  )_)
            )___))___))___)\
           )____)____)_____)\\
         _____|____|____|____\\\__
---------\                   /---------
  ^^^^^ ^^^^^^^^^^^^^^^^^^^^^
    ^^^^      ^^^^     ^^^    ^^
         ^^^^      ^^^
`;

function Index({ user }) {
  const classes = useStyles();

  return (
    <main className={classes.main}>
      <Top />
      <Layout>
        <pre className={classes.ship}>{ship}</pre>
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
