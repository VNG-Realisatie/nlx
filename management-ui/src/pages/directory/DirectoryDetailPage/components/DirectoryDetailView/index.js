// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../components/DetailView'
import CostsSection from '../../../../../components/CostsSection'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import usePolling from '../../../../../hooks/use-polling'
import ExternalLinkSection from './ExternalLinkSection'
import ContactSection from './ContactSection'
import {
  OutwaysWithoutAccessSection,
  OutwaysWithAccessSection,
} from './AccessSections'

const DirectoryDetailView = ({ service }) => {
  const [isLoading, setIsLoading] = useState(false)
  const [pauseFetchPolling, continueFetchPolling] = usePolling(async () => {
    if (isLoading) {
      return
    }

    setIsLoading(true)

    await service.syncOutgoingAccessRequests()
    await service.fetch()

    setIsLoading(false)
  })

  const onShowConfirmRequestAccessModalHandler = () => {
    pauseFetchPolling()
  }

  const onHideConfirmRequestAccessModalHandler = () => {
    continueFetchPolling()
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
