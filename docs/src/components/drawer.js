import React from "react";
import PropTypes from "prop-types";
import { useStaticQuery, graphql } from "gatsby";
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Drawer from '@material-ui/core/Drawer';
import { Link } from "gatsby";
import Hidden from '@material-ui/core/Hidden';
import { makeStyles } from '@material-ui/core/styles';

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

// Design a drawer for navigating content
const NavDrawer = ({ edges }) => {
  // Prepare our CSS styles
  const classes = useStyles();
  return (
    <div>
      <div className={classes.toolbar} />
      <List>
        {
          // Sort the site data by menuPosition
          edges.sort(function (a, b) {
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
};

const ResponsiveDrawer = ({ showSidebar, location, handleDrawerToggle, mobileOpen, children }) => {
  // Prepare our CSS styles
  const classes = useStyles();
  // Request site data
  const data = useStaticQuery(graphql`
    query {
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
  `);
  // Handle hide/show of menu on mobile
  const drawer = <NavDrawer edges={data.allMdx.edges} />
  // Show draw responsively
  return (
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
  );
};

ResponsiveDrawer.propTypes = {
  showSidebar: PropTypes.node.isRequired,
  location: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
}

ResponsiveDrawer.defaultProps = {
  showSidebar: true,
}

export default ResponsiveDrawer;
