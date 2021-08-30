// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Button } from '@commonground/design-system'

import { Section, StyledIconExternalLink } from './index.styles'

const ExternalLinkSection = ({ service }) => {
  const { documentationUrl, organization, name, apiType } = service

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
        href={`/documentation/${organization}/${name}`}
        role="button"
        rel="noreferrer"
        aria-disabled={!apiType}
        disabled={!apiType}
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
    organization: string,
    name: string,
    apiType: string,
  }),
}

export default ExternalLinkSection
