import React, { Component } from 'react'
import { QRCode } from 'react-qr-svg'

import { connect } from 'react-redux'

import Spinner from '../../components/Spinner'
// import logGroup from '../../utils/logGroup'
import * as actionType from '../../store/actions'

export class IrmaAuthPage extends Component {
    /**
     * Send user to error page
     * @param {number} error.number error number
     * @param {string} error.description error description
     */

    showError = (error) => {
        let { history, match } = this.props
        // let url = `/organization/${match.params.name}/error`
        let url = match.path.replace('login', 'error')
        history.push(url, {
            error,
        })
    }

    startLogin = () => {
        let {
            dispatch,
            qrCode,
            loginInProgress,
            info,
            error,
            history,
            jwt,
        } = this.props

        if (error) {
            // debugger
            this.showError(error)
        } else if (qrCode && !loginInProgress && jwt) {
            // debugger
            let url = `/organization/${info.name}/view`
            history.push(url)
        } else if (qrCode && !loginInProgress && !jwt) {
            // debugger
            dispatch({
                type: actionType.IRMA_LOGIN_START,
            })
        }
    }

    stopLogin = () => {
        let { dispatch, loginInProgress, error } = this.props

        if (error || loginInProgress) {
            // leaving login page due to forward to error page
            // or user navigating away from login in progress
            // debugger
            dispatch({
                type: actionType.RESET_ORGANIZATION,
            })
        } else {
            // debugger
        }
    }

    render() {
        let { qrCode } = this.props

        // logGroup({
        //     title: 'IrmaAuthPage',
        //     method: 'render',
        //     props: this.props,
        //     state: this.state,
        // })

        if (qrCode) {
            return (
                <div>
                    Please scan QR code with Irma app
                    <br />
                    <br />
                    <QRCode
                        bgColor="#FFFFFF"
                        fgColor="#000000"
                        level="Q"
                        style={{ width: 256 }}
                        value={qrCode}
                    />
                </div>
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
