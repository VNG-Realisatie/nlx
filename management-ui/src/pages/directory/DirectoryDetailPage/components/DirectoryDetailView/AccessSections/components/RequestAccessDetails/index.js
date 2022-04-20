// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, arrayOf, instanceOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'
import {
  IconOrganization,
  IconOutway,
  IconServices,
} from '../../../../../../../../icons'
import { SectionHeader, SectionContent, StyledIcon } from './index.styles'

const costFormatter = new Intl.NumberFormat('nl-NL', {
  style: 'currency',
  currency: 'EUR',
})

const RequestAccessDetails = ({
  service,
  publicKeyFingerprint,
  outwayNames,
}) => {
  const { t } = useTranslation()

  return (
    <>
      <p>{t('You are requesting access to a service')}.</p>

      <section>
        <SectionHeader>
          <StyledIcon as={IconOrganization} inline /> {t('Organization')}
        </SectionHeader>
        <SectionContent>{service.organization.name}</SectionContent>

        <SectionHeader withoutIcon>{t('OIN')}</SectionHeader>
        <SectionContent>{service.organization.serialNumber}</SectionContent>

        <SectionHeader>
          <StyledIcon as={IconServices} inline /> {t('Service')}
        </SectionHeader>
        <SectionContent>{service.serviceName}</SectionContent>

        {service.oneTimeCosts ? (
          <>
            <SectionHeader withoutIcon>{t('One time costs')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(service.oneTimeCosts)}
            </SectionContent>
          </>
        ) : null}

        {service.monthlyCosts ? (
          <>
            <SectionHeader withoutIcon>{t('Monthly costs')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(service.monthlyCosts)}
            </SectionContent>
          </>
        ) : null}

        {service.requestCosts ? (
          <>
            <SectionHeader withoutIcon>{t('Cost per request')}</SectionHeader>
            <SectionContent>
              {costFormatter.format(service.requestCosts)}
            </SectionContent>
          </>
        ) : null}

        <SectionHeader>
          <StyledIcon as={IconOutway} inline />{' '}
          {t('Outways', { count: outwayNames.length })}
        </SectionHeader>
        <SectionContent>{outwayNames.join(', ')}</SectionContent>

        <SectionHeader withoutIcon>{t('Public Key Fingerprint')}</SectionHeader>
        <SectionContent>{publicKeyFingerprint}</SectionContent>
      </section>
    </>
  )
}

RequestAccessDetails.propTypes = {
  service: instanceOf(DirectoryServiceModel),
  publicKeyFingerprint: string,
  outwayNames: arrayOf(string),
}

RequestAccessDetails.defaultProps = {
  outwayNames: [],
}

export default observer(RequestAccessDetails)
