// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
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
import CancelRequestAccessDetails from '../../components/WithdrawlRequestAccessDetails'

const Row = ({
  publicKeyFingerprint,
  publicKeyPem,
  outways,
  service,
  onShowConfirmRequestAccessModalHandler,
  onHideConfirmRequestAccessModalHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const [isLoading, setIsLoading] = useState(false)
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

  const [WithdrawRequestAccessConfirmationModal, confirmWithdrawRequestAccess] =
    useConfirmationModal({
      title: t('Withdraw access request'),
      okText: t('Confirm'),
      children: (
        <CancelRequestAccessDetails
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
        setIsLoading(true)
        await service.requestAccess(publicKeyPem)
      } catch (err) {
        let message = err.message

        if (err.response && err.response.status === 403) {
          message = t(`You don't have the required permission.`)
        }

        showToast({
          title: t('Failed to request access'),
          body: message,
          variant: 'error',
        })
      }

      setIsLoading(false)
    }

    onHideConfirmRequestAccessModalHandler()
  }

  const { accessRequest, accessProof } =
    service.getAccessStateFor(publicKeyFingerprint)

  const onWithdrawRequestAccess = async () => {
    onShowConfirmRequestAccessModalHandler()

    if (await confirmWithdrawRequestAccess()) {
      try {
        setIsLoading(true)
        await service.withdrawAccessRequest(accessRequest.publicKeyFingerprint)

        showToast({
          title: t('Access withdrawn'),
          variant: 'success',
        })
      } catch (err) {
        let message = err.message

        if (err.response && err.response.status === 403) {
          message = t(`You don't have the required permission.`)
        }

        showToast({
          title: t('Failed to withdraw request access'),
          body: message,
          variant: 'error',
        })
      }

      setIsLoading(false)
    }

    onHideConfirmRequestAccessModalHandler()
  }

  const onRetryRequestAccess = () => {
    return onRequestAccess()
  }

  const onWithdrawAccessButtonClick = () => {
    return onWithdrawRequestAccess()
  }

  return (
    <Table.Tr key={publicKeyFingerprint}>
      <Table.Td>
        <Outways>
          {outways.map((outway) => (
            <OutwayName key={outway.name}>{outway.name}</OutwayName>
          ))}
        </Outways>

        <AccessState
          isLoading={isLoading}
          accessRequest={accessRequest}
          accessProof={accessProof}
          onRequestAccess={onRequestAccess}
          onRetryRequestAccess={onRetryRequestAccess}
          onWithdrawAccessButtonClick={onWithdrawAccessButtonClick}
        />

        <RequestConfirmationModal />
        <WithdrawRequestAccessConfirmationModal />
      </Table.Td>
    </Table.Tr>
  )
}

Row.propTypes = {
  publicKeyFingerprint: string,
  publicKeyPem: string,
  outways: arrayOf(instanceOf(OutwayModel)),
  service: instanceOf(DirectoryServiceModel),
  onShowConfirmRequestAccessModalHandler: func,
  onHideConfirmRequestAccessModalHandler: func,
}

export default observer(Row)
