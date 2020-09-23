// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, bool, func, shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Link, useLocation } from 'react-router-dom'
import { Table } from '@commonground/design-system'

import Amount from '../../../../components/Amount'
import Collapsible from '../../../../components/Collapsible'
import EditButton from '../../../../components/EditButton'
import {
  DetailHeading,
  DetailHeadingLight,
  SectionGroup,
  StyledCollapsibleBody,
  StyledCollapsibleEmptyBody,
} from '../../../../components/DetailView'
import { IconHidden, IconInway, IconVisible } from '../../../../icons'
import { showServiceVisibilityAlert } from '../../../../components/ServiceVisibilityAlert'
import {
  StyledActionsBar,
  StyledInwayName,
  StyledRemoveButton,
  StyledServiceVisibilityAlert,
} from './index.styles'

const ServiceDetailView = ({ service, removeHandler }) => {
  const { internal, inways } = service
  const { t } = useTranslation()
  const location = useLocation()

  const handleRemove = () => {
    if (window.confirm(t('Do you want to remove the service?'))) {
      removeHandler()
    }
  }
  return (
    <>
      {showServiceVisibilityAlert({ internal, inways }) ? (
        <StyledServiceVisibilityAlert />
      ) : null}

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
        <DetailHeadingLight data-testid="service-published">
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
        </DetailHeadingLight>

        <Collapsible
          title={
            <DetailHeading data-testid="service-inways">
              <IconInway />
              {t('Inways')}
              <Amount value={inways.length} />
            </DetailHeading>
          }
          ariaLabel={t('Inways')}
        >
          <StyledCollapsibleBody>
            <Table data-testid="service-inways-list" role="grid" withLinks>
              <tbody>
                {inways.length ? (
                  inways.map((inway, i) => (
                    <Table.Tr
                      key={i}
                      data-testid={`service-inway-${i}`}
                      to={`/inways/${inway}`}
                    >
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
      </SectionGroup>
    </>
  )
}

ServiceDetailView.propTypes = {
  service: shape({
    endpointURL: string,
    documentationURL: string,
    apiSpecificationURL: string,
    internal: bool.isRequired,
    techSupportContact: string,
    publicSupportContact: string,
    inways: arrayOf(string),
  }).isRequired,
  removeHandler: func.isRequired,
}

ServiceDetailView.defaultProps = {
  removeHandler: () => {},
}

export default observer(ServiceDetailView)
