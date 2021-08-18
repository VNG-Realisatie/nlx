// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { RedocStandalone } from 'redoc'
import theme from '../../styling/theme'

const Documentation = ({ organizationName, serviceName }) => {
  const urlSafeOrganization = encodeURIComponent(organizationName)
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
  organizationName: string.isRequired,
  serviceName: string.isRequired,
}

export default Documentation
