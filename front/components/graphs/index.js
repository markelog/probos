import { useState, useEffect } from "react";

import fetch from "isomorphic-unfetch";

import Chart from "../chart";

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

  console.log(repository, branch);

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
      <div style={{ height: "500px", width: "500px" }}>
        <h1>{name}</h1>
        <Chart data={sizes} />
      </div>
    );
  });
};

export default Graphs;
