// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer } from '@commonground/design-system'

import Table from '../SpecTable'
import { SubHeader, StyledIconInway } from './index.styles'

// Note: if inway- & outway details are interchangable, we can rename this to GatewayDetails
const InwayDetails = ({ inway }) => {
  const { t } = useTranslation()
  const { name, hostname, selfAddress } = inway

  return (
    <>
      <Drawer.Header
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="service-name"
      />

      <Drawer.Content>
        <SubHeader data-testid="gateway-type">
          <StyledIconInway /> inway
        </SubHeader>

        <Table data-testid="inway-specs" role="grid" valueAlignRight>
          <tbody>
            <Table.Tr>
              <Table.Td>{t('IP-address')}</Table.Td>
              <Table.Td>123.456.789.1</Table.Td>
            </Table.Tr>
            <Table.Tr>
              <Table.Td>{t('Hostname')}</Table.Td>
              <Table.Td>{hostname}</Table.Td>
            </Table.Tr>
            <Table.Tr>
              <Table.Td>{t('Self address')}</Table.Td>
              <Table.Td>{selfAddress}</Table.Td>
            </Table.Tr>
          </tbody>
        </Table>
      </Drawer.Content>
    </>
  )
}

InwayDetails.propTypes = {
  inway: shape({
    name: string.isRequired,
    hostname: string,
    selfAddress: string,
    version: string,
  }),
}

InwayDetails.defaultProps = {}

export default InwayDetails
