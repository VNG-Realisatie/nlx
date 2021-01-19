// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { object } from 'prop-types'
import { Collapsible } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import useStores from '../../../../../hooks/use-stores'
import usePolling from '../../../../../hooks/use-polling'
import CollapsibleHeader from './CollapsibleHeader'
import CollapsibleBody from './CollapsibleBody'

export const POLLING_INTERVAL = 3000

const AccessRequestsSection = ({ service }) => {
  const { t } = useTranslation()
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

  const onLoadIncomingDataHandler = async () => {
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

  return (
    <Collapsible
      onToggle={(isCollapsibleOpen) => setIsOpen(isCollapsibleOpen)}
      title={
        <CollapsibleHeader counter={service.incomingAccessRequests.length} />
      }
      ariaLabel={t('Access requests')}
    >
      <CollapsibleBody
        accessRequests={service.incomingAccessRequests}
        showLoadIncomingDataButton={showUpdateUiButton}
        onClickLoadIncomingDataHandler={onLoadIncomingDataHandler}
      />
    </Collapsible>
  )
}

AccessRequestsSection.propTypes = {
  service: object,
}

export default observer(AccessRequestsSection)
