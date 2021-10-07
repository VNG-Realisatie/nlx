// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Drawer } from '@commonground/design-system'
import StateIndicator from '../../../../../components/StateIndicator'
import { SubTitle, Summary } from './index.styles'

const DrawerHeader = ({ service }) => {
  const { name, status, apiType, organization } = service

  return (
    <header data-testid="directory-detail-header">
      <Drawer.Header title={name} closeButtonLabel="Close" />
      <SubTitle>{organization.name}</SubTitle>
      <Summary>
        <StateIndicator state={status} showText />
        {apiType && <span>{apiType}</span>}
        <span>Serienummer {organization.serialNumber}</span>
      </Summary>
    </header>
  )
}

DrawerHeader.propTypes = {
  service: shape({
    name: string.isRequired,
    organization: shape({
      name: string.isRequired,
      serialNumber: string.isRequired,
    }),
    status: string.isRequired,
    apiType: string,
  }),
}

export default DrawerHeader
