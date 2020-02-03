import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Link from '@material-ui/core/Link';
import Avatar from '@material-ui/core/Avatar';

const useStyles = makeStyles(theme => ({
  avatar: {
    width: theme.spacing(4),
    height: theme.spacing(4),
    margin: theme.spacing(1),
    float: 'left'
  },
  link: {
    margin: '5px 0 0 -5px',
    display: 'block',
    float: 'left'
  },
  button: {
    textTransform: 'inherit',
    fontSize: '1rem'
  }
}));

function User({ user }) {
  const classes = useStyles();
  const { username, avatar } = user;
  return (
    <>
      <Avatar alt={username} src={avatar} className={classes.avatar} />
      <Link href="/api/auth/logout" color="inherit" className={classes.link}>
        <Button className={classes.button} color="inherit">
          Sign out
        </Button>
      </Link>
    </>
  );
}

function SignIn() {
  return (
    <Link href="/api/auth/github" color="inherit">
      sign in
    </Link>
  );
}

export default function SignInOut({ user }) {
  const classes = useStyles();
  return user ? <User user={user} /> : <SignIn />;
}
