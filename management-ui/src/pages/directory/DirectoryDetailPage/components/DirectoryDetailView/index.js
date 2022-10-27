// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { SectionGroup } from '../../../../../components/DetailView'
import CostsSection from '../../../../../components/CostsSection'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import usePollingEffect from '../../../../../hooks/use-polling-effect'
import ExternalLinkSection from './ExternalLinkSection'
import ContactSection from './ContactSection'
import {
  OutwaysWithoutAccessSection,
  OutwaysWithAccessSection,
} from './AccessSections'

const DirectoryDetailView = ({ service }) => {
  const { showToast } = useContext(ToasterContext)
  const { t } = useTranslation()
  const [isPollingPaused, setIsPollingPaused] = useState(false)

  usePollingEffect(
    async () => {
      if (isPollingPaused) {
        return
      }

      try {
        await service.fetch()
      } catch (error) {
        showToast({
          title: t('Failed to retrieve service details'),
          variant: 'error',
        })
      }

      try {
        await service.syncOutgoingAccessRequests()
      } catch (error) {
        showToast({
          title: t('Failed to synchronize access states'),
          body: t('The organization (Inway) might be unavailable.'),
          variant: 'error',
        })
      }
    },
    [isPollingPaused],
    {
      interval: 3000,
    },
  )

  const onShowConfirmRequestAccessModalHandler = () => {
    setIsPollingPaused(true)
  }

  const onHideConfirmRequestAccessModalHandler = () => {
    setIsPollingPaused(false)
  }

  return (
    <>
      <ExternalLinkSection service={service} />

      <SectionGroup>
        <ContactSection service={service} />

        <OutwaysWithoutAccessSection
          service={service}
          onShowConfirmRequestAccessModalHandler={
            onShowConfirmRequestAccessModalHandler
          }
          onHideConfirmRequestAccessModalHandler={
            onHideConfirmRequestAccessModalHandler
          }
        />

        <OutwaysWithAccessSection service={service} />

        <CostsSection
          oneTimeCosts={service.oneTimeCosts}
          monthlyCosts={service.monthlyCosts}
          requestCosts={service.requestCosts}
        />
      </SectionGroup>
    </>
  )
}

DirectoryDetailView.propTypes = {
  service: instanceOf(DirectoryServiceModel),
}

export default observer(DirectoryDetailView)
