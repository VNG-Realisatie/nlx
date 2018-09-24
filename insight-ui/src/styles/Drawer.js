const drawerWidth = 248;

const styles = theme => ({
	root: {
		minHeight: '100vh',
		display: 'flex'
	},
	appBar: {
		backgroundColor: theme.palette.background.default,
		zIndex: theme.zIndex.drawer + 1,
	},
	navIconHide: {
		[theme.breakpoints.up('md')]: {
			display: 'none',
		},
	},
	toolbar: theme.mixins.toolbar,
	drawerPaper: {
		width: drawerWidth,
	},
	content: {
		flexGrow: 1,
		backgroundColor: theme.palette.background.default,
		[theme.breakpoints.down('md')]: {
			padding: theme.spacing.unit * 2,
		},
		[theme.breakpoints.down('xs')]: {
			padding: theme.spacing.unit,
		},
		[theme.breakpoints.up('md')]: {
			padding: theme.spacing.unit * 3,
		}
	}
});

export default styles;