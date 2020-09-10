// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, arrayOf } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Table } from '@commonground/design-system'

import Amount from '../../../../components/Amount'
import Collapsible from '../../../../components/Collapsible'
import {
  DetailHeading,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
  SectionGroup,
} from '../../../../components/DetailView'

import { IconServices } from '../../../../icons'
import { SubHeader, StyledIconInway, StyledSpecList } from './index.styles'

// Note: if inway- & outway details are interchangable, we can rename this to GatewayDetails
const InwayDetails = ({ inway }) => {
  const { t } = useTranslation()
  const { ipAddress, hostname, selfAddress, version, services } = inway

  return (
    <>
      <SubHeader data-testid="gateway-type">
        <StyledIconInway /> inway
      </SubHeader>

      <StyledSpecList data-testid="inway-specs" alignValuesRight>
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
                  services.map(({ name }) => (
                    <Table.Tr key={name} to={`/services/${name}`}>
                      <Table.Td>{name}</Table.Td>
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
    services: arrayOf(
      shape({
        name: string,
      }),
    ),
  }),
}

InwayDetails.defaultProps = {}

export default observer(InwayDetails)
