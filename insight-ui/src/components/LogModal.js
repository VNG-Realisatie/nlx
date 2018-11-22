import React from 'react'
import PropTypes from 'prop-types'

import { Link } from 'react-router-dom'

import {
    withStyles,
    Modal,
    Paper,
    IconButton,
    Typography,
} from '@material-ui/core'

import modalStyles from '../styles/SimpleModal'
import { Close, CalendarToday } from '@material-ui/icons'
import ClockIcon from '@material-ui/icons/AccessTimeOutlined'

class SimpleModal extends React.Component {
    render() {
        const { data, classes, open, closeModal } = this.props

        const d = new Date(data['created'])
        const localDate = d.toLocaleDateString()
        const localTime = d.toLocaleTimeString()
        return (
            <Modal
                aria-labelledby="simple-modal-title"
                aria-describedby="simple-modal-description"
                open={open}
                disableAutoFocus={true}
            >
                <Paper className={classes.paper} style={{ outline: 'none' }}>
                    <IconButton
                        onClick={closeModal}
                        className={classes.closeButton}
                    >
                        <Close style={{ fontSize: 18 }} />
                    </IconButton>

                    <Typography
                        variant="h6"
                        color="primary"
                        style={{ marginLeft: -1, marginBottom: 5 }}
                    >
                        {data['data']['doelbinding-data-elements']
                            ? data['data']['doelbinding-data-elements']
                            : 'Geen attribuut opgevraagd.'}
                    </Typography>
                    <div
                        style={{
                            display: 'flex',
                            justifyContent: 'space-between',
                            flexWrap: 'wrap',
                        }}
                    >
                        <Typography variant="caption">
                            # {data['logrecord-id']}
                        </Typography>
                        <Typography variant="caption">
                            <CalendarToday className={classes.calendarIcon} />
                            {localDate}
                            &nbsp;&nbsp;&nbsp;
                            <ClockIcon className={classes.clockIcon} />
                            {localTime}
                        </Typography>
                    </div>
                    <br />
                    <div
                        style={{
                            display: 'flex',
                            justifyContent: 'space-between',
                        }}
                    >
                        <div>
                            <Typography variant="caption">
                                Opgevraagd door
                            </Typography>
                            <Link to="">{data['source_organization']}</Link>
                        </div>
                        <div>
                            <Typography variant="caption" align="right">
                                Opgevraagd bij
                            </Typography>
                            <Typography align="right">
                                <Link to="">
                                    {data['destination_organization']}
                                </Link>
                            </Typography>
                        </div>
                    </div>
                    <br />
                    <Typography variant="caption">Reden</Typography>
                    {data['data']['doelbinding-process-id']
                        ? data['data']['doelbinding-process-id']
                        : 'Geen reden opgegeven.'}
                </Paper>
            </Modal>
        )
    }
}

SimpleModal.propTypes = {
    classes: PropTypes.object.isRequired,
}

export default withStyles(modalStyles)(SimpleModal)
