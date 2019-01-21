import React from 'react'

import { InfoOutlined } from '@material-ui/icons'

import './NoDataMessage.scss'

const NoDataMessage = (props) => {
    return (
        <p className="NoDataMessage">
            <InfoOutlined /> <br />
            {props.msg ? props.msg : 'No logs to show'}
        </p>
    )
}

export default NoDataMessage
