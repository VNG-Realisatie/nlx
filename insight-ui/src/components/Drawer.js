import React from 'react';
import PropTypes from 'prop-types';
import { Link, Switch, Route } from 'react-router-dom'

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
import QRPage from './QRPage';
import IrmaPage from './IrmaTest';

class ResponsiveDrawer extends React.Component {
	state = {
		mobileOpen: false
	};

	handleDrawerToggle = () => {
		this.setState(state => ({ mobileOpen: !state.mobileOpen }));
	};

	render() {
		const { classes, theme /*,appTitle*/} = this.props;

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
						{/* <Typography variant="title" color="primary" noWrap>
							{appTitle}
						</Typography> */}
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

						<OrganizationList />
					</Drawer>
				</Hidden>
				<Hidden smDown implementation="css">
					<Drawer variant="permanent" className={classes.drawerPaper}>
						<div className={classes.toolbar} />
						<OrganizationList />
					</Drawer>
				</Hidden>
				<main className={classes.content}>
					<div className={classes.toolbar} />
					<Switch>
						<Route exact path="/" component={Home} />
						<Route exact path="/irma" component={IrmaPage} />
						<Route path="/organization/:cid" component={OrganizationPage} />
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

export default withStyles(styles, { withTheme: true })(ResponsiveDrawer);
