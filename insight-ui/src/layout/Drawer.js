import React from 'react'
import PropTypes from 'prop-types'
import { Link, Switch, Route, BrowserRouter, Redirect } from 'react-router-dom'

import { connect } from 'react-redux'
import { compose } from 'redux'
import { Spinner } from '@commonground/design-system'

import {
    withStyles,
    Drawer,
    AppBar,
    Toolbar,
    IconButton,
    Hidden,
} from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'

import styles from '../styles/Drawer'
import HomePage from '../pages/HomePage/HomePage'
import OrganizationPage from '../pages/organization/OrganizationPage'
import OrganizationList from '../components/OrganizationList'
import Logo from '../components/Logo'
import ErrorPage from '../pages/ErrorPage/ErrorPage'

class ResponsiveDrawer extends React.Component {
    state = {
        mobileOpen: false,
    }

    handleDrawerToggle = () => {
        this.setState((state) => ({ mobileOpen: !state.mobileOpen }))
    }

    getAppBar = () => {
        const { classes } = this.props
        return (
            <AppBar position="fixed" className={classes.appBar}>
                <Toolbar>
                    <IconButton
                        color="inherit"
                        aria-label="Open drawer"
                        onClick={this.handleDrawerToggle}
                        className={classes.navIconHide}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Hidden smDown implementation="css">
                        <Link to="/">
                            <Logo />
                        </Link>
                    </Hidden>
                </Toolbar>
            </AppBar>
        )
    }

    getDrawer = () => {
        const { classes, theme, organizations } = this.props
        return (
            <React.Fragment>
                <Hidden mdUp>
                    <Drawer
                        variant="temporary"
                        anchor={theme.direction === 'rtl' ? 'right' : 'left'}
                        open={this.state.mobileOpen}
                        onClose={this.handleDrawerToggle}
                        classes={{
                            paper: classes.drawerPaper,
                        }}
                        ModalProps={{
                            keepMounted: true, // Better open performance on mobile.
                        }}
                    >
                        <Toolbar>
                            <Logo />
                        </Toolbar>

                        <OrganizationList organizations={organizations} />
                    </Drawer>
                </Hidden>
                <Hidden smDown implementation="css">
                    <Drawer variant="permanent" className={classes.drawerPaper}>
                        <div className={classes.toolbar} />
                        <OrganizationList organizations={organizations} />
                    </Drawer>
                </Hidden>
            </React.Fragment>
        )
    }

    getMain = () => {
        let { classes, loading, error } = this.props

        if (error) {
            return <ErrorPage />
        }
        if (loading) {
            return <Spinner />
        }
        return (
            <main className={classes.content}>
                <div className={classes.toolbar} />
                <Switch>
                    <Route
                        exact
                        path="/"
                        component={HomePage}
                        {...this.props}
                    />
                    <Route
                        exact
                        path="/home"
                        component={HomePage}
                        {...this.props}
                    />
                    <Route
                        path="/organization/:name"
                        component={OrganizationPage}
                        {...this.props}
                    />
                    <Redirect to="/home" />
                </Switch>
            </main>
        )
    }

    render() {
        const { classes } = this.props

        return (
            <BrowserRouter>
                <div className={classes.root}>
                    {this.getAppBar()}

                    {this.getDrawer()}

                    {this.getMain()}
                </div>
            </BrowserRouter>
        )
    }
}

ResponsiveDrawer.propTypes = {
    classes: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
}

const mapStateToProps = (state) => {
    return {
        loading: state.loader.show,
        error: state.organizations.error,
        organizations: state.organizations.list,
    }
}

export default compose(
    withStyles(styles, { withTheme: true }),
    connect(mapStateToProps),
)(ResponsiveDrawer)
