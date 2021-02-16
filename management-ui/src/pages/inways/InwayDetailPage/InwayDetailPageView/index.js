// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, arrayOf } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import Table from '../../../../components/Table'
import Amount from '../../../../components/Amount'
import {
  DetailHeading,
  StyledCollapsibleBody,
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
        <StyledIconInway inline />
        inway
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
            {services.length ? (
              <Table data-testid="inway-services-list" role="grid" withLinks>
                <tbody>
                  {services.map(({ name }) => (
                    <Table.Tr key={name} to={`/services/${name}`}>
                      <Table.Td>{name}</Table.Td>
                    </Table.Tr>
                  ))}
                </tbody>
              </Table>
            ) : (
              <small>{t('No services have been connected')}</small>
            )}
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
