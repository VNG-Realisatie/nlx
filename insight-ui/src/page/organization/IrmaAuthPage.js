import React, { Component } from 'react'
import { QRCode } from 'react-qr-svg'
import { Typography } from '@material-ui/core'
import { connect } from 'react-redux'
import { Spinner } from '@commonground/design-system'

import * as actionType from '../../store/actions'

export class IrmaAuthPage extends Component {
    /**
     * Send user to error page
     * @param {number} error.number error number
     * @param {string} error.description error description
     */

    showError = (error) => {
        const { history, match } = this.props
        const url = match.path.replace('login', 'error')
        history.push(url, {
            error,
        })
    }

    startLogin = () => {
        const {
            dispatch,
            qrCode,
            loginInProgress,
            info,
            error,
            history,
            jwt,
        } = this.props

        if (error) {
            this.showError(error)
        } else if (qrCode && !loginInProgress && jwt) {
            let url = `/organization/${info.name}/view`
            history.push(url)
        } else if (qrCode && !loginInProgress && !jwt) {
            dispatch({
                type: actionType.IRMA_LOGIN_START,
            })
        }
    }

    stopLogin = () => {
        const { dispatch, loginInProgress, error } = this.props

        if (error || loginInProgress) {
            // leaving login page due to forward to error page
            // or user navigating away from login in progress
            dispatch({
                type: actionType.RESET_ORGANIZATION,
            })
        }
    }

    render() {
        const { qrCode } = this.props

        if (qrCode) {
            return (
                <React.Fragment>
                    <Typography variant="h6" color="default" gutterBottom>
                        Scan QR code with IRMA app
                    </Typography>
                    <QRCode
                        bgColor="#FFFFFF"
                        fgColor="#000000"
                        level="Q"
                        style={{ width: 256 }}
                        value={qrCode}
                    />
                </React.Fragment>
            )
        } else {
            return <Spinner />
        }
    }

    componentDidMount() {
        this.startLogin()
    }
    componentDidUpdate() {
        this.startLogin()
    }
    componentWillUnmount() {
        this.stopLogin()
    }
}

// -------------- REDUX CONNECTION ---------------------
/**
 * Map redux store states to local component properties
 * @param state: object, redux store object
 */
const mapStateToProps = (state) => {
    // debugger
    let props = {
        loading: true,
        info: state.organization.info,
        irma: state.organization.irma,
        qrCode: state.organization.irma.qrCode,
        loginInProgress: state.organization.irma.inProgress,
        jwt: state.organization.irma.jwt,
        error: state.organization.irma.error,
    }
    if (
        state.organization.irma.qrCode ||
        state.organization.irma.error ||
        state.organization.irma.jwt
    ) {
        props.loading = state.loader.show
    }
    return props
}

export default connect(mapStateToProps)(IrmaAuthPage)
