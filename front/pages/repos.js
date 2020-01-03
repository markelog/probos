import React, { useState, useEffect } from "react";

import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";

import { get as getUser } from "../lib/user";

import Layout from "../components/layout";

import Graphs from "../components/graphs";

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

const Repos = function({ user, repository }) {
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

Repos.getInitialProps = async ({ req, res, asPath }) => {
  const user = await getUser(req);

  // Redirect to login if user is not there
  if (user === null) {
    res.writeHead(302, {
      Location: "/api/login"
    });
    res.end();
    return;
  }

  return {
    user,
    repository: asPath.replace("/repos/", "")
  };
};

export default Repos;
