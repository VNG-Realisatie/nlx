// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, number } from 'prop-types'
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
    organization: string.isRequired,
    status: string.isRequired,
    oneTimeCosts: number,
    monthlyCosts: number,
    requestCosts: number,
  }),
}

export default observer(withDrawerStack(DirectoryDetailView))
