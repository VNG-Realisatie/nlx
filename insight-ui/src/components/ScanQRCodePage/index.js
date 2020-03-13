// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { string } from 'prop-types'
import { QRCode } from 'react-qr-svg'
import { StyledCard } from './index.styles'

const ScanQRCodePage = ({ qrCodeValue }) => (
  <StyledCard>
    <p>
      Scan this QR code with the{' '}
      <a
        href="https://privacybydesign.foundation/download-en/"
        target="_blank"
        rel="noopener noreferrer"
      >
        IRMA app
      </a>{' '}
      to get access to your logs.
    </p>

    <p style={{ textAlign: 'center' }}>
      <QRCode
        bgColor="#FFFFFF"
        fgColor="#000000"
        level="Q"
        style={{ width: 200 }}
        value={qrCodeValue}
      />
    </p>

    <p className="text-muted">
      Read more about IRMA and what it does{' '}
      <a
        href="https://privacybydesign.foundation/irma/"
        target="_blank"
        rel="noopener noreferrer"
      >
        here
      </a>
      .
    </p>
  </StyledCard>
)

ScanQRCodePage.propTypes = {
  qrCodeValue: string.isRequired,
}

export default ScanQRCodePage
