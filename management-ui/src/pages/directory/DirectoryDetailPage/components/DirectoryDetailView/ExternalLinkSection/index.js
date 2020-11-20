// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { Section, StyledIconExternalLink } from './index.styles'

const ExternalLinkSection = ({ service }) => {
  const { documentationURL } = service
  const { t } = useTranslation()

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
        href=""
        target="_blank"
        role="button"
        aria-disabled="true"
        disabled
      >
        {t('Specification')}
        <StyledIconExternalLink $disabled />
      </Button>
    </Section>
  )
}

ExternalLinkSection.propTypes = {
  service: shape({
    documentationURL: string,
  }),
}

export default ExternalLinkSection
