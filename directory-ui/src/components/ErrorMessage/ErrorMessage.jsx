import React from 'react'
import { StyledErrorMessage } from './ErrorMessage.styles'

const ErrorMessage = () => (
    <StyledErrorMessage>
        <h1>Failed to load information</h1>
        <p>
            Requested information is not available.
            <br />
            We apologize for any inconvenience.
        </p>
    </StyledErrorMessage>
)

export default ErrorMessage
