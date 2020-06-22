// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { shape, string } from 'prop-types'
import { Button, Spinner } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { AccessRequestContext } from '../../../DirectoryPage'
import AccessRequestMessage from '../../../DirectoryPage/components/AccessRequestMessage'
import { SectionGroup } from '../../../../components/DetailView'

import { IconKey } from '../../../../icons'
import { AccessSection, IconItem, StatusItem } from './index.styles'

const DirectoryDetailView = ({
  organizationName,
  serviceName,
  latestAccessRequest,
}) => {
  const { t } = useTranslation()
  const { handleRequestAccess, requestSentTo } = useContext(
    AccessRequestContext,
  )

  const requestAccess = () =>
    handleRequestAccess({ organizationName, serviceName })

  const isRequestSentForThisService =
    requestSentTo.organizationName === organizationName &&
    requestSentTo.serviceName === serviceName

  let icon = Spinner
  if (
    latestAccessRequest &&
    ['FAILED', 'SENT'].includes(latestAccessRequest.status)
  ) {
    icon = IconKey
  }

  return (
    <SectionGroup>
      <AccessSection data-testid="request-access-section">
        {latestAccessRequest ? (
          <>
            <IconItem as={icon} />
            <StatusItem>
              <AccessRequestMessage latestAccessRequest={latestAccessRequest} />
              {icon === Spinner && '...'}
            </StatusItem>
          </>
        ) : (
          <>
            <IconItem as={IconKey} />
            <StatusItem>{t('You have no access')}</StatusItem>
            <Button
              onClick={requestAccess}
              disabled={isRequestSentForThisService}
            >
              {t('Request Access')}
            </Button>
          </>
        )}
      </AccessSection>
    </SectionGroup>
  )
}

DirectoryDetailView.propTypes = {
  organizationName: string,
  serviceName: string,
  latestAccessRequest: shape({
    id: string,
    status: string,
    createdAt: string,
    updatedAt: string,
  }),
}

export default DirectoryDetailView
