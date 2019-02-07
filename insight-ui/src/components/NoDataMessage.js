import React from 'react'

import { InfoOutlined } from '@material-ui/icons'

import './NoDataMessage.scss'

const NoDataMessage = (props) => (
    <p className="NoDataMessage">
        <InfoOutlined /> <br />
        <span data-test="message">
            {props.msg ? props.msg : 'No logs to show'}
        </span>
    </p>
)

export default NoDataMessage
