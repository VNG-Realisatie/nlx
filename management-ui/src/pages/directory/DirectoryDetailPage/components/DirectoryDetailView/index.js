// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { SectionGroup } from '../../../../../components/DetailView'
import CostsSection from '../../../../../components/CostsSection'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import { useConfirmationModal } from '../../../../../components/ConfirmationModal'
import RequestAccessDetails from '../../../RequestAccessDetails'
import usePolling from '../../../../../hooks/use-polling'
import ExternalLinkSection from './ExternalLinkSection'
import ContactSection from './ContactSection'
import OutwaysWithoutAccessSection from './OutwaysWithoutAccessSection'
import OutwaysWithAccessSection from './OutwaysWithAccessSection'

const DirectoryDetailView = ({ service }) => {
  const { t } = useTranslation()

  const [RequestConfirmationModal, confirmRequest] = useConfirmationModal({
    title: t('Request access'),
    okText: t('Send'),
    children: (
      <RequestAccessDetails
        organization={service.organization}
        serviceName={service.serviceName}
        oneTimeCosts={service.oneTimeCosts}
        monthlyCosts={service.monthlyCosts}
        requestCosts={service.requestCosts}
      />
    ),
  })

  const [pauseFetchPolling, continueFetchPolling] = usePolling(service.fetch)

  const onRequestAccess = async (publicKeyFingerprint) => {
    pauseFetchPolling()

    if (await confirmRequest(publicKeyFingerprint)) {
      service.requestAccess(publicKeyFingerprint)
    }

    continueFetchPolling()
  }

  const onRetryRequestAccess = (publicKeyFingerprint) => {
    service.retryRequestAccess(publicKeyFingerprint)
  }

  return (
    <>
      <ExternalLinkSection service={service} />

      <SectionGroup>
        <OutwaysWithoutAccessSection
          service={service}
          requestAccessHandler={onRequestAccess}
          retryRequestAccessHandler={onRetryRequestAccess}
        />

        <OutwaysWithAccessSection service={service} />

        <ContactSection service={service} />
        <CostsSection
          oneTimeCosts={service.oneTimeCosts}
          monthlyCosts={service.monthlyCosts}
          requestCosts={service.requestCosts}
        />
      </SectionGroup>

      <RequestConfirmationModal />
    </>
  )
}

DirectoryDetailView.propTypes = {
  service: instanceOf(DirectoryServiceModel),
}

export default observer(DirectoryDetailView)
