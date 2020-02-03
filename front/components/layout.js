import Head from 'next/head';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';
import Box from '@material-ui/core/Box';

const useStyles = makeStyles(theme => ({
  main: {
    marginTop: 20
  },
  link: {
    margin: theme.spacing(1)
  }
}));

function Layout({ user, loading = false, children }) {
  const classes = useStyles();
  const preventDefault = event => event.preventDefault();

  return (
    <>
      <Container maxWidth="md">
        <main className={classes.main}>{children}</main>
      </Container>

      <Head>
        <title>Probos</title>
      </Head>
    </>
  );
}

export default Layout;
