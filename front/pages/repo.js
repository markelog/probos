import React, { useState, useEffect } from 'react';

import Link from '@material-ui/core/Link';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';

import { makeStyles } from '@material-ui/core/styles';

import Layout from '../components/layout';
import Graphs from '../components/graphs';
import Top from '../components/top';

const { API } = process.env;

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  },
  formControl: {
    marginTop: 10,
    minWidth: 120
  }
}));

function getData(repository) {
  const url = `${API}/repositories/${repository}`;
  return fetch(url)
    .then(response => response.json())
    .then(response => response.payload);
}

const SelectBranch = function({ current, branches, onChange }) {
  const classes = useStyles();
  const inputLabel = React.useRef(null);
  const [labelWidth, setLabelWidth] = React.useState(0);
  React.useEffect(() => {
    setLabelWidth(inputLabel.current.offsetWidth);
  }, []);

  return (
    <FormControl variant="outlined" className={classes.formControl}>
      <InputLabel ref={inputLabel}>branch</InputLabel>
      <Select value={current} onChange={onChange} labelWidth={labelWidth}>
        {branches.map(branch => {
          return <MenuItem value={branch}>{branch}</MenuItem>;
        })}
      </Select>
    </FormControl>
  );
};

const Repo = function(data) {
  const classes = useStyles();
  const { user, repository } = data;

  const [repo, setData] = useState({});
  const [branch, setBranch] = useState(undefined);
  const [branches, setBranches] = useState([]);
  const [name, setName] = useState('');

  const requestData = async () => {
    const data = await getData(repository);

    setBranch(data.defaultBranch);
    setBranches(data.branches);
    setName(data.name);
    setData(data);
  };

  useEffect(() => {
    requestData();
  }, []);

  return (
    <>
      <Top user={user} />
      <Layout user={user}>
        {branches.length === 0 ? null : (
          <Grid container>
            <Grid xs={10} item>
              <Typography
                className={classes.title}
                variant="h2"
                component="h2"
                gutterBottom={true}
              >
                {name}
              </Typography>
            </Grid>
            <Grid item>
              <SelectBranch
                className={classes.selectBranch}
                onChange={event => setBranch(event.target.value)}
                current={branch}
                branches={branches}
              />
            </Grid>
          </Grid>
        )}

        {branch === undefined ? null : (
          <Graphs repository={repository} branch={branch} />
        )}
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
