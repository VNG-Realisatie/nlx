// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { instanceOf, bool } from 'prop-types'
import { useTheme } from 'styled-components'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../components/Table'
import StateIndicator from '../../../../../components/StateIndicator'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import { IconCheck, IconWarning } from '../../../../../icons'
import { StyledIcon } from '../../../../../components/GlobalAlert/index.styles'
import {
  StyledTd,
  AccessMessageWrapper,
  StyledTdAccess,
  Message,
  WarnMessage,
} from './index.styles'

const DirectoryServiceRow = ({ service, ownService, ...props }) => {
  const theme = useTheme()
  const { t } = useTranslation()

  const failingAccessStates = service.getFailingAccessStates()
  const hasAccessWithAtLeastOneOutway =
    service.accessStatesWithAccess.length >= 1

  const ownServiceColor = theme.tokens.colorBrand1

  return (
    <Table.Tr
      to={`/directory/${service.organization.serialNumber}/${service.serviceName}`}
      name={`${service.organization.name} - ${service.serviceName}`}
      data-testid="directory-service-row"
      {...props}
    >
      <StyledTd color={ownService ? ownServiceColor : null}>
        {service.organization.name}
      </StyledTd>
      <Table.Td>{service.serviceName}</Table.Td>
      <Table.Td>
        <StateIndicator state={service.state} showText={false} />
      </Table.Td>
      <Table.Td>{service.apiSpecificationType}</Table.Td>
      <StyledTdAccess>
        <AccessMessageWrapper>
          {failingAccessStates.length ? (
            <WarnMessage>{t('Request could not be sent')}</WarnMessage>
          ) : service.outgoingAccessRequestsSyncError ? (
            <Message>
              <StyledIcon
                as={IconWarning}
                inline
                title={t(service.outgoingAccessRequestsSyncError.message)}
              />
            </Message>
          ) : hasAccessWithAtLeastOneOutway ? (
            <Message>
              <IconCheck title={t('Approved')} inline />
            </Message>
          ) : null}
        </AccessMessageWrapper>
      </StyledTdAccess>
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  selected: bool,
  ownService: bool,
}

export default observer(DirectoryServiceRow)
