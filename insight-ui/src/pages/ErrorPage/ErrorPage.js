import React, { Component } from 'react'

import './ErrorPage.css'

class ErrorPage extends Component {
    /**
     * Extract oError passed by router as state.
     */
    getStateContent = () => {
        if (!this.props.location) {
            return null
        } else if (
            this.props.location.state &&
            this.props.location.state.error
        ) {
            let { error } = this.props.location.state
            if (error && error.status) {
                return (
                    <h5>
                        Error {error.status}: {error.description}
                    </h5>
                )
            } else {
                return null
            }
        } else {
            return null
        }
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
                {this.getStateContent()}
                {this.props.children}
            </div>
        )
    }
}

export default ErrorPage
