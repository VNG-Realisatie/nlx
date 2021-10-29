// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { RedocStandalone } from 'redoc'
import theme from '../../theme'

const Documentation = ({ organizationSerialNumber, serviceName }) => {
  const urlSafeOrganization = encodeURIComponent(organizationSerialNumber)
  const urlSafeName = encodeURIComponent(serviceName)
  const specUrl = `/api/organizations/${urlSafeOrganization}/services/${urlSafeName}/api-spec`

  const style = {
    colors: { primary: { main: theme.tokens.colorBrand1 } },
    typography: { fontFamily: 'Source Sans Pro, sans-serif' },
  }

  return (
    <div style={{ background: '#ffffff' }}>
      <RedocStandalone
        options={{
          nativeScrollbars: true,
          theme: style,
          hideDownloadButton: false,
          hideLoading: true,
        }}
        specUrl={specUrl}
      />
    </div>
  )
}

Documentation.propTypes = {
  organizationSerialNumber: string.isRequired,
  serviceName: string.isRequired,
}

export default Documentation
