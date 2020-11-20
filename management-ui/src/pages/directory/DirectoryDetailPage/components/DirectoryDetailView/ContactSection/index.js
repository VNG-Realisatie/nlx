// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, text } from 'prop-types'
import { useTranslation } from 'react-i18next'

import {
  DetailHeading,
  DetailBody,
} from '../../../../../../components/DetailView'
import { IconMail } from '../../../../../../icons'
import { ContactLink } from './index.styles'

const ContactSection = ({ service }) => {
  const { t } = useTranslation()
  const email = service && service.publicSupportContact

  return (
    <section>
      <DetailHeading>
        <IconMail />
        Support
      </DetailHeading>

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
    </section>
  )
}

ContactSection.propTypes = {
  service: shape({
    publicSupportContact: text,
  }),
}

export default ContactSection
