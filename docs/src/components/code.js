import React from 'react';
import { makeStyles } from '@material-ui/core/styles';


const useStyles = makeStyles(() => ({
  pre: {
    maxHeight: "25em",
  },
}));

const Code = ({children}) => {
  // Prepare our CSS styles
  const classes = useStyles();
  // Return code
  return (
    <pre className={classes.pre}>{ children }</pre>
  );
};

export default Code;
