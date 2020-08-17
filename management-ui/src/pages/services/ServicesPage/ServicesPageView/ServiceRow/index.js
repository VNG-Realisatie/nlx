// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, array, string, bool } from 'prop-types'
import { Table } from '@commonground/design-system'

import {
  ServiceVisibilityMessage,
  showServiceVisibilityAlert,
} from '../../../../../components/ServiceVisibilityAlert'
import AuthorizationMode from './AuthorizationMode'
import { TdAlignRight } from './index.styles'

const ServiceRow = ({ service, ...props }) => {
  const { name, authorizationSettings, internal, inways } = service
  const { authorizations, mode } = authorizationSettings

  return (
    <Table.Tr
      to={`/services/${name}`}
      name={name}
      data-testid="service-row"
      {...props}
    >
      <Table.Td>{name}</Table.Td>
      <Table.Td>
        <AuthorizationMode authorizations={authorizations} mode={mode} />
      </Table.Td>
      <TdAlignRight data-testid="warning-cell">
        {showServiceVisibilityAlert({ internal, inways }) ? (
          <ServiceVisibilityMessage />
        ) : null}
      </TdAlignRight>
    </Table.Tr>
  )
}

ServiceRow.propTypes = {
  service: shape({
    name: string.isRequired,
    authorizationSettings: shape({
      mode: string,
      authorizations: array,
    }),
    internal: bool.isRequired,
    inways: array.isRequired,
  }),
}

export default ServiceRow
