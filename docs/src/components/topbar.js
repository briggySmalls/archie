import React from "react";
import PropTypes from "prop-types";
import { useStaticQuery, graphql } from "gatsby";
import IconButton from '@material-ui/core/IconButton';
import Toolbar from '@material-ui/core/Toolbar';
import Hidden from '@material-ui/core/Hidden';
import MenuIcon from '@material-ui/icons/Menu';
import { Link } from "gatsby";
import AppBar from '@material-ui/core/AppBar';
import MaterialLink from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
      display: 'none',
    },
  },
}));

// Topbar to hold site title, etc
const TopBar = ({ showSidebar, handleDrawerToggle }) => {
  // Prepare our CSS styles
  const classes = useStyles();
  // Request site data
  const data = useStaticQuery(graphql`
    query {
      site {
        siteMetadata {
          title
        }
      }
    }
  `);
  return (
    <AppBar position="sticky" className={classes.appBar}>
      <Toolbar>
        <Hidden smUp implementation="css">
          {/* Only show sidebar toggler if instructed */}
          {showSidebar &&
            <IconButton
              color="inherit"
              aria-label="open menu"
              edge="start"
              onClick={handleDrawerToggle}
            >
              <MenuIcon />
            </IconButton>
          }
        </Hidden>
        <MaterialLink variant="h6" color="inherit" noWrap component={Link} to='/'>
          {data.site.siteMetadata.title}
        </MaterialLink>
      </Toolbar>
    </AppBar>
  )
};

TopBar.propTypes = {
  showSidebar: PropTypes.node.isRequired,
}

TopBar.defaultProps = {
  showSidebar: true,
}

export default TopBar;
