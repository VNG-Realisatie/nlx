import React from 'react'
// import { Link } from 'react-router-dom';
import { Typography } from '@material-ui/core'

const Home = () => (
    <React.Fragment>
        <Typography variant="h6" color="primary" noWrap gutterBottom>
            Home
        </Typography>
        <Typography variant="subtitle1" color="default" noWrap gutterBottom>
            <p>Welcome to the NLX Insight App.</p>
            <ul>
                <li>
                    View logs by selecting an organizations using the menu items on the left.
                </li>
            </ul>
            <p>
                <strong>
                    You can only view organization logs if you have disclosed
                    required IRMA attributes.
                    <br />
                    If you haven&apos;t, you will be asked to do so.
                    <br />
                </strong>
            </p>
        </Typography>
    </React.Fragment>
)

export default Home
