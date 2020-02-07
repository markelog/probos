import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import Link from '@material-ui/core/Link';
import Avatar from '@material-ui/core/Avatar';

const useStyles = makeStyles(theme => ({
  avatar: {
    width: 20,
    height: 20,
    whiteSpace: 'nowrap',
    display: 'inline-block',
    top: 6,
    borderRadius: '30%'
  },
  content: {
    display: 'inline-block',
    marginLeft: 5
  }
}));

export default function({ url, username, avatar }) {
  const classes = useStyles();
  return (
    <>
      <Avatar alt={username} src={avatar} className={classes.avatar} />
      <b className={classes.content}>
        <Link href={url}>{username}</Link>
      </b>
    </>
  );
}
