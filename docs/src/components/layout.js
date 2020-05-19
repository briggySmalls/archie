/**
 * Layout component that queries for data
 * with Gatsby's useStaticQuery component
 *
 * See: https://www.gatsbyjs.org/docs/use-static-query/
 */

import React from "react";
import PropTypes from "prop-types";
import { makeStyles } from '@material-ui/core/styles';
import "./layout.css";
import "typeface-roboto";
import TopBar from './topbar';
import ResponsiveDrawer from './drawer';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  main: {
    fontFamily: 'Roboto',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
}));

const Layout = ({ showSidebar, location, children }) => {
  // Prepare our CSS styles
  const classes = useStyles();
  // Handle mobile visibility
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };
  return (
    <>
      <TopBar showSidebar={showSidebar} handleDrawerToggle={handleDrawerToggle} />
      {/* Only show drawer if instructed */}
      {showSidebar &&
        <nav className={classes.drawer} aria-label="site pages">
          <ResponsiveDrawer handleDrawerToggle={handleDrawerToggle} mobileOpen={mobileOpen} />
        </nav>
      }
      <div
        style={{
          margin: `0 auto`,
          maxWidth: 960,
          padding: `0 1.0875rem 1.45rem`,
        }}
      >
        <main className={classes.main}>{children}</main>
        <footer>
          © {new Date().getFullYear()}, Made with ❤️ by{" "}
          <a href="https://github.com/briggySmalls" target="_blank" rel="noopener noreferrer">
            Sam Briggs
          </a>
        </footer>
      </div>
    </>
  )
}

Layout.propTypes = {
  showSidebar: PropTypes.node.isRequired,
  location: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
}

Layout.defaultProps = {
  showSidebar: true,
}

export default Layout
