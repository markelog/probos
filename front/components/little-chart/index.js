import { ResponsiveBar } from '@nivo/bar';

import data from './data.js';
import config from './config.js';

const LittleChart = () => {
  return <ResponsiveBar data={data} {...config} />;
};

export default LittleChart;
