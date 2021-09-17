// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'
import StateIndicator from '../../../../../components/StateIndicator'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service }) => {
  const {
    serviceName,
    organizationName,
    state,
    apiSpecificationType,
    serialNumber,
  } = service
  const { t } = useTranslation()

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organizationName}</SubTitle>
      <Summary>
        <StateIndicator state={state} showText />
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <span>{t('Serial Number serialNumber', { serialNumber })}</span>
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    serviceName: string.isRequired,
    organizationName: string.isRequired,
    state: string.isRequired,
    apiSpecificationType: string,
    serialNumber: string.isRequired,
  }),
}

export default observer(DrawerHeader)
