// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Button } from '@commonground/design-system'

import { Section, StyledIconExternalLink } from './index.styles'

const getSpecificationUrl = (organizationSerialNumber, serviceName) => {
  const directoryHostname = `${window.location.protocol}//${window.location.hostname}`
  const specUri = `api/organizations/${organizationSerialNumber}/services/${serviceName}/api-spec`
  const directoryDocUrl = `${directoryHostname}/${specUri}`

  const redocUrl = new URL('https://redocly.github.io/redoc/')
  redocUrl.searchParams.set('url', directoryDocUrl)
  return redocUrl
}

const ExternalLinkSection = ({ service }) => {
  const { documentationUrl, organization, name, apiType } = service
  const specificationUrl = getSpecificationUrl(organization.serialNumber, name)

  return (
    <Section>
      <Button
        variant="secondary"
        as="a"
        href={documentationUrl}
        target="_blank"
        role="button"
        rel="noreferrer"
        aria-disabled={!documentationUrl}
        disabled={!documentationUrl}
      >
        Documentatie
        <StyledIconExternalLink $disabled={!documentationUrl} />
      </Button>

      <Button
        variant="secondary"
        as="a"
        href={apiType ? specificationUrl : ''}
        target="_blank"
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
    organization: shape({
      name: string,
      serialNumber: string,
    }),
    name: string,
    apiType: string,
  }),
}

export default ExternalLinkSection
