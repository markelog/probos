import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Grid from '@material-ui/core/Grid';

import Link from 'next/link';
import prettyBytes from 'pretty-bytes';

import LittleChart from '../little-chart';

const { API } = process.env;

const useStyles = makeStyles(({ palette }) => {
  return {
    card: {
      position: 'relative',
      minWidth: 275,
      height: 220
    },
    cardContent: {
      overflow: 'auto',
      width: '100%',
      height: 230
    },
    gridContainer: {
      position: 'relative'
    },
    chart: {
      position: 'absolute',
      top: 60,
      right: 0,
      bottom: 25,
      left: 0,
      opacity: 0.1
    },
    repo: {
      minWidth: '50%'
    },
    title: {
      fontSize: 14
    },
    link: {
      color: palette.secondary.dark,
      cursor: 'pointer'
    },
    pos: {
      marginBottom: 12
    },
    grid: {
      flexGrow: 1
    },
    list: {
      padding: 0
    },
    listHeader: {
      marginTop: 10
    },
    listItem: {
      padding: 0
    }
  };
});

function getData(username, page) {
  const url = `${API}/users/${username}/repos?page=${page}`;
  return fetch(url)
    .then(response => response.json())
    .then(response => response.payload);
}

function view(data) {
  const classes = useStyles();

  return (
    <Grid container spacing={3} justify="center">
      {data.map(repo => {
        return (
          <Grid xs item className={classes.repo}>
            {viewRepo(classes, repo)}
          </Grid>
        );
      })}
    </Grid>
  );
}

function viewRepo(classes, data, index) {
  const { name, repository, total } = data;
  const href = `/repos/${repository}`;

  return (
    <Card className={classes.card} key={index} variant="outlined">
      <CardContent className={classes.cardContent}>
        <div className={classes.chart}>
          <LittleChart total={total} />
        </div>
        <Typography variant="h5" component="h2">
          <Link href={href}>
            <a className={classes.link}>{name}</a>
          </Link>
        </Typography>
        <Grid
          container
          justify="flex-start"
          key={name}
          className={classes.gridContainer}
        >
          {data['last-report'].map(viewFiles.bind(null, classes))}
        </Grid>
      </CardContent>
    </Card>
  );
}

function viewFiles(classes, data) {
  const { name, sizes } = data;

  return (
    <Grid item xs={6}>
      <Typography className={classes.listHeader} variant="h6">
        {name}
      </Typography>
      <List className={classes.list}>
        {sizes.map((point, i) => {
          const { size, gzip, author, hash, message, date } = point;
          const prettySize = prettyBytes(size);
          const prettyGzip = prettyBytes(gzip);

          return (
            <>
              <ListItem className={classes.listItem}>
                <ListItemText primary="gzip" secondary={prettySize} />
              </ListItem>
              <ListItem className={classes.listItem}>
                <ListItemText primary="size" secondary={prettyGzip} />
              </ListItem>
            </>
          );
        })}
      </List>
    </Grid>
  );
}

export default function Repos({ user, page }) {
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

  return view(data);
}
