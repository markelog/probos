import React, { useState, useEffect } from 'react';
import { withRouter } from 'next/router';
import Router from 'next/router';
import cookies from 'js-cookie';

import Chip from '@material-ui/core/Chip';
import Link from '@material-ui/core/Link';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import LockOpenIcon from '@material-ui/icons/LockOpen';

import { makeStyles } from '@material-ui/core/styles';

import Layout from '../components/layout';
import Graphs from '../components/graphs';
import Top from '../components/top';

const { API } = process.env;

const useStyles = makeStyles(theme => ({
  link: {
    margin: theme.spacing(1)
  },
  chip: {
    marginTop: -40,
    marginLeft: 10,
    borderRadius: 0
  },
  formControl: {
    marginTop: 10,
    minWidth: 120
  }
}));

function getData(repository) {
  const url = `${API}/repositories/${repository}`;

  return fetch(url, {
    headers: {
      Authorization: `Bearer ${cookies.get('jwt')}`
    }
  })
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
  const { user, repository, branch, router } = data;

  const [repo, setData] = useState({});
  const [token, setToken] = useState(undefined);
  const [defaultBranch, setDefaultBranch] = useState(undefined);
  const [currentBranch, setCurrentBranch] = useState(branch);
  const [branches, setBranches] = useState([]);
  const [name, setName] = useState('');

  const requestData = async () => {
    const data = await getData(repository);

    setDefaultBranch(data.defaultBranch);
    setCurrentBranch(branch || data.defaultBranch);
    setToken(data.token);
    setBranches(data.branches);
    setName(data.name);
    setData(data);
  };

  const handleBranch = event => {
    const branch = event.target.value;

    if (branch === defaultBranch) {
      router.push(Router.pathname, `/repos/${repository}`, {
        shallow: true
      });
    } else {
      router.push(Router.pathname, `/repos/${repository}/branch/${branch}`, {
        shallow: true
      });
    }

    setCurrentBranch(branch);
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
                {token && (
                  <Chip
                    className={classes.chip}
                    label={'PROBOS_TOKEN=' + token}
                    icon={<LockOpenIcon />}
                  />
                )}
              </Typography>
            </Grid>
            <Grid item>
              <SelectBranch
                className={classes.selectBranch}
                onChange={handleBranch}
                current={currentBranch}
                branches={branches}
              />
            </Grid>
          </Grid>
        )}

        {currentBranch === undefined ? null : (
          <Graphs repository={repository} branch={currentBranch} />
        )}
      </Layout>
    </>
  );
};

Repo.getInitialProps = ({ query, router }) => {
  return {
    user: query.user,
    branch: query.branch,
    repository: query.repository
  };
};

export default withRouter(Repo);
