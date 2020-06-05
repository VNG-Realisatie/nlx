// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Link, useLocation } from 'react-router-dom'
import { arrayOf, bool, func, oneOf, shape, string } from 'prop-types'
import { Alert, Table } from '@commonground/design-system'

import Amount from '../Amount'
import Collapsible from '../Collapsible'
import EditButton from '../EditButton'
import {
  StyledActionsBar,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
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
  const { internal, authorizationSettings, inways } = service
  const [isRemoved, setIsRemoved] = useState(false)
  const { t } = useTranslation()
  const location = useLocation()

  const handleRemove = () => {
    if (window.confirm(t('Do you want to remove the service?'))) {
      removeHandler()
      setIsRemoved(true)
    }
  }
  return isRemoved ? (
    <Alert variant="success" data-testid="remove-success">
      {t('The service has been removed.')}
    </Alert>
  ) : (
    <>
      <StyledActionsBar>
        <EditButton
          as={Link}
          to={`${location.pathname}/edit-service`}
          data-testid="edit-button"
        />
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
          ariaLabel={t('Inways')}
        >
          <StyledCollapsibleBody>
            <Table data-testid="service-inways-list" role="grid">
              <tbody>
                {inways.length ? (
                  inways.map((inway, i) => (
                    <Table.Tr key={i} data-testid={`service-inway-${i}`}>
                      <Table.Td>
                        <StyledInwayName>{inway}</StyledInwayName>
                      </Table.Td>
                    </Table.Tr>
                  ))
                ) : (
                  <Table.Tr data-testid="service-no-inways">
                    <Table.Td>
                      <StyledCollapsibleEmptyBody>
                        {t('No inways have been added')}
                      </StyledCollapsibleEmptyBody>
                    </Table.Td>
                  </Table.Tr>
                )}
              </tbody>
            </Table>
          </StyledCollapsibleBody>
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
            ariaLabel={t('Whitelisted organizations')}
          >
            <StyledCollapsibleBody>
              <Table data-testid="service-authorizations-list">
                <tbody>
                  {authorizationSettings.authorizations.length ? (
                    authorizationSettings.authorizations.map(
                      ({ organizationName }, i) => (
                        <Table.Tr
                          key={i}
                          data-testid={`service-authorization-${i}`}
                        >
                          <Table.Td>{organizationName}</Table.Td>
                        </Table.Tr>
                      ),
                    )
                  ) : (
                    <Table.Tr data-testid="service-no-authorizations">
                      <Table.Td>
                        <StyledCollapsibleEmptyBody>
                          {t('No organizations have been added')}
                        </StyledCollapsibleEmptyBody>
                      </Table.Td>
                    </Table.Tr>
                  )}
                </tbody>
              </Table>
            </StyledCollapsibleBody>
          </Collapsible>
        ) : null}
      </SectionGroup>
    </>
  )
}

ServiceDetails.propTypes = {
  service: shape({
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
