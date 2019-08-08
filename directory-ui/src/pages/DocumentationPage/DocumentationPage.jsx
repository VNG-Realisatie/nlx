// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import Documentation from '../../components/Documentation/Documentation'

const DocumentationPage = ({ match }) => {
  const { organizationName, serviceName } = match.params
  return <Documentation serviceName={serviceName} organizationName={organizationName} />
}

export default DocumentationPage
