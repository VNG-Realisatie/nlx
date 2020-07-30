// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import { SubTitle, Summary, StyledStateIndicator } from './index.styles'

const DrawerHeader = ({ service }) => {
  const { serviceName, organizationName, state, apiSpecificationType } = service
  const { t } = useTranslation()

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={serviceName} closeButtonLabel={t('Close')} />
      <SubTitle>{organizationName}</SubTitle>
      <Summary>
        {apiSpecificationType && <span>{apiSpecificationType}</span>}
        <StyledStateIndicator state={state} showText />
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape(
    pick(directoryServicePropTypes, [
      'serviceName',
      'organizationName',
      'state',
      'apiSpecificationType',
    ]),
  ),
}

export default observer(DrawerHeader)
