// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { arrayOf, bool, func, oneOf, shape, string } from 'prop-types'
import { Alert } from '@commonground/design-system'

import Table from '../../components/Table'
import Amount from '../Amount'
import Collapsible from '../Collapsible'
import {
  StyledActionsBar,
  StyledDrawerHeading,
  StyledHeading,
  StyledInwayName,
  StyledLightHeading,
  StyledRemoveButton,
} from './index.styles'
import { ReactComponent as IconInway } from './inway.svg'
import { ReactComponent as IconWhitelist } from './whitelist.svg'
import { ReactComponent as IconVisible } from './visible.svg'
import { ReactComponent as IconHidden } from './hidden.svg'
import SectionGroup from './SectionGroup'

const ServiceDetails = ({ service, removeHandler }) => {
  const { name, internal, authorizationSettings, inways } = service
  const [isRemoved, setIsRemoved] = useState(false)
  const { t } = useTranslation()

  const handleRemove = () => {
    if (window.confirm(t('Do you want to remove the service?'))) {
      removeHandler()
      setIsRemoved(true)
    }
  }
  return (
    <>
      <StyledDrawerHeading data-testid="service-name">
        <h1>{name}</h1>
      </StyledDrawerHeading>
      {isRemoved ? (
        <Alert variant="success" data-testid="remove-success">
          {t('The service has been removed.')}
        </Alert>
      ) : (
        <>
          <StyledActionsBar>
            <StyledRemoveButton
              data-testid="remove-service"
              onClick={handleRemove}
            />
          </StyledActionsBar>
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
              {inways ? (
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
              ) : null}
            </Collapsible>
            {authorizationSettings.mode === 'whitelist' ? (
              <Collapsible
                title={
                  <StyledHeading data-testid="service-authorizations">
                    <IconWhitelist />
                    {t('Whitelisted organizations')}
                    <Amount
                      value={authorizationSettings.authorizations.length}
                    />
                  </StyledHeading>
                }
              >
                {authorizationSettings.authorizations ? (
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
                ) : null}
              </Collapsible>
            ) : null}
          </SectionGroup>
        </>
      )}
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
  removeHandler: func.isRequired,
}

ServiceDetails.defaultProps = {
  removeHandler: () => {},
}

export default ServiceDetails
