// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { withDrawerStack } from '@commonground/design-system'
import { SectionGroup } from '../../../../../components/DetailView'
import CostsSection from '../../../../../components/CostsSection'
import ExternalLinkSection from './ExternalLinkSection'
import ContactSection from './ContactSection'

const DirectoryDetailView = ({ service }) => {
  return (
    <>
      <ExternalLinkSection service={service} />

      <SectionGroup>
        {/* TODO: Contact section is missing tech support mail address */}
        {/* NOTE: This might be an issue for NLX management as well */}
        <ContactSection service={service} />
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
  service: shape({
    // TODO: fetch costs from API
    // oneTimeCosts: number,
    // monthlyCosts: number,
    // requestCosts: number,

    apiType: string,
    contactEmailAddress: string,
    documentationUrl: string,
    name: string.isRequired,
    organization: string.isRequired,
    status: string.isRequired,
  }),
}

export default observer(withDrawerStack(DirectoryDetailView))
