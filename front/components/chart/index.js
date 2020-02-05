import fetch from 'isomorphic-unfetch';

import { ResponsiveBar } from '@nivo/bar';

import { formatDistance, subDays } from 'date-fns';

import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';

import User from '../user';

import config from './config.js';

const useStyles = makeStyles({
  card: {
    minWidth: 275
  },
  title: {
    fontSize: 14
  },
  pos: {
    fontSize: 14,
    marginBottom: 12
  },
  container: {
    height: 300
  }
});

const Chart = ({ data }) => {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <ResponsiveBar
        {...config}
        data={data}
        axisBottom={null}
        tooltip={point => <Tooltip point={point} />}
      />
    </div>
  );
};

const Tooltip = ({ point }) => {
  const classes = useStyles();
  const { size, gzip, author, hash, message, date } = point.data;
  const formattedDate = formatDistance(subDays(new Date(), 3), new Date(date));

  return (
    <Card className={classes.card}>
      <CardContent>
        <Typography variant="h6" component="h2">
          {message}
        </Typography>
        <Typography className={classes.pos} color="textSecondary">
          <User
            username={author.username}
            avatar={author.avatar}
            url={author.url}
          />{' '}
          commited {formattedDate}
        </Typography>
        <Typography variant="body1">
          zip: <b>{size}</b>
          <br />
          gzip: <b>{gzip}</b>
        </Typography>
      </CardContent>
    </Card>
  );
};

export default Chart;
