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
import CollapsibleBody from './CollapsibleBody'
import CollapsibleHeader from './CollapsibleHeader'

const AccessGrantSection = ({ service }) => {
  const { t } = useTranslation()
  const rootStore = useStores()
  const [isOpen, setIsOpen] = useState(false)
  const [showUpdateUiButton, setShowUpdateUiButton] = useState(false)

  const [pausePollingClosed, startPollingClosed] = usePolling(() => {
    rootStore.accessGrantStore.fetchForService(service)
  })

  const [pausePollingOpen, startPollingOpen] = usePolling(async () => {
    const haveAccessGrantsChanged = await rootStore.accessGrantStore.haveChangedForService(
      service,
    )

    setShowUpdateUiButton(haveAccessGrantsChanged)

    if (haveAccessGrantsChanged) {
      pausePollingOpen()
    }
  })

  const loadIncomingData = async () => {
    setShowUpdateUiButton(false)

    await rootStore.accessGrantStore.fetchForService(service)

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

  const handleOnToggle = (isCollapsibleOpen) => {
    setIsOpen(isCollapsibleOpen)
  }

  return (
    <Collapsible
      onToggle={handleOnToggle}
      title={<CollapsibleHeader counter={service.accessGrants.length} />}
      ariaLabel={t('Organizations with access')}
    >
      <CollapsibleBody
        accessGrants={service.accessGrants}
        showLoadIncomingDataButton={showUpdateUiButton}
        onClickLoadIncomingDataHandler={loadIncomingData}
      />
    </Collapsible>
  )
}

AccessGrantSection.propTypes = {
  service: object.isRequired,
}

export default observer(AccessGrantSection)
