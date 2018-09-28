import React from 'react'
import PropTypes from 'prop-types'
import { Link, Switch, Route, withRouter } from 'react-router-dom'
import config from '../utils/config'
import axios from 'axios'

import {
	withStyles,
	Drawer,
	AppBar, Toolbar,
	IconButton,
	Hidden,
} from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'

import styles from '../styles/Drawer';
import Home from './Home';
import OrganizationPage from './OrganizationPage';
import OrganizationList from './OrganizationList';
import Logo from './Logo';

class ResponsiveDrawer extends React.Component {
	state = {
		mobileOpen: false,
		loading: true,
		error: false,
		organizations: []
	};

    componentDidMount() {
        this.getOrganizations()
    }

    getOrganizations() {
		const parser = document.createElement('a')
		parser.href = window.location.href

		let directoryProtocol, directoryHostname
		if (parser.hostname.indexOf("minikube") !== -1) {
			directoryProtocol = "http://"
			directoryHostname = "directory.dev.nlx.minikube:30080"
		} else if (parser.hostname.indexOf("localhost") !== -1) {
			directoryProtocol = "http://"
			directoryHostname = "directory.dev.nlx.minikube:30080"
		} else if (parser.hostname.indexOf("test.nlx.io") !== -1) {
			directoryProtocol = "https://"
			directoryHostname = "directory.test.nlx.io"
		} else {
			directoryProtocol = "https://"
			directoryHostname = "directory.demo.nlx.io"
		}

        axios.get(`${directoryProtocol}//${directoryHostname}/api/directory/list-organizations`)
        .then(response => {
            this.setState({
				loading: false,
                organizations: response.data.organizations
            })
        },(e)=>{
			console.error(e)
			this.setState({ error: true })
        })
    }


	handleDrawerToggle = () => {
		this.setState(state => ({ mobileOpen: !state.mobileOpen }));
	};

	render() {
		const { classes, theme } = this.props;

		if (this.state.error) {
			return (
				<div>Could not load data, please try again</div>
			)
		}

		return (
			<div className={classes.root}>
				<AppBar
					position="fixed"
					className={classes.appBar}>
					<Toolbar>
						<IconButton
							color="inherit"
							aria-label="Open drawer"
							onClick={this.handleDrawerToggle}
							className={classes.navIconHide}>
							<MenuIcon />
						</IconButton>
						<Hidden smDown implementation="css">
							<Link to="/">
								<Logo />
							</Link>
						</Hidden>
					</Toolbar>
				</AppBar>
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
						}}>

						<Toolbar>
							<Logo />
						</Toolbar>

						<OrganizationList organizations={this.state.organizations} />
					</Drawer>
				</Hidden>
				<Hidden smDown implementation="css">
					<Drawer variant="permanent" className={classes.drawerPaper}>
						<div className={classes.toolbar} />
						<OrganizationList organizations={this.state.organizations} />
					</Drawer>
				</Hidden>
				<main className={classes.content}>
					<div className={classes.toolbar} />
					<Switch>
						<Route exact path="/" component={Home} />
						<Route exact path="/home" component={Home} />
						<Route path="/organization/:name" render={() => <OrganizationPage organizations={this.state.organizations} />} />
					</Switch>
				</main>
			</div>
		);
	}
}

ResponsiveDrawer.propTypes = {
	classes: PropTypes.object.isRequired,
	theme: PropTypes.object.isRequired,
};

export default withStyles(styles, { withTheme: true })(withRouter(ResponsiveDrawer));
