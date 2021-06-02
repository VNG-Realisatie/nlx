// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { RedocStandalone } from 'redoc'

const Documentation = ({ organizationName, serviceName }) => {
  const urlSafeOrganization = encodeURIComponent(organizationName)
  const urlSafeName = encodeURIComponent(serviceName)
  const specUrl = `/api/organizations/${urlSafeOrganization}/services/${urlSafeName}/api-spec`

  return (
    <div style={{ background: '#ffffff' }}>
      <RedocStandalone
        specUrl={specUrl}
        options={{
          hideDownloadButton: false,
          hideLoading: true,
        }}
      />
    </div>
  )
}

Documentation.propTypes = {
  organizationName: string.isRequired,
  serviceName: string.isRequired,
}

export default Documentation
