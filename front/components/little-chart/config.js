import { makeStyles } from '@material-ui/core/styles';
import lightGreen from '@material-ui/core/colors/lightGreen';
import grey from '@material-ui/core/colors/grey';

export default {
  keys: ['size', 'gzip'],
  indexBy: 'date',
  theme: {
    tooltip: {
      container: {
        display: 'none'
      }
    }
  },
  colors: [grey[300], '#f3665c'],
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
