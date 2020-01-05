import React, { useState, useEffect } from "react";

import { makeStyles } from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import { useIdentity } from "../lib/withIdentity";

import Layout from "../components/layout";

import Graphs from "../components/graphs";

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  }
}));

function Index(data) {
  const classes = useStyles();
  const identity = useIdentity();

  return (
    <main>
      <h1>{JSON.stringify(identity)}</h1>
      <Link href="/api/auth/logout" color="inherit" className={classes.link}>
        logout
      </Link>
      <Layout>
        <Graphs repository="github.com/markelog/adit" branch="master" />
      </Layout>
    </main>
  );
}

export default Index;
