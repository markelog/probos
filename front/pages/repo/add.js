import React, { useState, useEffect } from 'react';
import Octokit from '@octokit/rest';

import { makeStyles } from '@material-ui/core/styles';

const octokit = new Octokit();

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

const getRepos = async (username) => {
  const repos = await octokit.repos.listForUser({
    per_page: 100,
    direction: 'desc',
    username
  });

  return repos.data;
};

const Add = function ({ user, repos }) {
  return <h1>{JSON.stringify(repos)}</h1>;
};

Add.getInitialProps = async ({ query }) => {
  const { user } = query;
  const repos = await getRepos(user.username);

  return {
    user,
    repos
  };
};

export default Add;
