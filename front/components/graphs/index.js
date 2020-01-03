import { useState, useEffect } from "react";

import Typography from "@material-ui/core/Typography";

import fetch from "isomorphic-unfetch";

import Chart from "../chart";
import Table from "../table";

const API = process.env.API;

function getData(branch, repository) {
  const url = `${API}/reports?repository=${repository}&branch=${branch}`;
  return fetch(url)
    .then(response => {
      return response.json();
    })
    .then(response => {
      return response.payload;
    });
}

const Graphs = ({ repository, branch }) => {
  const [data, setData] = useState([]);

  const requestData = async () => {
    const data = await getData(branch, repository);

    setData(data);
  };

  useEffect(() => {
    requestData();
  }, [setData]);

  return data.map(result => {
    const { name, sizes } = result;

    return (
      <div>
        <Typography variant="h3" component="h3" gutterBottom="true">
          {name}
        </Typography>
        <Chart data={sizes} />
        <Table data={sizes} />
      </div>
    );
  });
};

export default Graphs;
