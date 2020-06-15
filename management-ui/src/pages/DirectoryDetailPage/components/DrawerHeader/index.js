// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'

import StatusIndicator from '../../../../components/StatusIndicator'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service }) => {
  const {
    serviceName,
    organizationName,
    status,
    apiSpecificationType,
  } = service
  const { t } = useTranslation()

  return (
    <header data-testid="service-organisation-name">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organizationName}</SubTitle>
      <Summary>
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <StatusIndicator status={status} showText />
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    serviceName: string,
    organizationName: string,
    status: string,
    apiSpecificationType: string,
  }),
}

export default DrawerHeader
