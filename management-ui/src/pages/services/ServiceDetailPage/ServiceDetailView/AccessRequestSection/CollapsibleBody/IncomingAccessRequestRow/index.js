// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, func, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../../../components/Table'
import { useConfirmationModal } from '../../../../../../../components/ConfirmationModal'
import { IconCheck, IconClose } from '../../../../../../../icons'
import { TdActions, StyledButton } from './index.styles'
import AccessDetails from './components/AccessDetails'

const IncomingAccessRequestRow = ({
  accessRequest,
  approveHandler,
  rejectHandler,
}) => {
  const { t } = useTranslation()
  const { id, organization, serviceName, publicKeyFingerprint } = accessRequest

  const [ConfirmApproveModal, confirmApprove] = useConfirmationModal({
    title: t('Grant access'),
    okText: t('Approve'),
    children: (
      <AccessDetails
        subTitle={t(
          'Approving this access request will grant this organization access to the service. Are you sure?',
        )}
        organization={organization}
        serviceName={serviceName}
        publicKeyFingerprint={publicKeyFingerprint}
      />
    ),
  })

  const [ConfirmRejectModal, confirmReject] = useConfirmationModal({
    title: t('Reject access'),
    okText: t('Reject'),
    children: (
      <AccessDetails
        subTitle={t(
          'Rejecting this access request will refuse the organization access to the service. Are you sure?',
        )}
        organization={organization}
        serviceName={serviceName}
        publicKeyFingerprint={publicKeyFingerprint}
      />
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
      <Table.Td>
        {organization.name}
        <br />
        <small>
          {t('Serial Number serialNumber', {
            serialNumber: organization.serialNumber,
          })}
        </small>
        <br />
        <small>
          {t('Public Key Fingerprint publicKeyFingerprint', {
            publicKeyFingerprint: publicKeyFingerprint,
          })}
        </small>
      </Table.Td>
      <TdActions>
        <StyledButton size="small" variant="link" onClick={approve}>
          <IconCheck title={t('Approve')} />
        </StyledButton>
        <StyledButton size="small" variant="link" onClick={reject}>
          <IconClose title={t('Reject')} />
        </StyledButton>

        <ConfirmApproveModal />
        <ConfirmRejectModal />
      </TdActions>
    </Table.Tr>
  )
}

IncomingAccessRequestRow.propTypes = {
  accessRequest: shape({
    id: string,
    organization: shape({
      serialNumber: string.isRequired,
      name: string.isRequired,
    }).isRequired,
    serviceName: string.isRequired,
    publicKeyFingerprint: string.isRequired,
  }).isRequired,
  approveHandler: func.isRequired,
  rejectHandler: func.isRequired,
}

export default IncomingAccessRequestRow
