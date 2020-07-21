// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'

import StateIndicator from '../../../../components/StateIndicator'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service }) => {
  const { serviceName, organizationName, state, apiSpecificationType } = service
  const { t } = useTranslation()

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organizationName}</SubTitle>
      <Summary>
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <StateIndicator state={state} showText />
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    serviceName: string,
    organizationName: string,
    state: string,
    apiSpecificationType: string,
  }),
}

export default DrawerHeader
