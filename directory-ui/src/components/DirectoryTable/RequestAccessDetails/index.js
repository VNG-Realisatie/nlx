// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number, string } from 'prop-types'
import {
  SectionHeader,
  SectionContentService,
  SectionContent,
  // StyledIconServices,
  ServiceData,
} from './index.styles'

const costFormatter = new Intl.NumberFormat('nl-NL', {
  style: 'currency',
  currency: 'EUR',
})

const RequestAccessDetails = ({
  organizationName,
  serviceName,
  oneTimeCosts,
  monthlyCosts,
  requestCosts,
}) => {
  return (
    <>
      <p>You are requesting access to a service</p>

      <section>
        <SectionHeader>Service</SectionHeader>
        <SectionContentService>
          {/* <StyledIconServices /> */}
          <ServiceData>
            <strong>{serviceName}</strong>
            <span>{organizationName}</span>
          </ServiceData>
        </SectionContentService>

        {oneTimeCosts ? (
          <>
            <SectionHeader>One time costs</SectionHeader>
            <SectionContent>
              {costFormatter.format(oneTimeCosts)}
            </SectionContent>
          </>
        ) : null}

        {monthlyCosts ? (
          <>
            <SectionHeader>Monthly costs</SectionHeader>
            <SectionContent>
              {costFormatter.format(monthlyCosts)}
            </SectionContent>
          </>
        ) : null}

        {requestCosts ? (
          <>
            <SectionHeader>Cost per request</SectionHeader>
            <SectionContent>
              {costFormatter.format(requestCosts)}
            </SectionContent>
          </>
        ) : null}
      </section>
    </>
  )
}

RequestAccessDetails.propTypes = {
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  oneTimeCosts: number,
  monthlyCosts: number,
  requestCosts: number,
}

export default RequestAccessDetails
