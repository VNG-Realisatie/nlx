// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'

import {
  SubTitle,
  Summary,
  StyledSpinner,
  StyledStateIndicator,
} from './index.styles'

const DrawerHeader = ({ service }) => {
  const {
    serviceName,
    organizationName,
    state,
    apiSpecificationType,
    isLoading,
  } = service
  const { t } = useTranslation()

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organizationName}</SubTitle>
      <Summary>
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <StyledStateIndicator state={state} showText />
        {isLoading && <StyledSpinner />}
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

export default observer(DrawerHeader)
