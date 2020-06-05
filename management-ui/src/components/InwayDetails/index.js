// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, arrayOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Drawer, Table } from '@commonground/design-system'

import Amount from '../Amount'
import Collapsible from '../Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
  SectionGroup,
} from '../DetailView'

import { ReactComponent as IconServices } from './services.svg'
import { SubHeader, StyledIconInway, StyledSpecList } from './index.styles'

// Note: if inway- & outway details are interchangable, we can rename this to GatewayDetails
const InwayDetails = ({ inway }) => {
  const { t } = useTranslation()
  const { name, ipAddress, hostname, selfAddress, version, services } = inway

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

        <StyledSpecList data-testid="inway-specs" role="grid" alignValuesRight>
          <StyledSpecList.Item title={t('IP-address')} value={ipAddress} />
          <StyledSpecList.Item title={t('Hostname')} value={hostname} />
          <StyledSpecList.Item title={t('Self address')} value={selfAddress} />
          <StyledSpecList.Item title={t('Version')} value={version} />
        </StyledSpecList>

        <SectionGroup>
          <Collapsible
            title={
              <DetailHeading data-testid="inway-services">
                <IconServices />
                {t('Connected services')}
                <Amount value={services.length} />
              </DetailHeading>
            }
            ariaLabel={t('Connected services')}
          >
            <StyledCollapsibleBody>
              <Table data-testid="service-inways-list" role="grid" withLinks>
                <tbody>
                  {services.length ? (
                    services.map((service) => (
                      <Table.Tr key={service} to={`/services/${service}`}>
                        <Table.Td>{service}</Table.Td>
                      </Table.Tr>
                    ))
                  ) : (
                    <Table.Tr data-testid="service-no-inways">
                      <Table.Td>
                        <StyledCollapsibleEmptyBody>
                          {t('No services have been connected')}
                        </StyledCollapsibleEmptyBody>
                      </Table.Td>
                    </Table.Tr>
                  )}
                </tbody>
              </Table>
            </StyledCollapsibleBody>
          </Collapsible>
        </SectionGroup>
      </Drawer.Content>
    </>
  )
}

InwayDetails.propTypes = {
  inway: shape({
    name: string.isRequired,
    ipAddress: string,
    hostname: string,
    selfAddress: string,
    version: string,
    services: arrayOf(string),
  }),
}

InwayDetails.defaultProps = {}

export default InwayDetails
