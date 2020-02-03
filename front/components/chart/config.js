import theme from '../theme';
import lightGreen from '@material-ui/core/colors/lightGreen';
import grey from '@material-ui/core/colors/grey';

export default {
  keys: ['size', 'gzip'],
  margin: {
    top: 50,
    right: 130,
    bottom: 50,
    left: 60
  },
  padding: 0.3,
  colors: [grey[300], '#f3665c'],
  axisLeft: null,
  axisTop: null,
  axisRight: null,
  indexBy: 'date',
  enableGridX: false,
  enableGridY: false,
  legends: [
    {
      dataFrom: 'keys',
      anchor: 'bottom-right',
      direction: 'column',
      justify: false,
      translateX: 120,
      translateY: 0,
      itemsSpacing: 2,
      itemWidth: 100,
      itemHeight: 20,
      itemDirection: 'left-to-right',
      itemTextColor: '#fff'
    }
  ]
};
