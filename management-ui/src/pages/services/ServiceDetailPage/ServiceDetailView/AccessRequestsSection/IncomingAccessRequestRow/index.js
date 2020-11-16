// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { incomingAccessRequestPropTypes } from '../../../../../../models/IncomingAccessRequestModel'
import { IconCheck, IconClose } from '../../../../../../icons'
import { TdActions, StyledButton } from './index.styles'

const IncomingAccessRequestRow = ({
  accessRequest,
  approveHandler,
  rejectHandler,
}) => {
  const { t } = useTranslation()
  const { id, organizationName, serviceName } = accessRequest

  const handleApproveButtonClick = (event) => {
    event.stopPropagation()
    approve()
  }

  const handleRejectButtonClick = (event) => {
    event.stopPropagation()
    reject()
  }

  const approve = () => {
    const confirmed = window.confirm(
      t(
        'Approving this access request will grant organizationName access to serviceName. Are you sure?',
        {
          organizationName,
          serviceName,
        },
      ),
    )

    if (confirmed) {
      approveHandler(accessRequest)
    }
  }

  const reject = () => {
    const confirmed = window.confirm(
      t(
        'Rejecting this access request will refuse organizationName access to serviceName. Are you sure?',
        {
          organizationName,
          serviceName,
        },
      ),
    )

    if (confirmed) {
      rejectHandler(accessRequest)
    }
  }

  return (
    <Table.Tr data-testid={`service-incoming-accessrequest-${id}`}>
      <Table.Td>{organizationName}</Table.Td>
      <TdActions>
        <StyledButton
          size="small"
          variant="link"
          onClick={handleApproveButtonClick}
          title={t('Approve')}
        >
          <IconCheck />
        </StyledButton>
        <StyledButton
          size="small"
          variant="link"
          onClick={handleRejectButtonClick}
          title={t('Reject')}
        >
          <IconClose />
        </StyledButton>
      </TdActions>
    </Table.Tr>
  )
}

IncomingAccessRequestRow.propTypes = {
  accessRequest: shape(incomingAccessRequestPropTypes).isRequired,
  approveHandler: func.isRequired,
  rejectHandler: func.isRequired,
}

export default IncomingAccessRequestRow
