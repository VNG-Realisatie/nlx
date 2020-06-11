// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, func } from 'prop-types'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { DetailHeading, SectionGroup } from '../../../../components/DetailView'

import { Authorization } from './index.styles'
import { ReactComponent as IconKey } from './whitelist.svg'

const DirectoryDetailView = ({ service, onRequestAccess }) => {
  const { t } = useTranslation()

  return (
    <SectionGroup>
      <Authorization>
        <DetailHeading>
          <IconKey />
          {t('Authorization')}
        </DetailHeading>
        <Button onClick={onRequestAccess}>{t('Request Access')}</Button>
      </Authorization>
    </SectionGroup>
  )
}

DirectoryDetailView.propTypes = {
  service: shape({
    serviceName: string,
    organizationName: string,
    apiSpecificationType: string,
    status: string,
  }),
  onRequestAccess: func,
}

export default DirectoryDetailView
