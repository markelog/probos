import fetch from "isomorphic-unfetch";

import { ResponsiveLine } from "@nivo/line";

import { formatDistance, subDays } from "date-fns";

import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";

const useStyles = makeStyles({
  card: {
    minWidth: 275
  },
  title: {
    fontSize: 14
  },
  pos: {
    fontSize: 14,
    marginBottom: 12
  }
});

const CustomSymbol = ({ size, color, borderWidth, borderColor }) => (
  <g>
    <circle
      fill="#fff"
      r={size / 2}
      strokeWidth={borderWidth}
      stroke={borderColor}
    />
    <circle
      r={size / 5}
      strokeWidth={borderWidth}
      stroke={borderColor}
      fill={color}
      fillOpacity={0.35}
    />
  </g>
);

const Chart = ({ data }) => {
  const classes = useStyles();

  const commonProperties = {
    width: 900,
    height: 400,
    margin: { top: 20, right: 20, bottom: 60, left: 80 },
    data,
    animate: true,
    enableSlices: "x"
  };

  return (
    <ResponsiveLine
      {...commonProperties}
      data={data}
      xScale={{
        type: "time",
        format: "%Y-%m-%dT%H:%M",
        precision: "minute",
        stacked: true
      }}
      xFormat="time:%Y-%m-%dT%H:%M"
      yScale={{
        type: "linear",
        stacked: false
      }}
      axisLeft={{
        legend: "linear scale",
        legendOffset: 12
      }}
      axisBottom={{
        format: "%b %d",
        tickValues: "every 2 days",
        legend: "time scale",
        legendOffset: -12
      }}
      curve="linear"
      enablePointLabel={true}
      enableGridX={false}
      enableGridY={false}
      enableArea={true}
      areaOpacity={0.9}
      pointSymbol={CustomSymbol}
      pointSize={16}
      pointBorderWidth={1}
      tooltip={stuff => {
        const { serieId } = stuff.point;
        const { x, y, commit } = stuff.point.data;
        const { author, hash, message } = commit;
        const date = formatDistance(subDays(new Date(), 3), x);

        return (
          <Card className={classes.card}>
            <CardContent>
              <Typography variant="h6" component="h2">
                {message}
              </Typography>
              <Typography className={classes.pos} color="textSecondary">
                <b>{author}</b> commited {date}
              </Typography>
              <Typography variant="body2">
                {serieId}: <b>{y}</b>
              </Typography>
            </CardContent>
          </Card>
        );

        return <h1>commited {date}</h1>;
      }}
      pointBorderColor={{
        from: "color",
        modifiers: [["darker", 0.3]]
      }}
      useMesh={true}
      enableSlices={false}
    />
  );
};

export default Chart;
