// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, number } from 'prop-types'
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
    apiType: string,
    contactEmailAddress: string,
    documentationUrl: string,
    name: string.isRequired,
    organization: shape({
      serialNumber: string.isRequired,
      name: string.isRequired,
    }).isRequired,
    status: string.isRequired,
    oneTimeCosts: number,
    monthlyCosts: number,
    requestCosts: number,
  }),
}

export default withDrawerStack(DirectoryDetailView)
