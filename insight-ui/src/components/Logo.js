import React from 'react'

import { withStyles } from '@material-ui/core'

const styles = (theme) => {
    return {
        root: {
            display: 'flex',
        },
    }
}

const Logo = (props) => {
    let { classes } = props
    return (
        <div className={classes.root}>
            <img src="/logo.svg" className="app-logo-drawer" alt="logo" />
        </div>
    )
}

export default withStyles(styles)(Logo)
