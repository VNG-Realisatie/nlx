// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { arrayOf, bool, oneOf, shape, string } from 'prop-types'
import Table from '../../pages/ServicesPage/Table'
import Amount from '../Amount'
import Collapsible from '../Collapsible'
import {
  StyledDrawerHeading,
  StyledHeading,
  StyledInwayName,
  StyledLightHeading,
} from './index.styles'
import { ReactComponent as IconInway } from './inway.svg'
import { ReactComponent as IconWhitelist } from './whitelist.svg'
import { ReactComponent as IconVisible } from './visible.svg'
import { ReactComponent as IconHidden } from './hidden.svg'
import SectionGroup from './SectionGroup'

const ServiceDetails = ({ service }) => {
  const { name, internal, authorizationSettings, inways } = service
  const { t } = useTranslation()

  return (
    <>
      <StyledDrawerHeading data-testid="service-name">
        <h1>{name}</h1>
      </StyledDrawerHeading>
      <SectionGroup>
        <StyledLightHeading data-testid="service-published">
          {internal ? (
            <>
              <IconHidden />
              {t('Not visible in central directory')}
            </>
          ) : (
            <>
              <IconVisible />
              {t('Published in central directory')}
            </>
          )}
        </StyledLightHeading>
        <Collapsible
          title={
            <StyledHeading data-testid="service-inways">
              <IconInway />
              {t('Inways')}
              <Amount value={inways.length} />
            </StyledHeading>
          }
        >
          <Table data-testid="service-inways-list" role="grid">
            <tbody>
              {inways.map((inway, i) => (
                <Table.Tr key={i} data-testid={`service-inway-${i}`}>
                  <Table.Td>
                    <StyledInwayName>{inway}</StyledInwayName>
                  </Table.Td>
                </Table.Tr>
              ))}
            </tbody>
          </Table>
        </Collapsible>
        {authorizationSettings.mode === 'whitelist' ? (
          <Collapsible
            title={
              <StyledHeading data-testid="service-authorizations">
                <IconWhitelist />
                {t('Whitelisted organizations')}
                <Amount value={authorizationSettings.authorizations.length} />
              </StyledHeading>
            }
          >
            <Table data-testid="service-authorizations-list">
              <tbody>
                {authorizationSettings.authorizations.map(
                  ({ organizationName }, i) => (
                    <Table.Tr
                      key={i}
                      data-testid={`service-authorization-${i}`}
                    >
                      <Table.Td>{organizationName}</Table.Td>
                    </Table.Tr>
                  ),
                )}
              </tbody>
            </Table>
          </Collapsible>
        ) : null}
      </SectionGroup>
    </>
  )
}

ServiceDetails.propTypes = {
  service: shape({
    name: string.isRequired,
    endpointURL: string,
    documentationURL: string,
    apiSpecificationURL: string,
    internal: bool.isRequired,
    techSupportContact: string,
    publicSupportContact: string,
    authorizationSettings: shape({
      mode: oneOf(['whitelist', 'none']),
      authorizations: arrayOf(
        shape({ organizationName: string, publicKeyHash: string }),
      ),
    }).isRequired,
    inways: arrayOf(string),
  }).isRequired,
}

export default ServiceDetails
