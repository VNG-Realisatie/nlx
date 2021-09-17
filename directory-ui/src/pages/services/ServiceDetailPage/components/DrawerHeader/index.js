// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Drawer } from '@commonground/design-system'
import StateIndicator from '../../../../../components/StateIndicator'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service }) => {
  const { name, organization, status, apiType, serialNumber } = service

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={name} closeButtonLabel="Close" />
      <SubTitle>{organization}</SubTitle>
      <Summary>
        <StateIndicator state={status} showText />
        {apiType && <span>{apiType}</span>}
        <span>Serienummer {serialNumber}</span>
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    name: string.isRequired,
    organization: string.isRequired,
    status: string.isRequired,
    apiType: string,
  }),
}

export default DrawerHeader
