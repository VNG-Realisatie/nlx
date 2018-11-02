import React from 'react'
import PropTypes from 'prop-types'

import { withStyles, Modal, Paper, IconButton } from '@material-ui/core'

import modalStyles from '../styles/SimpleModal'
import { Close } from '@material-ui/icons'

class SimpleModal extends React.Component {
    render() {
        const { children, classes, open, closeModal } = this.props
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
                    {children}
                </Paper>
            </Modal>
        )
    }
}

SimpleModal.propTypes = {
    classes: PropTypes.object.isRequired,
}

export default withStyles(modalStyles)(SimpleModal)
