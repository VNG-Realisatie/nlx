// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useConfirmationModal } from '../../../../../../../components/ConfirmationModal'
import { incomingAccessRequestPropTypes } from '../../../../../../../stores/models/IncomingAccessRequestModel'
import { IconCheck, IconClose } from '../../../../../../../icons'
import { TdActions, StyledButton } from './index.styles'

const IncomingAccessRequestRow = ({
  accessRequest,
  approveHandler,
  rejectHandler,
}) => {
  const { t } = useTranslation()
  const { id, organizationName, serviceName } = accessRequest

  const [ConfirmApproveModal, confirmApprove] = useConfirmationModal({
    okText: t('Approve'),
    children: (
      <p>
        {t(
          'Approving this access request will grant organizationName access to serviceName. Are you sure?',
          {
            organizationName,
            serviceName,
          },
        )}
      </p>
    ),
  })

  const [ConfirmRejectModal, confirmReject] = useConfirmationModal({
    okText: t('Reject'),
    children: (
      <p>
        {t(
          'Rejecting this access request will refuse organizationName access to serviceName. Are you sure?',
          {
            organizationName,
            serviceName,
          },
        )}
      </p>
    ),
  })

  const approve = async () => {
    if (await confirmApprove()) {
      approveHandler(accessRequest)
    }
  }

  const reject = async () => {
    if (await confirmReject()) {
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
          onClick={approve}
          title={t('Approve')}
        >
          <IconCheck />
        </StyledButton>
        <StyledButton
          size="small"
          variant="link"
          onClick={reject}
          title={t('Reject')}
        >
          <IconClose />
        </StyledButton>

        <ConfirmApproveModal />
        <ConfirmRejectModal />
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
