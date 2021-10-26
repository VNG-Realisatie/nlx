// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import OutwayModel from '../../../stores/models/OutwayModel'
import OutwayDetailPageView from './OutwayDetailPageView'

const OutwayDetailPage = ({ parentUrl, outway }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const close = () => history.push(parentUrl)

  return (
    <Drawer noMask closeHandler={close} data-testid="outway-detail-page">
      <Drawer.Header
        as="header"
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="outway-name"
      />

      <Drawer.Content>
        {outway ? (
          <OutwayDetailPageView outway={outway} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the details for this outway', { name })}
          </Alert>
        )}
      </Drawer.Content>
    </Drawer>
  )
}

OutwayDetailPage.propTypes = {
  parentUrl: string,
  outway: instanceOf(OutwayModel),
}

OutwayDetailPage.defaultProps = {
  parentUrl: '/inways-and-outways',
}

export default OutwayDetailPage
