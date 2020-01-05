import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TablePagination from "@material-ui/core/TablePagination";
import TableRow from "@material-ui/core/TableRow";

import MoreIcon from "@material-ui/icons/AddCircleOutline";
import LessIcon from "@material-ui/icons/RemoveCircleOutline";
import red from "@material-ui/core/colors/red";
import green from "@material-ui/core/colors/green";

import prettyBytes from "pretty-bytes";
import { formatDistance, subDays } from "date-fns";

const useStyles = makeStyles({
  root: {},
  stiff: {
    whiteSpace: "nowrap"
  },
  MoreIcon: {
    position: "relative",
    top: 5,
    color: red[500],
    fontSize: 20
  },
  LessIcon: {
    position: "relative",
    top: 5,
    color: green[500],
    fontSize: 20
  },
  container: {
    height: 300
  }
});

const diffFormat = (value, classes) => {
  if (value.diff === "0") {
    return <span className={classes.stiff}>{value.diff}</span>;
  }

  if (value.increased) {
    return (
      <span className={classes.stiff}>
        <MoreIcon className={classes.MoreIcon} /> {value.diff} %
      </span>
    );
  }

  return (
    <span className={classes.stiff}>
      <LessIcon className={classes.LessIcon} /> {value.diff} %
    </span>
  );
};

const columns = [
  {
    id: "size",
    label: "size",
    align: "right",
    format: (value, classes) => {
      return <span className={classes.stiff}>{prettyBytes(value)}</span>;
    }
  },
  {
    id: "gzip",
    label: "gzip",
    align: "right",
    format: (value, classes) => {
      return <span className={classes.stiff}>{prettyBytes(value)}</span>;
    }
  },
  {
    id: "sizeDiff",
    label: "size diff",
    align: "right",
    format: diffFormat
  },
  {
    id: "gzipDiff",
    label: "gzip diff",
    align: "right",
    format: diffFormat
  },
  {
    id: "author",
    label: "author",
    align: "right"
  },
  {
    id: "message",
    label: "message",
    align: "right",
    format: (value, classes) => {
      return <span className={classes.stiff}>{value}</span>;
    }
  },
  {
    id: "date",
    label: "date",
    align: "right",
    format: (value, classes) => {
      const distanceDate = formatDistance(
        subDays(new Date(), 3),
        new Date(value)
      );

      return <span className={classes.stiff}>{distanceDate}</span>;
    }
  }
];

function getDiff(previous, current) {
  const result = {
    sizeDiff: {
      increased: false,
      diff: "0"
    },
    gzipDiff: {
      increased: false,
      diff: "0"
    }
  };

  if (previous === undefined) {
    return result;
  }

  if (current.size > previous.size) {
    result.sizeDiff.diff = 100 - (previous.size * 100) / current.size;
    result.sizeDiff.increased = true;
  } else {
    result.sizeDiff.diff = 100 - (current.size * 100) / previous.size;
  }

  if (current.gzip > previous.gzip) {
    result.gzipDiff.diff = 100 - (current.gzip * 100) / previous.gzip;
    result.gzipDiff.increased = true;
  } else {
    result.gzipDiff.diff = 100 - (current.gzip * 100) / previous.gzip;
  }

  if (result.sizeDiff.diff !== 0) {
    result.sizeDiff.diff = result.sizeDiff.diff.toFixed(2);
  }

  if (result.gzipDiff.diff !== 0) {
    result.gzipDiff.diff = result.gzipDiff.diff.toFixed(2);
  }

  result.sizeDiff.diff += "";
  result.gzipDiff.diff += "";

  return result;
}

function process(data) {
  for (let i = 0; i < data.length; i++) {
    const current = data[i];
    const previous = data[i - 1];

    Object.assign(data[i], getDiff(previous, current));
  }

  return data.reverse();
}

export default function Sizes({ data }) {
  const classes = useStyles();
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(3);

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  return (
    <Paper className={classes.root}>
      <TableContainer className={classes.container}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              {columns.map(column => (
                <TableCell key={column.id} align={column.align}>
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {process(data, classes)
              .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              .map((item, i) => {
                return (
                  <TableRow hover role="checkbox" tabIndex={-1} key={i}>
                    {columns.map(column => {
                      const value = item[column.id];

                      return (
                        <TableCell key={column.date} align={column.align}>
                          {column.format
                            ? column.format(value, classes)
                            : value}
                        </TableCell>
                      );
                    })}
                  </TableRow>
                );
              })}
          </TableBody>
        </Table>
      </TableContainer>

      <TablePagination
        rowsPerPageOptions={[]}
        component="div"
        count={data.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onChangePage={handleChangePage}
        onChangeRowsPerPage={handleChangeRowsPerPage}
      />
    </Paper>
  );
}
