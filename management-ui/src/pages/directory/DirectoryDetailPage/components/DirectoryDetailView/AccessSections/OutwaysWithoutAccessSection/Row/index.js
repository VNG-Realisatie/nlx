// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { arrayOf, func, instanceOf, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import { ToasterContext } from '@commonground/design-system'
import Table from '../../../../../../../../components/Table'
import { OutwayName, Outways } from '../../components/index.styles'
import AccessState from '../../components/AccessState'
import OutwayModel from '../../../../../../../../stores/models/OutwayModel'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'
import { useConfirmationModal } from '../../../../../../../../components/ConfirmationModal'
import RequestAccessDetails from '../../components/RequestAccessDetails'

const Row = ({
  publicKeyFingerprint,
  publicKeyPEM,
  outways,
  service,
  onShowConfirmRequestAccessModalHandler,
  onHideConfirmRequestAccessModalHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestAccessDetails
        service={service}
        publicKeyFingerprint={publicKeyFingerprint}
        outwayNames={outways.map((outway) => outway.name)}
      />
    ),
  })

  const onRequestAccess = async () => {
    onShowConfirmRequestAccessModalHandler()

    if (await confirmRequest()) {
      try {
        await service.requestAccess(publicKeyPEM)
      } catch (err) {
        showToast({
          title: t('Failed to request access'),
          body:
            err.response && err.response.status === 403
              ? t(`You don't have the required permission.`)
              : err.message,
          variant: 'error',
        })
      }
    }

    onHideConfirmRequestAccessModalHandler()
  }

  const onRetryRequestAccess = () => {
    service.retryRequestAccess(publicKeyFingerprint)
  }

  const { accessRequest, accessProof } =
    service.getAccessStateFor(publicKeyFingerprint)

  return (
    <Table.Tr key={publicKeyFingerprint}>
      <Table.Td>
        <Outways>
          {outways.map((outway) => (
            <OutwayName key={outway.name}>{outway.name}</OutwayName>
          ))}
        </Outways>

        <AccessState
          accessRequest={accessRequest}
          accessProof={accessProof}
          onRequestAccess={onRequestAccess}
          onRetryRequestAccess={onRetryRequestAccess}
        />

        <RequestConfirmationModal />
      </Table.Td>
    </Table.Tr>
  )
}

Row.propTypes = {
  publicKeyFingerprint: string,
  publicKeyPEM: string,
  outways: arrayOf(instanceOf(OutwayModel)),
  service: instanceOf(DirectoryServiceModel),
  onShowConfirmRequestAccessModalHandler: func,
  onHideConfirmRequestAccessModalHandler: func,
}

export default observer(Row)
