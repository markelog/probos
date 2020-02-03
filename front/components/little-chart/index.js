import { ResponsiveBar } from '@nivo/bar';

import config from './config.js';

const LittleChart = ({ total }) => {
  return <ResponsiveBar data={total} {...config} />;
};

export default LittleChart;
