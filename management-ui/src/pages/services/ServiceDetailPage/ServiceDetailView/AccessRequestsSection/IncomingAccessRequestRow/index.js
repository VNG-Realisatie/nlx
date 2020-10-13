// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { incomingAccessRequestPropTypes } from '../../../../../../models/IncomingAccessRequestModel'
import ButtonWithIcon from '../../../../../../components/ButtonWithIcon'
import { IconCheck } from '../../../../../../icons'
import { TdActions } from './index.styles'

const IncomingAccessRequestRow = ({ accessRequest, approveHandler }) => {
  const { t } = useTranslation()
  const { id, organizationName, serviceName } = accessRequest

  const handleApproveButtonClick = (event) => {
    event.stopPropagation()
    approve()
  }

  const approve = () => {
    const confirmed = window.confirm(
      t(
        'Approving this access request will grant {{organizationName}} access to {{serviceName}}. Are you sure?',
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

  return (
    <Table.Tr data-testid={`service-incoming-accessrequest-${id}`}>
      <Table.Td>{organizationName}</Table.Td>
      <TdActions>
        <ButtonWithIcon
          size="small"
          variant="link"
          onClick={handleApproveButtonClick}
        >
          <IconCheck />
          {t('Approve')}
        </ButtonWithIcon>
      </TdActions>
    </Table.Tr>
  )
}

IncomingAccessRequestRow.propTypes = {
  accessRequest: shape(incomingAccessRequestPropTypes).isRequired,
  approveHandler: func.isRequired,
}

export default IncomingAccessRequestRow
