// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service, ...props }) => {
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
        {apiSpecificationType && <p>{apiSpecificationType}</p>}
        <p>{status}</p>
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    serviceName: string.isRequired,
    organizationName: string.isRequired,
  }),
}

DrawerHeader.defaultProps = {}

export default DrawerHeader
