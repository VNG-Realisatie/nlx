// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Collapsible } from '@commonground/design-system'
import {
  DetailHeading,
  DetailBody,
  StyledCollapsibleBody,
} from '../../../../../../components/DetailView'
import { IconMail } from '../../../../../../icons'
import { ContactLink } from './index.styles'

const Heading = () => {
  const { t } = useTranslation()

  return (
    <DetailHeading>
      <IconMail />
      {t('Support')}
    </DetailHeading>
  )
}

const ContactSection = ({ service }) => {
  const { t } = useTranslation()
  const email = service && service.publicSupportContact

  return (
    <Collapsible title={<Heading />} ariaLabel={t('Support')}>
      <StyledCollapsibleBody>
        {email ? (
          <DetailBody>
            <small>Public support</small>
            <br />
            <ContactLink href={`mailto:${email}`}>{email}</ContactLink>
          </DetailBody>
        ) : (
          <DetailBody>
            <small>{t('No contact details available')}</small>
          </DetailBody>
        )}
      </StyledCollapsibleBody>
    </Collapsible>
  )
}

ContactSection.propTypes = {
  service: shape({
    publicSupportContact: string,
  }),
}

export default ContactSection
