// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { Section, StyledIconExternalLink } from './index.styles'

const ExternalLinkSection = ({ service }) => {
  const { documentationUrl } = service
  const { t } = useTranslation()

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
        {t('Documentation')}
        <StyledIconExternalLink $disabled={!documentationUrl} />
      </Button>
    </Section>
  )
}

ExternalLinkSection.propTypes = {
  service: shape({
    documentationUrl: string,
    organization: shape({
      serialNumber: string,
      name: string,
    }),
  }),
}

export default ExternalLinkSection
