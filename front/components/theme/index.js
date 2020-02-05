import { createMuiTheme } from '@material-ui/core/styles';
import { amber } from '@material-ui/core/colors';

// Create a theme instance.
const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: amber
  }
});

export default theme;
