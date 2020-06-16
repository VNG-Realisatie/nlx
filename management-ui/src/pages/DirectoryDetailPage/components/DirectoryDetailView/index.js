// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, bool } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { DetailHeading, SectionGroup } from '../../../../components/DetailView'

import { Authorization } from './index.styles'
import { ReactComponent as IconKey } from './whitelist.svg'

const DirectoryDetailView = ({ onRequestAccess, isAccessRequested }) => {
  const { t } = useTranslation()

  return (
    <SectionGroup data-testid="request-access-section">
      <Authorization>
        <DetailHeading>
          <IconKey />
          {t('Authorization')}
        </DetailHeading>
        <Button onClick={onRequestAccess} disabled={isAccessRequested}>
          {t('Request Access')}
        </Button>
      </Authorization>
    </SectionGroup>
  )
}

DirectoryDetailView.propTypes = {
  onRequestAccess: func,
  isAccessRequested: bool.isRequired,
}

export default DirectoryDetailView
