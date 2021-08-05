// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, text } from 'prop-types'

import {
  DetailHeading,
  DetailBody,
} from '../../../../../../components/DetailView'
import { IconMail } from '../../../../../../icons'
import { ContactLink } from './index.styles'

const ContactSection = ({ service }) => {
  const email = service && service.contactEmailAddress

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
          <small>Geen contactgegevens beschikbaar</small>
        </DetailBody>
      )}
    </section>
  )
}

ContactSection.propTypes = {
  service: shape({
    contactEmailAddress: text,
  }),
}

export default ContactSection
