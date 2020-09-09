// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { inwayModelPropTypes } from '../../../models/InwayModel'
import InwayDetailPageView from './InwayDetailPageView'

const InwayDetailPage = ({ parentUrl, inway }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const close = () => history.push(parentUrl)

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        as="header"
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="gateway-name"
      />

      <Drawer.Content>
        {inway ? (
          <InwayDetailPageView inway={inway} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the details for this inway.', { name })}
          </Alert>
        )}
      </Drawer.Content>
    </Drawer>
  )
}

InwayDetailPage.propTypes = {
  parentUrl: string,
  inway: shape(inwayModelPropTypes),
}

InwayDetailPage.defaultProps = {
  parentUrl: '/inways',
}

export default InwayDetailPage
