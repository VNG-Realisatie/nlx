// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { object } from 'prop-types'
import Documentation from '../../components/Documentation'

const DocumentationPage = ({ match }) => {
  const { organizationSerialNumber, serviceName } = match.params
  return (
    <Documentation
      serviceName={serviceName}
      organizationSerialNumber={organizationSerialNumber}
    />
  )
}

DocumentationPage.propTypes = {
  match: object,
}

export default DocumentationPage
