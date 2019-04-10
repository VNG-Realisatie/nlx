import React from 'react'
import { StyledCard } from './index.styles'

const LoginPage = () =>
  <StyledCard>
    <p>
      Scan this QR code with the <a href="https://privacybydesign.foundation/download-en/" target="_blank" rel="noopener noreferrer">IRMA app</a> to get access to your logs.
    </p>

    <p>
      TODO: QR code
    </p>

    <p className="text-muted">
      Read more about IRMA and what it does <a href="https://privacybydesign.foundation/irma/" target="_blank" rel="noopener noreferrer">here</a>.
    </p>
  </StyledCard>

export default LoginPage
