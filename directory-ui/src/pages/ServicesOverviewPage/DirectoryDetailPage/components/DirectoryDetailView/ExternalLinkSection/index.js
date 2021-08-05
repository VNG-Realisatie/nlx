// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Button } from '@commonground/design-system'

import { Section, StyledIconExternalLink } from './index.styles'

const ExternalLinkSection = ({ service }) => {
  const { documentationUrl } = service

  return (
    <Section>
      <Button
        variant="secondary"
        as="a"
        href={documentationUrl}
        target="_blank"
        role="button"
        aria-disabled={!documentationUrl}
        disabled={!documentationUrl}
      >
        Documentatie
        <StyledIconExternalLink $disabled={!documentationUrl} />
      </Button>

      <Button
        variant="secondary"
        as="a"
        href=""
        target="_blank"
        role="button"
        aria-disabled="true"
        disabled
      >
        Specificatie
        <StyledIconExternalLink $disabled />
      </Button>
    </Section>
  )
}

ExternalLinkSection.propTypes = {
  service: shape({
    documentationUrl: string,
  }),
}

export default ExternalLinkSection
