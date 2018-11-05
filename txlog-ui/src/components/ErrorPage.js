import React, { Component } from 'react'

import './ErrorPage.css'

class ErrorPage extends Component {
    render() {
        return (
            <div className="app-error-page">
                <h1>Failed to load information</h1>
                <p>
                    Requested information is not available.
                    <br />
                    We apologize for any inconvenience.
                </p>
            </div>
        )
    }
}

export default ErrorPage
