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

const useStyles = makeStyles(theme => ({}));

function User({ user }) {
  const { username, avatar } = user;
  return (
    <>
      <Avatar alt={username} src={avatar} />
      <Link href="/api/auth/logout" color="inherit">
        sign out
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
  return (
    <Button color="inherit">{user ? <User user={user} /> : <SignIn />}</Button>
  );
}
