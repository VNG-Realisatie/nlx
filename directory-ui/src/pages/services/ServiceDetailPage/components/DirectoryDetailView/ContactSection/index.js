// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Icon } from '@commonground/design-system'
import {
  DetailHeading,
  DetailBody,
} from '../../../../../../components/DetailView'
import { IconMail } from '../../../../../../icons'
import { ContactLink } from './index.styles'

const ContactSection = ({ service }) => {
  const email = service && service.contactEmailAddress

  return (
    <>
      <section>
        <DetailHeading>
          <Icon as={IconMail} />
          Support
        </DetailHeading>

        {email ? (
          <DetailBody>
            <small>Publieke support</small>
            <br />
            <ContactLink href={`mailto:${email}`}>{email}</ContactLink>
          </DetailBody>
        ) : (
          <DetailBody>
            <small>Geen contactgegevens beschikbaar</small>
          </DetailBody>
        )}
      </section>
    </>
  )
}

ContactSection.propTypes = {
  service: shape({
    contactEmailAddress: string,
  }),
}

export default ContactSection
