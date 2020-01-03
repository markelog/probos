export default {
  keys: ["size", "gzip"],
  margin: {
    top: 50,
    right: 130,
    bottom: 50,
    left: 60
  },
  padding: 0.3,
  colors: ["#f45b87", "#6eadd4"],
  axisLeft: null,
  axisTop: null,
  axisRight: null,
  indexBy: "date",
  enableGridX: false,
  enableGridY: false,
  legends: [
    {
      dataFrom: "keys",
      anchor: "bottom-right",
      direction: "column",
      justify: false,
      translateX: 120,
      translateY: 0,
      itemsSpacing: 2,
      itemWidth: 100,
      itemHeight: 20,
      itemDirection: "left-to-right"
    }
  ]
};
