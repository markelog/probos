import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({}));
const { API } = process.env;

function getData(username, page) {
  const url = `${API}/users/${username}/repos?page=${page}`;
  return fetch(url)
    .then(response => {
      return response.json();
    })
    .then(response => {
      return response.payload;
    });
}

export default function Repos({ username, page }) {
  const classes = useStyles();
  const [data, setData] = useState([]);

  const requestData = async () => {
    const data = await getData(username, page);

    setData(data);
  };

  useEffect(() => {
    requestData();
  }, [page]);

  return <h1>{JSON.stringify(data)}</h1>;
}
