// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'

import './ErrorPage.css'

const errorStateIsDefined = (location) =>
    location &&
    location.state &&
    location.state.error &&
    location.state.error.status

class ErrorPage extends Component {
    getErrorMessageFromState = () => {
        if (!errorStateIsDefined(this.props.location)) {
            return null
        }

        let { error } = this.props.location.state
        return (
            <h5>
                Error {error.status}: {error.description}
            </h5>
        )
    }

    render() {
        return (
            <div className="ErrorPage">
                <h1>Failed to load information</h1>
                <p>
                    Requested information is not available.
                    <br />
                    We apologize for any inconvenience.
                </p>
                {this.getErrorMessageFromState()}
                {this.props.children}
            </div>
        )
    }
}

export default ErrorPage
