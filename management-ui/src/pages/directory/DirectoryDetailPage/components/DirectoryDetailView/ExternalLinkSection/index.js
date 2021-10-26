// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, func } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import EnvironmentRepository from '../../../../../../domain/environment-repository'
import usePromise from '../../../../../../hooks/use-promise'
import { Section, StyledIconExternalLink } from './index.styles'

const getSpecificationUrl = (
  baseUrl,
  organizationSerialNumber,
  serviceName,
) => {
  const directoryHostname = `https://${baseUrl}`
  const specUri = `api/organizations/${organizationSerialNumber}/services/${serviceName}/api-spec`
  const directoryDocUrl = `${directoryHostname}/${specUri}`

  const redocUrl = new URL('https://redocly.github.io/redoc/')
  redocUrl.searchParams.set('url', directoryDocUrl)
  return redocUrl
}

const ExternalLinkSection = ({ service, getEnv }) => {
  const { documentationURL, organization, serviceName, apiSpecificationType } =
    service

  const { result } = usePromise(getEnv)
  const { t } = useTranslation()

  if (!result) {
    return null
  }

  const specificationUrl = getSpecificationUrl(
    result.directoryInspectionAddress,
    organization.serialNumber,
    serviceName,
  )

  const shouldShowSpecificationButton = !!apiSpecificationType

  return (
    <Section>
      <Button
        variant="secondary"
        as="a"
        href={documentationURL}
        target="_blank"
        role="button"
        aria-disabled={!documentationURL}
        disabled={!documentationURL}
      >
        {t('Documentation')}
        <StyledIconExternalLink $disabled={!documentationURL} />
      </Button>

      <Button
        variant="secondary"
        as="a"
        href={specificationUrl}
        target="_blank"
        role="button"
        aria-disabled={!shouldShowSpecificationButton}
        disabled={!shouldShowSpecificationButton}
      >
        {t('Specification')}
        <StyledIconExternalLink $disabled />
      </Button>
    </Section>
  )
}

ExternalLinkSection.propTypes = {
  getEnv: func,
  service: shape({
    documentationURL: string,
    serviceName: string,
    organization: shape({
      serialNumber: string,
      name: string,
    }),
    apiSpecificationType: string,
  }),
}

ExternalLinkSection.defaultProps = {
  getEnv: EnvironmentRepository.getCurrent,
}

export default ExternalLinkSection
