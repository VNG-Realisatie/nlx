// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number, string, shape, arrayOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import {
  SectionHeader,
  SectionContent,
  StyledIconServices,
  ServiceData,
  SectionContentWithPadding,
  StyledIconOutways,
} from './index.styles'

const costFormatter = new Intl.NumberFormat('nl-NL', {
  style: 'currency',
  currency: 'EUR',
})

const RequestAccessDetails = ({
  organization,
  serviceName,
  oneTimeCosts,
  monthlyCosts,
  requestCosts,
  publicKeyFingerprint,
  outwayNames,
}) => {
  const { t } = useTranslation()

  return (
    <>
      <p>{t('You are requesting access to a service')}.</p>

      <section>
        <SectionHeader>{t('Service')}</SectionHeader>
        <SectionContentWithPadding>
          <StyledIconServices />
          <ServiceData>
            <strong>{serviceName}</strong>
            <span>
              {organization.name} ({organization.serialNumber})
            </span>
          </ServiceData>
        </SectionContentWithPadding>

        {oneTimeCosts ? (
          <>
            <SectionHeader>{t('One time costs')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(oneTimeCosts)}
            </SectionContent>
          </>
        ) : null}

        {monthlyCosts ? (
          <>
            <SectionHeader>{t('Monthly costs')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(monthlyCosts)}
            </SectionContent>
          </>
        ) : null}

        {requestCosts ? (
          <>
            <SectionHeader>{t('Cost per request')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(requestCosts)}
            </SectionContent>
          </>
        ) : null}

        <SectionHeader>{t('Outways')}</SectionHeader>
        <SectionContentWithPadding>
          <StyledIconOutways />
          <ServiceData>
            <p>
              <strong>{outwayNames.join(', ')}</strong>
            </p>
            <p>{publicKeyFingerprint}</p>
          </ServiceData>
        </SectionContentWithPadding>
      </section>
    </>
  )
}

RequestAccessDetails.propTypes = {
  organization: shape({
    serialNumber: string.isRequired,
    name: string.isRequired,
  }).isRequired,
  serviceName: string.isRequired,
  oneTimeCosts: number,
  monthlyCosts: number,
  requestCosts: number,
  publicKeyFingerprint: string,
  outwayNames: arrayOf(string),
}

RequestAccessDetails.defaultProps = {
  outwayNames: [],
}

export default RequestAccessDetails
