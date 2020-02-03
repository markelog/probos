import { useState, useEffect } from 'react';

import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';

import fetch from 'isomorphic-unfetch';

import Chart from '../chart';
import Table from '../table';

const { API } = process.env;

const useStyles = makeStyles({
  container: {
    marginBottom: 20
  }
});

function getData(branch, repository) {
  const url = `${API}/reports?repository=${repository}&branch=${branch}`;
  return fetch(url)
    .then(response => response.json())
    .then(response => response.payload);
}

const Graphs = ({ repository, branch }) => {
  const classes = useStyles();
  const [data, setData] = useState([]);

  const requestData = async () => {
    const data = await getData(branch, repository);

    setData(data);
  };

  useEffect(() => {
    requestData();
  }, []);

  return data.map(result => {
    const { name, sizes } = result;

    return (
      <div className={classes.container}>
        <Typography
          className={classes.title}
          variant="h3"
          component="h3"
          gutterBottom="true"
        >
          {name}
        </Typography>
        <Chart data={sizes} />
        <Table data={sizes} />
      </div>
    );
  });
};

export default Graphs;
