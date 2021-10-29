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
  const { serviceName, organization, state, apiSpecificationType } = service
  const { t } = useTranslation()

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organization.name}</SubTitle>
      <Summary>
        <StateIndicator state={state} showText />
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <span>
          {t('Serial Number serialNumber', {
            serialNumber: organization.serialNumber,
          })}
        </span>
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    serviceName: string.isRequired,
    organization: shape({
      serialNumber: string.isRequired,
      name: string.isRequired,
    }).isRequired,
    state: string.isRequired,
    apiSpecificationType: string,
  }),
}

export default observer(DrawerHeader)
