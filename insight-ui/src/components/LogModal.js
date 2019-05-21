// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import PropTypes from 'prop-types'

import {
    withStyles,
    Modal,
    Paper,
    IconButton,
    Typography,
} from '@material-ui/core'

import modalStyles from '../styles/LogModal'
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

                    <Typography variant="caption">Tijdstip</Typography>
                    <Typography variant="body2" gutterBottom>
                        <CalendarToday className={classes.calendarIcon} />
                        {localDate}
                        &nbsp;&nbsp;&nbsp;
                        <ClockIcon className={classes.clockIcon} />
                        {localTime}
                    </Typography>

                    <div
                        style={{
                            display: 'flex',
                            justifyContent: 'space-between',
                        }}
                    >
                        <div>
                            <Typography variant="caption">
                                Applicatie
                            </Typography>
                            <Typography variant="body2" gutterBottom>
                                {data['data']['doelbinding-application-id']
                                    ? data['data']['doelbinding-application-id']
                                    : 'Geen applicatie opgegeven.'}
                            </Typography>
                        </div>
                        <div>
                            <Typography variant="caption" align="right">
                                Process
                            </Typography>
                            <Typography
                                variant="body2"
                                align="right"
                                gutterBottom
                            >
                                {data['data']['doelbinding-process-id']
                                    ? data['data']['doelbinding-process-id']
                                    : 'Geen process opgegeven.'}
                            </Typography>
                        </div>
                    </div>

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
                            <Typography variant="body2" gutterBottom>
                                {data['source_organization']}
                            </Typography>
                        </div>
                        <div>
                            <Typography variant="caption" align="right">
                                Opgevraagd bij
                            </Typography>
                            <Typography
                                variant="body2"
                                align="right"
                                gutterBottom
                            >
                                {data['destination_organization']}
                            </Typography>
                        </div>
                    </div>
                    <div
                        style={{
                            display: 'flex',
                            justifyContent: 'space-between',
                        }}
                    >
                        <div>
                            <Typography variant="caption">
                                Data elementen
                            </Typography>
                            <Typography variant="body2" gutterBottom>
                                {data['data']['doelbinding-data-elements']
                                    ? data['data']['doelbinding-data-elements']
                                    : 'Geen data element opgevraagd.'}
                            </Typography>
                        </div>
                        <div>
                            <Typography variant="caption" align="right">
                                Logrecord id
                            </Typography>
                            <Typography
                                variant="body2"
                                align="right"
                                gutterBottom
                            >
                                {data['logrecord-id']}
                            </Typography>
                        </div>
                    </div>
                </Paper>
            </Modal>
        )
    }
}

SimpleModal.propTypes = {
    classes: PropTypes.object.isRequired,
}

export default withStyles(modalStyles)(SimpleModal)
