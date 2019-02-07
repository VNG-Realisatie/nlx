import React from 'react'
import { Typography } from '@material-ui/core'

const HomePage = () => (
    <React.Fragment>
        <Typography variant="h4" color="primary" noWrap gutterBottom>
            Home
        </Typography>
        <Typography variant="h6" color="default" gutterBottom>
            Welcome to the NLX Insight App
        </Typography>
        <Typography variant="body1" color="default" gutterBottom>
            View logs by selecting an organization using the menu items on the
            left.
        </Typography>
        <Typography variant="body2" color="default" gutterBottom>
            You can only view organization logs if you have disclosed required
            IRMA attributes. <br />
            If you haven&apos;t, you will be asked to do so.
        </Typography>
    </React.Fragment>
)

export default HomePage
