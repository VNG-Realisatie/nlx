import React from 'react'

import { InfoOutlined } from '@material-ui/icons'

import './NoDataMessage.scss'

const NoDataMessage = (props) => {
    return (
        <p data-test-id="no-data-msg" className="NoDataMessage">
            <InfoOutlined /> <br />
            {props.msg ? props.msg : 'No logs to show'}
        </p>
    )
}

export default NoDataMessage
