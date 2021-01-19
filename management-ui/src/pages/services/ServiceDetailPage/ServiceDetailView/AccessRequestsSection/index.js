// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { object, func } from 'prop-types'
import { Table, ToasterContext, Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import {
  DetailHeading,
  StyledCollapsibleBody,
} from '../../../../../components/DetailView'
import Amount from '../../../../../components/Amount'
import { IconKey } from '../../../../../icons'
import useStores from '../../../../../hooks/use-stores'
import usePolling from '../../../../../hooks/use-polling'
import IncomingAccessRequestRow from './IncomingAccessRequestRow'
import { StyledUpdateUiButton } from './index.styles'

export const POLLING_INTERVAL = 3000

const AccessRequestsSection = ({
  service,
  onApproveOrRejectCallbackHandler,
}) => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const rootStore = useStores()
  const [isOpen, setIsOpen] = useState(false)
  const [showUpdateUiButton, setShowUpdateUiButton] = useState(false)
  const [pausePollingClosed, startPollingClosed] = usePolling(() => {
    rootStore.incomingAccessRequestsStore.fetchForService({
      name: service.name,
    })
  }, POLLING_INTERVAL)

  const [pausePollingOpen, startPollingOpen] = usePolling(async () => {
    const haveAccessRequestsChanged = await rootStore.incomingAccessRequestsStore.haveChangedForService(
      service,
    )

    setShowUpdateUiButton(haveAccessRequestsChanged)

    if (haveAccessRequestsChanged) {
      pausePollingOpen()
    }
  }, POLLING_INTERVAL)

  const onClickUpdateIncomingData = async () => {
    setShowUpdateUiButton(false)

    await rootStore.incomingAccessRequestsStore.fetchForService({
      name: service.name,
    })

    if (isOpen) {
      startPollingOpen()
    }
  }

  useEffect(() => {
    if (isOpen) {
      pausePollingClosed()
      startPollingOpen()
    } else {
      pausePollingOpen()
      startPollingClosed()
    }
  }, [
    isOpen,
    pausePollingOpen,
    pausePollingClosed,
    startPollingOpen,
    startPollingClosed,
  ])

  const approveHandler = async (accessRequest) => {
    try {
      await accessRequest.approve()

      showToast({
        title: t('Access request approved'),
        body: t('Organization has access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })

      await onApproveOrRejectCallbackHandler()
    } catch (error) {
      showToast({
        title: t('Failed to approve access request'),
        body: t('Please try again'),
        variant: 'error',
      })
    }
  }

  const rejectHandler = async (accessRequest) => {
    try {
      await accessRequest.reject()

      showToast({
        title: t('Access request rejected'),
        body: t('Organization has been denied access to service', {
          organizationName: accessRequest.organizationName,
          serviceName: accessRequest.serviceName,
        }),
        variant: 'success',
      })

      await onApproveOrRejectCallbackHandler()
    } catch (error) {
      showToast({
        title: t('Failed to reject access request'),
        body: t('Please try again'),
        variant: 'error',
      })
    }
  }

  return (
    <Collapsible
      onToggle={(isCollapsibleOpen) => setIsOpen(isCollapsibleOpen)}
      title={
        <DetailHeading data-testid="service-incoming-accessrequests">
          <IconKey />
          {t('Access requests')}
          {service.incomingAccessRequests.length > 0 ? (
            <Amount
              data-testid="service-incoming-accessrequests-amount-accented"
              value={service.incomingAccessRequests.length}
              isAccented
            />
          ) : (
            <Amount
              data-testid="service-incoming-accessrequests-amount"
              value={service.incomingAccessRequests.length}
            />
          )}
        </DetailHeading>
      }
      ariaLabel={t('Access requests')}
    >
      <StyledCollapsibleBody>
        {service.incomingAccessRequests.length ? (
          <Table data-testid="service-incoming-accessrequests-list">
            <tbody>
              {service.incomingAccessRequests.map((accessRequest) => (
                <IncomingAccessRequestRow
                  key={accessRequest.id}
                  accessRequest={accessRequest}
                  approveHandler={approveHandler}
                  rejectHandler={rejectHandler}
                />
              ))}
            </tbody>
          </Table>
        ) : (
          <small>{t('There are no access requests')}</small>
        )}

        {showUpdateUiButton ? (
          <StyledUpdateUiButton onClick={onClickUpdateIncomingData}>
            Nieuwe verzoeken
          </StyledUpdateUiButton>
        ) : null}
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

AccessRequestsSection.propTypes = {
  service: object,
  onApproveOrRejectCallbackHandler: func.isRequired,
}

export default observer(AccessRequestsSection)
