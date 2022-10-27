// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, shape } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { observer } from 'mobx-react'
import {
  IconOrganization,
  IconServices,
} from '../../../../../../../../../icons'
import { SectionHeader, SectionContent, StyledIcon } from './index.styles'

const AccessDetails = ({
  subTitle,
  serviceName,
  publicKeyFingerprint,
  organization,
}) => {
  const { t } = useTranslation()

  return (
    <>
      <p>{subTitle}</p>

      <section>
        <SectionHeader>
          <StyledIcon as={IconOrganization} inline /> {t('Organization')}
        </SectionHeader>
        <SectionContent>{organization.name}</SectionContent>

        <SectionHeader withoutIcon>{t('OIN')}</SectionHeader>
        <SectionContent>{organization.serialNumber}</SectionContent>

        <SectionHeader withoutIcon>{t('Public Key Fingerprint')}</SectionHeader>
        <SectionContent>{publicKeyFingerprint}</SectionContent>

        <SectionHeader>
          <StyledIcon as={IconServices} inline /> {t('Service')}
        </SectionHeader>
        <SectionContent>{serviceName}</SectionContent>
      </section>
    </>
  )
}

AccessDetails.propTypes = {
  subTitle: string.isRequired,
  serviceName: string.isRequired,
  publicKeyFingerprint: string.isRequired,
  organization: shape({
    serialNumber: string.isRequired,
    name: string.isRequired,
  }),
}

export default observer(AccessDetails)
