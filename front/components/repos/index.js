import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({}));
const { API } = process.env;

function getData(username, page) {
  const url = `${API}/users/${username}/repos?page=${page}`;
  return fetch(url)
    .then(response => response.json())
    .then(response => response.payload);
}

export default function Repos({ user, page }) {
  const classes = useStyles();
  const [data, setData] = useState([]);

  if (user === undefined) {
    return null;
  }

  const { username } = user;

  const requestData = async () => {
    const data = await getData(username, page);

    setData(data);
  };

  useEffect(() => {
    requestData();
  }, [page]);

  return <h1>{JSON.stringify(data)}</h1>;
}
