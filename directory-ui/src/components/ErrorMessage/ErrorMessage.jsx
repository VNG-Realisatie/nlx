import React from 'react'

import './ErrorMessage.scss'

const ErrorMessage = () => (
  <div className="ErrorMessage">
    <h1>Failed to load information</h1>
    <p>
      Requested information is not available.
      <br />
      We apologize for any inconvenience.
    </p>
  </div>
)

export default ErrorMessage
