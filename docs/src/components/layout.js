/**
 * Layout component that queries for data
 * with Gatsby's useStaticQuery component
 *
 * See: https://www.gatsbyjs.org/docs/use-static-query/
 */

import React from "react"
import PropTypes from "prop-types"
import { useStaticQuery, graphql } from "gatsby"
import AppBar from '@material-ui/core/AppBar';
import MenuIcon from '@material-ui/icons/Menu';
import IconButton from '@material-ui/core/IconButton';
import Toolbar from '@material-ui/core/Toolbar';
import { makeStyles } from '@material-ui/core/styles';
import "./layout.css"
import Drawer from '@material-ui/core/Drawer';
import Hidden from '@material-ui/core/Hidden';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { Link } from "gatsby"
import MaterialLink from '@material-ui/core/Link';
import "typeface-roboto";

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  main: {
    fontFamily: 'Roboto',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
      display: 'none',
    },
  },
  // necessary for content to be below app bar
  toolbar: theme.mixins.toolbar,
  drawerPaper: {
    width: drawerWidth,
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const Layout = ({ showSidebar, children }) => {
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
      allMdx(filter: {frontmatter: {menuPosition: {ne: null}}}) {
        edges {
          node {
            frontmatter {
              title
              menuPosition
            }
            fields {
              slug
            }
          }
        }
      }
    }
  `)
  // Handle hide/show of menu on mobile
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  // Design a drawer for navigating content
  const drawer = (
    <div>
      <div className={classes.toolbar} />
      <List>
        {
          // Sort the site data by menuPosition
          data.allMdx.edges.sort(function (a, b) {
            return a.node.frontmatter.menuPosition - b.node.frontmatter.menuPosition
          }).map((edge, index) => {
            const title = edge.node.frontmatter.title
            const slug = edge.node.fields.slug
            return (
              <ListItem button key={title} component={Link} to={slug}>
                <ListItemText primary={title} />
              </ListItem>
            )
          })
        }
      </List>
    </div>
  )

  // Show draw responsively
  const responsiveDrawer = (
    <>
      {/* Drawer for mobile */}
      <Hidden smUp implementation="css">
        <Drawer
          variant="temporary"
          anchor='left'
          open={mobileOpen}
          onClose={handleDrawerToggle}
          classes={{
            paper: classes.drawerPaper,
          }}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
        >
          {drawer}
        </Drawer>
      </Hidden>
      {/* Drawer for desktop */}
      <Hidden xsDown implementation="css">
        <Drawer
          classes={{
            paper: classes.drawerPaper,
          }}
          variant="permanent"
          open
        >
          {drawer}
        </Drawer>
      </Hidden>
    </>
  )

  // Topbar to hold site title, etc
  const appBar = (
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

  return (
    <>
      <header>
        {appBar}
      </header>
      {/* Only show drawer if instructed */}
      {showSidebar &&
        <nav className={classes.drawer} aria-label="site pages">
          {responsiveDrawer}
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
  children: PropTypes.node.isRequired,
}

Layout.defaultProps = {
  showSidebar: true,
}

export default Layout
