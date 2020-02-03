import { makeStyles } from '@material-ui/core/styles';
import theme from '../theme';

export default {
  keys: ['hot dog', 'burger'],
  indexBy: 'country',
  theme: {
    tooltip: {
      container: {
        display: 'none'
      }
    }
  },
  colors: ['rgb(235, 237, 239)', theme.palette.secondary.dark],
  // colors: [theme.palette.primary.main, theme.palette.info.dark],
  // colors: { scheme: 'nivo' },
  defs: [
    {
      id: 'dots',
      type: 'patternDots',
      background: 'inherit',
      color: '#38bcb2',
      size: 4,
      padding: 1,
      stagger: true
    },
    {
      id: 'lines',
      type: 'patternLines',
      background: 'inherit',
      color: '#eed312',
      rotation: -45,
      lineWidth: 6,
      spacing: 10
    }
  ],
  fill: [
    {
      match: {
        id: 'fries'
      },
      id: 'dots'
    },
    {
      match: {
        id: 'sandwich'
      },
      id: 'lines'
    }
  ],
  borderColor: { from: 'color', modifiers: [['darker', 1.6]] },
  axisTop: null,
  axisRight: null,
  enableGridX: false,
  enableGridY: false,
  label: null,
  labelSkipWidth: 12,
  labelSkipHeight: 12,
  labelTextColor: { from: 'color', modifiers: [['darker', 1.6]] },
  animate: true,
  motionStiffness: 90,
  motionDamping: 15
};
