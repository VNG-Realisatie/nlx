import React from 'react';

import { withStyles } from '@material-ui/core'

import logo from '../styles/logo.svg';

const styles = theme => {
	return {
		root: {
			display: 'flex'
		}
	}
};

const Logo = props => {
	let { classes } = props
	return (
		<div className={classes.root}>
			<img src={logo} className="app-logo-drawer" alt="logo" />
		</div>
	)
}

export default withStyles(styles)(Logo);